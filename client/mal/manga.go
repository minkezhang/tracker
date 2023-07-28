package mal

import (
	"context"
	"fmt"
	"net/http"
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
	}

	queuedLookup = map[mal.MangaStatus]bool{
		mal.MangaStatusReading:    true,
		mal.MangaStatusPlanToRead: true,
	}
)

type Manga struct {
	client *mal.Client
}

func NewManga(o O) *Manga {
	return &Manga{
		client: mal.NewClient(
			&http.Client{
				Transport: &transport{
					ClientID: o.ClientID,
				},
			},
		),
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

	return &model.APIData{
		API:    model.APITypeAPIMal,
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
		Queued: queuedLookup[m.MyListStatus.Status],
		Score:  &m.Mean,
		Aux: &model.AuxManga{
			Authors: authors,
			Artists: artists,
		},
		Tags: []string{
			m.MediaType,
		},
	}
}

func (c *Manga) Get(ctx context.Context, id string) (*model.APIData, error) {
	parts := strings.Split(id, "/")
	malID, err := strconv.Atoi(parts[len(parts)-1])
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
