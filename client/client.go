package client

import (
	"context"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type WO interface {
	Add(ctx context.Context, epb *dpb.Entry) (*dpb.Entry, error)
	Put(ctx context.Context, epb *dpb.Entry) (*dpb.Entry, error)
	Delete(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error)
}

type RO interface {
	Get(ctx context.Context, id *dpb.LinkedID, opts any) (*dpb.Entry, error)
	Search(ctx context.Context, query any) ([]*dpb.Entry, error)
}

type RW interface {
	RO
	WO
}
