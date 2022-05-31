package flagset

import (
	"flag"
	"testing"
	"unsafe"

	"github.com/google/go-cmp/cmp"
	"github.com/minkezhang/truffle/truffle/flag/entry"
	"google.golang.org/protobuf/testing/protocmp"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

func TestSetFlags(t *testing.T) {
	f := flag.NewFlagSet("", flag.ContinueOnError)
	got := &entry.E{}

	(*ID)(unsafe.Pointer(got)).SetFlags(f)
	(*Title)(unsafe.Pointer(got)).SetFlags(f)
	(*Corpus)(unsafe.Pointer(got)).SetFlags(f)
	(*Body)(unsafe.Pointer(got)).SetFlags(f)

	if err := f.Parse(
		[]string{
			"--corpus=manga",
			"--id=123",
			"--title=ABC",
			"--provider=crunchyroll",
		}); err != nil {
		t.Errorf("Parse() = %v, want = nil", err)
	}

	want := &entry.E{
		ID: &dpb.LinkedID{
			Id:  "123",
			Api: dpb.API_API_TRUFFLE,
		},
		Corpus: dpb.Corpus_CORPUS_MANGA,
		Titles: []string{"ABC"},
		Providers: []dpb.Provider{
			dpb.Provider_PROVIDER_CRUNCHYROLL,
		},
	}

	if diff := cmp.Diff(want, got, protocmp.Transform()); diff != "" {
		t.Errorf("SetFlags() mismatch (-want +got):\n%v", diff)
	}
}
