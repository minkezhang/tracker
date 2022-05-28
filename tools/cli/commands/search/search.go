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
)

type C struct {
	db *database.DB

	title  string
	corpus string
}

func New(db *database.DB) *C { return &C{db: db} }

func (c *C) Name() string     { return "search" }
func (c *C) Synopsis() string { return "search across multiple databases with matching parameters" }
func (c *C) Usage() string    { return c.Synopsis() }

func (c *C) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.title, "title", "", "entry title substring")
	f.StringVar(&c.corpus, "corpus", "unknown", "optional corpus hint for the entry")
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	corpus := dpb.Corpus(
		dpb.Corpus_value[utils.ToEnum("CORPUS", c.corpus)])
	entries := c.db.Search(database.O{Title: c.title, Corpus: corpus})

	e := &ce.E{Format: ce.FormatShort}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, epb := range entries {
		e.Dump(epb)
		fmt.Fprint(w, string(e.Data))
	}
	w.Flush()

	return subcommands.ExitSuccess
}
