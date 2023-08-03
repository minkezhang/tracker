package mal

import (
	"context"
	"fmt"
	"strconv"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client"
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
	client.Base

	client *mal.Client
}

func NewAnime(c *mal.Client, auth client.AuthType) *Anime {
	return &Anime{
		Base: *client.New(client.O{
			API:  model.APITypeAPIMal,
			Auth: auth,
			Corpora: []model.CorpusType{
				model.CorpusTypeCorpusAnime,
				model.CorpusTypeCorpusAnimeFilm,
			},
		}),

		client: c,
	}
}

func (c *Anime) APIData(a *mal.Anime) *model.APIData {
	var studios []string

	for _, s := range a.Studios {
		studios = append(studios, s.Name)
	}

	var genres []string
	for _, g := range a.Genres {
		genres = append(genres, g.Name)
	}

	tags := append(genres, a.MediaType)
	score := a.Mean

	if c.Auth().Check(client.AuthTypePrivateRead) {
		tags = append(tags, a.MyListStatus.Tags...)
		if a.MyListStatus.Score > 0 {
			score = float64(a.MyListStatus.Score)
		}
	}

	return &model.APIData{
		API:    c.API(),
		ID:     fmt.Sprintf("anime/%d", a.ID),
		Cached: true,

		// TODO(minkezhang): Conditionally handle anime films.
		Corpus: model.CorpusTypeCorpusAnime,

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
		Score:  &score,
		Aux: &model.AuxAnime{
			Studios: studios,
		},
		Tags: tags,
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
