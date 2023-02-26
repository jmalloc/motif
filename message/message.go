package message

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/jmalloc/motif/internal/crypto/chipcompat"
	crypto "github.com/jmalloc/motif/internal/crypto/mappingv1"
	"golang.org/x/exp/slices"
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
	if len(m.MessageExtensions) > math.MaxUint16 {
		return nil, errors.New("message extensions are too long")
	}

	var data []byte

	data = append(data, 0) // message flags
	data = binary.LittleEndian.AppendUint16(data, m.SessionID)
	data = append(data, 0) // security flags
	data = binary.LittleEndian.AppendUint32(data, m.MessageCounter)

	setSecurityFlag(data, sessionTypeGroup, m.IsGroupSession)
	setMessageFlag(data, securityFlagC, m.IsControlMessage)
	setSecurityFlag(data, securityFlagP, m.UsePrivacyExtensions)

	if m.SourceNodeID != 0 {
		setMessageFlag(data, messageFlagS, true)
		data = binary.LittleEndian.AppendUint64(data, m.SourceNodeID)
	}

	if m.DestinationNodeID != 0 {
		setDestinationType(data, destinationNode)
		data = binary.LittleEndian.AppendUint64(data, m.DestinationNodeID)
	} else if m.DestinationGroupID != 0 {
		setDestinationType(data, destinationGroup)
		data = binary.LittleEndian.AppendUint16(data, m.DestinationGroupID)
	}

	headerSize := len(data)

	if size := uint16(len(m.MessageExtensions)); size != 0 {
		setSecurityFlag(data, securityFlagMX, true)
		data = binary.LittleEndian.AppendUint16(data, size)
		data = append(data, m.MessageExtensions...)
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
			ciphertext := crypto.PrivacyEncrypt(
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
	header, destination, payload, err := splitPacket(data)
	if err != nil {
		return err
	}

	m.IsGroupSession = hasSecurityFlag(header, sessionTypeGroup)
	m.IsControlMessage = hasSecurityFlag(header, securityFlagC)
	m.UsePrivacyExtensions = hasSecurityFlag(header, securityFlagP)

	m.SessionID = binary.LittleEndian.Uint16(header[sessionIDOffset:])

	var encryptionKey crypto.SymmetricKey
	if m.SessionID != 0 {
		encryptionKey, err = key(m.SessionID)
		if err != nil {
			return err
		}

		if m.UsePrivacyExtensions {
			plaintext := crypto.PrivacyDecrypt(
				derivePrivacyKey(encryptionKey),
				header[messageCounterOffset:],
				privacyNonce(data),
			)
			copy(header[messageCounterOffset:], plaintext)
		}

		payload, err = crypto.AEADDecrypt(
			encryptionKey,
			payload,
			header,
			securityNonce(data),
		)
		if err != nil {
			return err
		}
	}

	m.MessageCounter = binary.LittleEndian.Uint32(header[messageCounterOffset:])

	if hasMessageFlag(header, messageFlagS) {
		m.SourceNodeID = binary.LittleEndian.Uint64(header[sourceAndDestinationOffset:])
	} else {
		m.SourceNodeID = 0
	}

	switch destinationType(header) {
	case destinationNode:
		m.DestinationNodeID = binary.LittleEndian.Uint64(destination)
		m.DestinationGroupID = 0
	case destinationGroup:
		m.DestinationNodeID = 0
		m.DestinationGroupID = binary.LittleEndian.Uint16(destination)
	default:
		m.DestinationNodeID = 0
		m.DestinationGroupID = 0
	}

	m.MessageExtensions = nil
	if hasSecurityFlag(header, securityFlagMX) {
		size := int(binary.LittleEndian.Uint16(payload))
		payload = payload[extensionsLengthSize:]

		if len(payload) < size {
			return fmt.Errorf("message extensions length is %d, but only %d bytes are available", size, len(payload))
		}

		if size != 0 {
			m.MessageExtensions = slices.Clone(payload[:size])
		}

		payload = payload[size:]
	}

	if len(payload) > 0 {
		m.MessagePayload = slices.Clone(payload)
	} else {
		m.MessagePayload = nil
	}

	return nil
}

// derivePrivacyKey derives a privacy key from an encryption key.
func derivePrivacyKey(encryptionKey crypto.SymmetricKey) crypto.SymmetricKey {
	return chipcompat.DeriveKey(
		encryptionKey,
		nil, // salt
		[]byte("PrivacyKey"),
	)
}
