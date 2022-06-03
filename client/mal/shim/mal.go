// Package shim wraps the underlying MAL API client and return database entries.
//
// N.B.: While the caller may set Manga.List and Anime.List limits, this limit
// is actually a minimum -- the actual number of returned results varies between
// the given limit and the API max of 100.
//
// N.B.: MAL HTTP response includes a NextOffset field -- this field is actually
// reset to 0 upon reaching the end if the query, and will need to be explicitly
// handled (as otherwise any loop checking the length of returned results will
// loop forever).
//
// TODO(github.com/nstratos/go-myanimelist/pull/8): Add support for configurable NSFW searches.
package shim

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/nstratos/go-myanimelist/mal"

	cpb "github.com/minkezhang/truffle/api/go/config"
	dpb "github.com/minkezhang/truffle/api/go/database"
)

type EndpointT string

const (
	EndpointAnime EndpointT = "anime"
	EndpointManga           = "manga"
)

var (
	SupportedEndpoints = map[EndpointT]bool{
		EndpointAnime: true,
		EndpointManga: true,
	}
)

var (
	// lookup is a reverse lookup table from MAL "media_type" to an
	// associated dpb.Corpus value.
	//
	// See https://myanimelist.net/apiconfig/references/api/v2 for potential
	// media_type valus.
	lookup = map[string]dpb.Corpus{
		"tv":      dpb.Corpus_CORPUS_ANIME,
		"ova":     dpb.Corpus_CORPUS_ANIME,
		"ona":     dpb.Corpus_CORPUS_ANIME,
		"special": dpb.Corpus_CORPUS_ANIME,
		"movie":   dpb.Corpus_CORPUS_ANIME_FILM,

		// MAL lists the "novel" type but experimentally, this is
		// "light_novel" instead.
		"light_novel": dpb.Corpus_CORPUS_BOOK,

		"manga":     dpb.Corpus_CORPUS_MANGA,
		"one_shot":  dpb.Corpus_CORPUS_MANGA,
		"doujinshi": dpb.Corpus_CORPUS_MANGA,
		"manhua":    dpb.Corpus_CORPUS_MANGA,
		"manhwa":    dpb.Corpus_CORPUS_MANGA,
		"oel":       dpb.Corpus_CORPUS_MANGA,
	}

	animeFields = mal.Fields{
		"media_type",
		"popularity",
		"title",
		"mean",
		"studios",
	}
	mangaFields = mal.Fields{
		"media_type",
		"popularity",
		"title",
		"alternative_titles",
		"mean",
		"authors{first_name,last_name}",
	}
)

type a mal.Anime

func (r a) PB() *dpb.Entry {
	epb := &dpb.Entry{
		Id: &dpb.LinkedID{
			Id:  fmt.Sprintf("%v/%v", EndpointAnime, strconv.FormatInt(int64(r.ID), 10)),
			Api: dpb.API_API_MAL,
		},
		Titles: []string{r.Title},
		Score:  float32(r.Mean),
		Corpus: lookup[r.MediaType],
	}

	var studios []string
	for _, s := range r.Studios {
		studios = append(studios, s.Name)
	}
	epb.AuxData = &dpb.Entry_AuxDataVideo{
		AuxDataVideo: &dpb.AuxDataVideo{
			Studios: studios,
		},
	}

	return epb
}

type m mal.Manga

func (r m) PB() *dpb.Entry {
	epb := &dpb.Entry{
		Id: &dpb.LinkedID{
			Id:  fmt.Sprintf("%v/%v", EndpointManga, strconv.FormatInt(int64(r.ID), 10)),
			Api: dpb.API_API_MAL,
		},
		Titles: []string{r.Title},
		Score:  float32(r.Mean),
		Corpus: lookup[r.MediaType],
	}

	for _, t := range r.AlternativeTitles.Synonyms {
		epb.Titles = append(epb.GetTitles(), t)
	}

	var authors []string
	for _, a := range r.Authors {
		var names []string
		for _, n := range []string{a.Person.FirstName, a.Person.LastName} {
			if n != "" {
				names = append(names, n)
			}
		}
		authors = append(authors, strings.Join(names, " "))
	}
	epb.AuxData = &dpb.Entry_AuxDataBook{
		AuxDataBook: &dpb.AuxDataBook{
			Authors: authors,
		},
	}
	return epb
}

type C struct {
	client mal.Client
	config *cpb.MALConfig
}

func New(config *cpb.MALConfig) *C {
	return &C{
		client: *mal.NewClient(
			&http.Client{
				Transport: t{clientID: config.GetClientId()},
			}),
		config: config,
	}
}

func (c *C) AnimeGet(ctx context.Context, id int) (*dpb.Entry, error) {
	result, _, err := c.client.Anime.Details(ctx, id, animeFields)
	if err != nil {
		return nil, err
	}
	return a(*result).PB(), nil
}

func (c *C) MangaGet(ctx context.Context, id int) (*dpb.Entry, error) {
	result, _, err := c.client.Manga.Details(ctx, id, mangaFields)
	if err != nil {
		return nil, err
	}
	return m(*result).PB(), nil
}

func (c *C) AnimeSearch(ctx context.Context, title string, corpus dpb.Corpus) ([]*dpb.Entry, error) {
	f := func(r *mal.Response) ([]mal.Anime, *mal.Response, error) {
		if r != nil && r.NextOffset == 0 {
			return nil, nil, nil
		}
		var offset int
		if r != nil {
			offset = r.NextOffset
		}
		results, r, err := c.client.Anime.List(
			ctx, title, animeFields,
			mal.Limit(math.Min(100, float64(c.config.GetSearchMaxResults()))),
			mal.Offset(offset),
		)
		return results, r, err
	}

	var results []mal.Anime
	var err error

	for page, r, err := f(nil); err == nil && page != nil && len(results) <= int(c.config.GetSearchMaxResults()); page, r, err = f(r) {
		results = append(results, page...)
	}
	if err != nil {
		return nil, err
	}

	var epbs []*dpb.Entry
	for _, r := range results {
		// Trim obscure series.
		if popularity := c.config.GetPopularityCutoff(); popularity >= 0 && r.Popularity >= int(popularity) {
			continue
		}
		if corpus != dpb.Corpus_CORPUS_UNKNOWN && corpus != lookup[r.MediaType] {
			continue
		}

		epbs = append(epbs, a(r).PB())
	}

	return epbs, nil
}

func (c *C) MangaSearch(ctx context.Context, title string, corpus dpb.Corpus) ([]*dpb.Entry, error) {
	f := func(r *mal.Response) ([]mal.Manga, *mal.Response, error) {
		if r != nil && r.NextOffset == 0 {
			return nil, nil, nil
		}
		var offset int
		if r != nil {
			offset = r.NextOffset
		}
		results, r, err := c.client.Manga.List(
			ctx, title, mangaFields,
			mal.Limit(math.Min(100, float64(c.config.GetSearchMaxResults()))),
			mal.Offset(offset),
		)
		return results, r, err
	}

	var results []mal.Manga
	var err error

	for page, r, err := f(nil); err == nil && page != nil && len(results) <= int(c.config.GetSearchMaxResults()); page, r, err = f(r) {
		results = append(results, page...)
	}
	if err != nil {
		return nil, err
	}

	var epbs []*dpb.Entry
	for _, r := range results {
		// Trim obscure series.
		if popularity := c.config.GetPopularityCutoff(); popularity >= 0 && r.Popularity >= int(popularity) {
			continue
		}
		if corpus != dpb.Corpus_CORPUS_UNKNOWN && corpus != lookup[r.MediaType] {
			continue
		}

		epbs = append(epbs, m(r).PB())
	}

	return epbs, nil
}

type t struct {
	clientID string
}

func (t t) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-MAL-CLIENT-ID", t.clientID)
	return http.DefaultTransport.RoundTrip(req)
}
