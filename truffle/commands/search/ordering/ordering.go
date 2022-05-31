package ordering

import (
	"flag"
	"sort"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	dpb "github.com/minkezhang/truffle/api/go/database"
	cf "github.com/minkezhang/truffle/truffle/flag"
)

type T int

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
)

type O struct {
	Orderings cf.MultiString
}

func (o *O) SetFlags(f *flag.FlagSet) {
	f.Var(&o.Orderings, "orderings", "list of fields to order by, e.g. \"title\"")
}

func Validate(o O) error {
	for _, ordering := range o.Orderings {
		if L[strings.ToLower(ordering)] == OrderingUnknown {
			return status.Errorf(codes.InvalidArgument, "invalid ordering field specified %v", ordering)
		}
	}
	return nil
}

var (
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

type C func(a *dpb.Entry, b *dpb.Entry) int
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

func Order(entries []*dpb.Entry, o O) ([]*dpb.Entry, error) {
	if err := Validate(o); err != nil {
		return nil, err
	}

	var priorities []C
	for _, ordering := range o.Orderings {
		if p := F[L[ordering]]; p != nil {
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
