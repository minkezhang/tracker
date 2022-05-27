package cache

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/formats/minkezhang/columns"
	"github.com/minkezhang/tracker/formats/minkezhang/lookup"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type E [11]string

func (e E) Corpus() dpb.Corpus { return lookup.Corpus[e[columns.Category]] }
func (e E) Titles() []string   { return []string{e[columns.Title]} }
func (e E) Queued() bool       { return e[columns.Queued] == "TRUE" }
func (e E) Studios() []string  { return []string{e[columns.Studio]} }

func (e E) Directors() []string {
	var directors []string
	for _, v := range strings.Split(e[columns.Directors], ",") {
		directors = append(directors, strings.Trim(v, " "))
	}
	return directors
}

func (e E) Writers() []string {
	var writers []string
	for _, v := range strings.Split(e[columns.Writers], ",") {
		writers = append(writers, strings.Trim(v, " "))
	}
	return writers
}

func (e E) Score() float32 {
	if f, err := strconv.ParseFloat(e[columns.Rating], 32); err != nil {
		return 0
	} else {
		return float32(f)
	}
}

func (e E) Providers() []dpb.Provider {
	var dedupe = map[dpb.Provider]bool{}
	dedupe[lookup.Provider[e[columns.Distributor1]]] = true
	dedupe[lookup.Provider[e[columns.Distributor2]]] = true

	var providers []dpb.Provider
	for k := range dedupe {
		if k != dpb.Provider_PROVIDER_UNKNOWN {
			providers = append(providers, k)
		}
	}
	return providers
}

func (e E) TrackerBook() *dpb.TrackerBook {
	reg := regexp.MustCompile(`v?(?P<volume>\d+)?c(?P<chapter>\d+)`)
	m := reg.FindStringSubmatch(e[columns.Bookmark])
	if m == nil {
		return &dpb.TrackerBook{}
	}
	result := map[string]int32{}
	for i, name := range reg.SubexpNames() {
		if i != 0 && name != "" {
			v, err := strconv.ParseInt(m[i], 10, 32)
			if err != nil {
				v = 0
			}
			result[name] = int32(v)
		}
	}
	return &dpb.TrackerBook{
		Volume:  result["volume"],
		Chapter: result["chapter"],
	}
}

func (e E) TrackerVideo() *dpb.TrackerVideo {
	reg := regexp.MustCompile(`s?(?P<season>\d+)?e(?P<episode>\d+)`)
	m := reg.FindStringSubmatch(e[columns.Bookmark])
	if m == nil {
		return &dpb.TrackerVideo{}
	}
	result := map[string]int32{}
	for i, name := range reg.SubexpNames() {
		if i != 0 && name != "" {
			v, err := strconv.ParseInt(m[i], 10, 32)
			if err != nil {
				v = 0
			}
			result[name] = int32(v)
		}
	}
	return &dpb.TrackerVideo{
		Season:  result["season"],
		Episode: result["episode"],
	}
}

// ProtoBuf returns a PB object for the given row.
//
// TODO(minkezhang): Refactor to proto.Unmarshal API.
func (e E) ProtoBuf() *dpb.Entry {
	epb := &dpb.Entry{
		Corpus:    e.Corpus(),
		Titles:    e.Titles(),
		Queued:    e.Queued(),
		Score:     e.Score(),
		Providers: e.Providers(),
	}

	switch utils.TrackerL[epb.GetCorpus()] {
	case utils.TrackerVideo:
		epb.Tracker = &dpb.Entry_TrackerVideo{
			TrackerVideo: e.TrackerVideo(),
		}
	case utils.TrackerBook:
		epb.Tracker = &dpb.Entry_TrackerBook{
			TrackerBook: e.TrackerBook(),
		}
	}

	switch utils.AuxDataL[epb.GetCorpus()] {
	case utils.AuxDataVideo:
		epb.AuxData = &dpb.Entry_AuxDataVideo{
			AuxDataVideo: &dpb.AuxDataVideo{
				Studios:   e.Studios(),
				Directors: e.Directors(),
				Writers:   e.Writers(),
			},
		}
	case utils.AuxDataGame:
		epb.AuxData = &dpb.Entry_AuxDataGame{
			AuxDataGame: &dpb.AuxDataGame{
				Studios:   e.Studios(),
				Directors: e.Directors(),
				Writers:   e.Writers(),
			},
		}
	case utils.AuxDataBook:
		epb.AuxData = &dpb.Entry_AuxDataBook{
			AuxDataBook: &dpb.AuxDataBook{
				Authors: e.Writers(),
			},
		}
	case utils.AuxDataAudio:
		epb.AuxData = &dpb.Entry_AuxDataAudio{
			AuxDataAudio: &dpb.AuxDataAudio{
				Composers: e.Writers(),
			},
		}
	}

	return epb
}
