package entry

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"

	dpb "github.com/minkezhang/tracker/api/go/database"
	cf "github.com/minkezhang/tracker/tools/cli/flag"
)

type E struct {
	Corpus    string
	Titles    cf.MultiString
	Providers cf.MultiString

	Directors cf.MultiString
	Studios   cf.MultiString
	Writers   cf.MultiString

	Score  float64
	Queued bool

	Season  int
	Episode int
}

func (e E) Load() (proto.Message, error) {
	corpus := dpb.Corpus(dpb.Corpus_value[fmt.Sprintf("CORPUS_%v", strings.ToUpper(e.Corpus))])

	var providers []dpb.Provider
	for _, p := range e.Providers {
		providers = append(providers, dpb.Provider(dpb.Provider_value[fmt.Sprintf("PROVIDER_%v", strings.ToUpper(p))]))
	}

	epb := &dpb.Entry{
		Titles:    e.Titles,
		Providers: providers,
		Corpus:    corpus,
		Queued:    e.Queued,
		Score:     float32(e.Score),
	}

	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_TV:         true,
		dpb.Corpus_CORPUS_ANIME:      true,
		dpb.Corpus_CORPUS_FILM:       true,
		dpb.Corpus_CORPUS_ANIME_FILM: true,
	}[corpus] {
		epb.AuxData = &dpb.Entry_AuxDataVideo{
			AuxDataVideo: &dpb.AuxDataVideo{
				Studios:   e.Studios,
				Directors: e.Directors,
				Writers:   e.Writers,
			},
		}
	}

	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_MANGA:       true,
		dpb.Corpus_CORPUS_BOOK:        true,
		dpb.Corpus_CORPUS_SHORT_STORY: true,
	}[corpus] {
	}

	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_GAME: true,
	}[corpus] {
	}

	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_ALBUM: true,
	}[corpus] {
	}
	return epb, nil
}
