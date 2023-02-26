package protocol

import "errors"

const (
	exchangeFlagsOffset = 0
	exchangeFlagsSize   = 1

	protocolOpCodeOffset = 1
	protocolOpCodeSize   = 1

	exchangeIDOffset = 2
	exchangeIDSize   = 2

	protocolIDOffset = 4
	protocolIDSize   = 2

	protocolVendorIDOffset = 6
	protocolVendorIDSize   = 2

	messageCounterSize = 4

	minHeaderSize        = exchangeFlagsSize + protocolOpCodeSize + exchangeIDSize + protocolIDSize
	extensionsLengthSize = 2
)

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

// hasExchangeFlag returns true if the given message flag is set in the header.
func hasExchangeFlag(header []byte, flag uint8) bool {
	return (header[exchangeFlagsOffset] & flag) != 0
}

// setExchangeFlag sets the given message flag in the header.
func setExchangeFlag(header []byte, flag uint8, on bool) {
	if on {
		header[exchangeFlagsOffset] |= flag
	} else {
		header[exchangeFlagsOffset] &^= flag
	}
}

// splitPacket splits a packet into separate header and payload chunks.
//
// It returns an error if data is definitely not long enough to contain the
// header and a valid message. A nil error does not guarantee that the message
// is long enough.
func splitPacket(data []byte) (header, ack, payload []byte, err error) {
	n := minHeaderSize

	if len(data) < n {
		return nil, nil, nil, errors.New("protocol message header is too short")
	}

	if hasExchangeFlag(data, exchangeFlagV) {
		n += protocolVendorIDSize
	}

	ackOffset := n

	if hasExchangeFlag(data, exchangeFlagA) {
		n += messageCounterSize
	}

	payloadOffset := n

	if hasExchangeFlag(data, exchangeFlagSX) {
		n += extensionsLengthSize
	}

	if len(data) < n {
		return nil, nil, nil, errors.New("message header is too short")
	}

	return data[:payloadOffset],
		data[ackOffset:payloadOffset],
		data[payloadOffset:],
		nil
}
