package entry

import (
	entry "github.com/minkezhang/tracker/formats"
)

var (
	_ entry.Importer = &Body{}
	_ entry.Importer = &Title{}
	_ entry.Importer = &Titles{}
	_ entry.Importer = &Corpus{}
	_ entry.Importer = &ID{}
)
