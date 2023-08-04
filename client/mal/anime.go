package mal

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client"
	"github.com/minkezhang/truffle/util"
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
	animeCorpusLookup = map[string]model.CorpusType{
		"tv":      model.CorpusTypeCorpusAnime,
		"ova":     model.CorpusTypeCorpusAnime,
		"ona":     model.CorpusTypeCorpusAnime,
		"special": model.CorpusTypeCorpusAnime,
		"movie":   model.CorpusTypeCorpusAnimeFilm,
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

func NewAnime(c *mal.Client, auth client.AuthType, config util.Config) *Anime {
	return &Anime{
		Base: *client.New(client.O{
			API:  model.APITypeAPIMal,
			Auth: auth,
			Corpora: []model.CorpusType{
				model.CorpusTypeCorpusAnime,
				model.CorpusTypeCorpusAnimeFilm,
			},
			Config: config,
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

	tags := genres
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
		Corpus: animeCorpusLookup[a.MediaType],
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

	d, resp, err := c.client.Anime.Details(ctx, malID, animeFields)
	if err != nil {
		return nil, fmt.Errorf("cannot get %s:%v (%d)", c.API(), id, resp.StatusCode)
	}
	return c.APIData(d), nil
}

func (c *Anime) List(ctx context.Context, query *model.ListInput) ([]*model.APIData, error) {
	var q string
	if query.Title != nil {
		q = *query.Title
	}

	f := func(r *mal.Response) ([]mal.Anime, *mal.Response, error) {
		// Handle EOF case.
		if r != nil && r.NextOffset == 0 {
			return nil, nil, nil
		}

		var offset int
		if r != nil {
			offset = r.NextOffset
		}

		var nsfw bool
		if query.Mal != nil {
			nsfw = query.Mal.Nsfw
		}

		ds, r, err := c.client.Anime.List(
			ctx,
			q,
			animeFields,
			mal.Limit(math.Min(100, float64(c.Config().MAL.SearchMaxResults))),
			mal.Offset(offset),
			mal.NSFW(nsfw),
		)
		return ds, r, err
	}

	var ds []mal.Anime

	var page []mal.Anime
	var r *mal.Response
	var err error
	for page, r, err = f(nil); err == nil && page != nil && len(ds) <= c.Config().MAL.SearchMaxResults; page, r, err = f(r) {
		ds = append(ds, page...)
	}
	if err != nil {
		return nil, err
	}

	var data []*model.APIData
	for _, d := range ds {
		// Trim obscure series.
		if popularity := c.Config().MAL.PopularityCutoff; popularity >= 0 && d.Popularity >= popularity {
			continue
		}

		data = append(data, c.APIData(&d))
	}
	return data, nil
}
