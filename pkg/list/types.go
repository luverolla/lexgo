package list

import "github.com/luverolla/lexgo/pkg/types"

type List[T any] interface {
	types.Collection[T]
	Append(...T)
	Prepend(...T)
	Insert(int, T)
	RemoveFirst(T) error
	RemoveAll(T) error
	RemoveAt(int) T
	IndexOf(T) int
	LastIndexOf(T) int
	Slice(int, int) List[T]
	Sort(types.Comparator[T]) List[T]
	Sublist(types.Filter[T]) List[T]
}
