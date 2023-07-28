package mal

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/minkezhang/truffle/api/graphql/model"
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

	mangaQueuedLookup = map[mal.MangaStatus]bool{
		mal.MangaStatusReading:    true,
		mal.MangaStatusPlanToRead: true,
	}
)

type Manga struct {
	client *mal.Client
}

func NewManga(c *mal.Client) *Manga {
	return &Manga{
		client: c,
	}
}

func (c *Manga) API() model.APIType { return model.APITypeAPIMal }

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

	return &model.APIData{
		API:    c.API(),
		ID:     fmt.Sprintf("manga/%d", m.ID),
		Cached: true,
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
		Score:  &m.Mean,
		Aux: &model.AuxManga{
			Authors: authors,
			Artists: artists,
		},
		Tags: append(genres, m.MediaType),
	}
}

func (c *Manga) Get(ctx context.Context, id string) (*model.APIData, error) {
	malID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	m, resp, err := c.client.Manga.Details(ctx, malID, mangaFields)
	if err != nil {
		return nil, fmt.Errorf("cannot get %s:%v (%d)", c.API(), id, resp.StatusCode)
	}
	return c.APIData(m), nil
}

func (c *Manga) List(ctx context.Context, q *model.ListInput) ([]*model.APIData, error) {
	return nil, fmt.Errorf("unimplemented")
}
