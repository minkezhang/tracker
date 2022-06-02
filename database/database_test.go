package database

import (
	"github.com/minkezhang/truffle/client"
)

var (
	_ client.RW = &DB{}
)
