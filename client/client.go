package client

import (
	"context"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type C interface {
	API() model.APIType

	// Get all info from the client.
	Get(ctx context.Context, id string) (*model.APIData, error)

	List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error)

	// Put will update user-related client info.
	Put(ctx context.Context, d *model.APIData) error

	// Remove will delete user-related client info, but will not touch the
	// read-only client data-source.
	Remove(ctx context.Context, id string) error
}
