package mal

import (
	"context"

	"github.com/minkezhang/truffle/clients/mal"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

const (
	cutoff = 2000
)

type S struct {
	client *mal.C

	title  string
	corpus dpb.Corpus
}

func New(title string, corpus dpb.Corpus) *S {
	return &S{
		client: mal.New(),
		title:  title,
		corpus: corpus,
	}
}

func (s *S) Search(ctx context.Context) ([]*dpb.Entry, error) {
	var candidates []*dpb.Entry
	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_ANIME:      true,
		dpb.Corpus_CORPUS_ANIME_FILM: true,
	}[s.corpus] {
		cs, err := s.client.AnimeSearch(ctx, s.title, s.corpus, cutoff)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cs...)
	}
	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_BOOK:  true,
		dpb.Corpus_CORPUS_MANGA: true,
	}[s.corpus] {
		cs, err := s.client.MangaSearch(ctx, s.title, s.corpus, cutoff)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cs...)
	}
	return candidates, nil
}
