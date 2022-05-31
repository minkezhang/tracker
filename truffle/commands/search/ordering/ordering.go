package ordering

import (
	"sort"
	"strings"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

type T int
type C func(a *dpb.Entry, b *dpb.Entry) int

const (
	OrderingUnknown T = iota
	OrderingTitles
	OrderingQueued
	OrderingScore
	OrderingCorpus
)

var (
	L = map[string]T{
		"titles": OrderingTitles,
		"queued": OrderingQueued,
		"score":  OrderingScore,
		"corpus": OrderingCorpus,
	}

	F = map[T]C{
		OrderingTitles: func(a *dpb.Entry, b *dpb.Entry) int {
			var at string
			var bt string
			// Use default titles for string comparison, as later
			// titles are less important.
			if len(a.GetTitles()) > 0 {
				at = strings.ToLower(a.GetTitles()[0])
			}
			if len(b.GetTitles()) > 0 {
				bt = strings.ToLower(b.GetTitles()[0])
			}
			return strings.Compare(at, bt)
		},
		OrderingCorpus: func(a *dpb.Entry, b *dpb.Entry) int {
			return strings.Compare(a.GetCorpus().String(), b.GetCorpus().String())
		},
		OrderingQueued: func(a *dpb.Entry, b *dpb.Entry) int {
			if a.GetQueued() && !b.GetQueued() {
				return -1
			}
			if b.GetQueued() && !a.GetQueued() {
				return 1
			}
			return 0
		},
		OrderingScore: func(a *dpb.Entry, b *dpb.Entry) int {
			if a.GetScore() > b.GetScore() {
				return -1
			}
			if a.GetScore() < b.GetScore() {
				return 1
			}
			return 0
		},
	}
)

type S struct {
	priorities []C

	entries []*dpb.Entry
}

func (s *S) Len() int      { return len(s.entries) }
func (s *S) Swap(i, j int) { s.entries[i], s.entries[j] = s.entries[j], s.entries[i] }

func (s *S) Less(i, j int) bool {
	for _, p := range s.priorities {
		switch cmp := p(s.entries[i], s.entries[j]); cmp {
		case -1:
			return true
		case 1:
			return false
		}
	}
	return false
}

func Order(entries []*dpb.Entry, orderings []T) ([]*dpb.Entry, error) {
	var priorities []C
	for _, ordering := range orderings {
		if p := F[ordering]; p != nil {
			priorities = append(priorities, p)
		}
	}

	s := &S{
		priorities: priorities,
		entries:    entries,
	}
	sort.Sort(s)
	return s.entries, nil
}
