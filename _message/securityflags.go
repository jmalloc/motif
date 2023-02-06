package message

// SecurityFlags is a bit-field that contains information about the structure of
// a Matter message.
//
// Matter v1.0 Core § 4.4.1.4
type SecurityFlags uint8

const (
	// SecurityFlagsPMask is a bit-mask that isolates the "P" sub-field of the
	// SecurityFlags bit-field. The "P" sub-field is a boolean that indicates
	// whether the message is encoded with privacy enhancements.
	//
	// Matter v1.0 Core § 4.8.3
	SecurityFlagsPMask SecurityFlags = 0b10000000

	// SecurityFlagsCMask is a bit-mask that isolates the "C" sub-field of the
	// SecurityFlags bit-field. The "C" sub-field is a boolean that indicates
	// whether the message is a control message.
	//
	// Matter v1.0 Core § 4.16
	// Matter v1.0 Core § 4.7.1.1
	SecurityFlagsCMask SecurityFlags = 0b01000000

	// SecurityFlagsMXMask is a bit-mask that isolates the "MX" sub-field of the
	// SecurityFlags bit-field. The "MX" sub-field is a boolean that indicates
	// whether the message has extensions.
	//
	// Matter v1.0 Core § 4.4.1.8
	SecurityFlagsMXMask SecurityFlags = 0b00100000

	// SecurityFlagsSessionTypeMask is a bit-mask that isolates the "session
	// type" sub-field of the SecurityFlags bit-field. The "session type"
	// sub-field is an enumeration that indicates the type of session associated
	// with the message.
	SecurityFlagsSessionTypeMask SecurityFlags = 0b00000011
)

// HasPrivacyEnhancements returns true if the message is encoded with privacy
// enhancements.
func (f SecurityFlags) HasPrivacyEnhancements() bool {
	return f&SecurityFlagsPMask != 0
}

// IsControlMessage returns true if the message is a control message.
func (f SecurityFlags) IsControlMessage() bool {
	return f&SecurityFlagsCMask != 0
}

// HasExtensions returns true if the message has extensions.
func (f SecurityFlags) HasExtensions() bool {
	return f&SecurityFlagsMXMask != 0
}

// SessionType returns the type of session associated with the message.
func (f SecurityFlags) SessionType() SessionType {
	return SessionType(f & SecurityFlagsSessionTypeMask)
}

// SessionType is an enumeration that indicates the type of session associated
// with a Matter message.
type SessionType uint8

const (
	// SessionTypeUnicast indicates that the message is sent to a specific node.
	SessionTypeUnicast SessionType = 0b00

	// SessionTypeGroup indicates that the message is sent to all nodes in a
	// specific group.
	SessionTypeGroup SessionType = 0b01
)
