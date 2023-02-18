package message_test

import (
	. "github.com/jmalloc/motif/message"
	"github.com/jmalloc/motif/optional"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Message", func() {
	Describe("func MarshalBinary() and UnmarshalBinary()", func() {
		DescribeTable(
			"it encodes/decodes unsecured messages correctly",
			func(m Message, data []byte) {
				d, err := m.MarshalBinary(nil)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(d).To(Equal(data))

				var u Message
				err = u.UnmarshalBinary(nil, data)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(u).To(Equal(m))
			},
			Entry(
				"empty",
				Message{},
				[]byte{
					0x00,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0x00, 0x00, 0x00, 0x00, // message counter
				},
			),
			Entry(
				"basic",
				Message{
					MessageCounter: 0xdeadbeef,
					MessagePayload: []byte("<payload>"),
				},
				[]byte{
					0x00,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0xef, 0xbe, 0xad, 0xde, // message counter
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
			Entry(
				"with source node ID",
				Message{
					SourceNodeID:   optional.With[uint64](0xbaadf00d),
					MessageCounter: 0xdeadbeef,
					MessagePayload: []byte("<payload>"),
				},
				[]byte{
					0x04,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0xef, 0xbe, 0xad, 0xde, // message counter
					0x0d, 0xf0, 0xad, 0xba, 0x00, 0x00, 0x00, 0x00, // source node ID
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
			Entry(
				"with destination node ID",
				Message{
					DestinationNodeID: optional.With[uint64](0xbaadf00d),
					MessageCounter:    0xdeadbeef,
					MessagePayload:    []byte("<payload>"),
				},
				[]byte{
					0x01,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0xef, 0xbe, 0xad, 0xde, // message counter
					0x0d, 0xf0, 0xad, 0xba, 0x00, 0x00, 0x00, 0x00, // destination node ID
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
			Entry(
				"with destination group ID",
				Message{
					DestinationGroupID: optional.With[uint16](0xf00d),
					MessageCounter:     0xdeadbeef,
					MessagePayload:     []byte("<payload>"),
				},
				[]byte{
					0x02,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0xef, 0xbe, 0xad, 0xde, // message counter
					0x0d, 0xf0, // destination group ID
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
			Entry(
				"with source and destination node IDs",
				Message{
					SourceNodeID:      optional.With[uint64](0xbaadf00d),
					DestinationNodeID: optional.With[uint64](0xdeadbeef),
					MessageCounter:    0,
					MessagePayload:    []byte("<payload>"),
				},
				[]byte{
					0x05,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0x00, 0x00, 0x00, 0x00, // message counter
					0x0d, 0xf0, 0xad, 0xba, 0x00, 0x00, 0x00, 0x00, // source node ID
					0xef, 0xbe, 0xad, 0xde, 0x00, 0x00, 0x00, 0x00, // destination node ID
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
			Entry(
				"with source node ID and destination group ID",
				Message{
					SourceNodeID:       optional.With[uint64](0xbaadf00d),
					DestinationGroupID: optional.With[uint16](0xbeef),
					MessageCounter:     0,
					MessagePayload:     []byte("<payload>"),
				},
				[]byte{
					0x06,       // message flags
					0x00, 0x00, // session ID
					0x00,                   // security flags
					0x00, 0x00, 0x00, 0x00, // message counter
					0x0d, 0xf0, 0xad, 0xba, 0x00, 0x00, 0x00, 0x00, // source node ID
					0xef, 0xbe, // destination node ID
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
			Entry(
				"with message extensions",
				Message{
					MessageExtensions: []byte("<extensions>"),
					MessagePayload:    []byte("<payload>"),
				},
				[]byte{
					0x00,       // message flags
					0x00, 0x00, // session ID
					0x20,                   // security flags
					0x00, 0x00, 0x00, 0x00, // message counter
					0x0c, 0x00, '<', 'e', 'x', 't', 'e', 'n', 's', 'i', 'o', 'n', 's', '>', // message extensions
					'<', 'p', 'a', 'y', 'l', 'o', 'a', 'd', '>', // message payload
				},
			),
		)
	})
})
