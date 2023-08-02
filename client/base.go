package client

import (
	"context"
	"fmt"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type Base struct {
	api  model.APIType
	auth AuthType
}

func New(api model.APIType, auth AuthType) *Base {
	return &Base{
		auth: auth,
	}
}

func (c Base) API() model.APIType { return c.api }
func (c Base) Auth() AuthType     { return c.auth }

func (c Base) Put(ctx context.Context, d *model.APIData) error {
	if !c.Auth().Check(AuthTypePrivateWrite) {
		return nil
	}

	return fmt.Errorf("unimplemented")
}

func (c Base) List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error) {
	if !c.Auth().Check(AuthTypePublic) {
		return nil, nil
	}

	return nil, fmt.Errorf("unimplemented")
}

func (c Base) Remove(ctx context.Context, id string) error {
	if !c.Auth().Check(AuthTypePrivateWrite) {
		return nil
	}

	return fmt.Errorf("unimplemented")
}
