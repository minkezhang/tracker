package bump

import (
	"context"
	"flag"
	"fmt"
	"math"
	"unsafe"

	"github.com/google/subcommands"
	"github.com/minkezhang/truffle/api/go/database/utils"
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

	major bool
}

func New(db *database.DB, common common.O) *C {
	return &C{
		common: common,
		db:     db,
		entry:  &entry.E{},
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
		fmt.Fprintln(c.common.Error, err)
		return subcommands.ExitFailure
	}

	epb, err = get.Get(ctx, c.db, epb, []dpb.API{
		dpb.API_API_TRUFFLE,
	})
	if err != nil {
		fmt.Fprintln(c.common.Error, err)
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
		fmt.Fprintln(c.common.Error, fmt.Errorf("Cannot bump version for a non-trackable entry."))
		return subcommands.ExitFailure
	}

	// Most likely use-case for bumping an entry tracker is after reading
	// the first chapter or watching the first episode.
	if c.major {
		season = int32(math.Max(2, float64(season)+1))
		episode = 1
	} else {
		episode = int32(math.Max(2, float64(episode)+1))
	}

	switch utils.TrackerL[epb.GetCorpus()] {
	case utils.TrackerVideo:
		epb.Tracker = &dpb.Entry_TrackerVideo{
			TrackerVideo: &dpb.TrackerVideo{
				Season:  season,
				Episode: episode,
			},
		}
	case utils.TrackerBook:
		epb.Tracker = &dpb.Entry_TrackerBook{
			TrackerBook: &dpb.TrackerBook{
				Volume:  season,
				Chapter: episode,
			},
		}
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
