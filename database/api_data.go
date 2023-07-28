package database

import (
	"context"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client"
)

type APIData struct {
	data   map[string]*model.APIData
	client client.C
}

func NewAPIData(c client.C) *APIData {
	return &APIData{
		data:   map[string]*model.APIData{},
		client: c,
	}
}

func (db *APIData) API() model.APIType { return db.client.API() }

func (db *APIData) Get(ctx context.Context, id string) (*model.APIData, error) {
	if d, ok := db.data[id]; !ok || ok && !d.Cached {
		d, err := db.client.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		db.data[id] = d
		return d, nil
	} else {
		return d, nil
	}
}
