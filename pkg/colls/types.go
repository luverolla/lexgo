// This package contains interfaces [tau.Cmp] for various collections.
// The generic type parameter T allows any time that fits into the
// [tau.Cmp] comparing function, therefore limited, at run-time, by
//   - types under [constraints.Ordered] constraint
//   - types implementing [tau.Base]
package colls

import (
	"github.com/luverolla/lexgo/pkg/tau"
)

// Generic list with index access
type List[T any] interface {
	tau.IdxedCollection[T]
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
	Sort(tau.Comparator[T]) List[T]
	// Returns a new list containing the elements for which the given filter returns true
	// It makes a copy of the list, so the original one is not modified
	Sublist(tau.Filter[T]) List[T]
}

// Generic (D)ouble-(e)nded (que)ue
// It allows both FIFO (queue-like) and LIFO (stack-like) access
type Deque[T any] interface {
	tau.Collection[T]
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
	FIFOIter() tau.Iterator[T]
	// Returns an iterator that iterates over the queue in LIFO order
	LIFOIter() tau.Iterator[T]
}

// Generic node for a binary tree
// It is allowed to be nil, hence must be implemented as a pointer
type BSTreeNode[T any] interface {
	tau.Box[T]
	// Returns the left child node or nil if it does not exist
	Left() BSTreeNode[T]
	// Returns the right child node or nil if it does not exist
	Right() BSTreeNode[T]
}

// Generic binary search tree
type BSTree[T any] interface {
	tau.Collection[T]
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
	PreOrder() tau.Iterator[T]
	// Returns an iterator that iterates over the tree in in-order
	InOrder() tau.Iterator[T]
	// Returns an iterator that iterates over the tree in post-order
	PostOrder() tau.Iterator[T]
}

// Generic map
// The key MUST be a primary type (under [constraints.Ordered])
// while the type parameter can be really anything
type Map[K any, V any] interface {
	tau.Collection[K]
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
	Keys() tau.Iterator[K]
	// Returns an iterator that iterates over the values of the map
	Values() tau.Iterator[V]
}

// Generic graph node
type GraphNode[T any] tau.Box[T]

// Generic graph edge.
//
// It's suitable for both directed and undirected edges.
// The first value of the underlying [Pair] is the source node,
// while the second one is the destination node.
//
// The implementation may contain a direction flag or one edge
// for each direction. For unweighted graphs, the weight method
type GraphEdge[T any] interface {
	tau.Pair[GraphNode[T], GraphNode[T]]
	Weight() float32
}

// Generic interface for graph-like structures.
//
// It allows both directed and undirected edges as well as
// both weighted and unweighted edges.
//
// For directed edges, the implementor may choose to add
// a direction flag or to create two edges (one for each direction).
// For unweigthted edges, the [tau.GraphNode.Weight] method must be
// implemented anyway, but it can be written to return a common value
// such as zero or one.
type Graph[T any] interface {
	tau.Collection[T]
	// Returns the node containing the given value
	// Returns nil if the value does not exist in the graph
	GetNode(T) GraphNode[T]
	// Add a new node with the given value
	AddNode(T) GraphNode[T]
	// Removes the node containing the given value
	// Returns an error if the node is not found
	RemoveNode(T) error
	// Returns true if the graph has an edge between the two given values
	HasEdge(T, T) bool
	// Returns true if the graph has a directed edge between the two given values
	HasDirEdge(T, T) bool
	// Adds an undirected edge between the two given values
	// Returns an error if an edge already exists or if the nodes
	// do not exist in the graph
	AddEdge(T, T) error
	// Adds a directed edge between the two given values
	// Returns an error if the directed edge already exists or if the nodes
	// do not exist in the graph
	AddDirEdge(T, T) error
	// Removes the edge between the two given values
	// Returns an error if either the edge or the nodes
	// do not exist in the graph
	RemoveEdge(T, T) error
	// Removes the directed edge between the two given values
	// Returns an error if either a directed edge or the nodes
	// do not exist in the graph
	RemoveDirEdge(T, T) error
	// Returns an iterator that iterates over the nodes of the graph
	Nodes() tau.Iterator[T]
	// Returns an iterator that iterates over the edges of the graph
	Edges() tau.Iterator[T]
	// Returns an iterator that iterates over the directed edges of the graph
	DirEdges() tau.Iterator[T]
}
