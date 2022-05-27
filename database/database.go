package database

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/minkezhang/tracker/api/go/database/validator"
	"github.com/minkezhang/tracker/database/ids"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/prototext"

	dpb "github.com/minkezhang/tracker/api/go/database"
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
		if err := db.AddEntry(epb); err != nil {
			panic(fmt.Sprintf("could not create database: %v", err))
		}
	}

	return db
}

func (db *DB) AddEntry(epb *dpb.Entry) error {
	if err := validator.Validate(epb); err != nil {
		return status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	eid := ids.New()
	for ; db.db.GetEntries()[eid] != nil; eid = ids.New() {
	}

	epb.Id = eid
	etag, err := ETag(epb)
	if err != nil {
		return err
	}
	epb.Etag = etag

	db.db.GetEntries()[eid] = epb
	return nil
}

func (db *DB) GetEntry(id string) (*dpb.Entry, error) {
	epb, ok := db.db.GetEntries()[id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "cannot find entry with id %v", id)
	}
	return epb, nil
}

func (db *DB) PutEntry(epb *dpb.Entry) error {
	if err := validator.Validate(epb); err != nil {
		return status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	fpb, ok := db.db.GetEntries()[epb.GetId()]
	if !ok {
		return status.Errorf(codes.NotFound, "cannot find entry with id %v", epb.GetId())
	}
	if !reflect.DeepEqual(epb.GetEtag(), fpb.GetEtag()) {
		return status.Errorf(codes.InvalidArgument, "cannot update entry with mismatching ETag values: %v != %v", epb.GetEtag(), fpb.GetEtag())
	}
	etag, err := ETag(epb)
	if err != nil {
		return err
	}
	epb.Etag = etag

	db.db.GetEntries()[epb.GetId()] = epb
	return nil
}

func (db *DB) DeleteEntry(id string) error {
	delete(db.db.GetEntries(), id)
	return nil
}

type O struct {
	Title  string
	Corpus dpb.Corpus
}

func (db *DB) Search(opts O) []*dpb.Entry {
	var candidates []*dpb.Entry
	for _, epb := range db.db.GetEntries() {
		for _, t := range epb.GetTitles() {
			if strings.Contains(strings.ToLower(t), strings.ToLower(opts.Title)) && ((epb.GetCorpus() == opts.Corpus) || (epb.GetCorpus() == dpb.Corpus_CORPUS_UNKNOWN) || (opts.Corpus == dpb.Corpus_CORPUS_UNKNOWN)) {
				candidates = append(candidates, epb)
				continue
			}
		}
	}
	return candidates
}

func ETag(epb *dpb.Entry) ([]byte, error) {
	fpb := proto.Clone(epb).(*dpb.Entry)
	fpb.Id = ""
	fpb.Etag = nil

	data, err := prototext.Marshal(fpb)
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
		epb.Id = eid
		etag, err := ETag(epb)
		if err != nil {
			return nil, err
		}
		epb.Etag = etag
	}
	return &DB{db: pb}, nil
}
