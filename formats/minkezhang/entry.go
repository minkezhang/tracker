package entry

import (
	"bytes"
	"encoding/csv"

	"github.com/minkezhang/tracker/formats/minkezhang/cache"
	"google.golang.org/protobuf/proto"
)

type E struct{}

func (e E) Unmarshal(b []byte, m proto.Message) error {
	l, err := csv.NewReader(bytes.NewReader(b)).Read()
	if err != nil {
		return err
	}

	r := (*cache.E)(l)

	proto.Merge(m, r.ProtoBuf())
	return nil
}
