package list

import (
	"fmt"
	"sort"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/types"
)

type ArrayList[T any] struct {
	data []T
}

// --- Constructors ---
func NewArrayList[T any](data ...T) *ArrayList[T] {
	list := new(ArrayList[T])
	list.data = make([]T, len(data))
	copy(list.data, data)
	return list
}

// --- Methods from Collection[T] ---
func (list *ArrayList[T]) String() string {
	s := "ArrayList["
	for index, value := range list.data {
		if index != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", value)
	}
	s += "]"
	return s
}

func (list *ArrayList[T]) ValueOf() uint32 {
	return uint32(len(list.data))
}

func (list *ArrayList[T]) Iter() types.Iterator[T] {
	return newIterator[T](list)
}

func (list *ArrayList[T]) Size() int {
	return len(list.data)
}

func (list *ArrayList[T]) Empty() bool {
	return len(list.data) == 0
}

func (list *ArrayList[T]) Clear() {
	list.data = make([]T, 0)
}

func (list *ArrayList[T]) Contains(data T) bool {
	return list.IndexOf(data) != -1
}

func (list *ArrayList[T]) ContainsAll(other types.Collection[T]) bool {
	for data, ok := other.Iter().Next(); ok; data, ok = other.Iter().Next() {
		if !list.Contains(data) {
			return false
		}
	}
	return true
}

func (list *ArrayList[T]) ContainsAny(other types.Collection[T]) bool {
	for data, ok := other.Iter().Next(); ok; data, ok = other.Iter().Next() {
		if list.Contains(data) {
			return true
		}
	}
	return false
}

// --- Methods from List[T] ---
func (list *ArrayList[T]) Append(data ...T) {
	list.data = append(list.data, data...)
}

func (list *ArrayList[T]) Prepend(data ...T) {
	list.data = append(data, list.data...)
}

func (list *ArrayList[T]) Insert(index int, data T) {
	list.data = append(list.data[:index], append([]T{data}, list.data[index:]...)...)
}

func (list *ArrayList[T]) RemoveFirst(data T) error {
	index := list.IndexOf(data)
	if index == -1 {
		return errs.NotFound()
	}
	list.RemoveAt(index)
	return nil
}

func (list *ArrayList[T]) RemoveAll(data T) error {
	index := list.IndexOf(data)
	if index == -1 {
		return errs.NotFound()
	}
	for index != -1 {
		list.RemoveAt(index)
		index = list.IndexOf(data)
	}
	return nil
}

func (list *ArrayList[T]) RemoveAt(index int) T {
	data := list.data[index]
	list.data = append(list.data[:index], list.data[index+1:]...)
	return data
}

func (list *ArrayList[T]) IndexOf(data T) int {
	for index, value := range list.data {
		if types.Eq(value, data) {
			return index
		}
	}
	return -1
}

func (list *ArrayList[T]) LastIndexOf(data T) int {
	for index := len(list.data) - 1; index >= 0; index-- {
		if types.Eq(list.data[index], data) {
			return index
		}
	}
	return -1
}

func (list *ArrayList[T]) Slice(start, end int) List[T] {
	return NewArrayList(list.data[start:end]...)
}

// perform quicksort on the list's data and create a new list with the sorted data
// sort according to the comparator function
func (list *ArrayList[T]) Sort(comparator types.Comparator[T]) List[T] {
	data := make([]T, len(list.data))
	copy(data, list.data)
	sort.Slice(data, func(i, j int) bool {
		return comparator(data[i], data[j]) < 0
	})
	return NewArrayList(data...)
}

// create a new list with the data that satisfies the filter function
func (list *ArrayList[T]) Sublist(filter types.Filter[T]) List[T] {
	data := make([]T, 0)
	for _, value := range list.data {
		if filter(value) {
			data = append(data, value)
		}
	}
	return NewArrayList(data...)
}

// --- Iterator struct and constructor ---
type iterator[T any] struct {
	list  *ArrayList[T]
	index int
}

func newIterator[T any](list *ArrayList[T]) *iterator[T] {
	iterator := new(iterator[T])
	iterator.list = list
	iterator.index = -1
	return iterator
}

// --- Methods from Iterator[T] ---
func (iterator *iterator[T]) Next() (T, bool) {
	iterator.index++
	if iterator.index >= len(iterator.list.data) {
		return types.Empty.(T), false
	}
	return iterator.list.data[iterator.index], true
}

func (iterator *iterator[T]) Each(f func(T)) {
	for data, ok := iterator.Next(); ok; data, ok = iterator.Next() {
		f(data)
	}
}
