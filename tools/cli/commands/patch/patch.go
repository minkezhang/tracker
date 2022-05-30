package patch

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/tools/cli/commands/patch/common"

	ce "github.com/minkezhang/truffle/formats/cli"
	se "github.com/minkezhang/truffle/formats/cli/struct"
)

type C struct {
	db *database.DB

	title *se.Title
	id    *se.ID

	body *se.Body
}

func New(db *database.DB) *C {
	return &C{
		db: db,

		title: &se.Title{},
		body:  &se.Body{},
		id:    &se.ID{},
	}
}

func (c *C) Name() string     { return "patch" }
func (c *C) Synopsis() string { return "patch entry with matching query parameters" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	c.title.SetFlags(f)
	c.body.SetFlags(f)
	c.id.SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := common.Patch(ctx, common.O{
		DB:    c.db,
		Title: c.title.Title,
		ID:    c.id.ID,
		Body:  c.body,
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
