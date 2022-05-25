package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/minkezhang/tracker/database"
	"github.com/minkezhang/tracker/tools/cli/commands/get"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

const (
	format = "Title: %v\nCorpus: %v\nScore: %v\n"
)

func main() {
	fp, _ := os.Open("data/database.textproto")
	defer fp.Close()

	data, _ := ioutil.ReadAll(fp)
	db, _ := database.Unmarshal(data)

	e, _ := get.O{DB: db, Title: "12 Angry Men", Corpus: dpb.Corpus_CORPUS_FILM}.Get()

	fmt.Printf(format, e.PB.GetTitles()[0], e.PB.GetCorpus(), e.PB.GetScore())
}
