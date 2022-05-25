package database

import (
	dpb "github.com/minkezhang/tracker/api/go/database"
)

type DB struct {
}

func (db *DB) AddEntry(epb *dpb.Entry) error            { return nil }
func (db *DB) GetEntry(id string) error                 { return nil }
func (db *DB) PutEntry(id string, epb *dpb.Entry) error { return nil }
func (db *DB) DeleteEntry(id string) error              { return nil }
func (db *DB) Search(title string) []*dpb.Entry         { return nil }
