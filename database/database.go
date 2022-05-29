package database

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/minkezhang/truffle/api/go/database/utils"
	"github.com/minkezhang/truffle/api/go/database/validator"
	"github.com/minkezhang/truffle/database/ids"
	"github.com/minkezhang/truffle/database/search"
	"github.com/minkezhang/truffle/database/search/mal"
	"github.com/minkezhang/truffle/database/search/truffle"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type DB struct {
	db *dpb.Database
}

func New(epbs []*dpb.Entry) *DB {
	db := &DB{
		db: &dpb.Database{
			Entries: map[string]*dpb.Entry{},
		},
	}
	for _, epb := range epbs {
		if _, err := db.Add(epb); err != nil {
			panic(fmt.Sprintf("could not create database: %v", err))
		}
	}

	return db
}

func (db *DB) Add(epb *dpb.Entry) (*dpb.Entry, error) {
	epb = proto.Clone(epb).(*dpb.Entry)
	if err := validator.Validate(epb); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	eid := ids.New()
	for ; db.db.GetEntries()[eid] != nil; eid = ids.New() {
	}

	epb.Id = &dpb.LinkedID{
		Id:  eid,
		Api: dpb.API_API_TRUFFLE,
	}
	etag, err := ETag(epb)
	if err != nil {
		return nil, err
	}
	epb.Etag = etag

	db.db.GetEntries()[eid] = epb
	return proto.Clone(epb).(*dpb.Entry), nil
}

func (db *DB) Get(id string) (*dpb.Entry, error) {
	epb, ok := db.db.GetEntries()[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "cannot find entry with id %v", id)
	}

	return proto.Clone(epb).(*dpb.Entry), nil
}

func (db *DB) Put(epb *dpb.Entry) (*dpb.Entry, error) {
	epb = proto.Clone(epb).(*dpb.Entry)
	if err := validator.Validate(epb); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	fpb, ok := db.db.GetEntries()[epb.GetId().GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "cannot find entry with id %v", epb.GetId())
	}
	if !reflect.DeepEqual(epb.GetEtag(), fpb.GetEtag()) {
		return nil, status.Errorf(codes.InvalidArgument, "cannot update entry with mismatching ETag values: %s != %s", epb.GetEtag(), fpb.GetEtag())
	}
	etag, err := ETag(epb)
	if err != nil {
		return nil, err
	}
	epb.Etag = etag

	db.db.GetEntries()[epb.GetId().GetId()] = epb
	return epb, nil
}

func (db *DB) Delete(id string) (*dpb.Entry, error) {
	epb := db.db.GetEntries()[id]
	delete(db.db.GetEntries(), id)
	return proto.Clone(epb).(*dpb.Entry), nil
}

type O struct {
	Context context.Context

	Title  string
	Corpus dpb.Corpus

	APIs []dpb.API

	// MAL contains MAL-specific API options.
	MAL mal.O
}

func (db *DB) Search(opts O) ([]*dpb.Entry, error) {
	s := map[dpb.API]search.S{
		dpb.API_API_TRUFFLE: truffle.S{
			DB:     db.db,
			Title:  opts.Title,
			Corpus: opts.Corpus,
		},
		dpb.API_API_MAL: mal.New(opts.Title, opts.Corpus, opts.MAL.Cutoff),
	}

	apis := map[dpb.API]bool{}
	for _, api := range opts.APIs {
		apis[api] = true
	}

	duplicates := map[string]bool{}
	for _, epb := range db.db.GetEntries() {
		for _, id := range epb.GetLinkedIds() {
			duplicates[utils.ID(id)] = true
		}
	}

	var candidates []*dpb.Entry
	for api, _ := range apis {
		if sf, ok := s[api]; ok {
			cs, err := sf.Search(opts.Context)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "error while executing search operation: %v", err)
			}
			candidates = append(candidates, cs...)
		}
	}

	var unique []*dpb.Entry
	for _, epb := range candidates {
		// Skip reporting entries which are already accounted for in the
		// user DB (if we are returning user DB results).
		if apis[dpb.API_API_TRUFFLE] && duplicates[utils.ID(epb.GetId())] {
			continue
		}
		unique = append(unique, epb)
	}
	return unique, nil
}

func ETag(epb *dpb.Entry) ([]byte, error) {
	epb = proto.Clone(epb).(*dpb.Entry)

	epb.Id = nil
	epb.Etag = nil

	data, err := prototext.Marshal(epb)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot marshal input proto: %v", err)
	}

	s := md5.New()
	io.WriteString(s, string(data))

	// b64 string of the etag is easier to read.
	return []byte(
		base64.URLEncoding.EncodeToString(
			s.Sum(nil),
		),
	), nil
}

func Marshal(db *DB) ([]byte, error) {
	data, err := prototext.MarshalOptions{
		Multiline: true,
	}.Marshal(db.db)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot marshal DB entry: %v", err)
	}
	return data, nil
}

func Unmarshal(data []byte) (*DB, error) {
	pb := &dpb.Database{}
	if err := prototext.Unmarshal(data, pb); err != nil {
		return nil, err
	}
	for eid, epb := range pb.GetEntries() {
		epb.Id = &dpb.LinkedID{
			Id:  eid,
			Api: dpb.API_API_TRUFFLE,
		}
		etag, err := ETag(epb)
		if err != nil {
			return nil, err
		}
		epb.Etag = etag
	}

	if pb.GetEntries() == nil {
		pb.Entries = map[string]*dpb.Entry{}
	}
	return &DB{db: pb}, nil
}
