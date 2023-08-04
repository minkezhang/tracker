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
	client.Base

	manga *Manga
	anime *Anime
}

func New(o O) *MAL {
	c := mal.NewClient(
		&http.Client{
			Transport: o.transport,
		},
	)

	manga := NewManga(c, o.auth)
	anime := NewAnime(c, o.auth)
	var corpora []model.CorpusType
	for c := range manga.Corpora() {
		corpora = append(corpora, c)
	}
	for c := range anime.Corpora() {
		corpora = append(corpora, c)
	}

	return &MAL{
		Base: *client.New(client.O{
			API:     model.APITypeAPIMal,
			Auth:    o.auth,
			Corpora: corpora,
		}),

		manga: manga,
		anime: anime,
	}
}

func (c *MAL) Get(ctx context.Context, id string) (*model.APIData, error) {
	if !c.Auth().Check(client.AuthTypePublic) {
		return nil, nil
	}

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

func (c *MAL) List(ctx context.Context, query *model.ListInput) ([]*model.APIData, error) {
	if !c.Auth().Check(client.AuthTypePublic) {
		return nil, nil
	}

	// TODO(minkezhang): Handle paging.
	d, err := c.manga.List(ctx, query)
	if err != nil {
		return nil, err
	}
	e, err := c.anime.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return append(d, e...), nil
}
