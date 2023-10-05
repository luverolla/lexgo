// This package contains interfaces for various collections.
// The generic type parameter T allows any time that fits into the
// [uni.Cmp] comparing function, therefore limited, at run-time, by
//   - types under [constraints.Ordered] constraint
//   - types implementing [types.Base] interface
package colls

import (
	"github.com/luverolla/lexgo/pkg/types"
	"golang.org/x/exp/constraints"
)

// Generic list with indexwise access. It allows duplicates
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
type List[T any] interface {
	types.Collection[T]
	Get(int) (*T, error)
	// Replaces the existing value at the given index with the given one
	// Does not add new elements and hence does not change the size of the list
	Set(int, T)
	Append(...T)
	Prepend(...T)
	// Inserts a new "block" with given value at the given index
	// The existing "block" at that index and all the following ones are shifted to the right
	Insert(int, T)
	// Removes the first occurrence of the given value
	// Returns an error if the value is not found
	RemoveFirst(T) error
	// Removes all the occurrences of the given value
	// Returns an error if the value is not found
	RemoveAll(T) error
	// Removes the "block" at the given index
	// The following "blocks" are shifted to the left
	// Returns an error if the list is empty
	RemoveAt(int) (*T, error)
	// Returns the index of the first occurrence of the given value
	// Returns -1 if the value is not found
	IndexOf(T) int
	// Returns the index of the last occurrence of the given value
	// Returns -1 if the value is not found
	LastIndexOf(T) int
	// Returns a new list containing the elements in the range [start, end)
	// It makes a copy of involved elements, so the original list is not modified
	// Obviously, a slice of an empty list is an empty list itself
	//
	// The following checks are performed:
	// 	- start == end: returns an empty list
	// 	- start > end: start and end are swapped
	// After these checks, the aforementioned index sanification is applied
	Slice(int, int) List[T]
	Sort(types.Comparator[T]) List[T]
	Sublist(types.Filter[T]) List[T]
}

// Generic (D)ouble-(e)nded (que)ue
// It allows both FIFO (queue-like) and LIFO (stack-like) access
type Deque[T any] interface {
	types.Collection[T]
	PushFront(...T)
	PushBack(...T)
	PopFront() (*T, error)
	PopBack() (*T, error)
	Front() (*T, error)
	Back() (*T, error)
	FIFOIter() types.Iterator[T]
	LIFOIter() types.Iterator[T]
}

// Generic node for a binary tree
type BSTreeNode[T any] interface {
	Value() T
	Left() BSTreeNode[T]
	Right() BSTreeNode[T]
}

// Generic binary search tree
type BSTree[T any] interface {
	types.Collection[T]
	Get(T) BSTreeNode[T]
	Root() BSTreeNode[T]
	Insert(T) BSTreeNode[T]
	Remove(T) BSTreeNode[T]
	Min() BSTreeNode[T]
	Max() BSTreeNode[T]
	Pred(T) BSTreeNode[T]
	Succ(T) BSTreeNode[T]
	PreOrder() types.Iterator[T]
	InOrder() types.Iterator[T]
	PostOrder() types.Iterator[T]
}

// Generic map
// The key MUST be a primary type (under [constraints.Ordered])
// while the type parameter can be really anything
type Map[K constraints.Ordered, V any] interface {
	types.Collection[K]
	Put(K, V)
	Get(K) (*V, error)
	Remove(K) (*V, error)
	HasKey(K) bool
	Keys() types.Iterator[K]
	Values() types.Iterator[V]
}
