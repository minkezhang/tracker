package resolver

import (
	"github.com/minkezhang/truffle/api/graphql/generated/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
}

func NewEntry(q *model.MutateEntryInput) (*model.Entry, error) {
	m := &model.Entry{
		Corpus: q.Corpus,
		Metadata: &model.Metadata{
			Truffle: &model.APIData{
				API:       model.APITypeAPITruffle,
				Providers: q.Providers,
			},
		},
	}

	if q.ID != nil {
		m.ID = *q.ID
		m.Metadata.Truffle.ID = *q.ID
	}

	if q.Queued != nil {
		m.Queued = *q.Queued
	}

	if q.Score != nil {
		m.Metadata.Truffle.Score = q.Score
	}

	for _, t := range q.Titles {
		m.Metadata.Truffle.Titles = append(
			m.Metadata.Truffle.Titles,
			&model.Title{
				Language: t.Language,
				Title:    t.Title,
			},
		)
	}

	return m, nil
}
