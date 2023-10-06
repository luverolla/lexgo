// This package contains implementation for the interface [colls.Deque]
package deque

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Double-ended queue implemented with a dynamic array
type ArrDeque[T any] struct {
	data []T
	size int
}

// Creates a new empty deque implemented with a dynamic array
func Arr[T any](data ...T) *ArrDeque[T] {
	return &ArrDeque[T]{data, len(data)}
}

// --- Methods from Collection[T] ---
func (deque *ArrDeque[T]) String() string {
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

func (deque *ArrDeque[T]) Cmp(other any) int {
	otherDeque, ok := other.(*ArrDeque[T])
	if !ok {
		return -1
	}
	if deque.size != otherDeque.size {
		return deque.size - otherDeque.size
	}
	for index, value := range deque.data {
		cmp := tau.Cmp(value, otherDeque.data[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (deque *ArrDeque[T]) Iter() tau.Iterator[T] {
	return deque.FIFOIter()
}

func (deque *ArrDeque[T]) Size() int {
	return deque.size
}

func (deque *ArrDeque[T]) Empty() bool {
	return deque.size == 0
}

func (deque *ArrDeque[T]) Clear() {
	deque.data = make([]T, 0)
	deque.size = 0
}

func (deque *ArrDeque[T]) Contains(val T) bool {
	for _, value := range deque.data {
		if tau.Cmp(value, val) == 0 {
			return true
		}
	}
	return false
}

func (deque *ArrDeque[T]) ContainsAll(c tau.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !deque.Contains(*data) {
			return false
		}
	}
	return true
}

func (deque *ArrDeque[T]) ContainsAny(c tau.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if deque.Contains(*data) {
			return true
		}
	}
	return false
}

func (deque *ArrDeque[T]) Clone() tau.Collection[T] {
	return Arr[T](deque.data...)
}

// --- Methods from Deque[T] ---
func (deque *ArrDeque[T]) PushFront(data ...T) {
	deque.data = append(data, deque.data...)
	deque.size += len(data)
}

func (deque *ArrDeque[T]) PushBack(data ...T) {
	deque.data = append(deque.data, data...)
	deque.size += len(data)
}

func (deque *ArrDeque[T]) PopFront() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	val := deque.data[0]
	deque.data = deque.data[1:]
	deque.size--
	return &val, nil
}

func (deque *ArrDeque[T]) PopBack() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	val := deque.data[deque.size-1]
	deque.data = deque.data[:deque.size-1]
	deque.size--
	return &val, nil
}

func (deque *ArrDeque[T]) Front() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	return &deque.data[0], nil
}

func (deque *ArrDeque[T]) Back() (*T, error) {
	if deque.size == 0 {
		return nil, errs.Empty()
	}
	return &deque.data[deque.size-1], nil
}

func (deque *ArrDeque[T]) FIFOIter() tau.Iterator[T] {
	return newAdqIter[T](deque, false)
}

func (deque *ArrDeque[T]) LIFOIter() tau.Iterator[T] {
	return newAdqIter[T](deque, true)
}

// --- Iterator ---
type adqIter[T any] struct {
	deque *ArrDeque[T]
	lifo  bool
}

func newAdqIter[T any](deque *ArrDeque[T], lifo bool) *adqIter[T] {
	return &adqIter[T]{deque, lifo}
}

func (iter *adqIter[T]) Next() (*T, bool) {
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

func (iter *adqIter[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}
