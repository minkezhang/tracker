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

func (c *C) Name() string             { return "add" }
func (c *C) Synopsis() string         { return "add entry to database" }
func (c *C) Usage() string            { return c.Synopsis() }
func (c *C) SetFlags(f *flag.FlagSet) { c.e.SetFlags(f) }

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
