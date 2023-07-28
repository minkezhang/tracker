package client

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type Client interface {
	Get(ctx context.Context, s *model.APIData) (*model.APIData, error)
	List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error)
}

type Cache struct {
}

func (c Cache) Get(ctx context.Context, s *model.APIData) (*model.APIData, error) {
	return s, nil
}

func (c Cache) List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error) {
	return nil, fmt.Errorf("unimplemented")
}
