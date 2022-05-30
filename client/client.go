package client

import (
	"context"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type RO[T any] interface {
	Get(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error)
	Search(ctx context.Context, query T) ([]*dpb.Entry, error)
}
