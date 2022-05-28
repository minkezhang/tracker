package search

import (
	"context"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type S interface {
	Search(ctx context.Context) ([]*dpb.Entry, error)
}
