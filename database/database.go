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

func (db *DB) Add(ctx context.Context, epb *dpb.Entry) (*dpb.Entry, error) {
	return db.truffle.Add(ctx, epb)
}

func (db *DB) Get(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error) {
	return db.truffle.Get(ctx, id)
}

func (db *DB) Put(ctx context.Context, epb *dpb.Entry) (*dpb.Entry, error) {
	return db.truffle.Put(ctx, epb)
}

func (db *DB) Delete(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error) {
	return db.truffle.Delete(ctx, id)
}

type SearchOpts struct {
	Title  string
	Corpus dpb.Corpus

	APIs []dpb.API

	// MAL contains MAL-specific API options.
	MAL mal.SearchOpts
}

func (db *DB) Search(ctx context.Context, query SearchOpts) ([]*dpb.Entry, error) {
	apis := map[dpb.API]bool{}
	for _, api := range query.APIs {
		apis[api] = true
	}

	var candidates []*dpb.Entry

	if apis[dpb.API_API_TRUFFLE] {
		if cs, err := db.truffle.Search(
			ctx,
			truffle.SearchOpts{
				Title:  query.Title,
				Corpus: query.Corpus,
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
			ctx,
			mal.SearchOpts{
				Title:  query.Title,
				Corpus: query.Corpus,
				Cutoff: query.MAL.Cutoff,
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
