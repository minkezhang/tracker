package search

import (
	dpb "github.com/minkezhang/tracker/api/go/database"
)

type S interface {
	Search() ([]*dpb.Entry, error)
}
