package client

import (
	"context"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type C interface {
	API() model.APIType
	Get(ctx context.Context, id string) (*model.APIData, error)
}
