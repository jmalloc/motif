package tlv

// Element is an interface for a TLV element.
type Element interface {
	Components() (Tag, Value)
}

// Container is an interface for elements that contain other elements.
type Container interface {
	Members() []Element
}

// Root is the element at the root of a TLV element tree.
type Root struct {
	Tag   Tag
	Value Value
}

// Components returns the tag and value of the element.
func (r Root) Components() (Tag, Value) {
	t := r.Tag
	if t == nil {
		t = AnonymousTag
	}

	v := r.Value
	if v == nil {
		v = Null
	}

	return t, v
}
