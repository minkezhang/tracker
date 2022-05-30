package entry

import (
	"strings"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/tools/cli/flag"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

// E centralizes flagset storage.
type E struct {
	// ID is a string of the truffle ID.
	ID string

	// LinkedIDs is a list of linked API IDs; each ID is of the form
	// {API}:{API_ID}, e.g. "mal:123"
	LinkedIDs flag.MultiString

	Titles flag.MultiString

	// Corpus is a sanitized dpb.Corpus enum string, e.g. "manga".
	Corpus string

	// Providers is a list of sanitized dpb.Provider enum strings, e.g.
	// "crunchyroll".
	Providers flag.MultiString

	Score  float64
	Queued bool

	Directors flag.MultiString
	Studios   flag.MultiString
	Writers   flag.MultiString

	Season  int
	Episode int

	ETag string
}

func ID(id string) *dpb.LinkedID {
	api, lid, ok := strings.Cut(id, ":")
	// Assume a solid string is the ID, i.e. "123" and not the API.
	if !ok {
		api, lid = lid, api
	}
	return &dpb.LinkedID{
		Id: lid,
		Api: dpb.API(
			dpb.API_value[utils.ToEnum("API", api)]),
	}
}

func (e E) PB() (*dpb.Entry, error) {
	epb := &dpb.Entry{
		Id: ID(e.ID),
		Corpus: dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", e.Corpus)]),
		Score:  float32(e.Score),
		Queued: e.Queued,
		Etag:   []byte(e.ETag),
	}
	for _, id := range e.LinkedIDs {
		epb.LinkedIds = append(epb.GetLinkedIds(), ID(id))
	}
	for _, p := range e.Providers {
		epb.Providers = append(epb.GetProviders(),
			dpb.Provider(
				dpb.Provider_value[utils.ToEnum("PROVIDER", p)]))
	}

	switch utils.AuxDataL[epb.GetCorpus()] {
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

	switch utils.TrackerL[epb.GetCorpus()] {
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
