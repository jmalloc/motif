package message

import (
	"bytes"
	"io"

	"github.com/jmalloc/motif/internal/crypto/mappingv1"
	"github.com/jmalloc/motif/internal/wire"
)

// KeyProvider is a function that locates the encryption/decryption key to use
// based on the session ID.
type KeyProvider func(sessionID uint16) ([]byte, error)

// Message is a Matter message frame.
type Message struct {
	SessionID              uint16
	IsGroupSession         bool
	SourceNodeID           uint64
	DestinationNodeID      uint64
	DestinationGroupID     uint16
	MessageCounter         uint32
	MessagePayload         []byte
	IsControlMessage       bool
	HasPrivacyEnhancements bool
	MessageExtensions      []byte
}

// MarshalBinary returns the binary representation of m.
func (m Message) MarshalBinary(p KeyProvider) ([]byte, error) {
	header, err := m.header()
	if err != nil {
		return nil, err
	}

	payload, err := m.payload()
	if err != nil {
		return nil, err
	}

	if m.SessionID != 0 {
		key, err := p(m.SessionID)
		if err != nil {
			return nil, err
		}

		nonce, err := m.nonce()
		if err != nil {
			return nil, err
		}

		payload, err = mappingv1.Encrypt(key, payload, header, nonce)
		if err != nil {
			return nil, err
		}
	}

	w := &bytes.Buffer{}
	w.Write(header)
	w.Write(payload)

	return w.Bytes(), nil
}

// UnmarshalBinary sets m to the value represented by data.
func (m *Message) UnmarshalBinary(p KeyProvider, data []byte) error {
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

	m.IsGroupSession = securityFlags&sessionTypeGroup != 0
	m.IsControlMessage = securityFlags&securityFlagC != 0

	if err := wire.AssignInt(r, &m.MessageCounter); err != nil {
		return err
	}

	if messageFlags&messageFlagS == 0 {
		m.SourceNodeID = 0
	} else if err := wire.AssignInt(r, &m.SourceNodeID); err != nil {
		return err
	}

	if messageFlags&messageFlagDSIZNodeID == 0 {
		m.DestinationNodeID = 0
	} else if wire.AssignInt(r, &m.DestinationNodeID); err != nil {
		return err
	}

	if messageFlags&messageFlagDSIZGroupID == 0 {
		m.DestinationGroupID = 0
	} else if err := wire.AssignInt(r, &m.DestinationGroupID); err != nil {
		return err
	}

	if m.SessionID != 0 {
		offset, err := r.Seek(0, io.SeekCurrent) // tell
		if err != nil {
			return err
		}
		header := data[:offset]

		key, err := p(m.SessionID)
		if err != nil {
			return err
		}

		ciphertext, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		nonce, err := m.nonce()
		if err != nil {
			return err
		}

		plaintext, err := mappingv1.Decrypt(key, ciphertext, header, nonce)
		if err != nil {
			return err
		}

		r.Reset(plaintext)
	}

	if securityFlags&securityFlagMX == 0 {
		m.MessageExtensions = nil
	} else if err := wire.AssignString[uint16](r, &m.MessageExtensions); err != nil {
		return err
	}

	return wire.AssignRemaining(r, &m.MessagePayload)
}

func (m Message) messageFlags() byte {
	var f byte

	if m.SourceNodeID != 0 {
		f |= messageFlagS
	}

	if m.DestinationNodeID != 0 {
		f |= messageFlagDSIZNodeID
	}

	if m.DestinationGroupID != 0 {
		if m.DestinationNodeID != 0 {
			panic("cannot use both destination node ID and destination group ID")
		}

		f |= messageFlagDSIZGroupID
	}

	return f
}

func (m Message) securityFlags() byte {
	var f uint8

	if m.HasPrivacyEnhancements {
		f |= securityFlagP
	}

	if m.IsControlMessage {
		f |= securityFlagC
	}

	if len(m.MessageExtensions) > 0 {
		f |= securityFlagMX
	}

	if m.IsGroupSession {
		f |= sessionTypeGroup
	} else {
		f |= sessionTypeUnicast
	}

	return f
}

func (m Message) header() ([]byte, error) {
	w := &bytes.Buffer{}

	if err := wire.WriteInt(w, m.messageFlags()); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.SessionID); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.securityFlags()); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.MessageCounter); err != nil {
		return nil, err
	}

	if m.SourceNodeID != 0 {
		if err := wire.WriteInt(w, m.SourceNodeID); err != nil {
			return nil, err
		}
	}

	if m.DestinationNodeID != 0 {
		if err := wire.WriteInt(w, m.DestinationNodeID); err != nil {
			return nil, err
		}
	} else if m.DestinationGroupID != 0 {
		if err := wire.WriteInt(w, m.DestinationGroupID); err != nil {
			return nil, err
		}
	}

	return w.Bytes(), nil
}

func (m Message) payload() ([]byte, error) {
	w := &bytes.Buffer{}

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

func (m Message) nonce() ([]byte, error) {
	w := &bytes.Buffer{}

	if err := wire.WriteInt(w, m.securityFlags()); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.MessageCounter); err != nil {
		return nil, err
	}

	if err := wire.WriteInt(w, m.SourceNodeID); err != nil {
		return nil, err
	}

	return w.Bytes(), nil
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
