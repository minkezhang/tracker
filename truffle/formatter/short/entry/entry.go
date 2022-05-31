package entry

import (
	"fmt"
	"strings"

	"github.com/minkezhang/truffle/api/go/database/utils"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type row [10]string
type c int

const (
	columnTitle c = iota
	columnCorpus
	columnID
	columnQueued
	columnScore
	columnBookmark
	columnDirector
	columnWriter
	columnStudio
	columnProvider
)

func bookmark(epb *dpb.Entry) string {
	type d struct {
		SeasonLabel  string
		Season       int32
		EpisodeLabel string
		Episode      int32
	}
	var data d

	switch utils.TrackerL[epb.GetCorpus()] {
	case utils.TrackerVideo:
		tracker := epb.GetTrackerVideo()
		data = d{
			SeasonLabel:  "s",
			Season:       tracker.GetSeason(),
			EpisodeLabel: "e",
			Episode:      tracker.GetEpisode(),
		}
	case utils.TrackerBook:
		tracker := epb.GetTrackerBook()
		data = d{
			SeasonLabel:  "v",
			Season:       tracker.GetVolume(),
			EpisodeLabel: "c",
			Episode:      tracker.GetChapter(),
		}
	}
	var s string
	if data.Season > 0 {
		s = fmt.Sprintf("%v%v", data.SeasonLabel, data.Season)
	}
	if data.Episode > 0 {
		s = fmt.Sprintf("%v%v%v", s, data.EpisodeLabel, data.Episode)
	}
	return s
}

func aux_data(epb *dpb.Entry) []string {
	type d struct {
		Director string
		Writer   string
		Studio   string
	}
	data := &d{}

	switch utils.AuxDataL[epb.GetCorpus()] {
	case utils.AuxDataBook:
		aux := epb.GetAuxDataBook()
		if len(aux.GetAuthors()) > 0 {
			data.Writer = aux.GetAuthors()[0]
		}
	case utils.AuxDataVideo:
		aux := epb.GetAuxDataVideo()
		if len(aux.GetDirectors()) > 0 {
			data.Director = aux.GetDirectors()[0]
		}
		if len(aux.GetWriters()) > 0 {
			data.Writer = aux.GetWriters()[0]
		}
		if len(aux.GetStudios()) > 0 {
			data.Studio = aux.GetStudios()[0]
		}
	case utils.AuxDataAudio:
		aux := epb.GetAuxDataAudio()
		if len(aux.GetComposers()) > 0 {
			data.Writer = aux.GetComposers()[0]
		}
	case utils.AuxDataGame:
		aux := epb.GetAuxDataGame()
		if len(aux.GetDirectors()) > 0 {
			data.Director = aux.GetDirectors()[0]
		}
		if len(aux.GetWriters()) > 0 {
			data.Writer = aux.GetWriters()[0]
		}
		if len(aux.GetStudios()) > 0 {
			data.Studio = aux.GetStudios()[0]
		}
	}

	return []string{data.Director, data.Writer, data.Studio}
}

func format(epb *dpb.Entry) ([]string, error) {
	r := [10]string{}

	if len(epb.GetTitles()) > 0 {
		r[columnTitle] = epb.GetTitles()[0]
	}

	r[columnCorpus] = epb.GetCorpus().String()
	r[columnID] = utils.ID(epb.GetId())
	r[columnQueued] = map[bool]string{
		true:  "Q",
		false: "",
	}[epb.GetQueued()]
	if s := epb.GetScore(); s > 0 {
		r[columnScore] = fmt.Sprintf("%.1f", s)
	}

	r[columnBookmark] = bookmark(epb)

	aux := aux_data(epb)
	r[columnDirector] = aux[0]
	r[columnWriter] = aux[1]
	r[columnStudio] = aux[2]

	if len(epb.GetProviders()) > 0 {
		r[columnProvider] = epb.GetProviders()[0].String()
	}

	return r[:], nil
}

func Format(epb *dpb.Entry) ([]byte, error) {
	data, err := format(epb)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("%v\n", strings.Join(data, "\t"))), nil
}
