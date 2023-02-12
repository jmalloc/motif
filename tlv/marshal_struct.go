package tlv

const (
	structType     = 0b000_10101
	endOfContainer = 0b000_11000

	contextSpecificTagForm = 0b001_00000
)

func (m marshaler) VisitStruct(s Struct) {
	m.WriteControl(structType)
	for _, sm := range s {
		marshal(m.Buffer, sm)
	}
	m.WriteByte(endOfContainer)
}
