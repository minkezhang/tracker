package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type Entry struct {
	data map[string]*model.Entry
	fn   string
}

func NewEntry(fn string) *Entry {
	db := &Entry{
		data: map[string]*model.Entry{},
		fn:   fn,
	}
	if err := db.load(db.fn); err != nil {
		panic(err)
	}
	return db
}

func (db *Entry) Get(ctx context.Context, id string) (*model.Entry, error) {
	if e, ok := db.data[id]; !ok {
		return nil, fmt.Errorf("cannot find entry: %s", id)
	} else {
		return e, nil
	}
}

func (db *Entry) Put(ctx context.Context, e *model.Entry) (*model.Entry, error) {
	db.data[e.ID] = e
	return e, db.dump(db.fn)
}

func (db *Entry) Delete(ctx context.Context, id string) (*model.Entry, error) {
	e := db.data[id]
	delete(db.data, id)
	return e, db.dump(db.fn)
}

func (db *Entry) List(ctx context.Context, query *model.ListInput) ([]*model.Entry, error) {
	corpora := map[model.CorpusType]bool{}
	for _, c := range query.Corpora {
		corpora[c] = true
	}

	var matches []*model.Entry
	for _, e := range db.data {
		if _, ok := corpora[e.Metadata.Truffle.Corpus]; !ok {
			continue
		}
		for _, t := range e.Metadata.Truffle.Titles {
			if query.Title != nil && strings.Contains(t.Title, *query.Title) {
				matches = append(matches, e)
			}
		}
	}
	return matches, nil
}

func (db *Entry) dump(fn string) error {
	data, err := db.marshal()
	if err != nil {
		return err
	}

	return os.WriteFile(fn, data, 0666)
}

func (db *Entry) load(fn string) error {
	data, err := os.ReadFile(fn)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}

	return db.unmarshal(data)
}

func (db *Entry) marshal() ([]byte, error)    { return json.MarshalIndent(db.data, "", "  ") }
func (db *Entry) unmarshal(data []byte) error { return json.Unmarshal(data, &db.data) }
