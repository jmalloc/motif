package tlv

import (
	"bytes"
	"context"
)

// Marshal returns the TLV encoded form of the given element.
func Marshal(e Element) ([]byte, error) {
	var m marshaler

	ctx := visitContext{
		Context: context.Background(),
	}

	if err := e.AcceptVisitor(ctx, m); err != nil {
		return nil, err
	}

	return m.data.Bytes(), nil
}

type marshaler struct {
	data    bytes.Buffer
	control byte
	buffer  []byte
}
