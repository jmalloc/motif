package message

import (
	crypto "github.com/jmalloc/motif/internal/crypto/mappingv1"
)

const (
	minHeaderSize = 4
	maxHeaderSize = 24
)

const (
	messageFlagsOffset         = 0
	sessionIDOffset            = 1
	securityFlagsOffset        = 3
	messageCounterOffset       = 4
	sourceAndDestinationOffset = 8
)

const (
	// messageFlagS is a bit-mask that isolates the "S" sub-field of the
	// "message flags" bit-field. The "S" sub-field is a boolean that indicates
	// whether the message has a "source node ID" field.
	messageFlagS = 0b000_0_1_00

	// messageFlagDSIZMask is a bit-mask that isolates the "DSIZ" sub-field of
	// the "message flags" bit-field.
	messageFlagDSIZMask = 0b000_0_0_11

	// messageFlagDSIZNodeID is a value of the "DSIZ" sub-field of the "message
	// flags" bit-field. It indicates that the message has a "destination node
	// ID" field.
	messageFlagDSIZNodeID = 0b000_0_0_01

	// messageFlagDSIZGroupID is a value of the "DSIZ" sub-field of the "message
	// flags" bit-field. It indicates that the message has a "destination group
	// ID" field.
	messageFlagDSIZGroupID = 0b000_0_0_10
)

// hasMessageFlag returns true if the given message flag is set in the header.
func hasMessageFlag(header []byte, flag uint8) bool {
	return (header[messageFlagsOffset] & flag) != 0
}

// setMessageFlag sets the given message flag in the header.
func setMessageFlag(header []byte, flag uint8) {
	header[messageFlagsOffset] |= flag
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
func setSecurityFlag(header []byte, flag uint8) {
	header[securityFlagsOffset] |= flag
}

// securityNonce returns the nonce to use for encrypting/decrypting the message
// payload.
func securityNonce(data []byte) crypto.Nonce {
	var nonce crypto.Nonce

	// security flags & message counter
	copy(nonce[:], data[3:8])

	// source node ID, if present
	if hasMessageFlag(data, messageFlagS) {
		copy(nonce[5:], data[8:])
	}

	return nonce
}

// privacyNonce returns the nonce to use for encrypting/decrypting the part of
// the message header that is privatized.
func privacyNonce(data []byte) crypto.Nonce {
	var nonce crypto.Nonce

	// sessionID in big-endian format (usual encoding is little-endian)
	nonce[0] = data[2]
	nonce[1] = data[1]

	// remainder of nonce is the least-significant end of the MIC
	const micFragmentSize = crypto.MICSize - crypto.NonceSize + 2
	mic := crypto.ExtractMIC(data)
	copy(nonce[2:], mic[micFragmentSize:])

	return nonce
}
