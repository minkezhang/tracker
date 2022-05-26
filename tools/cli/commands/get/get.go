package get

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/database"
	"strings"

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
			dpb.Corpus_value[fmt.Sprintf("CORPUS_%v", strings.ToUpper(c.corpus))])

		entries = append(entries, c.db.Search(database.O{Title: c.title, Corpus: corpus})...)
	}

	for _, epb := range entries {
		data, _ := ce.E{}.Marshal(epb)
		fmt.Printf("%s\n", data)
	}

	return subcommands.ExitSuccess
}
