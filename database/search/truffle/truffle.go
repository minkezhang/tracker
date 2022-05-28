package truffle

import (
	"strings"

	dpb "github.com/minkezhang/tracker/api/go/database"
)

type S struct {
	DB *dpb.Database

	Title  string
	Corpus dpb.Corpus
}

func (s S) Search() ([]*dpb.Entry, error) {
	var candidates []*dpb.Entry
	for _, epb := range s.DB.GetEntries() {
		for _, t := range epb.GetTitles() {
			m := map[dpb.Corpus]bool{
				epb.GetCorpus():           true,
				dpb.Corpus_CORPUS_UNKNOWN: true,
			}
			if strings.Contains(strings.ToLower(t), strings.ToLower(s.Title)) && (m[s.Corpus] || (epb.GetCorpus() == dpb.Corpus_CORPUS_UNKNOWN)) {
				candidates = append(candidates, epb)
				continue
			}
		}
	}
	return candidates, nil
}
