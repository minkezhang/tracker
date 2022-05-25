package entry

import (
	"google.golang.org/protobuf/proto"
)

type E interface {
	// TODO(minkezhang): Implement the Marshal function.
	//
	// Marshal(m proto.Message) ([]byte, error)

	Unmarshal(b []byte, m proto.Message) error
}
