package optional

// Value represents an optional value of type T.
type Value[T any] struct {
	v T
	p bool
}

// With returns an Optional[T] that contains the given value.
func With[T any](v T) Value[T] {
	return Value[T]{v, true}
}

// None returns an Optional[T] that does not contain a value.
func None[T any]() Value[T] {
	return Value[T]{}
}

// IsPresent returns true if the optional value is present.
func (o Value[T]) IsPresent() bool {
	return o.p
}

// Value returns the value.
//
// It panics if the value is not present.
func (o Value[T]) Value() T {
	if o.p {
		return o.v
	}

	panic("value is not present")
}
