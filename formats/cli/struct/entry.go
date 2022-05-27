package entry

import (
	"fmt"
	"strings"

	"github.com/minkezhang/tracker/api/go/database/utils"
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

	switch utils.AuxDataL[corpus] {
	case utils.AuxDataVideo:
		epb.AuxData = &dpb.Entry_AuxDataVideo{
			AuxDataVideo: &dpb.AuxDataVideo{
				Studios:   e.Studios,
				Directors: e.Directors,
				Writers:   e.Writers,
			},
		}
	case utils.AuxDataBook:
		epb.AuxData = &dpb.Entry_AuxDataBook{
			AuxDataBook: &dpb.AuxDataBook{
				Authors: e.Writers,
			},
		}
	case utils.AuxDataGame:
		epb.AuxData = &dpb.Entry_AuxDataGame{
			AuxDataGame: &dpb.AuxDataGame{
				Studios:   e.Studios,
				Directors: e.Directors,
				Writers:   e.Writers,
			},
		}
	case utils.AuxDataAudio:
		epb.AuxData = &dpb.Entry_AuxDataAudio{
			AuxDataAudio: &dpb.AuxDataAudio{
				Composers: e.Writers,
			},
		}
	}

	switch utils.TrackerL[corpus] {
	case utils.TrackerVideo:
		epb.Tracker = &dpb.Entry_TrackerVideo{
			TrackerVideo: &dpb.TrackerVideo{
				Season:  int32(e.Season),
				Episode: int32(e.Episode),
			},
		}
	case utils.TrackerBook:
		epb.Tracker = &dpb.Entry_TrackerBook{
			TrackerBook: &dpb.TrackerBook{
				Volume:  int32(e.Season),
				Chapter: int32(e.Episode),
			},
		}
	}

	return epb, nil
}
