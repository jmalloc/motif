package tlv

const (
	structType     = 0b000_10101
	endOfContainer = 0b000_11000
)

func (m *marshaler) VisitStruct(v Struct) {
	m.control |= structType

	for _, member := range v.members {
		member.AcceptElementVisitor(m)
	}

	m.payload.WriteByte(endOfContainer)
}
