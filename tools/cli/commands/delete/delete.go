package del

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/tools/cli/commands/get/common"

	ce "github.com/minkezhang/truffle/formats/cli"
	se "github.com/minkezhang/truffle/formats/cli/struct"
)

type C struct {
	db *database.DB

	title  *se.Title
	id     *se.ID
	corpus *se.Corpus
}

func New(db *database.DB) *C {
	return &C{
		db: db,

		title:  &se.Title{},
		id:     &se.ID{},
		corpus: &se.Corpus{},
	}
}

func (c *C) Name() string     { return "delete" }
func (c *C) Synopsis() string { return "delete matching entry from database" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	c.title.SetFlags(f)
	c.id.SetFlags(f)
	c.corpus.SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := common.Get(common.O{
		DB:     c.db,
		ID:     c.id.ID,
		Title:  c.title.Title,
		Corpus: c.corpus.Corpus,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	epb, err = c.db.Delete(epb.GetId().GetId())
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
