package message

import (
	"bytes"
	"io"

	"github.com/jmalloc/motif/internal/wire"
)

// ProtocolMessage is the payload of a protocol message.
type ProtocolMessage struct {
	ExchangeID         uint16
	ProtocolVendorID   uint16
	ProtocolID         uint16
	ProtocolOpCode     uint8
	ApplicationPayload []byte
	IsFromInitiator    bool
	RequiresAck        bool
	AckMessageCounter  Optional[uint32]
	SecuredExtensions  []byte
}

// MarshalBinary returns the binary representation of m.
func (m ProtocolMessage) MarshalBinary() ([]byte, error) {
	w := &bytes.Buffer{}

	var flags byte

	if m.IsFromInitiator {
		flags |= exchangeFlagI
	}

	if m.RequiresAck {
		flags |= exchangeFlagR
	}

	if m.ProtocolVendorID != 0 {
		flags |= exchangeFlagV
	}

	if m.AckMessageCounter.ok {
		flags |= exchangeFlagA
	}

	if len(m.SecuredExtensions) != 0 {
		flags |= exchangeFlagSX
	}

	if err := wire.WriteInt(w, flags); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.ProtocolOpCode); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.ExchangeID); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.ProtocolID); err != nil {
		return nil, err
	}

	if m.ProtocolVendorID != 0 {
		if err := wire.WriteInt(w, m.ProtocolVendorID); err != nil {
			return nil, err
		}
	}

	if m.AckMessageCounter.ok {
		if err := wire.WriteInt(w, m.AckMessageCounter.value); err != nil {
			return nil, err
		}
	}

	if n := len(m.SecuredExtensions); n != 0 {
		if err := wire.WriteString[uint16](w, m.SecuredExtensions); err != nil {
			return nil, err
		}
	}

	if _, err := w.Write(m.ApplicationPayload); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

// UnmarshalBinary sets m to the value represented by data.
func (m *ProtocolMessage) UnmarshalBinary(data []byte) error {
	r := bytes.NewReader(data)

	flags, err := wire.ReadInt[byte](r)
	if err != nil {
		return err
	}

	m.IsFromInitiator = flags&exchangeFlagI != 0
	m.RequiresAck = flags&exchangeFlagR != 0

	if err := wire.AssignInt(r, &m.ProtocolOpCode); err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.ExchangeID); err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.ProtocolID); err != nil {
		return err
	}

	if flags&exchangeFlagV != 0 {
		if err := wire.AssignInt(r, &m.ProtocolVendorID); err != nil {
			return err
		}
	}

	if flags&exchangeFlagA != 0 {
		c, err := wire.ReadInt[uint32](r)
		if err != nil {
			return err
		}

		m.AckMessageCounter = With(c)
	}

	if flags&exchangeFlagSX != 0 {
		if err := wire.AssignString[uint16](r, &m.SecuredExtensions); err != nil {
			return err
		}
	}

	m.IsFromInitiator = flags&exchangeFlagI != 0
	m.RequiresAck = flags&exchangeFlagR != 0

	if r.Len() == 0 {
		return nil
	}

	m.ApplicationPayload, err = io.ReadAll(r)

	return err
}

const (
	// ExchangeFlagI is a bit-mask that isolates the "I" sub-field of the
	// "exchange flags" bit-field. The "I" sub-field is a boolean that indicates
	// whether the message was sent by the initiator of the exchange.
	exchangeFlagI = 0b00000001

	// ExchangeFlagA is a bit-mask that isolates the "A" sub-field of the
	// "exchange flags" bit-field. The "A" sub-field is a boolean that indicates
	// whether the message serves as an acknowledgement of a previously received
	// message.
	exchangeFlagA = 0b00000010

	// ExchangeFlagR is a bit-mask that isolates the "R" sub-field of the
	// "exchange flags" bit-field. The "R" sub-field is a boolean that indicates
	// whether the sender requires an acknowledgement of the message.
	exchangeFlagR = 0b00000100

	// ExchangeFlagSX is a bit-mask that isolates the "SX" sub-field of the
	// "exchange flags" bit-field. The "SX" sub-field is a boolean that
	// indicates whether the message has secured extensions.
	exchangeFlagSX = 0b00001000

	// ExchangeFlagV is a bit-mask that isolates the "V" sub-field of the
	// "exchange flags" bit-field. The "V" sub-field is a boolean that indicates
	// whether the message has a "vendor ID" field.
	exchangeFlagV = 0b00010000
)
