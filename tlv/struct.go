package tlv

// Struct is a TLV structure element.
type Struct []StructMember

// AcceptVisitor dispatches to the method on v that corresponds to the concrete
// type the method's receiver.
func (s Struct) AcceptVisitor(v ValueVisitor) {
	v.VisitStruct(s)
}

// StructMember is an element that is a member of a structure.
type StructMember struct {
	Tag   Tag
	Value Value
}

// Components returns the tag and value of the element.
func (m StructMember) Components() (Tag, Value) {
	return m.Tag, m.Value
}

func (m marshaler) VisitStruct(s Struct) {
	m.WriteControl(structType)
	for _, sm := range s {
		marshal(m.Buffer, sm)
	}
	m.WriteByte(endOfContainer)
}
