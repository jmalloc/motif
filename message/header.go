package message

import (
	"encoding/binary"
	"errors"

	crypto "github.com/jmalloc/motif/internal/crypto/mappingv1"
)

const (
	messageFlagsOffset = 0
	messageFlagsSize   = 1

	sessionIDOffset = 1
	sessionIDSize   = 2

	securityFlagsOffset = 3
	securityFlagsSize   = 1

	messageCounterOffset = 4
	messageCounterSize   = 4

	sourceAndDestinationOffset = 8
	nodeIDSize                 = 8
	groupIDSize                = 2

	minHeaderSize        = messageFlagsSize + sessionIDSize + securityFlagsSize + messageCounterSize
	extensionsLengthSize = 2
)

const (
	// messageFlagS is a bit-mask that isolates the "S" sub-field of the
	// "message flags" bit-field. The "S" sub-field is a boolean that indicates
	// whether the message has a "source node ID" field.
	messageFlagS = 0b000_0_1_00

	// messageFlagDSIZMask is a bit-mask that isolates the "DSIZ" sub-field of
	// the "message flags" bit-field.
	messageFlagDSIZMask = 0b000_0_0_11
)

const (
	// destinationNone is a value of the "DSIZ" sub-field of the "message flags"
	// bit-field. It indicates that the message has no destination field.
	destinationNone = 0b00

	// destinationNode is a value of the "DSIZ" sub-field of the "message flags"
	// bit-field. It indicates that the message has a "destination node ID"
	// field.
	destinationNode = 0b01

	// destinationGroup is a value of the "DSIZ" sub-field of the "message
	// flags" bit-field. It indicates that the message has a "destination group
	// ID" field.
	destinationGroup = 0b10

	// destinationReserved is a value of the "DSIZ" sub-field of the "message
	// flags" bit-field. This value is reserved for future use.
	destinationReserved = 0b11
)

// hasMessageFlag returns true if the given message flag is set in the header.
func hasMessageFlag(header []byte, flag uint8) bool {
	return (header[messageFlagsOffset] & flag) != 0
}

// setMessageFlag sets the given message flag in the header.
func setMessageFlag(header []byte, flag uint8, on bool) {
	if on {
		header[messageFlagsOffset] |= flag
	} else {
		header[messageFlagsOffset] &^= flag
	}
}

// destinationType returns the value of the "DSIZ" sub-field of the "message
// flags" bit-field.
func destinationType(header []byte) uint8 {
	return header[messageFlagsOffset] & messageFlagDSIZMask
}

// setDestinationType sets the value of the "DSIZ" sub-field of the "message
// flags" bit-field.
func setDestinationType(header []byte, t uint8) {
	header[messageFlagsOffset] &^= messageFlagDSIZMask
	header[messageFlagsOffset] |= t
}

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

// hasSecurityFlag returns true if the given security flag is set in the header.
func hasSecurityFlag(header []byte, flag uint8) bool {
	return (header[securityFlagsOffset] & flag) != 0
}

// setSecurityFlag sets the given security flag in the header.
func setSecurityFlag(header []byte, flag uint8, on bool) {
	if on {
		header[securityFlagsOffset] |= flag
	} else {
		header[securityFlagsOffset] &^= flag
	}
}

// securityNonce returns the nonce to use for encrypting/decrypting the message
// payload.
func securityNonce(data []byte) crypto.Nonce {
	var nonce crypto.Nonce

	// security flags & message counter
	copy(
		nonce[:],
		data[securityFlagsOffset:messageCounterOffset+messageCounterSize],
	)

	// source node ID, if present
	if hasMessageFlag(data, messageFlagS) {
		copy(
			nonce[messageCounterOffset+messageCounterSize-securityFlagsOffset:],
			data[sourceAndDestinationOffset:],
		)
	}

	return nonce
}

// privacyNonce returns the nonce to use for encrypting/decrypting the part of
// the message header that is privatized.
func privacyNonce(data []byte) crypto.Nonce {
	var nonce crypto.Nonce

	// session ID in big-endian format
	binary.BigEndian.PutUint16(
		nonce[:],
		binary.LittleEndian.Uint16(data[sessionIDOffset:]),
	)

	mic := crypto.ExtractMIC(data)

	copy(
		nonce[sessionIDSize:],
		mic[crypto.MICSize-crypto.NonceSize+sessionIDSize:],
	)

	return nonce
}

// splitPacket splits a packet into separate header and payload chunks.
//
// It returns an error if data is definitely not long enough to contain the
// header and a valid message. A nil error does not guarantee that the message
// is long enough.
func splitPacket(data []byte) (header, destination, payload []byte, err error) {
	n := minHeaderSize

	if len(data) < n {
		return nil, nil, nil, errors.New("message header is too short")
	}

	if hasMessageFlag(data, messageFlagS) {
		n += nodeIDSize
	}

	destinationOffset := n

	switch destinationType(data) {
	case destinationReserved:
		return nil, nil, nil, errors.New("unsupported (reserved) destination type")
	case destinationNode:
		n += nodeIDSize
	case destinationGroup:
		n += groupIDSize
	}

	payloadOffset := n

	if hasSecurityFlag(data, securityFlagMX) {
		n += extensionsLengthSize
	}

	if len(data) < n {
		return nil, nil, nil, errors.New("message header is too short")
	}

	return data[:payloadOffset],
		data[destinationOffset:payloadOffset],
		data[payloadOffset:],
		nil
}
