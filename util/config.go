package util

import (
	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client/mal"
)

func GenerateConfig(c *model.Config) {
	if c.Mal == nil {
		c.Mal = &model.ConfigMal{
			ClientID:         mal.CLIENT_ID,
			PopularityCutoff: 1500,
			SearchMaxResults: 50,
		}
	}
}
