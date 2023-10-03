package tree

import "github.com/luverolla/lexgo/pkg/types"

type BSTree[T any] interface {
	types.Collection[T]
	Insert(T)
	Remove(T)
	Min() T
	Max() T
	Pred(T) T
	Succ(T) T
	PreOrder() types.Iterator[T]
	InOrder() types.Iterator[T]
	PostOrder() types.Iterator[T]
}
