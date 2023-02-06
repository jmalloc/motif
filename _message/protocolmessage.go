package message

import (
	"io"

	"github.com/jmalloc/motif/internal/wire"
)

// ProtocolMessage is the payload of a protocol message.
type ProtocolMessage struct {
	ExchangeFlags      uint8
	ProtocolOpCode     uint8
	ExchangeID         uint16
	ProtocolID         uint16
	ProtocolVendorID   uint16
	AckMessageCounter  uint32
	SecuredExtensions  []byte
	ApplicationPayload []byte
}

const (
	// ExchangeFlagI is a bit-mask that isolates the "I" sub-field of the
	// "exchange flags" bit-field. The "I" sub-field is a boolean that indicates
	// whether the message was sent by the initiator of the exchange.
	ExchangeFlagI = 0b00000001

	// ExchangeFlagA is a bit-mask that isolates the "A" sub-field of the
	// "exchange flags" bit-field. The "A" sub-field is a boolean that indicates
	// whether the message serves as an acknowledgement of a previously received
	// message.
	ExchangeFlagA = 0b00000010

	// ExchangeFlagR is a bit-mask that isolates the "R" sub-field of the
	// "exchange flags" bit-field. The "R" sub-field is a boolean that indicates
	// whether the sender requires an acknowledgement of the message.
	ExchangeFlagR = 0b00000100

	// ExchangeFlagSX is a bit-mask that isolates the "SX" sub-field of the
	// "exchange flags" bit-field. The "SX" sub-field is a boolean that
	// indicates whether the message has secured extensions.
	ExchangeFlagSX = 0b00001000

	// ExchangeFlagV is a bit-mask that isolates the "V" sub-field of the
	// "exchange flags" bit-field. The "V" sub-field is a boolean that indicates
	// whether the message has a "vendor ID" field.
	ExchangeFlagV = 0b00010000
)

// IsFromInitiator returns true if the message was sent by the initiator of the
// exchange.
func (m ProtocolMessage) IsFromInitiator() bool {
	return m.ExchangeFlags&ExchangeFlagI != 0
}

// IsAck returns true if the message serves as an acknowledgement of a previously
// received message.
func (m ProtocolMessage) IsAck() bool {
	return m.ExchangeFlags&ExchangeFlagA != 0
}

// RequiresAck returns true if the sender requires an acknowledgement of the
// message.
func (m ProtocolMessage) RequiresAck() bool {
	return m.ExchangeFlags&ExchangeFlagR != 0
}

// HasSecuredExtensions returns true if the message has secured extensions.
func (m ProtocolMessage) HasSecuredExtensions() bool {
	return m.ExchangeFlags&ExchangeFlagSX != 0
}

// HasVendorID returns true if the message has a "vendor ID" field.
func (m ProtocolMessage) HasVendorID() bool {
	return m.ExchangeFlags&ExchangeFlagV != 0
}

// Unmarshal reads a protocol message from r into *m.
func (m *ProtocolMessage) Unmarshal(r io.Reader) error {
	if err := wire.UnmarshalLittleEndian(r, &m.ExchangeFlags); err != nil {
		return err
	}

	if err := wire.UnmarshalLittleEndian(r, &m.ProtocolOpCode); err != nil {
		return err
	}

	if err := wire.UnmarshalLittleEndian(r, &m.ExchangeID); err != nil {
		return err
	}

	if err := wire.UnmarshalLittleEndian(r, &m.ProtocolID); err != nil {
		return err
	}

	if m.HasVendorID() {
		if err := wire.UnmarshalLittleEndian(r, &m.ProtocolVendorID); err != nil {
			return err
		}
	}

	if m.IsAck() {
		if err := wire.UnmarshalLittleEndian(r, &m.AckMessageCounter); err != nil {
			return err
		}
	}

	if m.HasSecuredExtensions() {
		if err := wire.UnmarshalOctetString[uint16](r, &m.SecuredExtensions); err != nil {
			return err
		}
	}

	payload, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	m.ApplicationPayload = payload

	return nil
}
