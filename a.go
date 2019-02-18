package vet

import (
	"bytes"
	"io"
)

func F() {
	_ = bytes.Buffer{}
}

func G() io.Writer {
	return nil
}
