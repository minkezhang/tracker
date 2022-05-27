package add

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
	ce "github.com/minkezhang/tracker/formats/cli"
	se "github.com/minkezhang/tracker/formats/cli/struct"
)

type C struct {
	db *database.DB

	e *se.E
}

func New(db *database.DB) *C {
	return &C{
		db: db,
		e:  &se.E{},
	}
}

func (c *C) Name() string     { return "add" }
func (c *C) Synopsis() string { return "add entry to database" }
func (c *C) Usage() string    { return c.Synopsis() }
func (c *C) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.e.Corpus, "corpus", "unknown", "corpus for the entry")
	f.Var(&c.e.Titles, "titles", "title of the entry")
	f.Var(&c.e.Providers, "providers", "providers of the entry")

	f.Float64Var(&c.e.Score, "score", 0, "score of the entry")
	f.BoolVar(&c.e.Queued, "queued", false, "if the entry is on the current watchlist")

	f.Var(&c.e.Directors, "directors", "directors of the entry for visual or game entries")
	f.Var(&c.e.Studios, "studios", "producing studios of the entry for visual or game entries")
	f.Var(&c.e.Writers, "writers", "writers of the entry for visual or game entries")
	f.Var(&c.e.Writers, "composers", "composers of the entry for album entries")
	f.Var(&c.e.Writers, "authors", "authors of the entry for literary entries")

	f.IntVar(&c.e.Season, "season", 0, "season of the entry for visual entries")
	f.IntVar(&c.e.Season, "volume", 0, "volume of the entry for literary entries")
	f.IntVar(&c.e.Episode, "episode", 0, "episode of the entry for visual entries")
	f.IntVar(&c.e.Episode, "chapter", 0, "chapter of the entry for literary entries")
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	s, err := c.e.Load()
	if err != nil {
		fmt.Printf("Could not load flags into data struct: %v\n", err)
		return subcommands.ExitFailure
	}

	epb := s.(*dpb.Entry)
	if err := c.db.AddEntry(epb); err != nil {
		fmt.Printf("Could not add data to database: %v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Printf("%s\n", e.Data)

	return subcommands.ExitSuccess
}
