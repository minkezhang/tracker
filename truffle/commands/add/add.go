package add

import (
	"context"
	"flag"
	"fmt"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/truffle/flag/entry"
	"github.com/minkezhang/truffle/truffle/flag/flagset"

	formatter "github.com/minkezhang/truffle/truffle/formatter/full/entry"
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

func (c *C) Name() string     { return "add" }
func (c *C) Synopsis() string { return "add entry to database" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	(*flagset.Corpus)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Titles)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Body)(unsafe.Pointer(c.entry)).SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := c.entry.PB()
	if err != nil {
		fmt.Printf("Could not load flags into data struct: %v\n", err)
		return subcommands.ExitFailure
	}

	epb, err = c.db.Add(ctx, epb)
	if err != nil {
		fmt.Printf("Could not add data to database: %v\n", err)
		return subcommands.ExitFailure
	}

	data, err := formatter.Format(epb)
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}
	fmt.Print(string(data))

	return subcommands.ExitSuccess
}
