package wire

import (
	"io"
)

// ReadRemaining reads all remaining bytes from r.
//
// It differs from io.ReadAll in that it returns nil if the resulting slice
// would be empty.
func ReadRemaining(r io.Reader) ([]byte, error) {
	var data []byte
	return data, AssignRemaining(r, &data)
}

// AssignRemaining reads all remaining bytes from r and assigns them to *data.
//
// It always assigns nil if there is no data (never an empty slice).
func AssignRemaining(r io.Reader, data *[]byte) error {
	if r, ok := r.(interface{ Len() int }); ok {
		if r.Len() == 0 {
			*data = nil
			return nil
		}
	}

	d, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if len(d) == 0 {
		*data = nil
	} else {
		*data = d
	}

	return nil
}
