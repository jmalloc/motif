package message

import (
	"bytes"

	"github.com/jmalloc/motif/internal/wire"
	"github.com/jmalloc/motif/optional"
)

// Encrypter is an interface for performing encryption operations on unencrypted
// message data.
type Encrypter interface{}

// Decrypter is an interface for performing decryption operations on encrypted
// message data.
type Decrypter interface{}

// Message is a Matter message frame.
type Message struct {
	SessionID              uint16
	IsGroupSession         bool
	SourceNodeID           optional.Value[uint64]
	DestinationNodeID      optional.Value[uint64]
	DestinationGroupID     optional.Value[uint16]
	MessageCounter         uint32
	MessagePayload         []byte
	IsControlMessage       bool
	HasPrivacyEnhancements bool
	MessageExtensions      []byte
}

// MarshalBinary returns the binary representation of m.
func (m Message) MarshalBinary(e Encrypter) ([]byte, error) {
	var messageFlags, securityFlags uint8

	if m.SourceNodeID.IsPresent() {
		messageFlags |= messageFlagS
	}

	if m.DestinationNodeID.IsPresent() {
		messageFlags |= messageFlagDSIZNodeID
	}

	if m.DestinationGroupID.IsPresent() {
		messageFlags |= messageFlagDSIZGroupID

		if m.DestinationNodeID.IsPresent() {
			panic("cannot use both destination node ID and destination group ID")
		}
	}

	if m.HasPrivacyEnhancements {
		securityFlags |= securityFlagP
	}

	if m.IsControlMessage {
		securityFlags |= securityFlagC
	}

	if len(m.MessageExtensions) > 0 {
		securityFlags |= securityFlagMX
	}

	if m.IsGroupSession {
		securityFlags |= sessionTypeGroup
	} else {
		securityFlags |= sessionTypeUnicast
	}

	w := &bytes.Buffer{}

	if err := wire.WriteInt(w, messageFlags); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.SessionID); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, securityFlags); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.MessageCounter); err != nil {
		return nil, err
	}

	if m.SourceNodeID.IsPresent() {
		if err := wire.WriteInt(w, m.SourceNodeID.Value()); err != nil {
			return nil, err
		}
	}

	if m.DestinationNodeID.IsPresent() {
		if err := wire.WriteInt(w, m.DestinationNodeID.Value()); err != nil {
			return nil, err
		}
	} else if m.DestinationGroupID.IsPresent() {
		if err := wire.WriteInt(w, m.DestinationGroupID.Value()); err != nil {
			return nil, err
		}
	}

	if len(m.MessageExtensions) > 0 {
		if err := wire.WriteString[uint16](w, m.MessageExtensions); err != nil {
			return nil, err
		}
	}

	if _, err := w.Write(m.MessagePayload); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
}

// UnmarshalBinary sets m to the value represented by data.
func (m *Message) UnmarshalBinary(d Decrypter, data []byte) error {
	r := bytes.NewReader(data)

	messageFlags, err := wire.ReadInt[uint8](r)
	if err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.SessionID); err != nil {
		return err
	}

	securityFlags, err := wire.ReadInt[uint8](r)
	if err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.MessageCounter); err != nil {
		return err
	}

	if err := wire.AssignOptionalInt(
		r,
		messageFlags&messageFlagS != 0,
		&m.SourceNodeID,
	); err != nil {
		return err
	}

	if err := wire.AssignOptionalInt(
		r,
		messageFlags&messageFlagDSIZNodeID != 0,
		&m.DestinationNodeID,
	); err != nil {
		return err
	}

	if err := wire.AssignOptionalInt(
		r,
		messageFlags&messageFlagDSIZGroupID != 0,
		&m.DestinationGroupID,
	); err != nil {
		return err
	}

	if securityFlags&securityFlagMX == 0 {
		m.MessageExtensions = nil
	} else if err := wire.AssignString[uint16](r, &m.MessageExtensions); err != nil {
		return err
	}

	return wire.AssignRemaining(r, &m.MessagePayload)
}

const (
	// messageFlagS is a bit-mask that isolates the "S" sub-field of the
	// "message flags" bit-field. The "S" sub-field is a boolean that indicates
	// whether the message has a "source node ID" field.
	messageFlagS = 0b000_0_1_00

	// messageFlagDSIZNodeID is a value of the "DSIZ" sub-field of the "message
	// flags" bit-field. It indicates that the message has a "destination node
	// ID" field.
	messageFlagDSIZNodeID = 0b000_0_0_01

	// messageFlagDSIZGroupID is a value of the "DSIZ" sub-field of the "message
	// flags" bit-field. It indicates that the message has a "destination group
	// ID" field.
	messageFlagDSIZGroupID = 0b000_0_0_10
)

const (
	// securityFlagP is a bit-mask that isolates the "P" sub-field of the
	// "security flags" bit-field. The "P" sub-field is a boolean that indicates
	// whether the message uses privacy enhancements.
	securityFlagP = 0b100_000_00

	// securityFlagC is a bit-mask that isolates the "C" sub-field of the
	// "security flags" bit-field. The "C" sub-field is a boolean that indicates
	// whether the message is a control message.
	securityFlagC = 0b010_000_00

	// securityFlagMX is a bit-mask that isolates the "MX" sub-field of the
	// "security flags" bit-field. The "MX" sub-field is a boolean that
	// indicates whether the message has a "message extensions" field.
	securityFlagMX = 0b001_000_00

	// sessionTypeUnicast is a value of the "session type" sub-field of the
	// "security flags" bit-field. It indicates that the message is being sent
	// within a unicast session.
	sessionTypeUnicast = 0b000_000_00

	// sessionTypeGroup is a value of the "session type" sub-field of the
	// "security flags" bit-field. It indicates that the message is being sent
	// within a group session.
	sessionTypeGroup = 0b000_000_01
)
