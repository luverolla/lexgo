package tau

// Generic list with index access
type List[T any] interface {
	IdxedColl[T]
	Append(...T)
	Prepend(...T)
	// Removes the first occurrence of the given value
	// Returns an error if the value is not found
	RemoveFirst(T) error
	// Removes all the occurrences of the given value
	// Returns an error if the value is not found
	RemoveAll(T) error
	// Sorts the list using the given comparator
	// A copy of the list is made, so the original list is not modified
	Sort(Comparator[T]) List[T]
	// Returns a new list containing the elements for which the given filter returns true
	// It makes a copy of the list, so the original one is not modified
	Sublist(Filter[T]) List[T]
}

// Generic (D)ouble-(e)nded (que)ue
// It allows both FIFO (queue-like) and LIFO (stack-like) access
type Deque[T any] interface {
	Collection[T]
	// Adds the given value to the front of the queue
	PushFront(...T)
	// Adds the given value to the back of the queue
	PushBack(...T)
	// Removes the first value from the queue and returns it
	// Returns an error if the queue is empty
	PopFront() (*T, error)
	// Removes the last value from the queue and returns it
	// Returns an error if the queue is empty
	PopBack() (*T, error)
	// Returns the first value from the queue without removing it
	// Returns an error if the queue is empty
	Front() (*T, error)
	// Returns the last value from the queue without removing it
	// Returns an error if the queue is empty
	Back() (*T, error)
	// Returns an iterator that iterates over the queue in FIFO order
	FIFOIter() Iterator[T]
	// Returns an iterator that iterates over the queue in LIFO order
	LIFOIter() Iterator[T]
}

// Generic node for a binary tree
// It is allowed to be nil, hence must be implemented as a pointer
type BSTreeNode[T any] interface {
	Box[T]
	// Returns the left child node or nil if it does not exist
	Left() BSTreeNode[T]
	// Returns the right child node or nil if it does not exist
	Right() BSTreeNode[T]
}

// Generic binary search tree
type BSTree[T any] interface {
	Collection[T]
	// Returns the node containing the given value
	Get(T) BSTreeNode[T]
	// Returns the root node
	Root() BSTreeNode[T]
	// Inserts a new node with the given value
	// The return value IS NOT the inserted node, but the root node
	// The root node can change after the insertion (e.g in case of rotations)
	Insert(T) BSTreeNode[T]
	// Removes the node containing the given value
	// The return value IS NOT the removed node, but the root node
	// The root node can change after the removal (e.g in case of rotations)
	Remove(T) BSTreeNode[T]
	// Returns the node containing the minimum value in the tree
	Min() BSTreeNode[T]
	// Returns the node containing the maximum value in the tree
	Max() BSTreeNode[T]
	// Returns the node containing the predecessor of the given value
	Pred(T) BSTreeNode[T]
	// Returns the node containing the successor of the given value
	Succ(T) BSTreeNode[T]
	// Returns an iterator that iterates over the tree in pre-order
	PreOrder() Iterator[T]
	// Returns an iterator that iterates over the tree in in-order
	InOrder() Iterator[T]
	// Returns an iterator that iterates over the tree in post-order
	PostOrder() Iterator[T]
}

// Generic map
// The key MUST be a primary type (under [constraints.Ordered])
// while the type parameter can be really anything
type Map[K any, V any] interface {
	Collection[K]
	// Associates the given value with the given key
	Put(K, V)
	// Returns the value associated with the given key
	// Returns an error if the key is not found
	Get(K) (*V, error)
	// Removes the value associated with the given key
	// Returns an error if the key is not found or if the map is empty
	Remove(K) (*V, error)
	// Returns true if the map contains the given key
	HasKey(K) bool
	// Returns an iterator that iterates over the keys of the map
	Keys() Iterator[K]
	// Returns an iterator that iterates over the values of the map
	Values() Iterator[V]
}
