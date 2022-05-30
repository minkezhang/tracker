package flagset

import (
	"flag"
	"fmt"
	"testing"
	"unsafe"

	entry "github.com/minkezhang/truffle/formats/cli/x"
)

func TestSetFlags(t *testing.T) {
	e := &entry.E{}

	f := flag.NewFlagSet("x", flag.ContinueOnError)

	(*Corpus)(unsafe.Pointer(e)).SetFlags(f)
	(*ID)(unsafe.Pointer(e)).SetFlags(f)
	(*Title)(unsafe.Pointer(e)).SetFlags(f)

	if err := f.Parse([]string{"--corpus=x", "--id=f", "--title=moooo"}); err != nil {
		t.Errorf("Parse() = %v, want = nil", err)
	}

	panic(fmt.Sprintf("corpus == %v, id == %v, titles == %v", e.Corpus, e.ID, e.Titles))
}
