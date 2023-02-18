package message

import (
	"bytes"

	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/optional"
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
	AckMessageCounter  optional.Value[uint32]
	SecuredExtensions  []byte
}

// MarshalBinary returns the binary representation of m.
func (m ProtocolMessage) MarshalBinary() ([]byte, error) {
	var exchangeFlags byte

	if m.IsFromInitiator {
		exchangeFlags |= exchangeFlagI
	}

	if m.RequiresAck {
		exchangeFlags |= exchangeFlagR
	}

	if m.ProtocolVendorID != 0 {
		exchangeFlags |= exchangeFlagV
	}

	if m.AckMessageCounter.IsPresent() {
		exchangeFlags |= exchangeFlagA
	}

	if len(m.SecuredExtensions) != 0 {
		exchangeFlags |= exchangeFlagSX
	}

	w := &bytes.Buffer{}

	if err := wire.WriteInt(w, exchangeFlags); err != nil {
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

	if m.AckMessageCounter.IsPresent() {
		if err := wire.WriteInt(w, m.AckMessageCounter.Value()); err != nil {
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

	exchangeFlags, err := wire.ReadInt[byte](r)
	if err != nil {
		return err
	}

	m.IsFromInitiator = exchangeFlags&exchangeFlagI != 0
	m.RequiresAck = exchangeFlags&exchangeFlagR != 0

	if err := wire.AssignInt(r, &m.ProtocolOpCode); err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.ExchangeID); err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.ProtocolID); err != nil {
		return err
	}

	if exchangeFlags&exchangeFlagV == 0 {
		m.ProtocolVendorID = 0
	} else if err := wire.AssignInt(r, &m.ProtocolVendorID); err != nil {
		return err
	}

	if err := wire.AssignOptionalInt(
		r,
		exchangeFlags&exchangeFlagA != 0,
		&m.AckMessageCounter,
	); err != nil {
		return err
	}

	if exchangeFlags&exchangeFlagSX == 0 {
		m.SecuredExtensions = nil
	} else if err := wire.AssignString[uint16](r, &m.SecuredExtensions); err != nil {
		return err
	}

	m.IsFromInitiator = exchangeFlags&exchangeFlagI != 0
	m.RequiresAck = exchangeFlags&exchangeFlagR != 0

	return wire.AssignRemaining(r, &m.ApplicationPayload)
}

const (
	// exchangeFlagI is a bit-mask that isolates the "I" sub-field of the
	// "exchange flags" bit-field. The "I" sub-field is a boolean that indicates
	// whether the message was sent by the initiator of the exchange.
	exchangeFlagI = 0b00000001

	// exchangeFlagA is a bit-mask that isolates the "A" sub-field of the
	// "exchange flags" bit-field. The "A" sub-field is a boolean that indicates
	// whether the message serves as an acknowledgement of a previously received
	// message.
	exchangeFlagA = 0b00000010

	// exchangeFlagR is a bit-mask that isolates the "R" sub-field of the
	// "exchange flags" bit-field. The "R" sub-field is a boolean that indicates
	// whether the sender requires an acknowledgement of the message.
	exchangeFlagR = 0b00000100

	// exchangeFlagSX is a bit-mask that isolates the "SX" sub-field of the
	// "exchange flags" bit-field. The "SX" sub-field is a boolean that
	// indicates whether the message has secured extensions.
	exchangeFlagSX = 0b00001000

	// exchangeFlagV is a bit-mask that isolates the "V" sub-field of the
	// "exchange flags" bit-field. The "V" sub-field is a boolean that indicates
	// whether the message has a "vendor ID" field.
	exchangeFlagV = 0b00010000
)
