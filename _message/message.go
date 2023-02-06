package message

import (
	"errors"
	"fmt"
	"io"

	"github.com/jmalloc/motif/internal/crypto"
	"github.com/jmalloc/motif/internal/wire"
)

// MatterMessage is a Matter message.
//
// Matter v1.0 Core § 4.4
type MatterMessage struct {
	Flags              uint8
	SessionID          uint16
	SecurityFlags      SecurityFlags
	MessageCounter     uint32
	SourceNodeID       uint64
	DestinationNodeID  uint64
	DestinationGroupID uint16
	Extensions         []byte
	Payload            []byte
	IntegrityCheck     []byte
}

const (
	// FlagsVersionMask is a bit-mask that isolates the "version" sub-field of
	// the Flags bit-field.
	FlagsVersionMask = 0b11110000

	// FlagS is a bit-mask that isolates the "S" sub-field of the Flags
	// bit-field. The "S" sub-field is a boolean that indicates whether the
	// message has a "source node ID" field.
	FlagS = 0b100

	// FlagsDSIZMask is a bit-mask that isolates the "DSIZ" sub-field of the
	// Flags bit-field. The "DSIZ" sub-field is an enumeration that indicates
	// the size and meaning of the message's "destination" field.
	FlagsDSIZMask = 0b11
)

// Version is a 4-bit field that indicates the version of a message's format.
type Version uint8

const (
	// Version1 refers to version 1.0 of the Matter message format.
	Version1 Version = 0b00000000
)

// DestinationFieldType is an enumeration that indicates the size and meaning of
// the message's "destination" field.
type DestinationFieldType uint8

const (
	// DestinationFieldNone indicates that the message has no "destination"
	// field.
	DestinationFieldNone DestinationFieldType = 0b00

	// DestinationFieldNodeID indicates that the message's "destination"
	// field is a 64-bit node ID.
	DestinationFieldNodeID DestinationFieldType = 0b01

	// DestinationFieldGroupID indicates that the message's "destination"
	// field is a 16-bit group ID.
	DestinationFieldGroupID DestinationFieldType = 0b10
)

// Version returns the version of the message format.
func (m MatterMessage) Version() Version {
	return Version(m.Flags & FlagsVersionMask)
}

// HasSourceField returns true if the message has a "source node ID" field.
func (m MatterMessage) HasSourceField() bool {
	return m.Flags&FlagS != 0
}

// DestinationFieldType returns the size and meaning of the message's
// "destination" field.
func (m MatterMessage) DestinationFieldType() DestinationFieldType {
	return DestinationFieldType(m.Flags & FlagsDSIZMask)
}

// IsUnsecured returns true if the message is unsecured.
func (m MatterMessage) IsUnsecured() bool {
	return m.SecurityFlags.SessionType() == SessionTypeUnicast &&
		m.SessionID == 0
}

// Unmarshal reads a Matter message from r into *m.
func (m *MatterMessage) Unmarshal(r io.Reader) error {
	if err := wire.UnmarshalLittleEndian(r, &m.Flags); err != nil {
		return err
	}

	if m.Version() != Version1 {
		return fmt.Errorf(
			"unsupported message version (%x)",
			m.Version(),
		)
	}

	if err := wire.UnmarshalLittleEndian(r, &m.SessionID); err != nil {
		return err
	}

	if err := wire.UnmarshalLittleEndian(r, &m.SecurityFlags); err != nil {
		return err
	}

	if err := wire.UnmarshalLittleEndian(r, &m.MessageCounter); err != nil {
		return err
	}

	if m.HasSourceField() {
		if err := wire.UnmarshalLittleEndian(r, &m.SourceNodeID); err != nil {
			return err
		}
	}

	switch m.DestinationFieldType() {
	case DestinationFieldNone:
	case DestinationFieldNodeID:
		if err := wire.UnmarshalLittleEndian(r, &m.DestinationNodeID); err != nil {
			return err
		}
	case DestinationFieldGroupID:
		if err := wire.UnmarshalLittleEndian(r, &m.DestinationGroupID); err != nil {
			return err
		}
	default:
		return fmt.Errorf(
			"unsupported DSIZ flag value (%x)",
			m.DestinationFieldType(),
		)
	}

	if m.SecurityFlags.HasExtensions() {
		if err := wire.UnmarshalOctetString[uint16](r, &m.Extensions); err != nil {
			return err
		}
	}

	payload, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if m.IsUnsecured() {
		m.Payload = payload
		return nil
	}

	if len(payload) < crypto.AEADMICLengthBytes {
		return errors.New("payload is shorter than integrity check")
	}

	n := len(payload) - crypto.AEADMICLengthBytes
	m.Payload = payload[:n]
	m.IntegrityCheck = payload[n:]

	return nil
}
