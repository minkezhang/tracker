package entry

import (
	"google.golang.org/protobuf/encoding/prototext"

	dpb "github.com/minkezhang/truffle/api/go/database"
)

func Format(epb *dpb.Entry) ([]byte, error) {
	return prototext.MarshalOptions{Multiline: true}.Marshal(epb)
}
