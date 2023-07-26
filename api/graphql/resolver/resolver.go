package resolver

import (
	"github.com/minkezhang/truffle/api/graphql/generated/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Entry map[string]*model.Entry
}
