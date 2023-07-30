package database

import (
	"context"
	"encoding/json"
	"os"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client"
)

type APIData struct {
	data   map[string]*model.APIData
	client client.C
	fn     string
}

func NewAPIData(c client.C, fn string) *APIData {
	db := &APIData{
		data:   map[string]*model.APIData{},
		client: c,
		fn:     fn,
	}
	if err := db.load(db.fn); err != nil {
		panic(err)
	}
	return db
}

func (db *APIData) API() model.APIType { return db.client.API() }

func (db *APIData) Get(ctx context.Context, id string) (*model.APIData, error) {
	if d, ok := db.data[id]; !ok || ok && !d.Cached {
		d, err := db.client.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		db.data[id] = d
		return d, db.dump(db.fn)
	} else {
		return d, nil
	}
}

func (db *APIData) dump(fn string) error {
	data, err := db.marshal()
	if err != nil {
		return err
	}

	return os.WriteFile(fn, data, 0666)
}

func (db *APIData) load(fn string) error {
	data, err := os.ReadFile(fn)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	return db.unmarshal(data)
}

func (db *APIData) marshal() ([]byte, error)    { return json.MarshalIndent(db.data, "", "  ") }
func (db *APIData) unmarshal(data []byte) error { return json.Unmarshal(data, db) }
