package tlv

// Element is an interface for a TLV element.
type Element interface {
	Tag() Tag
	Value() Value
}

// Root is the element at the root of a TLV element tree.
type Root struct {
	T Tag
	V Value
}

// Tag returns the element's tag.
func (r Root) Tag() Tag {
	if r.T == nil {
		return AnonymousTag
	}
	return r.T
}

// Value returns the element's value.
func (r Root) Value() Value {
	if r.V == nil {
		return Null
	}
	return r.V
}
