package entry

import (
	"google.golang.org/protobuf/proto"
)

type Importer interface {
	Unmarshal(b []byte, m proto.Message) error
}

type Exporter interface {
	Marshal(m proto.Message) ([]byte, error)
}
