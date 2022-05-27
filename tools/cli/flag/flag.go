package flag

import (
	"strings"
)

type MultiString []string

func (f *MultiString) String() string { return strings.Join(*f, ", ") }

func (f *MultiString) Set(v string) error {
	*f = append(*f, v)
	return nil
}
