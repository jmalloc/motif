package protocol

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/jmalloc/motif/optional"
	"golang.org/x/exp/slices"
)

// Message is the payload of a protocol message.
type Message struct {
	ExchangeID         uint16
	ProtocolVendorID   uint16
	ProtocolID         uint16
	ProtocolOpCode     uint8
	ApplicationPayload []byte
	IsFromInitiator    bool
	RequiresAck        bool
	AckMessageCounter  optional.Value[uint32]
	SecuredExtensions  []byte
}

// MarshalBinary returns the binary representation of m.
func (m Message) MarshalBinary() ([]byte, error) {
	if len(m.SecuredExtensions) > math.MaxUint16 {
		return nil, errors.New("secured extensions are too long")
	}

	var data []byte

	data = append(data, 0) // exchange flags
	setExchangeFlag(data, exchangeFlagI, m.IsFromInitiator)
	setExchangeFlag(data, exchangeFlagR, m.RequiresAck)

	data = append(data, m.ProtocolOpCode)
	data = binary.LittleEndian.AppendUint16(data, m.ExchangeID)
	data = binary.LittleEndian.AppendUint16(data, m.ProtocolID)

	if m.ProtocolVendorID != 0 {
		setExchangeFlag(data, exchangeFlagV, true)
		data = binary.LittleEndian.AppendUint16(data, m.ProtocolVendorID)
	}

	if m.AckMessageCounter.IsPresent() {
		setExchangeFlag(data, exchangeFlagA, true)
		data = binary.LittleEndian.AppendUint32(data, m.AckMessageCounter.Value())
	}

	if size := len(m.SecuredExtensions); size != 0 {
		setExchangeFlag(data, exchangeFlagSX, true)
		data = binary.LittleEndian.AppendUint16(data, uint16(size))
		data = append(data, m.SecuredExtensions...)
	}

	return append(data, m.ApplicationPayload...), nil
}

// UnmarshalBinary sets m to the value represented by data.
func (m *Message) UnmarshalBinary(data []byte) error {
	header, ack, payload, err := splitPacket(data)
	if err != nil {
		return err
	}

	m.IsFromInitiator = hasExchangeFlag(header, exchangeFlagI)
	m.RequiresAck = hasExchangeFlag(header, exchangeFlagR)

	m.ProtocolOpCode = header[protocolOpCodeOffset]
	m.ExchangeID = binary.LittleEndian.Uint16(header[exchangeIDOffset:])
	m.ProtocolID = binary.LittleEndian.Uint16(header[protocolIDOffset:])

	if hasExchangeFlag(header, exchangeFlagV) {
		m.ProtocolVendorID = binary.LittleEndian.Uint16(header[protocolVendorIDOffset:])
	} else {
		m.ProtocolVendorID = 0
	}

	if hasExchangeFlag(header, exchangeFlagA) {
		m.AckMessageCounter = optional.Some(binary.LittleEndian.Uint32(ack))
	} else {
		m.AckMessageCounter = optional.None[uint32]()
	}

	m.SecuredExtensions = nil
	if hasExchangeFlag(header, exchangeFlagSX) {
		size := int(binary.LittleEndian.Uint16(payload))
		payload = payload[extensionsLengthSize:]

		if len(payload) < size {
			return fmt.Errorf("secured extensions length is %d, but only %d bytes are available", size, len(payload))
		}

		if size != 0 {
			m.SecuredExtensions = slices.Clone(payload[:size])
		}

		payload = payload[size:]
	}

	if len(payload) > 0 {
		m.ApplicationPayload = slices.Clone(payload)
	} else {
		m.ApplicationPayload = nil
	}

	return nil
}
