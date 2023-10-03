package types

import "fmt"

type Base interface {
	fmt.Stringer
	Cmp(any) int
}

type Iterator[T any] interface {
	Next() (*T, bool)
	Each(func(T))
}

type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Collection[T any] interface {
	Base
	Iterable[T]
	Size() int
	Empty() bool
	Clear()
	Contains(T) bool
	ContainsAll(C Collection[T]) bool
	ContainsAny(C Collection[T]) bool
}

type Filter[T any] func(T, ...any) bool
type Comparator[T any] func(T, T) int
