package get

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/database"
	"github.com/minkezhang/tracker/tools/cli/commands/get/common"

	ce "github.com/minkezhang/tracker/formats/cli"
)

type C struct {
	c *common.C
}

func New(db *database.DB) *C {
	return &C{
		c: &common.C{DB: db},
	}
}

func (c *C) Name() string             { return "get" }
func (c *C) Synopsis() string         { return "get entry from database with matching parameters" }
func (c *C) Usage() string            { return c.Synopsis() }
func (c *C) SetFlags(f *flag.FlagSet) { c.c.SetFlags(f) }

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := c.c.Get()
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
