package entry

import (
	"flag"
	"fmt"
	"strings"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/truffle/api/go/database"
	cf "github.com/minkezhang/truffle/tools/cli/flag"
)

type Body struct {
	corpus *Corpus

	Providers cf.MultiString

	Directors cf.MultiString
	Studios   cf.MultiString
	Writers   cf.MultiString

	Score  float64
	Queued bool

	Season  int
	Episode int

	ETag string
}

func (b *Body) GetCorpus() string {
	if b.corpus == nil {
		return "unknown"
	}
	return b.corpus.Corpus
}

func (b *Body) SetCorpus(c dpb.Corpus) {
	_, corpus, _ := strings.Cut(c.String(), "_")
	b.corpus = &Corpus{
		Corpus: corpus,
	}
}

func (b *Body) SetFlags(f *flag.FlagSet) {
	b.corpus = &Corpus{}
	b.corpus.SetFlags(f)

	f.Var(&b.Providers, "providers", "distributors of the entry, e.g. \"google_play\"")

	f.Float64Var(&b.Score, "score", 0, "user score")
	f.BoolVar(&b.Queued, "queued", false, "indicates if the entry is on the user watchlist")

	f.Var(&b.Directors, "directors", "directors of game or visual-based entries")
	f.Var(&b.Studios, "studios", "studios of game or visual-based entries")
	f.Var(&b.Writers, "writers", "writers of game or visual-based entries")
	f.Var(&b.Writers, "composers", "composers for album-only entries")
	f.Var(&b.Writers, "authors", "authors for literature-based entries")

	f.IntVar(&b.Season, "season", 0, "current anime or tv show season")
	f.IntVar(&b.Season, "volume", 0, "current manga or book volume")
	f.IntVar(&b.Episode, "episode", 0, "current anime or tv show episode")
	f.IntVar(&b.Episode, "chapter", 0, "current manga or book chapter")

	f.StringVar(&b.ETag, "etag", "", "current etag of the entry; ignored if empty")
}

func (b *Body) Load() (proto.Message, error) {
	if b.corpus == nil {
		return nil, fmt.Errorf("cannot load body without a corpus defined")
	}

	epb := &dpb.Entry{}

	s, _ := b.corpus.Load()
	proto.Merge(epb, s.(*dpb.Entry))

	var providers []dpb.Provider
	for _, p := range b.Providers {
		providers = append(providers, dpb.Provider(dpb.Provider_value[utils.ToEnum("PROVIDER", p)]))
	}

	epb.Providers = providers
	epb.Queued = b.Queued
	epb.Score = float32(b.Score)
	epb.Etag = []byte(b.ETag)

	switch utils.AuxDataL[epb.GetCorpus()] {
	case utils.AuxDataVideo:
		epb.AuxData = &dpb.Entry_AuxDataVideo{
			AuxDataVideo: &dpb.AuxDataVideo{
				Studios:   b.Studios,
				Directors: b.Directors,
				Writers:   b.Writers,
			},
		}
	case utils.AuxDataBook:
		epb.AuxData = &dpb.Entry_AuxDataBook{
			AuxDataBook: &dpb.AuxDataBook{
				Authors: b.Writers,
			},
		}
	case utils.AuxDataGame:
		epb.AuxData = &dpb.Entry_AuxDataGame{
			AuxDataGame: &dpb.AuxDataGame{
				Studios:   b.Studios,
				Directors: b.Directors,
				Writers:   b.Writers,
			},
		}
	case utils.AuxDataAudio:
		epb.AuxData = &dpb.Entry_AuxDataAudio{
			AuxDataAudio: &dpb.AuxDataAudio{
				Composers: b.Writers,
			},
		}
	}

	switch utils.TrackerL[epb.GetCorpus()] {
	case utils.TrackerVideo:
		epb.Tracker = &dpb.Entry_TrackerVideo{
			TrackerVideo: &dpb.TrackerVideo{
				Season:  int32(b.Season),
				Episode: int32(b.Episode),
			},
		}
	case utils.TrackerBook:
		epb.Tracker = &dpb.Entry_TrackerBook{
			TrackerBook: &dpb.TrackerBook{
				Volume:  int32(b.Season),
				Chapter: int32(b.Episode),
			},
		}
	}

	return epb, nil
}

type Title struct {
	Title string
}

func (t *Title) SetFlags(f *flag.FlagSet) {
	f.StringVar(&t.Title, "title", "", "entry title, e.g. \"12 Angry Men\"")
}

func (t *Title) Load() (proto.Message, error) {
	return &dpb.Entry{
		Titles: []string{t.Title},
	}, nil
}

type Titles struct {
	Titles cf.MultiString
}

func (t *Titles) SetFlags(f *flag.FlagSet) {
	f.Var(&t.Titles, "titles", "entry titles, e.g. \"12 Angry Men\"")
}

func (t *Titles) Load() (proto.Message, error) {
	return &dpb.Entry{
		Titles: t.Titles,
	}, nil
}

type Corpus struct {
	Corpus string
}

func (c *Corpus) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.Corpus, "corpus", "unknown", "entry corpus, e.g. \"film\"")
}

func (c *Corpus) Load() (proto.Message, error) {
	return &dpb.Entry{
		Corpus: dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", c.Corpus)]),
	}, nil
}

type ID struct {
	ID string
}

func (id *ID) SetFlags(f *flag.FlagSet) {
	f.StringVar(&id.ID, "id", "", "entry ID")
}

func (id *ID) Load() (proto.Message, error) {
	return &dpb.Entry{
		Id: id.ID,
	}, nil
}
