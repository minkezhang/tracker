package mal

import (
	"context"
	"fmt"
	"strings"

	"github.com/minkezhang/truffle/api/graphql/model"
)

type MAL struct {
	manga *Manga
}

func New(o O) *MAL {
	return &MAL{
		manga: NewManga(o),
	}
}

func (c *MAL) API() model.APIType { return model.APITypeAPIMal }

func (c *MAL) Get(ctx context.Context, id string) (*model.APIData, error) {
	parts := strings.Split(id, "/")
	corpus := parts[0]
	if corpus == "manga" {
		return c.manga.Get(ctx, parts[1])
	}
	return nil, fmt.Errorf("unimplemented MAL corpus: %s", corpus)
}
