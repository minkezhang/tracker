package resolver

import (
	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *DB
}

type DB struct {
	Entry   *database.Entry
	APIData map[model.APIType]*database.APIData
}
