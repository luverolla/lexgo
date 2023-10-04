package deque

import (
	"github.com/luverolla/lexgo/pkg/list"
	"github.com/luverolla/lexgo/pkg/types"
)

type Linked[T any] struct {
	inner *list.Linked[T]
}

// --- Constructor ---
func NewLinked[T any](data ...T) *Linked[T] {
	return &Linked[T]{list.NewLinked[T](data...)}
}

// --- Methods from Collection[T] ---
func (deque *Linked[T]) String() string {
	return deque.inner.String()
}

func (deque *Linked[T]) Cmp(other any) int {
	return deque.inner.Cmp(other)
}

func (deque *Linked[T]) Iter() types.Iterator[T] {
	return deque.inner.Iter()
}

func (deque *Linked[T]) Size() int {
	return deque.inner.Size()
}

func (deque *Linked[T]) Empty() bool {
	return deque.inner.Empty()
}

func (deque *Linked[T]) Clear() {
	deque.inner.Clear()
}

func (deque *Linked[T]) Contains(val T) bool {
	return deque.inner.Contains(val)
}

func (deque *Linked[T]) ContainsAll(c types.Collection[T]) bool {
	return deque.inner.ContainsAll(c)
}

func (deque *Linked[T]) ContainsAny(c types.Collection[T]) bool {
	return deque.inner.ContainsAny(c)
}

// --- Methods from Deque[T] ---
func (deque *Linked[T]) PushFront(val T) {
	deque.inner.Prepend(val)
}

func (deque *Linked[T]) PushBack(val T) {
	deque.inner.Append(val)
}

func (deque *Linked[T]) PopFront() (*T, error) {
	return deque.inner.RemoveAt(0)
}

func (deque *Linked[T]) PopBack() (*T, error) {
	return deque.inner.RemoveAt(deque.Size() - 1)
}

func (deque *Linked[T]) Front() (*T, error) {
	return deque.inner.Get(0)
}

func (deque *Linked[T]) Back() (*T, error) {
	return deque.inner.Get(deque.Size() - 1)
}

func (deque *Linked[T]) FIFOIter() types.Iterator[T] {
	return newLdqIterator[T](deque, false)
}

func (deque *Linked[T]) LIFOIter() types.Iterator[T] {
	return newLdqIterator[T](deque, true)
}

// --- Iterator ---
type ldqIterator[T any] struct {
	deque *Linked[T]
	lifo  bool
	index int
}

func newLdqIterator[T any](deque *Linked[T], lifo bool) *ldqIterator[T] {
	return &ldqIterator[T]{deque, lifo, 0}
}

func (iter *ldqIterator[T]) Next() (*T, bool) {
	if iter.index == iter.deque.Size() {
		return nil, false
	}
	var actIdx int
	if iter.lifo {
		actIdx = iter.deque.Size() - iter.index - 1
	}
	val, _ := iter.deque.inner.Get(actIdx)
	iter.index++
	return val, true
}

func (iter *ldqIterator[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}
