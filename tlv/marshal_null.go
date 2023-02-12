package tlv

const nullType = 0b000_10100

func (m marshaler) VisitNull() {
	m.WriteControl(nullType)
}
