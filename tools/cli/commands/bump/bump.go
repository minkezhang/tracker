package bump

import (
	"context"
	"flag"
	"fmt"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/database"
	"github.com/minkezhang/truffle/database/helper/get"
	"github.com/minkezhang/truffle/database/helper/patch"
	"github.com/minkezhang/truffle/formats/cli/struct"
	"github.com/minkezhang/truffle/tools/cli/flag/flagset"

	ce "github.com/minkezhang/truffle/formats/cli"
)

type C struct {
	db *database.DB

	entry *entry.E
	major bool
}

func New(db *database.DB) *C {
	return &C{
		db:    db,
		entry: &entry.E{},
	}
}

func (c *C) Name() string     { return "bump" }
func (c *C) Synopsis() string { return "bump entry bookmark" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	(*flagset.ID)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Corpus)(unsafe.Pointer(c.entry)).SetFlags(f)
	(*flagset.Title)(unsafe.Pointer(c.entry)).SetFlags(f)

	f.BoolVar(&c.major, "major", false, "bump the bookmark season / volume instead of the episode / chapter")
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

	var season int32
	var episode int32

	switch utils.TrackerL[epb.GetCorpus()] {
	case utils.TrackerVideo:
		season = epb.GetTrackerVideo().GetSeason()
		episode = epb.GetTrackerVideo().GetEpisode()
	case utils.TrackerBook:
		season = epb.GetTrackerBook().GetVolume()
		episode = epb.GetTrackerBook().GetChapter()
	default:
		fmt.Printf("Cannot bump version for a non-trackable entry.\n")
		return subcommands.ExitFailure
	}

	if c.major {
		c.entry.Season = int(season) + 1
	} else {
		c.entry.Episode = int(episode) + 1
	}

	epb, err = c.entry.PB()
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	epb, err = patch.Patch(ctx, c.db, epb)
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Print(string(e.Data))

	return subcommands.ExitSuccess
}
