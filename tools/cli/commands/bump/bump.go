package bump

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/database"

	ce "github.com/minkezhang/tracker/formats/cli"
	se "github.com/minkezhang/tracker/formats/cli/struct"
	gc "github.com/minkezhang/tracker/tools/cli/commands/get/common"
	pc "github.com/minkezhang/tracker/tools/cli/commands/patch/common"
)

type C struct {
	db *database.DB

	title  *se.Title
	id     *se.ID
	corpus *se.Corpus

	major bool
}

func New(db *database.DB) *C {
	return &C{
		db: db,

		title:  &se.Title{},
		id:     &se.ID{},
		corpus: &se.Corpus{},
	}
}

func (c *C) Name() string     { return "bump" }
func (c *C) Synopsis() string { return "bump entry bookmark" }
func (c *C) Usage() string    { return fmt.Sprintf("%v\n", c.Synopsis()) }

func (c *C) SetFlags(f *flag.FlagSet) {
	c.title.SetFlags(f)
	c.id.SetFlags(f)
	c.corpus.SetFlags(f)

	f.BoolVar(&c.major, "major", false, "bump the bookmark season / volume instead of the episode / chapter")
}

func (c *C) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	epb, err := gc.Get(gc.O{
		DB:     c.db,
		ID:     c.id.ID,
		Title:  c.title.Title,
		Corpus: c.corpus.Corpus,
	})
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

	body := &se.Body{}
	if c.major {
		body.Season = int(season) + 1
	} else {
		body.Episode = int(episode) + 1
	}

	epb, err = pc.Patch(pc.O{
		DB:    c.db,
		ID:    c.id.ID,
		Title: c.title.Title,
		Body:  body,
	})
	if err != nil {
		fmt.Printf("%v\n", err)
		return subcommands.ExitFailure
	}

	e := &ce.E{}
	e.Dump(epb)
	fmt.Printf(string(e.Data))

	return subcommands.ExitSuccess
}
