// Package main runs the importer on a CSV file.
//
// The format of the CSV file is explicitly described in
// /formats/minkezhang/columns/
//
// Example
//
// go run github.com/minkezhang/truffle/tools/importer \
//   --input=/path/to/input.csv
//   --output=/path/to/output.textproto
package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"

	"github.com/minkezhang/truffle/client/truffle"

	dpb "github.com/minkezhang/truffle/api/go/database"
	entry "github.com/minkezhang/truffle/formats/minkezhang"
)

var (
	input  = flag.String("input", "/dev/stdin", "input CSV path, e.g. /path/to/database.csv")
	output = flag.String("output", "/dev/stdout", "output textproto path, e.g. /path/to/database.textproto")
)

func main() {
	flag.Parse()

	fp, err := os.Open(*input)
	if err != nil {
		log.Fatalf("cannot open file %v: %v", *input, err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Scan()

	db := truffle.New(&dpb.Database{})
	for scanner.Scan() {
		m, err := entry.New(scanner.Bytes()).Load()
		if err != nil {
			log.Fatalf("error while loading data: %v", err)
		}
		epb := m.(*dpb.Entry)
		if epb.GetCorpus() != dpb.Corpus_CORPUS_UNKNOWN {
			db.Add(context.Background(), epb)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading CSV file %v: %v", *input, err)
	}

	data, err := truffle.Marshal(db)
	if err != nil {
		log.Fatalf("error while marshalling proto: %v", err)
	}

	w, err := os.Create(*output)
	if err != nil {
		log.Fatalf("cannot open output file %v: %v", *output, err)
	}
	defer w.Close()
	if _, err := w.Write(data); err != nil {
		log.Fatalf("could not write to output file %v: %v", *output, err)
	}
}
