package common

import (
	"io"
)

type O struct {
	Output io.Writer
	Error  io.Writer
}
