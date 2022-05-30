package common

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/database"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type O struct {
	DB *database.DB

	ID     string
	Title  string
	Corpus string
}

func Get(ctx context.Context, opts O) (*dpb.Entry, error) {
	var entries []*dpb.Entry
	if opts.ID != "" {
		results, err := opts.DB.Get(ctx,
			&dpb.LinkedID{
				Id:  opts.ID,
				Api: dpb.API_API_TRUFFLE,
			})
		if err != nil {
			return nil, err
		}
		entries = append(entries, results)
	} else {
		corpus := dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", opts.Corpus)])

		candidates, err := opts.DB.Search(ctx, database.SearchOpts{
			Title:  opts.Title,
			Corpus: corpus,
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
