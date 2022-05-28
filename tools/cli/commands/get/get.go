package get

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/database"
	"github.com/minkezhang/tracker/tools/cli/commands/get/common"

	ce "github.com/minkezhang/tracker/formats/cli"
	se "github.com/minkezhang/tracker/formats/cli/struct"
)

type C struct {
	db *database.DB

	title  *se.Title
	id     *se.ID
	corpus *se.Corpus
}

func New(db *database.DB) *C {
	return &C{
		db:     db,
		title:  &se.Title{},
		id:     &se.ID{},
		corpus: &se.Corpus{},
	}
}

func (c *C) Name() string     { return "get" }
func (c *C) Synopsis() string { return "get entry from database with matching parameters" }
func (c *C) Usage() string    { return c.Synopsis() }

func (c *C) SetFlags(f *flag.FlagSet) {
	c.title.SetFlags(f)
	c.id.SetFlags(f)
	c.corpus.SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := common.Get(common.O{
		DB:     c.db,
		Title:  c.title.Title,
		ID:     c.id.ID,
		Corpus: c.corpus.Corpus,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
