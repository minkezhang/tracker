package utils

import (
	cpb "github.com/minkezhang/truffle/api/go/config"
)

type MAL struct {
	ClientID         string
	PopularityCutoff int
	NSFW             bool
	SearchMaxResults int
}

func (mal MAL) PB() (*cpb.MALConfig, error) {
	return &cpb.MALConfig{
		ClientId:         mal.ClientID,
		PopularityCutoff: int32(mal.PopularityCutoff),
		Nsfw:             mal.NSFW,
		SearchMaxResults: int32(mal.SearchMaxResults),
	}, nil
}

type Config struct {
	MAL MAL
}

func (c Config) PB() (*cpb.Config, error) {
	mal, err := c.MAL.PB()
	if err != nil {
		return nil, err
	}

	return &cpb.Config{
		Mal: mal,
	}, nil
}
