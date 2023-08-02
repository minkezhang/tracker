package client

import (
	"context"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type AuthType int

const (
	AuthTypeNone   AuthType = 0
	AuthTypePublic AuthType = 1 << iota
	AuthTypePrivateRead
	AuthTypePrivateWrite
)

func (a AuthType) Check(b AuthType) bool { return a&b == b }

type C interface {
	Auth() AuthType
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
