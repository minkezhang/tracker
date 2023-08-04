package mal

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/minkezhang/truffle/api/graphql/model"
	"github.com/minkezhang/truffle/client"
	"github.com/minkezhang/truffle/util"
	"github.com/nstratos/go-myanimelist/mal"
)

var (
	mangaFields = mal.Fields{
		"media_type",
		"popularity",
		"title",
		"alternative_titles",
		"mean",
		"authors{first_name,last_name}",
		"my_list_status",
		"genres",
	}
	mangaCorpusLookup = map[string]model.CorpusType{
		// MAL lists the "novel" type but from experimentation, this is
		// "light_novel" instead.
		"light_novel": model.CorpusTypeCorpusBook,

		"manga":     model.CorpusTypeCorpusManga,
		"one_shot":  model.CorpusTypeCorpusManga,
		"doujinshi": model.CorpusTypeCorpusManga,
		"manhua":    model.CorpusTypeCorpusManga,
		"manhwa":    model.CorpusTypeCorpusManga,
		"oel":       model.CorpusTypeCorpusManga,
	}

	mangaQueuedLookup = map[mal.MangaStatus]bool{
		mal.MangaStatusReading:    true,
		mal.MangaStatusPlanToRead: true,
	}
)

type Manga struct {
	client.Base

	client *mal.Client
}

func NewManga(c *mal.Client, auth client.AuthType, config util.Config) *Manga {
	return &Manga{
		Base: *client.New(client.O{
			API:  model.APITypeAPIMal,
			Auth: auth,
			Corpora: []model.CorpusType{
				model.CorpusTypeCorpusManga,
				model.CorpusTypeCorpusBook,
			},
			Config: config,
		}),

		client: c,
	}
}

func (c *Manga) APIData(m *mal.Manga) *model.APIData {
	var artists []string
	var authors []string

	for _, a := range m.Authors {
		s := strings.Join([]string{a.Person.FirstName, a.Person.LastName}, " ")
		if strings.Contains(a.Role, "Story") {
			authors = append(authors, s)
		}
		if strings.Contains(a.Role, "Art") {
			artists = append(artists, s)
		}
	}

	var genres []string
	for _, g := range m.Genres {
		genres = append(genres, g.Name)
	}

	tags := genres
	score := m.Mean
	var t *model.TrackerManga

	if c.Auth().Check(client.AuthTypePrivateRead) {
		tags = append(tags, m.MyListStatus.Tags...)
		if m.MyListStatus.Score > 0 {
			score = float64(m.MyListStatus.Score)
		}
		if m.MyListStatus.NumVolumesRead > 0 || m.MyListStatus.NumChaptersRead > 0 {
			v := fmt.Sprintf("%d", m.MyListStatus.NumVolumesRead)
			ch := fmt.Sprintf("%d", m.MyListStatus.NumChaptersRead)
			t = &model.TrackerManga{
				Volume:      &v,
				Chapter:     &ch,
				LastUpdated: &m.MyListStatus.UpdatedAt,
			}
		}
	}

	return &model.APIData{
		API:    c.API(),
		ID:     fmt.Sprintf("manga/%d", m.ID),
		Cached: true,
		Corpus: mangaCorpusLookup[m.MediaType],
		Titles: []*model.Title{
			&model.Title{
				Locale: "en",
				Title:  m.Title,
			},
			&model.Title{
				Locale: "en",
				Title:  m.AlternativeTitles.En,
			},
			&model.Title{
				Locale: "ja_jp",
				Title:  m.AlternativeTitles.Ja,
			},
		},
		Queued: mangaQueuedLookup[m.MyListStatus.Status],
		Score:  &score,
		Aux: &model.AuxManga{
			Authors: authors,
			Artists: artists,
		},
		Tags:    tags,
		Tracker: t,
	}
}

func (c *Manga) Get(ctx context.Context, id string) (*model.APIData, error) {
	malID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	d, resp, err := c.client.Manga.Details(ctx, malID, mangaFields)
	if err != nil {
		return nil, fmt.Errorf("cannot get %s:%v (%d)", c.API(), id, resp.StatusCode)
	}
	return c.APIData(d), nil
}

func (c *Manga) List(ctx context.Context, query *model.ListInput) ([]*model.APIData, error) {
	var q string
	if query.Title != nil {
		q = *query.Title
	}

	f := func(r *mal.Response) ([]mal.Manga, *mal.Response, error) {
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

		ds, r, err := c.client.Manga.List(
			ctx,
			q,
			mangaFields,
			mal.Limit(math.Min(100, float64(c.Config().MAL.SearchMaxResults))),
			mal.Offset(offset),
			mal.NSFW(nsfw),
		)
		return ds, r, err
	}

	var ds []mal.Manga

	var page []mal.Manga
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
