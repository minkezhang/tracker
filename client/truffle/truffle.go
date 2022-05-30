package truffle

import (
	"context"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/api/go/database/validator"
	"github.com/minkezhang/truffle/client/truffle/id"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type SearchOpts struct {
	Title  string
	Corpus dpb.Corpus
}

type C struct {
	db *dpb.Database
}

func New(db *dpb.Database) *C {
	db = proto.Clone(db).(*dpb.Database)
	if db == nil {
		db = &dpb.Database{}
	}
	if db.GetEntries() == nil {
		db.Entries = map[string]*dpb.Entry{}
	}
	return &C{db: db}
}

func (c *C) Add(ctx context.Context, epb *dpb.Entry) (*dpb.Entry, error) {
	epb = proto.Clone(epb).(*dpb.Entry)
	if err := validator.Validate(epb); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	eid := id.New()
	for ; c.db.GetEntries()[eid] != nil; eid = id.New() {
	}

	epb.Id = &dpb.LinkedID{
		Id:  eid,
		Api: dpb.API_API_TRUFFLE,
	}
	etag, err := utils.ETag(epb)
	if err != nil {
		return nil, err
	}
	epb.Etag = etag

	c.db.GetEntries()[eid] = epb
	return proto.Clone(epb).(*dpb.Entry), nil
}

func (c *C) Put(ctx context.Context, epb *dpb.Entry) (*dpb.Entry, error) {
	epb = proto.Clone(epb).(*dpb.Entry)
	if err := validator.Validate(epb); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	fpb, ok := c.db.GetEntries()[epb.GetId().GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "cannot find entry with id %v", epb.GetId())
	}
	if !reflect.DeepEqual(epb.GetEtag(), fpb.GetEtag()) {
		return nil, status.Errorf(codes.InvalidArgument, "cannot update entry with mismatching ETag values: %s != %s", epb.GetEtag(), fpb.GetEtag())
	}
	etag, err := utils.ETag(epb)
	if err != nil {
		return nil, err
	}
	epb.Etag = etag

	c.db.GetEntries()[epb.GetId().GetId()] = epb
	return epb, nil
}

func (c *C) Delete(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error) {
	epb, err := c.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	delete(c.db.GetEntries(), id.GetId())
	return epb, nil
}

func (c *C) Get(ctx context.Context, id *dpb.LinkedID) (*dpb.Entry, error) {
	if id.GetApi() != dpb.API_API_TRUFFLE {
		return nil, status.Errorf(codes.InvalidArgument, "cannot look up entry for given API %v", id.GetApi())
	}

	epb, ok := c.db.GetEntries()[id.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "entry with ID %v not found", utils.ID(id))
	}
	return proto.Clone(epb).(*dpb.Entry), nil
}

func (c *C) Search(ctx context.Context, query SearchOpts) ([]*dpb.Entry, error) {
	var candidates []*dpb.Entry
	for _, epb := range c.db.GetEntries() {
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

func Marshal(c *C) ([]byte, error) {
	data, err := prototext.MarshalOptions{
		Multiline: true,
	}.Marshal(c.db)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot marshal DB entry: %v", err)
	}
	return data, nil
}

func Unmarshal(data []byte) (*C, error) {
	pb := &dpb.Database{}
	if err := prototext.Unmarshal(data, pb); err != nil {
		return nil, status.Errorf(codes.Internal, "cannot unmarshal DB entry: %v", err)
	}

	if pb.GetEntries() == nil {
		pb.Entries = map[string]*dpb.Entry{}
	}

	for eid, epb := range pb.GetEntries() {
		epb.Id = &dpb.LinkedID{
			Id:  eid,
			Api: dpb.API_API_TRUFFLE,
		}
		etag, err := utils.ETag(epb)
		if err != nil {
			return nil, err
		}
		epb.Etag = etag
	}

	return &C{db: pb}, nil
}
