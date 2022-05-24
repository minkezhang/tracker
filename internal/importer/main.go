package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/minkezhang/tracker/internal/importer/lookup"
	"google.golang.org/protobuf/encoding/prototext"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

var (
	input  = flag.String("input", "/dev/stdin", "input CSV path, e.g. path/to/database.csv")
	output = flag.String("output", "/dev/stdout", "output textproto path, e.g. path/to/database.textproto")
)

type E [11]string

func (r E) Corpus() dpb.Corpus  { return lookup.Corpus[r[8]] }
func (r E) Title() string       { return r[0] }
func (r E) Queued() bool        { return r[2] == "TRUE" }
func (r E) Composers() []string { return r.Writers() }
func (r E) Authors() []string   { return r.Writers() }
func (r E) Studios() []string   { return []string{r[7]} }

func (r E) Directors() []string {
	var directors []string
	for _, v := range strings.Split(r[5], ",") {
		directors = append(directors, strings.Trim(v, " "))
	}
	return directors
}

func (r E) Writers() []string {
	var writers []string
	for _, v := range strings.Split(r[6], ",") {
		writers = append(writers, strings.Trim(v, " "))
	}
	return writers
}

func (r E) Score() float32 {
	if f, err := strconv.ParseFloat(r[3], 32); err != nil {
		return 0
	} else {
		return float32(f)
	}
}

func (r E) Providers() []dpb.Provider {
	var dedupe = map[dpb.Provider]bool{}
	dedupe[lookup.Provider[r[9]]] = true
	dedupe[lookup.Provider[r[10]]] = true

	var providers []dpb.Provider
	for k := range dedupe {
		if k != dpb.Provider_PROVIDER_UNKNOWN {
			providers = append(providers, k)
		}
	}
	return providers
}

func (r E) TrackerBook() *dpb.TrackerBook {
	reg := regexp.MustCompile(`v?(?P<volume>\d+)?c(?P<chapter>\d+)`)
	m := reg.FindStringSubmatch(r[4])
	if m == nil {
		return &dpb.TrackerBook{}
	}
	result := map[string]int32{}
	for i, name := range reg.SubexpNames() {
		if i != 0 && name != "" {
			v, err := strconv.ParseInt(m[i], 10, 32)
			if err != nil {
				v = 0
			}
			result[name] = int32(v)
		}
	}
	return &dpb.TrackerBook{
		Volume:  result["volume"],
		Chapter: result["chapter"],
	}
}

func (r E) TrackerVideo() *dpb.TrackerVideo {
	reg := regexp.MustCompile(`s?(?P<season>\d+)?e(?P<episode>\d+)`)
	m := reg.FindStringSubmatch(r[4])
	if m == nil {
		return &dpb.TrackerVideo{}
	}
	result := map[string]int32{}
	for i, name := range reg.SubexpNames() {
		if i != 0 && name != "" {
			v, err := strconv.ParseInt(m[i], 10, 32)
			if err != nil {
				v = 0
			}
			result[name] = int32(v)
		}
	}
	return &dpb.TrackerVideo{
		Season:  result["season"],
		Episode: result["episode"],
	}
}

// ProtoBuf returns a PB object for the given row.
//
// TODO(minkezhang): Refactor to proto.Unmarshal API.
func (r E) ProtoBuf() *dpb.Entry {
	e := &dpb.Entry{
		Corpus:    r.Corpus(),
		Title:     r.Title(),
		Queued:    r.Queued(),
		Score:     r.Score(),
		Providers: r.Providers(),
	}
	switch e.GetCorpus() {

	case dpb.Corpus_CORPUS_TV:
		fallthrough
	case dpb.Corpus_CORPUS_ANIME:
		e.Tracker = &dpb.Entry_TrackerVideo{
			TrackerVideo: r.TrackerVideo(),
		}
		fallthrough
	case dpb.Corpus_CORPUS_FILM:
		fallthrough
	case dpb.Corpus_CORPUS_ANIME_FILM:
		e.AuxData = &dpb.Entry_AuxDataVideo{
			AuxDataVideo: &dpb.AuxDataVideo{
				Studios:   r.Studios(),
				Directors: r.Directors(),
				Writers:   r.Writers(),
			},
		}

	case dpb.Corpus_CORPUS_GAME:
		e.AuxData = &dpb.Entry_AuxDataGame{
			AuxDataGame: &dpb.AuxDataGame{
				Studios:   r.Studios(),
				Directors: r.Directors(),
				Writers:   r.Writers(),
			},
		}

	case dpb.Corpus_CORPUS_MANGA:
		fallthrough
	case dpb.Corpus_CORPUS_BOOK:
		e.Tracker = &dpb.Entry_TrackerBook{
			TrackerBook: r.TrackerBook(),
		}
		fallthrough
	case dpb.Corpus_CORPUS_SHORT_STORY:
		e.AuxData = &dpb.Entry_AuxDataBook{
			AuxDataBook: &dpb.AuxDataBook{
				Authors: r.Authors(),
			},
		}

	case dpb.Corpus_CORPUS_ALBUM:
		e.AuxData = &dpb.Entry_AuxDataAudio{
			AuxDataAudio: &dpb.AuxDataAudio{
				Composers: r.Composers(),
			},
		}
	default:
	}
	return e
}

func main() {
	flag.Parse()

	fp, err := os.Open(*input)
	if err != nil {
		log.Fatalf("cannot open file %v: %v", *input, err)
	}
	defer fp.Close()

	scanner := csv.NewReader(fp)
	if _, err := scanner.Read(); err != nil {
		log.Fatalf("error while reading CSV file %v: %v", *input, err)
	}

	var entries []*dpb.Entry
	for l, err := scanner.Read(); err == nil; l, err = scanner.Read() {
		r := (*E)(l)
		if e := r.ProtoBuf(); e.Corpus != dpb.Corpus_CORPUS_UNKNOWN {
			entries = append(entries, e)
		}
	}

	if err != nil && err != io.EOF {
		log.Fatalf("error while reading CSV file %v: %v", *input, err)
	}

	db := &dpb.Database{
		Entries: entries,
	}

	data, err := prototext.MarshalOptions{
		Multiline: true,
	}.Marshal(db)
	if err != nil {
		log.Fatalf("error while marshalling proto: %v")
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
