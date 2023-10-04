package colls

import (
	"github.com/luverolla/lexgo/pkg/types"
	"golang.org/x/exp/constraints"
)

type List[T any] interface {
	types.Collection[T]
	Get(int) (*T, error)
	Set(int, T)
	Append(...T)
	Prepend(...T)
	Insert(int, T)
	RemoveFirst(T) error
	RemoveAll(T) error
	RemoveAt(int) (*T, error)
	IndexOf(T) int
	LastIndexOf(T) int
	Slice(int, int) List[T]
	Sort(types.Comparator[T]) List[T]
	Sublist(types.Filter[T]) List[T]
}

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

type BSTreeNode[T any] interface {
	Value() T
	Left() BSTreeNode[T]
	Right() BSTreeNode[T]
}

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

type Map[K constraints.Ordered, V any] interface {
	types.Collection[K]
	Put(K, V)
	Get(K) (*V, error)
	Remove(K) (*V, error)
	HasKey(K) bool
	Keys() types.Iterator[K]
	Values() types.Iterator[V]
}
