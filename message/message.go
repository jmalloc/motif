package message

import (
	"bytes"
	"io"

	"github.com/jmalloc/motif/internal/crypto/chipcompat"
	crypto "github.com/jmalloc/motif/internal/crypto/mappingv1"
	"github.com/jmalloc/motif/internal/wire"
)

// KeyProvider is a function that locates the encryption/decryption key to use
// based on the session ID.
type KeyProvider func(sessionID uint16) (crypto.SymmetricKey, error)

// Message is a Matter message frame.
type Message struct {
	SessionID            uint16
	IsGroupSession       bool
	IsControlMessage     bool
	UsePrivacyExtensions bool
	SourceNodeID         uint64
	DestinationNodeID    uint64
	DestinationGroupID   uint16
	MessageCounter       uint32
	MessagePayload       []byte
	MessageExtensions    []byte
}

// MarshalBinary returns the binary representation of m.
func (m Message) MarshalBinary(key KeyProvider) ([]byte, error) {
	var data []byte

	wire.AppendInt(&data, uint8(0)) // message flags
	wire.AppendInt(&data, m.SessionID)
	wire.AppendInt(&data, uint8(0)) // security flags
	wire.AppendInt(&data, m.MessageCounter)

	if m.IsGroupSession {
		setSecurityFlag(data, sessionTypeGroup)
	}

	if m.IsControlMessage {
		setMessageFlag(data, securityFlagC)
	}

	if m.SourceNodeID != 0 {
		setMessageFlag(data, messageFlagS)
		wire.AppendInt(&data, m.SourceNodeID)
	}

	if m.DestinationNodeID != 0 {
		setMessageFlag(data, messageFlagDSIZNodeID)
		wire.AppendInt(&data, m.DestinationNodeID)
	} else if m.DestinationGroupID != 0 {
		setMessageFlag(data, messageFlagDSIZGroupID)
		wire.AppendInt(&data, m.DestinationGroupID)
	}

	headerSize := len(data)

	if size := len(m.MessageExtensions); size > 0 {
		setSecurityFlag(data, securityFlagMX)
		if err := wire.AppendString[uint16](&data, m.MessageExtensions); err != nil {
			return nil, err
		}
	}

	data = append(data, m.MessagePayload...)

	if m.SessionID != 0 {
		encryptionKey, err := key(m.SessionID)
		if err != nil {
			return nil, err
		}

		ciphertext, err := crypto.AEADEncrypt(
			encryptionKey,
			data[headerSize:],
			data[:headerSize],
			securityNonce(data),
		)
		if err != nil {
			return nil, err
		}

		data = append(data[:headerSize], ciphertext...)

		if m.UsePrivacyExtensions {
			setSecurityFlag(data, securityFlagP)

			ciphertext = crypto.PrivacyEncrypt(
				derivePrivacyKey(encryptionKey),
				data[messageCounterOffset:headerSize],
				privacyNonce(data),
			)
			copy(data[messageCounterOffset:], ciphertext)
		}
	}

	return data, nil
}

// UnmarshalBinary sets m to the value represented by data.
func (m *Message) UnmarshalBinary(key KeyProvider, data []byte) error {
	r := bytes.NewReader(data)

	if _, err := r.Seek(sessionIDOffset, io.SeekStart); err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.SessionID); err != nil {
		return err
	}

	if _, err := r.Seek(messageCounterOffset, io.SeekStart); err != nil {
		return err
	}

	if err := wire.AssignInt(r, &m.MessageCounter); err != nil {
		return err
	}

	if !hasMessageFlag(data, messageFlagS) {
		m.SourceNodeID = 0
	} else if err := wire.AssignInt(r, &m.SourceNodeID); err != nil {
		return err
	}

	if !hasMessageFlag(data, messageFlagDSIZNodeID) {
		m.DestinationNodeID = 0
	} else if err := wire.AssignInt(r, &m.DestinationNodeID); err != nil {
		return err
	}

	if !hasMessageFlag(data, messageFlagDSIZGroupID) {
		m.DestinationGroupID = 0
	} else if err := wire.AssignInt(r, &m.DestinationGroupID); err != nil {
		return err
	}

	m.IsGroupSession = hasSecurityFlag(data, sessionTypeGroup)
	m.IsControlMessage = hasSecurityFlag(data, securityFlagC)
	m.UsePrivacyExtensions = hasSecurityFlag(data, securityFlagP)

	headerSize, _ := r.Seek(0, io.SeekCurrent)

	if m.SessionID != 0 {
		encryptionKey, err := key(m.SessionID)
		if err != nil {
			return err
		}

		if m.UsePrivacyExtensions {
			plaintext := crypto.PrivacyDecrypt(
				derivePrivacyKey(encryptionKey),
				data[messageCounterOffset:headerSize],
				privacyNonce(data),
			)
			copy(data[messageCounterOffset:], plaintext)
		}

		plaintext, err := crypto.AEADDecrypt(
			encryptionKey,
			data[headerSize:],
			data[:headerSize],
			securityNonce(data),
		)
		if err != nil {
			return err
		}

		r.Reset(plaintext)
	}

	if !hasSecurityFlag(data, securityFlagMX) {
		m.MessageExtensions = nil
	} else if err := wire.AssignString[uint16](r, &m.MessageExtensions); err != nil {
		return err
	}

	return wire.AssignRemaining(r, &m.MessagePayload)
}

// derivePrivacyKey derives a privacy key from an encryption key.
func derivePrivacyKey(encryptionKey crypto.SymmetricKey) crypto.SymmetricKey {
	return chipcompat.DeriveKey(
		encryptionKey,
		nil, // salt
		[]byte("PrivacyKey"),
	)
}
