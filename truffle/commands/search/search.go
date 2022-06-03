package search

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"text/tabwriter"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/truffle/commands/common"
	"github.com/minkezhang/truffle/truffle/commands/search/ordering"
	"github.com/minkezhang/truffle/truffle/flag/entry"
	"github.com/minkezhang/truffle/truffle/flag/flagset"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dpb "github.com/minkezhang/truffle/api/go/database"
	formatter "github.com/minkezhang/truffle/truffle/formatter/short/entry"
)

type C struct {
	common common.O
	db     *database.DB
	entry  *entry.E

	apis      []dpb.API
	orderings []ordering.T
}

func New(db *database.DB, common common.O) *C {
	return &C{
		common: common,
		db:     db,
		entry:  &entry.E{},
	}
}

func (c *C) Name() string     { return "search" }
func (c *C) Synopsis() string { return "search across multiple databases with matching parameters" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	f.Func("api", "APIs to use in the search operation, e.g. \"truffle\"", func(api string) error {
		c.apis = append(c.apis, dpb.API(
			dpb.API_value[utils.ToEnum("API", api)]))
		return nil
	})

	f.Func("order", "list of fields to order by, e.g. \"title\"", func(order string) error {
		if o := ordering.L[strings.ToLower(order)]; o == ordering.OrderingUnknown {
			return status.Errorf(codes.InvalidArgument, "invalid ordering field specified %v", order)
		} else {
			c.orderings = append(c.orderings, o)
		}
		return nil
	})

	(*flagset.Title)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Corpus)(unsafe.Pointer(c.entry)).SetFlags(f)
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := c.entry.PB()
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	if len(c.orderings) == 0 {
		c.orderings = []ordering.T{
			ordering.OrderingQueued,
			ordering.OrderingCorpus,
			ordering.OrderingScore,
			ordering.OrderingTitles,
		}
	}

	if len(c.apis) == 0 {
		c.apis = []dpb.API{dpb.API_API_TRUFFLE}
	}

	var title string
	if len(epb.GetTitles()) > 0 {
		title = epb.GetTitles()[0]
	}

	entries, err := c.db.Search(ctx, database.SearchOpts{
		Title:  title,
		Corpus: epb.GetCorpus(),
		APIs:   c.apis,
	})
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	entries, err = ordering.Order(entries, c.orderings)
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	w := tabwriter.NewWriter(c.common.Output, 0, 0, 2, ' ', 0)
	defer w.Flush()

	for _, epb := range entries {
		data, err := formatter.Format(epb)
		if err != nil {
			fmt.Fprintln(c.common.Error, err)
			return subcommands.ExitFailure
		}
		fmt.Fprint(w, string(data))
	}

	return subcommands.ExitSuccess
}
