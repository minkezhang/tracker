package common

import (
	"fmt"

	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type O struct {
	DB *database.DB

	ID     string
	Title  string
	Corpus string
}

func Get(opts O) (*dpb.Entry, error) {
	var entries []*dpb.Entry
	if opts.ID != "" {
		results, _ := opts.DB.GetEntry(opts.ID)
		entries = append(entries, results)
	} else {
		corpus := dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", opts.Corpus)])

		candidates, err := opts.DB.Search(database.O{
			Title:  opts.Title,
			Corpus: corpus,
			APIs:   []dpb.API{dpb.API_API_TRACKER},
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
		fmt.Printf("%s", entries)
		return nil, fmt.Errorf("Too many results returned. Please refine your search.")
	}

	return entries[0], nil
}
