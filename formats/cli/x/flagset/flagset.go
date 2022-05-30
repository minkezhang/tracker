// Package flagset contains shared flagsets for the CLI.
//
// TODO(minkezhang): Move to tools/cli/command/flag instead.
package flagset

import (
	"flag"

	entry "github.com/minkezhang/truffle/formats/cli/x"
)

type Corpus entry.E

func (set *Corpus) SetFlags(f *flag.FlagSet) {
	f.StringVar(&set.Corpus, "corpus", "unknown", "entry corpus, e.g. \"film\"")
}

type Title entry.E

func (set *Title) SetFlags(f *flag.FlagSet) {
	f.Func("title", "entry title, e.g. \"12 Angry Men\"", func(s string) error {
		set.Titles = append(set.Titles, s)
		return nil
	})
}

type Titles entry.E

func (set *Titles) SetFlags(f *flag.FlagSet) {
	f.Var(&set.Titles, "titles", "entry titles, e.g. \"12 Angry Men\"")
}

type ID entry.E

func (set *ID) SetFlags(f *flag.FlagSet) {
	f.StringVar(&set.ID, "id", "", "entry ID")
}

type Body entry.E

func (set *Body) SetFlags(f *flag.FlagSet) {
	f.Var(&set.Providers, "providers", "distributors of the entry, e.g. \"google_play\"")

	f.Float64Var(&set.Score, "score", 0, "user score")
	f.BoolVar(&set.Queued, "queued", false, "indicates if the entry is on the user watchlist")

	f.Var(&set.Directors, "directors", "directors of game or visual-based entries")
	f.Var(&set.Studios, "studios", "studios of game or visual-based entries")
	f.Var(&set.Writers, "writers", "writers of game or visual-based entries")
	f.Var(&set.Writers, "composers", "composers for album-only entries")
	f.Var(&set.Writers, "authors", "authors for literature-based entries")

	f.IntVar(&set.Season, "season", 0, "current anime or tv show season")
	f.IntVar(&set.Season, "volume", 0, "current manga or book volume")
	f.IntVar(&set.Episode, "episode", 0, "current anime or tv show episode")
	f.IntVar(&set.Episode, "chapter", 0, "current manga or book chapter")

	f.Var(&set.LinkedIDs, "links", "linked API IDs, e.g. \"mal:123\"")

	f.StringVar(&set.ETag, "etag", "", "current etag of the entry; ignored if empty")
}
