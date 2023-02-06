package tlv

import (
	"bytes"
	"context"
)

// Marshal returns the binary representation of e.
func Marshal(e Element) ([]byte, error) {
	data := &bytes.Buffer{}

	err := e.AcceptVisitor(
		context.Background(),
		&marshaler{data},
	)

	return data.Bytes(), err
}

type marshaler struct {
	Data *bytes.Buffer
}
