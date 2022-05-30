package get

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/database"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type O struct {
	ID     *dpb.LinkedID
	Title  string
	Corpus dpb.Corpus
}

func Get(ctx context.Context, db *database.DB, opts O) (*dpb.Entry, error) {
	var entries []*dpb.Entry

	if opts.ID != nil {
		results, err := db.Get(ctx, opts.ID)
		if err != nil {
			return nil, err
		}

		entries = append(entries, results)
	} else {
		candidates, err := db.Search(ctx, database.SearchOpts{
			Title:  opts.Title,
			Corpus: opts.Corpus,
			APIs:   []dpb.API{dpb.API_API_TRUFFLE},
		})
		if err != nil {
			return nil, err
		}

		entries = append(entries, candidates...)
	}

	if len(entries) == 0 {
		return nil, fmt.Errorf("Could not find result with the given input. Please refine your search.")
	}

	if len(entries) > 1 {
		return nil, fmt.Errorf("Too many results returned. Please refine your search.")
	}

	return entries[0], nil
}
