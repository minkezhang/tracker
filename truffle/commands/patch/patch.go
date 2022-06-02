package patch

import (
	"context"
	"flag"
	"fmt"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/database/helper/get"
	"github.com/minkezhang/truffle/database/helper/patch"
	"github.com/minkezhang/truffle/truffle/commands/common"
	"github.com/minkezhang/truffle/truffle/flag/entry"
	"github.com/minkezhang/truffle/truffle/flag/flagset"

	dpb "github.com/minkezhang/truffle/api/go/database"
	formatter "github.com/minkezhang/truffle/truffle/formatter/full/entry"
)

type C struct {
	common common.O
	db     *database.DB
	entry  *entry.E
}

func New(db *database.DB, common common.O) *C {
	return &C{
		common: common,
		db:     db,
		entry:  &entry.E{},
	}
}

func (c *C) Name() string     { return "patch" }
func (c *C) Synopsis() string { return "patch entry with matching query parameters" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	(*flagset.Title)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.ID)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Body)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Corpus)(unsafe.Pointer(c.entry)).SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := c.entry.PB()
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	fpb, err := get.Get(ctx, c.db, epb, []dpb.API{
		dpb.API_API_TRUFFLE,
	})
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	if !c.entry.SetQueued {
		epb.Queued = fpb.GetQueued()
	}

	epb, err = patch.Patch(ctx, c.db, epb)
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	data, err := formatter.Format(epb)
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}
	fmt.Fprint(c.common.Output, string(data))

	return subcommands.ExitSuccess
}
