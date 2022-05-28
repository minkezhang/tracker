package search

import (
	dpb "github.com/minkezhang/truffle/api/go/database"
)

type S interface {
	Search() ([]*dpb.Entry, error)
}
