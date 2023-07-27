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

func MergeEntry(q *model.MutateEntryInput, m *model.Entry) error {
	if q.ID == nil && m.Corpus == model.CorpusTypeCorpusNone {
		return fmt.Errorf("mandatory Corpus type is unspecified")
	}

	if m.Metadata == nil {
		m.Metadata = &model.Metadata{
			Truffle: &model.APIData{
				API: model.APITypeAPITruffle,
				ID:  m.ID,
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

	return nil
}
