// This package contains custom error and error messages
package errs

// This error is thrown when a method attempts to access an element,
// in a collection, that does not exist
type NotFoundErr struct{}

func NotFound() NotFoundErr {
	return NotFoundErr{}
}

func (err NotFoundErr) Error() string {
	return "Not found"
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
