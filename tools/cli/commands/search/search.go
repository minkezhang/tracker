package search

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/client/mal"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/formats/cli/struct"
	"github.com/minkezhang/truffle/tools/cli/commands/search/ordering"
	"github.com/minkezhang/truffle/tools/cli/flag/flagset"

	dpb "github.com/minkezhang/truffle/api/go/database"
	ce "github.com/minkezhang/truffle/formats/cli"
)

type C struct {
	db    *database.DB
	entry *entry.E

	apis     []dpb.API
	ordering *ordering.O
}

func New(db *database.DB) *C {
	return &C{
		db:    db,
		entry: &entry.E{},

		ordering: &ordering.O{},
	}
}

func (c *C) Name() string     { return "search" }
func (c *C) Synopsis() string { return "search across multiple databases with matching parameters" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	f.Func("apis", "APIs to use in the search operation, e.g. \"truffle\"", func(api string) error {
		c.apis = append(c.apis, dpb.API(
			dpb.API_value[utils.ToEnum("API", api)]))
		return nil
	})
	c.ordering.SetFlags(f)

	(*flagset.Title)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Corpus)(unsafe.Pointer(c.entry)).SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := c.entry.PB()
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	if len(c.ordering.Orderings) == 0 {
		c.ordering.Orderings = append(
			c.ordering.Orderings, "queued", "corpus", "score", "titles",
		)
	}

	if len(c.apis) == 0 {
		c.apis = []dpb.API{dpb.API_API_TRUFFLE}
	}

	var title string
	if len(epb.GetTitles()) > 0 {
		title = epb.GetTitles()[0]
	}

	entries, err := c.db.Search(ctx, database.SearchOpts{
		Title:  title,
		Corpus: epb.GetCorpus(),
		APIs:   c.apis,
		MAL: mal.SearchOpts{
			Cutoff: 2000,
		},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	entries, err = ordering.Order(entries, *c.ordering)
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{Format: ce.FormatShort}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, epb := range entries {
		e.Dump(epb)
		fmt.Fprint(w, string(e.Data))
	}
	w.Flush()

	return subcommands.ExitSuccess
}
