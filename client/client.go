package client

import (
	"context"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type C interface {
	Get(ctx context.Context, id string, corpus dpb.Corpus) (*dpb.Entry, error)
	Search(ctx context.Context, query string, corpus dpb.Corpus) ([]*dpb.Entry, error)
}
