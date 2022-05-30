package mal

import (
	"context"

	"github.com/minkezhang/truffle/client/mal"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type O struct {
	Cutoff int
}

type S struct {
	client *mal.C

	title  string
	corpus dpb.Corpus

	cutoff int
}

func New(title string, corpus dpb.Corpus, cutoff int) *S {
	return &S{
		client: mal.New(),
		title:  title,
		corpus: corpus,
		cutoff: cutoff,
	}
}

func (s *S) Search(ctx context.Context) ([]*dpb.Entry, error) {
	var candidates []*dpb.Entry
	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_ANIME:      true,
		dpb.Corpus_CORPUS_ANIME_FILM: true,
		dpb.Corpus_CORPUS_UNKNOWN:    true,
	}[s.corpus] {
		cs, err := s.client.AnimeSearch(ctx, s.title, s.corpus, s.cutoff)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cs...)
	}
	if map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_BOOK:    true,
		dpb.Corpus_CORPUS_MANGA:   true,
		dpb.Corpus_CORPUS_UNKNOWN: true,
	}[s.corpus] {
		cs, err := s.client.MangaSearch(ctx, s.title, s.corpus, s.cutoff)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, cs...)
	}
	return candidates, nil
}
