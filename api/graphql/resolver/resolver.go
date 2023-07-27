package resolver

import (
	"fmt"

	"github.com/minkezhang/truffle/api/graphql/generated/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
}

func PUTEntry(q *model.MutateEntryInput, m *model.Entry) error {
	if q.ID == nil {
		m.Corpus = *q.Corpus
	}

	if m.Corpus == model.CorpusTypeCorpusNone {
		return fmt.Errorf("mandatory Corpus type is unspecified")
	}

	if m.Metadata == nil {
		m.Metadata = &model.Metadata{
			Truffle: &model.APIData{
				API:    model.APITypeAPITruffle,
				ID:     m.ID,
				Cached: true,
			},
		}
	}

	if q.Aux != nil {
		m.Metadata.Truffle.Aux = nil
		switch m.Corpus {
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
		default:
		}
	}

	if q.Providers != nil {
		m.Metadata.Truffle.Providers = q.Providers
	}

	if q.Queued != nil {
		m.Queued = *q.Queued
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
					Language: t.Language,
					Title:    t.Title,
				},
			)
		}
	}

	if q.Tags != nil {
		m.Metadata.Truffle.Tags = q.Tags
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
