package tau

import "fmt"

// Stream-like interface, that allows to get values one by one.
// It is used to iterate over collections abstracting from their
// implementation details and internal structure
type Iterator[T any] interface {
	// Tries to get the next value from the iterator.
	// It returns a tuple (value, hasNext), where:
	// 	- value is the pointer to the next value in the iteration
	// 	- hasNext is a boolean that is true if values are still available
	// When hasNext is false, value is nil
	Next() (*T, bool)

	// This function emulates a for-each loop.
	// The given function is called for each available value.
	Each(func(T))
}

// Interface for objects that can be iterated over.
type Iterable[T any] interface {
	Iter() Iterator[T]
}

// Generic collection of objects. It can be iterated over and comparated
// with other collections.
//
// For the comparison to be possible, the elements of the collection must
// be of comparable trait, i.e. the [Cmp] function must not panic on them.
//
// The comparison criteria are the following
// 	- if the two collections have different size, the smaller is the lesser
// 	- else, elements are compared one by one and the comparison value of
//    the first different pair is returned
//
// The code is the following, where recv is the receiver
// and other is the argument:
//
// 	if recv.Size() != other.Size() {
// 		Cmp(recv.Size(), other.Size())
// 	}
// 	iter, otherIter := recv.Iter(), other.Iter()
// 	for next, hasNext = iter.Next(); hasNext; next, hasNext = iter.Next() {
// 		otherNext, _ := otherIter.Next()
// 		cmp := Cmp(*next, *otherNext)
// 		if cmp != 0 {
// 			return cmp
// 		}
// 	}
// 	return 0
//
// So, two collections are equal if of the same size and whose
// elements are equal in the same order. The two collections have
// not to be of the same kind (e.g. also a list and a set can be equal).
type Collection[T any] interface {
	fmt.Stringer
	Comparable
	Iterable[T]
	Size() int
	// It's a shorthand for Size() == 0. But, if there are
	// more efficient ways to check emptiness, it can be overridden.
	Empty() bool
	// Removes all the elements from the collection
	Clear()
	Contains(T) bool
	// Checks if the receiver contains ALL the elements of the other collection
	ContainsAll(Collection[T]) bool
	// Checks if the receiver contains ANY (even only one) of the elements of the other collection
	ContainsAny(Collection[T]) bool
	// Makes a copy of the collection
	Clone() Collection[T]
}

// Generic collection with indexwise access. It allows duplicates
//
// All implementation provides a circular access,
// like in other languages as Python. Meaning that:
//   - negative indices are allowed and will be interpreted
//     as their absolute value, but starting from the end of the list.
//   - after this check, a modulus operation is performed on the index
//     to make sure that it is in the range [0, size)
//
// These two operations together will be referred to as "index sanification"
// and are performed by a private function in the implementations.
// The sanification pseudo-code is the following:
//
//	INPUT(index, size)
//	if index < 0:
//		index = index + size
//	endif
//	index = index % size
//	OUTPUT(index)
type IdxedColl[T any] interface {
	Collection[T]
	// Returns the element at the given index
	// Returns an error if the collection is empty
	Get(int) (*T, error)
	// Replaces the existing value at the given index with the given one
	// Does not add new elements and hence does not change the size of the collection
	Set(int, T)
	// Inserts a new "block" with given value at the given index
	// The existing "block" at that index and all the following ones are rearranged
	// The rearrangement is implementation-dependent
	Insert(int, T)
	// Removes the "block" at the given index
	// The following "blocks" are rearranged
	// The rearrangement is implementation-dependent
	// Returns an error if the collection is empty
	RemoveAt(int) (*T, error)
	// Returns the index of the first occurrence of the given value
	// Returns -1 if the value is not found
	IndexOf(T) int
	// Returns the index of the last occurrence of the given value
	// Returns -1 if the value is not found
	LastIndexOf(T) int
	// Swaps the elements at the given indices
	Swap(int, int)
	// Returns a new collection containing the elements in the range [start, end)
	// It makes a copy of involved elements, so the original collection is not modified
	// Obviously, a slice of an empty collection is an empty collection itself
	//
	// The following checks are performed:
	// 	- start == end: returns an empty collection
	// 	- start > end: start and end are swapped
	// After these checks, the aforementioned index sanification is applied
	Slice(int, int) IdxedColl[T]
}

// Interface for filtering functions.
// The value argument is compulsory, while the others are optional.
type Filter[T any] func(T, ...any) bool

// Interface comparator functions. Allow definition of custom comparison criteria
// for sorting collections. The comparison must be the same as in [Cmp] and [Comparable]
type Comparator[T any] func(T, T) int
