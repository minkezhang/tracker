package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/minkezhang/tracker/database"

	dpb "github.com/minkezhang/tracker/api/go/database"
	entry "github.com/minkezhang/tracker/formats/minkezhang"
)

var (
	input  = flag.String("input", "/dev/stdin", "input CSV path, e.g. path/to/database.csv")
	output = flag.String("output", "/dev/stdout", "output textproto path, e.g. path/to/database.textproto")
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

	var entries []*dpb.Entry
	for scanner.Scan() {
		epb := &dpb.Entry{}
		if err := (entry.E{}).Unmarshal(scanner.Bytes(), epb); err != nil {
			log.Fatalf("error while unmarshalling data: %v", err)
		}
		if epb.GetCorpus() != dpb.Corpus_CORPUS_UNKNOWN {
			entries = append(entries, epb)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error while reading CSV file %v: %v", *input, err)
	}

	data, err := database.New(entries).Marshal()
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
