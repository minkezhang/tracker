package database

import (
	"fmt"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type DB struct {
	data map[string]*model.Entry
}

func New() *DB {
	return &DB{
		data: map[string]*model.Entry{},
	}
}

func (db *DB) Get(id string) (*model.Entry, error) {
	if e, ok := db.data[id]; !ok {
		return nil, fmt.Errorf("cannot find entry: %s", id)
	} else {
		return e, nil
	}
}

func (db *DB) Put(e *model.Entry) (*model.Entry, error) {
	db.data[e.ID] = e
	return e, nil
}

func (db *DB) Delete(id string) (*model.Entry, error) {
	e := db.data[id]
	delete(db.data, id)
	return e, nil
}

func (db *DB) List() (*model.Entry, error) {
	return nil, fmt.Errorf("unimplemented")
}
