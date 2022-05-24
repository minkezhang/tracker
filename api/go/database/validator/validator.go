package validator

import (
	"fmt"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type f func(e *dpb.Entry) error

var (
	validators = []f{
		corpus,
		aux_data,
		tracker,
	}
)

func Validate(db *dpb.Database) error {
	for _, e := range db.GetEntries() {
		for _, v := range validators {
			if err := v(e); err != nil {
				return err
			}
		}
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
	var supported map[dpb.Corpus]bool

	switch e.AuxData.(type) {
	case *dpb.Entry_AuxDataGame:
		supported = map[dpb.Corpus]bool{
			dpb.Corpus_CORPUS_GAME: true,
		}
	case *dpb.Entry_AuxDataBook:
		supported = map[dpb.Corpus]bool{
			dpb.Corpus_CORPUS_MANGA:       true,
			dpb.Corpus_CORPUS_BOOK:        true,
			dpb.Corpus_CORPUS_SHORT_STORY: true,
		}
	case *dpb.Entry_AuxDataVideo:
		supported = map[dpb.Corpus]bool{
			dpb.Corpus_CORPUS_ANIME:      true,
			dpb.Corpus_CORPUS_ANIME_FILM: true,
			dpb.Corpus_CORPUS_FILM:       true,
			dpb.Corpus_CORPUS_TV:         true,
		}
	case *dpb.Entry_AuxDataAudio:
		supported = map[dpb.Corpus]bool{
			dpb.Corpus_CORPUS_ALBUM: true,
		}
	default:
		supported = map[dpb.Corpus]bool{
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
	}

	if !supported[c] {
		return fmt.Errorf("invalid aux_data type for corpus type %v", c)
	}
	return nil
}

func tracker(e *dpb.Entry) error {
	c := e.GetCorpus()
	var supported map[dpb.Corpus]bool

	switch e.Tracker.(type) {
	case *dpb.Entry_TrackerVideo:
		supported = map[dpb.Corpus]bool{
			dpb.Corpus_CORPUS_ANIME: true,
			dpb.Corpus_CORPUS_TV:    true,
		}
	case *dpb.Entry_TrackerBook:
		supported = map[dpb.Corpus]bool{
			dpb.Corpus_CORPUS_MANGA: true,
			dpb.Corpus_CORPUS_BOOK:  true,
		}
	default:
		supported = map[dpb.Corpus]bool{
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
	}

	if !supported[c] {
		return fmt.Errorf("invalid tracker type for corpus type %v", c)
	}
	return nil
}
