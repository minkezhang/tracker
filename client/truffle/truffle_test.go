package truffle

import (
	"github.com/minkezhang/truffle/client"
)

var (
	_ client.RO[SearchOpts] = &C{}
)
