package common

import (
	"flag"
	"fmt"

	"github.com/minkezhang/tracker/api/go/database/utils"
	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type C struct {
	DB *database.DB

	ID     string
	Title  string
	Corpus string
}

func (c *C) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.ID, "id", "", "entry ID")
	f.StringVar(&c.Title, "title", "", "entry title substring")
	f.StringVar(&c.Corpus, "corpus", "unknown", "optional corpus hint for the entry")
}

func (c *C) Get() (*dpb.Entry, error) {
	var entries []*dpb.Entry
	if c.ID != "" {
		results, _ := c.DB.GetEntry(c.ID)
		entries = append(entries, results)
	} else {
		corpus := dpb.Corpus(
			dpb.Corpus_value[utils.ToEnum("CORPUS", c.Corpus)])

		candidates, err := c.DB.Search(database.O{
			Title:  c.Title,
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
