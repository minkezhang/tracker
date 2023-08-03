package util

import (
	"fmt"
	"time"

	"github.com/minkezhang/truffle/api/graphql/model"
)

func PatchEntry(m *model.Entry, q *model.PatchInput) error {
	if m.Metadata == nil && *q.Corpus == model.CorpusTypeCorpusNone {
		return fmt.Errorf("mandatory Corpus type is unspecified")
	}

	if m.Metadata == nil {
		m.Metadata = &model.Metadata{
			Truffle: &model.APIData{
				API:    model.APITypeAPITruffle,
				ID:     m.ID,
				Cached: true,
				Corpus: *q.Corpus,
			},
		}
	}

	if q.Aux != nil {
		m.Metadata.Truffle.Aux = nil
		switch m.Metadata.Truffle.Corpus {
		case model.CorpusTypeCorpusAnime:
			m.Metadata.Truffle.Aux = &model.AuxAnime{
				Studios: q.Aux.Studios,
			}
		case model.CorpusTypeCorpusAnimeFilm:
			m.Metadata.Truffle.Aux = &model.AuxAnimeFilm{
				Studios: q.Aux.Studios,
			}
		case model.CorpusTypeCorpusManga:
			m.Metadata.Truffle.Aux = &model.AuxManga{
				Authors: q.Aux.Authors,
			}
		case model.CorpusTypeCorpusBook:
			m.Metadata.Truffle.Aux = &model.AuxBook{
				Authors: q.Aux.Authors,
			}
		case model.CorpusTypeCorpusShortStory:
			m.Metadata.Truffle.Aux = &model.AuxShortStory{
				Authors: q.Aux.Authors,
			}
		case model.CorpusTypeCorpusAlbum:
			m.Metadata.Truffle.Aux = &model.AuxAlbum{
				Composers: q.Aux.Composers,
			}
		case model.CorpusTypeCorpusFilm:
			m.Metadata.Truffle.Aux = &model.AuxFilm{
				Directors: q.Aux.Directors,
			}
		case model.CorpusTypeCorpusGame:
			m.Metadata.Truffle.Aux = &model.AuxGame{
				Developers: q.Aux.Developers,
			}
		}
	}

	if q.Providers != nil {
		m.Metadata.Truffle.Providers = q.Providers
	}

	if q.Queued != nil {
		m.Metadata.Truffle.Queued = *q.Queued
	}

	if q.Score != nil {
		m.Metadata.Truffle.Score = q.Score
	}

	if q.Titles != nil {
		m.Metadata.Truffle.Titles = nil
		for _, t := range q.Titles {
			m.Metadata.Truffle.Titles = append(
				m.Metadata.Truffle.Titles,
				&model.Title{
					Locale: t.Locale,
					Title:  t.Title,
				},
			)
		}
	}

	if q.Tags != nil {
		m.Metadata.Truffle.Tags = q.Tags
	}

	if q.Tracker != nil {
		t := time.Now()
		m.Metadata.Truffle.Tracker = nil
		switch m.Metadata.Truffle.Corpus {
		case model.CorpusTypeCorpusAnime:
			m.Metadata.Truffle.Tracker = &model.TrackerAnime{
				Season:      q.Tracker.Season,
				Episode:     q.Tracker.Episode,
				LastUpdated: &t,
			}
		case model.CorpusTypeCorpusTv:
			m.Metadata.Truffle.Tracker = &model.TrackerTv{
				Season:      q.Tracker.Season,
				Episode:     q.Tracker.Episode,
				LastUpdated: &t,
			}
		case model.CorpusTypeCorpusManga:
			m.Metadata.Truffle.Tracker = &model.TrackerManga{
				Volume:      q.Tracker.Volume,
				Chapter:     q.Tracker.Chapter,
				LastUpdated: &t,
			}
		case model.CorpusTypeCorpusBook:
			m.Metadata.Truffle.Tracker = &model.TrackerBook{
				Volume:      q.Tracker.Volume,
				LastUpdated: &t,
			}
		}
	}

	if q.Sources != nil {
		m.Metadata.Sources = nil
		for _, l := range q.Sources {
			if l.API == model.APITypeAPINone || l.API == model.APITypeAPITruffle {
				return fmt.Errorf("invalid API type: %s", l.API)
			}
			m.Metadata.Sources = append(
				m.Metadata.Sources,
				&model.APIData{
					API: l.API,
					ID:  l.ID,
				},
			)
		}
	}

	return nil
}
