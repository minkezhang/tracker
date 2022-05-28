package search

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
	ce "github.com/minkezhang/tracker/formats/cli"
	se "github.com/minkezhang/tracker/formats/cli/struct"
	cf "github.com/minkezhang/tracker/tools/cli/flag"
)

type C struct {
	db *database.DB

	apis   cf.MultiString
	title  *se.Title
	corpus *se.Corpus
}

func New(db *database.DB) *C {
	return &C{
		db:     db,
		title:  &se.Title{},
		corpus: &se.Corpus{},
	}
}

func (c *C) Name() string     { return "search" }
func (c *C) Synopsis() string { return "search across multiple databases with matching parameters" }
func (c *C) Usage() string    { return c.Synopsis() }

func (c *C) SetFlags(f *flag.FlagSet) {
	f.Var(&c.apis, "apis", "APIs to use in the search operation, e.g. \"tracker\"")
	c.title.SetFlags(f)
	c.corpus.SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var apis []dpb.API
	if len(c.apis) == 0 {
		c.apis = append(c.apis, "tracker")
	}
	for _, api := range c.apis {
		apis = append(apis, dpb.API(
			dpb.API_value[utils.ToEnum("API", api)]))
	}

	s, _ := c.corpus.Load()
	entries, err := c.db.Search(database.O{
		Title:  c.title.Title,
		Corpus: s.(*dpb.Entry).GetCorpus(),
		APIs:   apis,
	})
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
