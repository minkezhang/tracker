package mal

import (
	"context"

	"github.com/minkezhang/truffle/client/mal/shim"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type SearchOpts struct {
	Cutoff int

	Title  string
	Corpus dpb.Corpus
}

type C struct {
	client *shim.C
}

func New() *C { return &C{client: shim.New()} }

func (c C) Get(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error) {
	return nil, status.Errorf(codes.Unimplemented, "")
}

func (c C) Search(ctx context.Context, query SearchOpts) ([]*dpb.Entry, error) {
	var candidates []*dpb.Entry
	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_ANIME:      true,
		dpb.Corpus_CORPUS_ANIME_FILM: true,
		dpb.Corpus_CORPUS_UNKNOWN:    true,
	}[query.Corpus] {
		cs, err := c.client.AnimeSearch(ctx, query.Title, query.Corpus, query.Cutoff)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cs...)
	}
	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_BOOK:    true,
		dpb.Corpus_CORPUS_MANGA:   true,
		dpb.Corpus_CORPUS_UNKNOWN: true,
	}[query.Corpus] {
		cs, err := c.client.MangaSearch(ctx, query.Title, query.Corpus, query.Cutoff)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cs...)
	}
	return candidates, nil
}
