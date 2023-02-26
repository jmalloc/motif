package message_test

import (
	crypto "github.com/jmalloc/motif/internal/crypto/mappingv1"
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
		func(m Message, key crypto.SymmetricKey, data []byte) {
			p := func(sessionID uint16) (crypto.SymmetricKey, error) {
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
			crypto.SymmetricKey{ // key
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
			crypto.SymmetricKey{ // key
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
		XEntry(
			"group message with privacy enhancements",
			Message{
				SessionID:            56189,
				IsGroupSession:       true,
				SourceNodeID:         1,
				DestinationGroupID:   2,
				MessageCounter:       305419897,
				MessagePayload:       []byte{0x01, 0x64, 0xee, 0x0e, 0x20, 0x7d},
				UsePrivacyExtensions: true,
			},
			crypto.SymmetricKey{ // key
				0xca, 0x92, 0xd7, 0xa0, 0x94, 0x2d, 0x1a, 0x51,
				0x1a, 0x0e, 0x26, 0xad, 0x07, 0x4f, 0x4c, 0x2f,
			},
			[]byte{ // ciphertext packet
				0x06,       // message flags
				0x7d, 0xdb, // session ID
				0x81,                   // security flags
				0xd9, 0x26, 0xaf, 0xce, // privatized data (message counter)
				0x24, 0xc8, 0xa0, 0x98, 0x1b, 0xdd, 0x44, 0xf4, // privatized data (source node ID)
				0xe7, 0x30, // privatized data (destination group ID)
				0x2b, 0x2f, 0x91, 0x5a, 0x66, 0xc9, // ciphertext payload
				0x59, 0x62, 0x90, 0xeb, 0xe4, 0x40, 0x82, 0x17, // mic
				0xb3, 0xc0, 0xc9, 0x21, 0xa2, 0xfc, 0xa4, 0xe1,
			},
		),
	)
})

// messageCounter = 0x12345679
// .plain     = "\x06 \x7d\xdb \x81 \x79\x56\x34\x12 \x01\x00\x00\x00\x00\x00\x00\x00 \x02\x00 \x01\x64\xee\x0e\x20\x7d",
// .encrypted = "\x06 \x7d\xdb \x81 \x79\x56\x34\x12 \x01\x00\x00\x00\x00\x00\x00\x00 \x02\x00 \x2b\x2f\x91\x5a\x66\xc9"
//              "\x59\x62\x90\xeb\xe4\x40\x82\x17\xb3\xc0\xc9\x21\xa2\xfc\xa4\xe1",
// .privacy =   "\x06 \x7d\xdb \x81 \xd9\x26\xaf\xce \x24\xc8\xa0\x98\x1b\xdd\x44\xf4 \xe7\x30 \x2b\x2f\x91\x5a\x66\xc9"
//              "\x59\x62\x90\xeb\xe4\x40\x82\x17\xb3\xc0\xc9\x21\xa2\xfc\xa4\xe1",

// .payloadLength   = 0,
// .plainLength     = 24,
// .encryptedLength = 40,
// .privacyLength   = 40,

// .encryptKey = "\xca\x92\xd7\xa0\x94\x2d\x1a\x51\x1a\x0e\x26\xad\x07\x4f\x4c\x2f",
// .privacyKey = "\xbf\xe9\xda\x01\x6a\x76\x53\x65\xf2\xdd\x97\xa9\xf9\x39\xe4\x25",
// .epochKey   = "\xb0\xb1\xb2\xb3\xb4\xb5\xb6\xb7\xb8\xb9\xba\xbb\xbc\xbd\xbe\xbf",

// .nonce        = "\x01\x79\x56\x34\x12\x01\x00\x00\x00\x00\x00\x00\x00",
// .privacyNonce = "\xdb\x7d\x40\x82\x17\xb3\xc0\xc9\x21\xa2\xfc\xa4\xe1",

// .sessionId    = 0xdb7d, // 56189
// .peerNodeId   = 0x0000000000000000ULL,
// .groupId      = 2,
// .sourceNodeId = 0x0000000000000002ULL,

// .expectedMessageCount = 1,
