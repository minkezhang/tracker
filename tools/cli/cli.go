package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/minkezhang/tracker/database"
	"github.com/minkezhang/tracker/tools/cli/commands/get"

	dpb "github.com/minkezhang/tracker/api/go/database"
	ce "github.com/minkezhang/tracker/formats/cli"
)

const (
	fn = "data/database.textproto"
)

func main() {
	fp, _ := os.Open(fn)
	defer fp.Close()

	data, _ := ioutil.ReadAll(fp)
	db, _ := database.Unmarshal(data)

	qs := []struct {
		db     *database.DB
		title  string
		corpus dpb.Corpus
	}{
		{
			db:     db,
			title:  "12 Angry Men",
			corpus: dpb.Corpus_CORPUS_FILM,
		},
		{
			db:     db,
			title:  "Bastion",
			corpus: dpb.Corpus_CORPUS_GAME,
		},
		{
			db:     db,
			title:  "Akira",
			corpus: dpb.Corpus_CORPUS_ANIME_FILM,
		},
		{
			db:     db,
			title:  "heart of my own",
			corpus: dpb.Corpus_CORPUS_ALBUM,
		},
		{
			db:     db,
			title:  "sense8",
			corpus: dpb.Corpus_CORPUS_TV,
		},
		{
			db:     db,
			title:  "Chrono Crusade",
			corpus: dpb.Corpus_CORPUS_MANGA,
		},
		{
			db:     db,
			title:  "Sabikui",
			corpus: dpb.Corpus_CORPUS_ANIME,
		},
	}

	for _, q := range qs {
		e, _ := get.O{DB: q.db, Title: q.title, Corpus: q.corpus}.Get()
		d, _ := ce.E{}.Marshal(e.PB)
		fmt.Println(string(d))
	}
}
