package utils

import (
	"fmt"
	"strings"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type TrackerT int
type AuxDataT int

const (
	AuxDataNone AuxDataT = iota
	AuxDataVideo
	AuxDataBook
	AuxDataGame
	AuxDataAudio

	TrackerNone TrackerT = iota
	TrackerVideo
	TrackerBook
)

var (
	APIPrefix = map[dpb.API]string{
		dpb.API_API_MAL: "mal",
	}
)

func ID(id *dpb.LinkedID) string {
	var parts []string
	if api := APIPrefix[id.GetApi()]; api != "" {
		parts = append(parts, api)
	}
	parts = append(parts, id.GetId())
	return strings.Join(parts, ":")
}

var (
	AuxDataL = map[dpb.Corpus]AuxDataT{
		dpb.Corpus_CORPUS_TV:         AuxDataVideo,
		dpb.Corpus_CORPUS_ANIME:      AuxDataVideo,
		dpb.Corpus_CORPUS_FILM:       AuxDataVideo,
		dpb.Corpus_CORPUS_ANIME_FILM: AuxDataVideo,

		dpb.Corpus_CORPUS_GAME: AuxDataGame,

		dpb.Corpus_CORPUS_MANGA:       AuxDataBook,
		dpb.Corpus_CORPUS_BOOK:        AuxDataBook,
		dpb.Corpus_CORPUS_SHORT_STORY: AuxDataBook,

		dpb.Corpus_CORPUS_ALBUM: AuxDataAudio,
	}

	TrackerL = map[dpb.Corpus]TrackerT{
		dpb.Corpus_CORPUS_ANIME: TrackerVideo,
		dpb.Corpus_CORPUS_TV:    TrackerVideo,

		dpb.Corpus_CORPUS_MANGA: TrackerBook,
		dpb.Corpus_CORPUS_BOOK:  TrackerBook,
	}
)

func ToEnum(prefix string, suffix string) string {
	return fmt.Sprintf("%v_%v", prefix, strings.ReplaceAll(strings.ToUpper(suffix), " ", "_"))
}
