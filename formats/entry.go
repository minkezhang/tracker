package entry

import (
	"google.golang.org/protobuf/proto"
)

type Importer interface {
	Load() (proto.Message, error)
}

type Exporter interface {
	Dump(m proto.Message) error
}
