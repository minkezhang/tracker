package mal

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client"
	"github.com/nstratos/go-myanimelist/mal"
)

var (
	_ client.C = &MAL{}
)

type MAL struct {
	manga *Manga
	anime *Anime
	auth  client.AuthType
}

func New(o O) *MAL {
	c := mal.NewClient(
		&http.Client{
			Transport: o.transport,
		},
	)
	return &MAL{
		manga: NewManga(c),
		anime: NewAnime(c),
		auth:  o.auth,
	}
}

func (c *MAL) Auth() client.AuthType { return c.auth }

func (c *MAL) API() model.APIType { return model.APITypeAPIMal }

func (c *MAL) Get(ctx context.Context, id string) (*model.APIData, error) {
	parts := strings.Split(id, "/")
	corpus := parts[0]
	if corpus == "manga" {
		return c.manga.Get(ctx, parts[1])
	}
	if corpus == "anime" {
		return c.anime.Get(ctx, parts[1])
	}
	return nil, fmt.Errorf("unimplemented MAL corpus: %s", corpus)
}

func (c *MAL) Put(ctx context.Context, d *model.APIData) error {
	return fmt.Errorf("unimplemented")
}

func (c *MAL) List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (c *MAL) Remove(ctx context.Context, id string) error {
	return fmt.Errorf("unimplemented")
}
