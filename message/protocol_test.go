package message_test

import (
	. "github.com/jmalloc/motif/message"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type ProtocolMessage", func() {
	Describe("func MarshalBinary() and UnmarshalBinary()", func() {
		DescribeTable(
			"it encodes/decodes protocol messages correctly",
			func(m ProtocolMessage, data []byte) {
				d, err := m.MarshalBinary()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(d).To(Equal(data))

				var u ProtocolMessage
				err = u.UnmarshalBinary(d)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(u).To(Equal(m))
			},
			Entry(
				"zero-value",
				ProtocolMessage{},
				[]byte{
					0x00,       // exchange flags
					0x00,       // protocol opcode
					0x00, 0x00, // exchange ID
					0x00, 0x00, // protocol ID
				},
			),
			Entry(
				"is from initiator",
				ProtocolMessage{
					IsFromInitiator: true,
				},
				[]byte{
					0x01,       // exchange flags
					0x00,       // protocol opcode
					0x00, 0x00, // exchange ID
					0x00, 0x00, // protocol ID
				},
			),
			Entry(
				"is acknolwedgement",
				ProtocolMessage{
					AckMessageCounter: With[uint32](0xdeadbeef),
				},
				[]byte{
					0x02,       // exchange flags
					0x00,       // protocol opcode
					0x00, 0x00, // exchange ID
					0x00, 0x00, // protocol ID
					0xef, 0xbe, 0xad, 0xde, // ack message counter
				},
			),
			Entry(
				"requires acknowledgement",
				ProtocolMessage{
					RequiresAck: true,
				},
				[]byte{
					0x04,       // exchange flags
					0x00,       // protocol opcode
					0x00, 0x00, // exchange ID
					0x00, 0x00, // protocol ID
				},
			),
			Entry(
				"with secured extensions",
				ProtocolMessage{
					SecuredExtensions: []byte("<extensions>"),
				},
				[]byte{
					0x08,       // exchange flags
					0x00,       // protocol opcode
					0x00, 0x00, // exchange ID
					0x00, 0x00, // protocol ID
					0x0c, 0x00, '<', 'e', 'x', 't', 'e', 'n', 's', 'i', 'o', 'n', 's', '>', // secured extensions
				},
			),
			Entry(
				"has protocol vendor ID",
				ProtocolMessage{
					ProtocolVendorID: 42020,
				},
				[]byte{
					0x10,       // exchange flags
					0x00,       // protocol opcode
					0x00, 0x00, // exchange ID
					0x00, 0x00, // protocol ID
					0x24, 0xa4, // protocol vendor ID
				},
			),
			Entry(
				"all fields populated",
				ProtocolMessage{
					ExchangeID:         41010,
					ProtocolVendorID:   42020,
					ProtocolID:         43030,
					ProtocolOpCode:     44,
					ApplicationPayload: []byte("<payload>"),
					IsFromInitiator:    true,
					RequiresAck:        true,
					AckMessageCounter:  With[uint32](0xbaadf00d),
					SecuredExtensions:  []byte("<extensions>"),
				},
				[]byte{
					0x1f,       // exchange flags
					0x2c,       // protocol opcode
					0x32, 0xa0, // exchange ID
					0x16, 0xa8, // protocol ID
					0x24, 0xa4, // protocol vendor ID
					0x0d, 0xf0, 0xad, 0xba, // ack message counter
					0x0c, 0x00, '<', 'e', 'x', 't', 'e', 'n', 's', 'i', 'o', 'n', 's', '>', // secured extensions
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // application payload
				},
			),
		)
	})
})
