package client

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/api/graphql/generated/model"
)

type Client interface {
	Get(ctx context.Context, s *model.APIData) (*model.APIData, error)
	Search(ctx context.Context, q *model.SearchInput) ([]*model.APIData, error)
}

type Cache struct {
}

func (c Cache) Get(ctx context.Context, s *model.APIData) (*model.APIData, error) {
	return s, nil
}

func (c Cache) Search(ctx context.Context, q *model.SearchInput) ([]*model.APIData, error) {
	return nil, fmt.Errorf("unimplemented")
}
