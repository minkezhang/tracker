package mal

import (
	"context"
	"fmt"
	"strconv"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/nstratos/go-myanimelist/mal"
)

var (
	animeFields = mal.Fields{
		"media_type",
		"alternative_titles",
		"popularity",
		"title",
		"mean",
		"studios",
		"genres",
		"my_list_status",
	}

	animeQueuedLookup = map[mal.AnimeStatus]bool{
		mal.AnimeStatusWatching:    true,
		mal.AnimeStatusPlanToWatch: true,
	}
)

type Anime struct {
	client *mal.Client
}

func NewAnime(c *mal.Client) *Anime {
	return &Anime{
		client: c,
	}
}

func (c *Anime) API() model.APIType { return model.APITypeAPIMal }

func (c *Anime) APIData(a *mal.Anime) *model.APIData {
	var studios []string

	for _, s := range a.Studios {
		studios = append(studios, s.Name)
	}

	var genres []string
	for _, g := range a.Genres {
		genres = append(genres, g.Name)
	}

	return &model.APIData{
		API:    c.API(),
		ID:     fmt.Sprintf("anime/%d", a.ID),
		Cached: true,
		Titles: []*model.Title{
			&model.Title{
				Locale: "en",
				Title:  a.Title,
			},
			&model.Title{
				Locale: "en",
				Title:  a.AlternativeTitles.En,
			},
			&model.Title{
				Locale: "ja_jp",
				Title:  a.AlternativeTitles.Ja,
			},
		},
		Queued: animeQueuedLookup[a.MyListStatus.Status],
		Score:  &a.Mean,
		Aux: &model.AuxAnime{
			Studios: studios,
		},
		Tags: append(genres, a.MediaType),
	}
}

func (c *Anime) Get(ctx context.Context, id string) (*model.APIData, error) {
	malID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	m, resp, err := c.client.Anime.Details(ctx, malID, animeFields)
	if err != nil {
		return nil, fmt.Errorf("cannot get %s:%v (%d)", c.API(), id, resp.StatusCode)
	}
	return c.APIData(m), nil
}

func (c *Anime) List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error) {
	return nil, fmt.Errorf("unimplemented")
}
