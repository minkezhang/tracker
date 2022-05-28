package entry

import (
	"bytes"
	"encoding/csv"

	"github.com/minkezhang/truffle/formats/minkezhang/cache"
	"google.golang.org/protobuf/proto"
)

type E struct {
	data []byte
}

func New(data []byte) *E { return &E{data: data} }
func (e E) Load() (proto.Message, error) {
	l, err := csv.NewReader(bytes.NewReader(e.data)).Read()
	if err != nil {
		return nil, err
	}

	r := (*cache.E)(l)

	return r.ProtoBuf(), nil
}
