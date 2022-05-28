package validator

import (
	"fmt"

	"github.com/minkezhang/truffle/api/go/database/utils"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type f func(e *dpb.Entry) error

var (
	validators = []f{
		score,
		corpus,
		aux_data,
		tracker,
	}
)

func Validate(epb *dpb.Entry) error {
	for _, v := range validators {
		if err := v(epb); err != nil {
			return err
		}
	}
	return nil
}

func score(e *dpb.Entry) error {
	if 0 > e.GetScore() || e.GetScore() > 10 {
		return fmt.Errorf("invalid score %v, must be in the range [0, 10]", e.GetScore())
	}
	return nil
}

func corpus(e *dpb.Entry) error {
	handled := map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_ANIME:       true,
		dpb.Corpus_CORPUS_ANIME_FILM:  true,
		dpb.Corpus_CORPUS_MANGA:       true,
		dpb.Corpus_CORPUS_BOOK:        true,
		dpb.Corpus_CORPUS_SHORT_STORY: true,
		dpb.Corpus_CORPUS_FILM:        true,
		dpb.Corpus_CORPUS_TV:          true,
		dpb.Corpus_CORPUS_ALBUM:       true,
		dpb.Corpus_CORPUS_GAME:        true,
	}

	c := e.GetCorpus()
	if !handled[c] {
		return fmt.Errorf("invalid corpus type %v", c)
	}
	return nil
}

func aux_data(e *dpb.Entry) error {
	c := e.GetCorpus()
	var t utils.AuxDataT

	switch e.AuxData.(type) {
	case *dpb.Entry_AuxDataGame:
		t = utils.AuxDataGame
	case *dpb.Entry_AuxDataBook:
		t = utils.AuxDataBook
	case *dpb.Entry_AuxDataVideo:
		t = utils.AuxDataVideo
	case *dpb.Entry_AuxDataAudio:
		t = utils.AuxDataAudio
	}

	if t != utils.AuxDataNone && t != utils.AuxDataL[c] {
		return fmt.Errorf("invalid aux_data type for corpus type %v", c)
	}
	return nil
}

func tracker(e *dpb.Entry) error {
	c := e.GetCorpus()
	var t utils.TrackerT

	switch e.Tracker.(type) {
	case *dpb.Entry_TrackerVideo:
		t = utils.TrackerVideo
	case *dpb.Entry_TrackerBook:
		t = utils.TrackerBook
	}

	if t != utils.TrackerNone && t != utils.TrackerL[c] {
		return fmt.Errorf("invalid tracker type for corpus type %v", c)
	}
	return nil
}
