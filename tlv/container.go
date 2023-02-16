package tlv

type (
	// Struct is a TLV structure element.
	Struct []StructMember

	// StructMember is an element that is a member of a structure.
	StructMember struct {
		T NonAnonymousTag
		V Value
	}

	// Array is a TLV array element.
	Array []Value

	// List is a TLV list element.
	List []ListMember

	// ListMember is an element that is a member of a list.
	ListMember struct {
		T Tag
		V Value
	}
)

func (v Struct) acceptVisitor(vis ValueVisitor) error { return vis.VisitStruct(v) }
func (v Array) acceptVisitor(vis ValueVisitor) error  { return vis.VisitArray(v) }
func (v List) acceptVisitor(vis ValueVisitor) error   { return vis.VisitList(v) }
