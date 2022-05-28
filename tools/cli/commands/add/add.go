package add

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/truffle/api/go/database"
	ce "github.com/minkezhang/truffle/formats/cli"
	se "github.com/minkezhang/truffle/formats/cli/struct"
)

type C struct {
	db *database.DB

	body   *se.Body
	titles *se.Titles
}

func New(db *database.DB) *C {
	return &C{
		db:     db,
		titles: &se.Titles{},
		body:   &se.Body{},
	}
}

func (c *C) Name() string     { return "add" }
func (c *C) Synopsis() string { return "add entry to database" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	c.titles.SetFlags(f)
	c.body.SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	s, err := c.titles.Load()
	if err != nil {
		fmt.Printf("Could not load flags into data struct: %v\n", err)
		return subcommands.ExitFailure
	}

	epb := s.(*dpb.Entry)

	s, err = c.body.Load()
	if err != nil {
		fmt.Printf("Could not load flags into data struct: %v\n", err)
	}

	proto.Merge(epb, s.(*dpb.Entry))

	if err := c.db.AddEntry(epb); err != nil {
		fmt.Printf("Could not add data to database: %v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
