package message_test

import (
	. "github.com/jmalloc/motif/message"
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
					SourceNodeID:   0xbaadf00d,
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
					DestinationNodeID: 0xbaadf00d,
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
					DestinationGroupID: 0xf00d,
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
					SourceNodeID:      0xbaadf00d,
					DestinationNodeID: 0xdeadbeef,
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
					SourceNodeID:       0xbaadf00d,
					DestinationGroupID: 0xbeef,
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

	DescribeTable(
		"it encodes/decodes secured messages correctly",
		func(m Message, key, data []byte) {
			p := func(sessionID uint16) ([]byte, error) {
				Expect(sessionID).To(Equal(m.SessionID))
				return key, nil
			}

			d, err := m.MarshalBinary(p)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(d).To(Equal(data))

			var u Message
			err = u.UnmarshalBinary(p, data)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(u).To(Equal(m))
		},
		Entry(
			"basic",
			Message{
				SessionID:      3000,
				MessageCounter: 12345,
				MessagePayload: []byte{0x05, 0x64, 0xee, 0x0e, 0x20, 0x7d},
			},
			[]byte{ // key
				0x5e, 0xde, 0xd2, 0x44, 0xe5, 0x53, 0x2b, 0x3c,
				0xdc, 0x23, 0x40, 0x9d, 0xba, 0xd0, 0x52, 0xd2,
			},
			[]byte{ // ciphertext packet
				0x00,       // message flags
				0xb8, 0x0b, // session ID
				0x00,                   // security flags
				0x39, 0x30, 0x00, 0x00, // message counter
				0x5a, 0x98, 0x9a, 0xe4, 0x2e, 0x8d, // ciphertext payload
				0x84, 0x7f, 0x53, 0x5c, 0x30, 0x07, 0xe6, 0x15, // mic
				0x0c, 0xd6, 0x58, 0x67, 0xf2, 0xb8, 0x17, 0xdb,
			},
		),
		Entry(
			"group message",
			Message{
				SessionID:          56189,
				IsGroupSession:     true,
				SourceNodeID:       1,
				DestinationGroupID: 2,
				MessageCounter:     305419896,
				MessagePayload:     []byte{0x01, 0x64, 0xee, 0x0e, 0x20, 0x7d},
			},
			[]byte{ // key
				0xca, 0x92, 0xd7, 0xa0, 0x94, 0x2d, 0x1a, 0x51,
				0x1a, 0x0e, 0x26, 0xad, 0x07, 0x4f, 0x4c, 0x2f,
			},
			[]byte{ // ciphertext packet
				0x06,       // message flags
				0x7d, 0xdb, // session ID
				0x01,                   // security flags
				0x78, 0x56, 0x34, 0x12, // message counter
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // source node ID
				0x02, 0x00, // destination group ID
				0x65, 0xc7, 0x67, 0xbc, 0x6c, 0xda, // ciphertext payload
				0x01, 0x06, 0xc9, 0x80, 0x13, 0x23, 0x90, 0x0e, // mic
				0x9b, 0x3c, 0xe6, 0xd4, 0xbb, 0x03, 0x27, 0xd6,
			},
		),

		// {
		// 	name: "private group message",
		// 	plain: "067ddb8179563412010000000000000002000164ee0e207d",
		// 	encrypted: "067ddb8179563412010000000000000002002b2f915a66c9596290ebe4408217b3c0c921a2fca4e1",
		// 	privacy:   "067ddb81d926afce24c8a0981bdd44f4e7302b2f915a66c9596290ebe4408217b3c0c921a2fca4e1",

		// 	encryptKey: "ca92d7a0942d1a511a0e26ad074f4c2f",
		// 	privacyKey: "bfe9da016a765365f2dd97a9f939e425",
		// 	privacyNonce: "db7d408217b3c0c921a2fca4e1",
		// 	nonce: "81795634120100000000000000",

		// 	// Extra fields for deriving the privacy key
		// 	compressedFabricId: "2906c908d115d362",
		// 	epochKey: "b0b1b2b3b4b5b6b7b8b9babbbcbdbebf",
		// 	sessionId: 0xdb7d, // 56157
		// },
		// {
		// 	name: "private group message (retry)",
		// 	plain: "067ddb8178563412010000000000000002000164ee0e207d",
		// 	encrypted: "067ddb8178563412010000000000000002005b458675d42d2486dad6944fca22328e6b1e44dc0468",
		// 	privacy:   "067ddb8131457fc65ec2edafaf0166a065425b458675d42d2486dad6944fca22328e6b1e44dc0468",

		// 	encryptKey: "ca92d7a0942d1a511a0e26ad074f4c2f",
		// 	privacyKey: "bfe9da016a765365f2dd97a9f939e425",
		// 	privacyNonce: "db7d4fca22328e6b1e44dc0468",
		// 	nonce: "81785634120100000000000000",

		// 	// Extra fields for deriving the privacy key
		// 	compressedFabricId: "2906c908d115d362",
		// 	epochKey: "b0b1b2b3b4b5b6b7b8b9babbbcbdbebf",
		// 	sessionId: 0xdb7d, // 56157
		// },

	)
})
