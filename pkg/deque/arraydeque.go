package deque

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/types"
)

type ArrayDeque[T any] struct {
	data []T
	size int
}

// --- Constructor ---
func NewArrayDeque[T any](data ...T) *ArrayDeque[T] {
	return &ArrayDeque[T]{data, len(data)}
}

// --- Methods from Collection[T] ---
func (deque *ArrayDeque[T]) String() string {
	s := "ArrayDeque["
	for index, value := range deque.data {
		if index != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", value)
	}
	s += "]"
	return s
}

func (deque *ArrayDeque[T]) Cmp(other any) int {
	otherDeque, ok := other.(*ArrayDeque[T])
	if !ok {
		return -1
	}
	if deque.size != otherDeque.size {
		return deque.size - otherDeque.size
	}
	for index, value := range deque.data {
		cmp := types.Cmp(value, otherDeque.data[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (deque *ArrayDeque[T]) Iter() types.Iterator[T] {
	return newAdqIterator[T](deque)
}

func (deque *ArrayDeque[T]) Size() int {
	return deque.size
}

func (deque *ArrayDeque[T]) Empty() bool {
	return deque.size == 0
}

func (deque *ArrayDeque[T]) Clear() {
	deque.data = make([]T, 0)
	deque.size = 0
}

func (deque *ArrayDeque[T]) Contains(val T) bool {
	for _, value := range deque.data {
		if types.Cmp(value, val) == 0 {
			return true
		}
	}
	return false
}

func (deque *ArrayDeque[T]) ContainsAll(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !deque.Contains(*data) {
			return false
		}
	}
	return true
}

func (deque *ArrayDeque[T]) ContainsAny(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if deque.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from Deque[T] ---
func (deque *ArrayDeque[T]) PushFront(data ...T) {
	deque.data = append(data, deque.data...)
	deque.size += len(data)
}

func (deque *ArrayDeque[T]) PushBack(data ...T) {
	deque.data = append(deque.data, data...)
	deque.size += len(data)
}

func (deque *ArrayDeque[T]) PopFront() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	val := deque.data[0]
	deque.data = deque.data[1:]
	deque.size--
	return &val, nil
}

func (deque *ArrayDeque[T]) PopBack() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	val := deque.data[deque.size-1]
	deque.data = deque.data[:deque.size-1]
	deque.size--
	return &val, nil
}

func (deque *ArrayDeque[T]) Front() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	return &deque.data[0], nil
}

func (deque *ArrayDeque[T]) Back() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	return &deque.data[deque.size-1], nil
}

// --- Iterator ---
type adqIterator[T any] struct {
	deque *ArrayDeque[T]
	index int
}

func newAdqIterator[T any](deque *ArrayDeque[T]) *adqIterator[T] {
	return &adqIterator[T]{deque, -1}
}

func (iter *adqIterator[T]) Next() (*T, bool) {
	iter.index++
	if iter.index >= iter.deque.size {
		return nil, false
	}
	return &iter.deque.data[iter.index], true
}

func (iter *adqIterator[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}
