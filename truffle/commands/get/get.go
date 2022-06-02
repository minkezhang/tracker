package get

import (
	"context"
	"flag"
	"fmt"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/database/helper/get"
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
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	var apis []dpb.API
	for a, _ := range dpb.API_name {
		if api := dpb.API(a); api != dpb.API_API_UNKNOWN {
			apis = append(apis, api)
		}
	}

	epb, err = get.Get(ctx, c.db, epb, apis)
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
