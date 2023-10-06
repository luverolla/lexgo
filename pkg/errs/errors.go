// This package contains custom error and error messages
package errs

import "fmt"

// This error is thrown when a method attempts to access an element,
// in a collection, that does not exist
type NotFoundErr struct {
	// The value of the element that was not found
	Value any
}

func NotFound(v any) NotFoundErr {
	return NotFoundErr{v}
}

func (err NotFoundErr) Error() string {
	return fmt.Sprintf("Element %v not found", err.Value)
}

// This error is thrown when a method attempts to Get/Peek/Pop/Remove
// from an empty collection
type EmptyErr struct{}

func Empty() EmptyErr {
	return EmptyErr{}
}

func (err EmptyErr) Error() string {
	return "Attempted to Get/Peek/Pop/Remove from an empty collection"
}
