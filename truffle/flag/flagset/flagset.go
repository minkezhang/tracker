// Package flagset contains shared flagsets for the CLI.
//
// TODO(minkezhang): Move to tools/cli/command/flag instead.
package flagset

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/truffle/flag/entry"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type Corpus entry.E

func (set *Corpus) SetFlags(f *flag.FlagSet) {
	g := func(corpus string) error {
		if corpus == "" {
			corpus = "unknown"
		}
		set.Corpus = dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", corpus)])
		return nil
	}
	f.Func("corpus", "entry corpus, e.g. \"film\"", g)
}

// Title is a flag setter which allows for a single title input.
type Title entry.E

func (set *Title) SetFlags(f *flag.FlagSet) {
	g := func(title string) error {
		set.Titles = []string{title}
		return nil
	}
	f.Func("title", "entry title, e.g. \"12 Angry Men\"", g)
}

type Titles entry.E

func (set *Titles) SetFlags(f *flag.FlagSet) {
	f.Var(&set.Titles, "title", "entry title, e.g. \"12 Angry Men\"")
}

type ID entry.E

func (set *ID) SetFlags(f *flag.FlagSet) {
	g := func(id string) error {
		prefix, id, ok := strings.Cut(id, ":")
		if !ok {
			prefix, id = id, prefix
		}

		api := dpb.API(dpb.API_value[utils.ToEnum("API", prefix)])
		if api == dpb.API_API_UNKNOWN {
			if prefix != "" {
				return fmt.Errorf("invalid ID API: %v", prefix)
			}
			api = dpb.API_API_TRUFFLE
		}

		set.ID = &dpb.LinkedID{
			Id:  id,
			Api: api,
		}
		return nil
	}
	f.Func("id", "entry ID, e.g. \"JxCF\"", g)
}

type Body entry.E

func (set *Body) SetFlags(f *flag.FlagSet) {
	f.Func("provider", "distributors of the entry, e.g. \"google_play\"", func(provider string) error {
		set.Providers = append(set.Providers, dpb.Provider(
			dpb.Provider_value[utils.ToEnum("PROVIDER", provider)]))
		return nil
	})

	f.Float64Var(&set.Score, "score", 0, "user score")
	f.Func("queued", "indicates if the entry is on the user watchlist", func(queued string) error {
		set.SetQueued = true
		if q, err := strconv.ParseBool(queued); err != nil {
			return err
		} else {
			set.Queued = q
		}
		return nil
	})

	f.Var(&set.Directors, "director", "directors of game or visual-based entries")
	f.Var(&set.Studios, "studio", "studios of game or visual-based entries")
	f.Var(&set.Writers, "writer", "writers of game or visual-based entries")
	f.Var(&set.Writers, "composer", "composers for album-only entries")
	f.Var(&set.Writers, "author", "authors for literature-based entries")

	f.IntVar(&set.Season, "season", 0, "current anime or tv show season")
	f.IntVar(&set.Season, "volume", 0, "current manga or book volume")
	f.IntVar(&set.Episode, "episode", 0, "current anime or tv show episode")
	f.IntVar(&set.Episode, "chapter", 0, "current manga or book chapter")

	f.Func("link", "linked API IDs, e.g. \"mal:123\"", func(link string) error {
		set.LinkedIDs = append(set.LinkedIDs, entry.ID(link))
		return nil
	})

	f.Func("etag", "current etag of the entry; ignored if empty", func(etag string) error {
		set.ETag = []byte(etag)
		return nil
	})
}
