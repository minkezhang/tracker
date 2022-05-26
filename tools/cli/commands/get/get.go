package get

import (
	"fmt"
	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type O struct {
	DB *database.DB

	ID     string
	Title  string
	Corpus dpb.Corpus
}

func (o O) Get() (*dpb.Entry, error) {
	if o.DB == nil {
		return nil, nil
	}

	if o.ID != "" {
		return o.DB.GetEntry(o.ID)
	}

	results := o.DB.Search(database.O{Title: o.Title, Corpus: o.Corpus})
	if len(results) == 0 {
		return nil, nil
	}
	if len(results) == 1 {
		return results[0], nil
	}
	return nil, fmt.Errorf("multiple records found for query")
}
