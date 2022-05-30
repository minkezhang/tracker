package truffle

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/minkezhang/truffle/api/go/database/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type SearchOpts struct {
	Title  string
	Corpus dpb.Corpus
}

type C struct {
	db *dpb.Database
}

func New(db *dpb.Database) *C { return &C{db: db} }

func (s *C) Get(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error) {
	if id.GetApi() != dpb.API_API_TRUFFLE {
		return nil, status.Errorf(codes.InvalidArgument, "cannot look up entry for given API %v", id.GetApi())
	}

	epb, ok := s.db.GetEntries()[id.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "entry with ID %v not found", utils.ID(id))
	}
	return proto.Clone(epb).(*dpb.Entry), nil
}

func (s *C) Search(ctx context.Context, query SearchOpts) ([]*dpb.Entry, error) {
	var candidates []*dpb.Entry
	for _, epb := range s.db.GetEntries() {
		epb = proto.Clone(epb).(*dpb.Entry)

		if query.Title == "" && len(epb.Titles) == 0 {
			candidates = append(candidates, epb)
			continue
		}
		for _, t := range epb.GetTitles() {
			m := map[dpb.Corpus]bool{
				epb.GetCorpus():           true,
				dpb.Corpus_CORPUS_UNKNOWN: true,
			}
			if strings.Contains(strings.ToLower(t), strings.ToLower(query.Title)) && (m[query.Corpus] || (epb.GetCorpus() == dpb.Corpus_CORPUS_UNKNOWN)) {
				candidates = append(candidates, epb)
				continue
			}
		}
	}
	return candidates, nil
}
