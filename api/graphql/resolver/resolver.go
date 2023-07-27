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
	if q.ID == nil && m.Corpus == model.CorpusTypeCorpusNone {
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

	if q.Links != nil {
		m.Metadata.Mal = nil
		m.Metadata.Spotify = nil
		m.Metadata.Kitsu = nil
		m.Metadata.Steam = nil
		for _, l := range q.Links {
			d := &model.APIData{
				API: l.API,
				ID:  l.ID,
			}
			switch l.API {
			case model.APITypeAPIMal:
				m.Metadata.Mal = d
			case model.APITypeAPISpotify:
				m.Metadata.Spotify = d
			case model.APITypeAPIKitsu:
				m.Metadata.Kitsu = d
			case model.APITypeAPISteam:
				m.Metadata.Steam = d
			default:
				return fmt.Errorf("invalid API type: %s", l.API)
			}
		}
	}

	return nil
}
