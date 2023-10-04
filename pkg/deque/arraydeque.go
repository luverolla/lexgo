package deque

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/types"
	"github.com/luverolla/lexgo/pkg/uni"
)

type Array[T any] struct {
	data []T
	size int
}

// --- Constructor ---
func NewArray[T any](data ...T) *Array[T] {
	return &Array[T]{data, len(data)}
}

// --- Methods from Collection[T] ---
func (deque *Array[T]) String() string {
	s := "Array["
	for index, value := range deque.data {
		if index != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", value)
	}
	s += "]"
	return s
}

func (deque *Array[T]) Cmp(other any) int {
	otherDeque, ok := other.(*Array[T])
	if !ok {
		return -1
	}
	if deque.size != otherDeque.size {
		return deque.size - otherDeque.size
	}
	for index, value := range deque.data {
		cmp := uni.Cmp(value, otherDeque.data[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (deque *Array[T]) Iter() types.Iterator[T] {
	return deque.FIFOIter()
}

func (deque *Array[T]) Size() int {
	return deque.size
}

func (deque *Array[T]) Empty() bool {
	return deque.size == 0
}

func (deque *Array[T]) Clear() {
	deque.data = make([]T, 0)
	deque.size = 0
}

func (deque *Array[T]) Contains(val T) bool {
	for _, value := range deque.data {
		if uni.Cmp(value, val) == 0 {
			return true
		}
	}
	return false
}

func (deque *Array[T]) ContainsAll(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !deque.Contains(*data) {
			return false
		}
	}
	return true
}

func (deque *Array[T]) ContainsAny(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if deque.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from Deque[T] ---
func (deque *Array[T]) PushFront(data ...T) {
	deque.data = append(data, deque.data...)
	deque.size += len(data)
}

func (deque *Array[T]) PushBack(data ...T) {
	deque.data = append(deque.data, data...)
	deque.size += len(data)
}

func (deque *Array[T]) PopFront() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	val := deque.data[0]
	deque.data = deque.data[1:]
	deque.size--
	return &val, nil
}

func (deque *Array[T]) PopBack() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	val := deque.data[deque.size-1]
	deque.data = deque.data[:deque.size-1]
	deque.size--
	return &val, nil
}

func (deque *Array[T]) Front() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	return &deque.data[0], nil
}

func (deque *Array[T]) Back() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	return &deque.data[deque.size-1], nil
}

func (deque *Array[T]) FIFOIter() types.Iterator[T] {
	return newAdqIterator[T](deque, false)
}

func (deque *Array[T]) LIFOIter() types.Iterator[T] {
	return newAdqIterator[T](deque, true)
}

// --- Iterator ---
type adqIterator[T any] struct {
	deque *Array[T]
	lifo  bool
}

func newAdqIterator[T any](deque *Array[T], lifo bool) *adqIterator[T] {
	return &adqIterator[T]{deque, lifo}
}

func (iter *adqIterator[T]) Next() (*T, bool) {
	if iter.deque.size == 0 {
		return nil, false
	}
	var val T
	if iter.lifo {
		val = iter.deque.data[iter.deque.size-1]
		iter.deque.data = iter.deque.data[:iter.deque.size-1]
	} else {
		val = iter.deque.data[0]
		iter.deque.data = iter.deque.data[1:]
	}
	iter.deque.size--
	return &val, true
}

func (iter *adqIterator[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}
