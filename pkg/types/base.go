package types

import "fmt"

// Interface for object from which is possible to get a hash value.
type Hashable interface {
	Hash() uint32
}

// Interface for object that are comparable with each other
type Comparable interface {

	// The comparison function returns
	// 	-1 if the receiver is less than the argument
	// 	 0 if the receiver is equal to the argument
	// 	 1 if the receiver is greater than the argument

	// It checks that the argument is of the same type as the receiver.
	// If it is not the case, it panics.
	Cmp(any) int
}

// Base "building-block" interface for many objects
// It allows comparison, hashing and string representation
type Base interface {
	fmt.Stringer
	Hashable
	Comparable
}
