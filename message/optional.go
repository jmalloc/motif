package message

// Optional represents an optional value of type T.
type Optional[T any] struct {
	value T
	ok    bool
}

// With returns an Optional[T] that contains the given value.
func With[T any](v T) Optional[T] {
	return Optional[T]{v, true}
}
