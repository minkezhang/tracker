package get

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
	ce "github.com/minkezhang/tracker/formats/cli"
)

type C struct {
	db *database.DB

	id     string
	title  string
	corpus string
}

func New(db *database.DB) *C { return &C{db: db} }

func (c *C) Name() string     { return "get" }
func (c *C) Synopsis() string { return "get entry from database with matching parameters" }
func (c *C) Usage() string    { return c.Synopsis() }

func (c *C) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.id, "id", "", "entry ID")
	f.StringVar(&c.title, "title", "", "entry title substring")
	f.StringVar(&c.corpus, "corpus", "unknown", "optional corpus hint for the entry")
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var entries []*dpb.Entry
	if c.id != "" {
		results, _ := c.db.GetEntry(c.id)
		entries = append(entries, results)
	} else {
		corpus := dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", c.corpus)])

		entries = append(entries, c.db.Search(database.O{Title: c.title, Corpus: corpus})...)
	}

	if len(entries) == 0 {
		fmt.Printf("Could not find result with the given input. Please refine your search.\n")
		return subcommands.ExitFailure
	}

	if len(entries) > 1 {
		fmt.Printf("Too many results returned. Please refine your search.\n")
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(entries[0])
	fmt.Printf("%s\n", e.Data)

	return subcommands.ExitSuccess
}
