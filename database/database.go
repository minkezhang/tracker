package database

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/client/mal"
	"github.com/minkezhang/truffle/client/truffle"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type DB struct {
	truffle *truffle.C
}

func New(epbs []*dpb.Entry) *DB {
	db := &DB{
		truffle: truffle.New(
			&dpb.Database{
				Entries: map[string]*dpb.Entry{},
			},
		),
	}
	for _, epb := range epbs {
		if _, err := db.truffle.Add(context.Background(), epb); err != nil {
			panic(fmt.Sprintf("could not create database: %v", err))
		}
	}

	return db
}

func (db *DB) Add(epb *dpb.Entry) (*dpb.Entry, error) {
	return db.truffle.Add(context.Background(), epb)
}

func (db *DB) Get(id string) (*dpb.Entry, error) {
	return db.truffle.Get(
		context.Background(),
		&dpb.LinkedID{
			Id:  id,
			Api: dpb.API_API_TRUFFLE,
		})
}

func (db *DB) Put(epb *dpb.Entry) (*dpb.Entry, error) {
	return db.truffle.Put(context.Background(), epb)
}

func (db *DB) Delete(id string) (*dpb.Entry, error) {
	return db.truffle.Delete(context.Background(), &dpb.LinkedID{
		Id:  id,
		Api: dpb.API_API_TRUFFLE,
	})
}

type O struct {
	Context context.Context

	Title  string
	Corpus dpb.Corpus

	APIs []dpb.API

	// MAL contains MAL-specific API options.
	MAL mal.SearchOpts
}

func (db *DB) Search(opts O) ([]*dpb.Entry, error) {
	apis := map[dpb.API]bool{}
	for _, api := range opts.APIs {
		apis[api] = true
	}

	var candidates []*dpb.Entry

	if apis[dpb.API_API_TRUFFLE] {
		if cs, err := db.truffle.Search(
			opts.Context,
			truffle.SearchOpts{
				Title:  opts.Title,
				Corpus: opts.Corpus,
			},
		); err != nil {
			return nil, err
		} else {
			candidates = append(candidates, cs...)
		}
	}

	duplicates := map[string]bool{}
	for _, epb := range candidates {
		for _, id := range epb.GetLinkedIds() {
			duplicates[utils.ID(id)] = true
		}
	}

	if apis[dpb.API_API_MAL] {
		if cs, err := mal.New().Search(
			opts.Context,
			mal.SearchOpts{
				Title:  opts.Title,
				Corpus: opts.Corpus,
				Cutoff: opts.MAL.Cutoff,
			},
		); err != nil {
			return nil, err
		} else {
			candidates = append(candidates, cs...)
		}
	}

	var unique []*dpb.Entry
	for _, epb := range candidates {
		// Skip reporting entries which are already accounted for in the
		// user DB (if we are returning user DB results).
		if apis[dpb.API_API_TRUFFLE] && duplicates[utils.ID(epb.GetId())] {
			continue
		}
		unique = append(unique, epb)
	}
	return unique, nil
}

func Marshal(db *DB) ([]byte, error) { return truffle.Marshal(db.truffle) }
func Unmarshal(data []byte) (*DB, error) {
	truffle, err := truffle.Unmarshal(data)
	if err != nil {
		return nil, err
	}
	return &DB{truffle: truffle}, nil
}
