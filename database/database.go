package database

import (
	"fmt"

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

func (db *DB) Marshal() ([]byte, error) {
	data, err := prototext.MarshalOptions{
		Multiline: true,
	}.Marshal(db.db)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot marshal DB entry: %v", err)
	}
	return data, nil
}

func (db *DB) AddEntry(epb *dpb.Entry) error {
	if err := validator.Validate(epb); err != nil {
		return status.Errorf(codes.InvalidArgument, "cannot add invalid entry: %v", err)
	}

	eid := ids.New()
	for ; db.db.GetEntries()[eid] != nil; eid = ids.New() {
	}
	db.db.GetEntries()[eid] = epb
	return nil
}

func (db *DB) ETag(id string) (string, error)                        { return "", nil }
func (db *DB) GetEntry(id string) (string, error)                    { return "", nil }
func (db *DB) PutEntry(id string, epb *dpb.Entry, etag string) error { return nil }
func (db *DB) DeleteEntry(id string) error                           { return nil }
func (db *DB) Search(title string) []*dpb.Entry                      { return nil }
