package mal

import (
	"context"
	"strconv"
	"strings"

	"github.com/minkezhang/truffle/client/mal/shim"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	cpb "github.com/minkezhang/truffle/api/go/config"
	dpb "github.com/minkezhang/truffle/api/go/database"
)

var (
	animeCorpora = map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_ANIME:      true,
		dpb.Corpus_CORPUS_ANIME_FILM: true,
		dpb.Corpus_CORPUS_UNKNOWN:    true,
	}

	mangaCorpora = map[dpb.Corpus]bool{
		dpb.Corpus_CORPUS_BOOK:    true,
		dpb.Corpus_CORPUS_MANGA:   true,
		dpb.Corpus_CORPUS_UNKNOWN: true,
	}
)

type SearchOpts struct {
	Title  string
	Corpus dpb.Corpus
}

type C struct {
	client *shim.C
}

func New(config *cpb.MALConfig) *C { return &C{client: shim.New(config)} }

func (c C) Get(ctx context.Context, id *dpb.LinkedID, opts interface{}) (*dpb.Entry, error) {
	if id.GetApi() != dpb.API_API_MAL {
		return nil, status.Errorf(codes.InvalidArgument, "cannot use the MAL client to look up non-MAL IDs")
	}
	endpoint, sid, ok := strings.Cut(id.GetId(), "/")
	if !ok || !shim.SupportedEndpoints[shim.EndpointT(endpoint)] {
		return nil, status.Errorf(codes.InvalidArgument, "invalid MAL endpoint")
	}
	lid, err := strconv.Atoi(sid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid MAL ID: %v", err)
	}

	if shim.EndpointT(endpoint) == shim.EndpointAnime {
		epb, err := c.client.AnimeGet(ctx, lid)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot get MAL entry: %v", err)
		}
		return epb, nil
	}
	if shim.EndpointT(endpoint) == shim.EndpointManga {
		epb, err := c.client.MangaGet(ctx, lid)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot get MAL entry: %v", err)
		}
		return epb, nil
	}
	return nil, nil
}

func (c C) Search(ctx context.Context, opts interface{}) ([]*dpb.Entry, error) {
	query, ok := opts.(SearchOpts)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid search opts provided")
	}

	var candidates []*dpb.Entry
	if animeCorpora[query.Corpus] {
		cs, err := c.client.AnimeSearch(ctx, query.Title, query.Corpus)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error while getting data from MAL: %v", err)
		}
		candidates = append(candidates, cs...)
	}
	if mangaCorpora[query.Corpus] {
		cs, err := c.client.MangaSearch(ctx, query.Title, query.Corpus)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error while getting data from MAL: %v", err)
		}
		candidates = append(candidates, cs...)
	}
	return candidates, nil
}
