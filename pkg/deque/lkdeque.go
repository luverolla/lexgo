package deque

import (
	"github.com/luverolla/lexgo/pkg/list"
	"github.com/luverolla/lexgo/pkg/types"
)

type LkDeque[T any] struct {
	inner *list.LkdList[T]
}

// --- Constructor ---
func Lkd[T any](data ...T) *LkDeque[T] {
	return &LkDeque[T]{list.Lkd[T](data...)}
}

// --- Methods from Collection[T] ---
func (deque *LkDeque[T]) String() string {
	return deque.inner.String()
}

func (deque *LkDeque[T]) Cmp(other any) int {
	return deque.inner.Cmp(other)
}

func (deque *LkDeque[T]) Iter() types.Iterator[T] {
	return deque.inner.Iter()
}

func (deque *LkDeque[T]) Size() int {
	return deque.inner.Size()
}

func (deque *LkDeque[T]) Empty() bool {
	return deque.inner.Empty()
}

func (deque *LkDeque[T]) Clear() {
	deque.inner.Clear()
}

func (deque *LkDeque[T]) Contains(val T) bool {
	return deque.inner.Contains(val)
}

func (deque *LkDeque[T]) ContainsAll(c types.Collection[T]) bool {
	return deque.inner.ContainsAll(c)
}

func (deque *LkDeque[T]) ContainsAny(c types.Collection[T]) bool {
	return deque.inner.ContainsAny(c)
}

// --- Methods from Deque[T] ---
func (deque *LkDeque[T]) PushFront(val T) {
	deque.inner.Prepend(val)
}

func (deque *LkDeque[T]) PushBack(val T) {
	deque.inner.Append(val)
}

func (deque *LkDeque[T]) PopFront() (*T, error) {
	return deque.inner.RemoveAt(0)
}

func (deque *LkDeque[T]) PopBack() (*T, error) {
	return deque.inner.RemoveAt(deque.Size() - 1)
}

func (deque *LkDeque[T]) Front() (*T, error) {
	return deque.inner.Get(0)
}

func (deque *LkDeque[T]) Back() (*T, error) {
	return deque.inner.Get(deque.Size() - 1)
}

func (deque *LkDeque[T]) FIFOIter() types.Iterator[T] {
	return newLdqIter[T](deque, false)
}

func (deque *LkDeque[T]) LIFOIter() types.Iterator[T] {
	return newLdqIter[T](deque, true)
}

// --- Iterator ---
type ldqIter[T any] struct {
	deque *LkDeque[T]
	lifo  bool
	index int
}

func newLdqIter[T any](deque *LkDeque[T], lifo bool) *ldqIter[T] {
	return &ldqIter[T]{deque, lifo, 0}
}

func (iter *ldqIter[T]) Next() (*T, bool) {
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

func (iter *ldqIter[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}
