package add

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/database"
	"google.golang.org/protobuf/encoding/prototext"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type C struct {
	db *database.DB

	entry []byte
}

func New(db *database.DB) *C { return &C{db: db} }

func (c *C) Name() string             { return "add" }
func (c *C) Synopsis() string         { return "add entry to database" }
func (c *C) Usage() string            { return c.Synopsis() }
func (c *C) SetFlags(f *flag.FlagSet) {}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if len(f.Args()) != 1 {
		fmt.Println(c.Usage())
		return subcommands.ExitUsageError
	}
	c.entry = []byte(f.Args()[0])

	epb := &dpb.Entry{}
	if err := prototext.Unmarshal(c.entry, epb); err != nil {
		fmt.Printf("Could not marshal input data: %v\n", err)
		return subcommands.ExitFailure
	}

	if err := c.db.AddEntry(epb); err != nil {
		fmt.Printf("Could not add data to database: %v\n", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
