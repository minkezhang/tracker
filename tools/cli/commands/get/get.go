package get

import (
	"context"
	"flag"
	"fmt"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/database/helper/get"
	"github.com/minkezhang/truffle/formats/cli/struct"
	"github.com/minkezhang/truffle/tools/cli/flag/flagset"

	ce "github.com/minkezhang/truffle/formats/cli"
)

type C struct {
	db    *database.DB
	entry *entry.E
}

func New(db *database.DB) *C {
	return &C{
		db:    db,
		entry: &entry.E{},
	}
}

func (c *C) Name() string     { return "get" }
func (c *C) Synopsis() string { return "get entry from database with matching parameters" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	(*flagset.Title)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.ID)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Corpus)(unsafe.Pointer(c.entry)).SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := c.entry.PB()
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	epb, err = get.Get(ctx, c.db, epb)
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
