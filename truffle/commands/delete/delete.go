package del

import (
	"context"
	"flag"
	"fmt"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/database/helper/get"
	"github.com/minkezhang/truffle/truffle/flag/entry"
	"github.com/minkezhang/truffle/truffle/flag/flagset"

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

func (c *C) Name() string     { return "delete" }
func (c *C) Synopsis() string { return "delete matching entry from database" }
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

	epb, err = c.db.Delete(ctx, epb.GetId())
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
