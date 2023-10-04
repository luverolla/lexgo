package list

import (
	"fmt"
	"sort"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/types"
	"github.com/luverolla/lexgo/pkg/uni"
)

type Array[T any] struct {
	data []T
}

// --- Constructors ---
func NewArray[T any](data ...T) *Array[T] {
	list := new(Array[T])
	list.data = make([]T, len(data))
	copy(list.data, data)
	return list
}

// --- Methods from Collection[T] ---
func (list *Array[T]) String() string {
	s := "Array["
	for index, value := range list.data {
		if index != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", value)
	}
	s += "]"
	return s
}

func (list *Array[T]) Cmp(other any) int {
	otherList, ok := other.(*Array[T])
	if !ok {
		return -1
	}
	if len(list.data) != len(otherList.data) {
		return len(list.data) - len(otherList.data)
	}
	for index, value := range list.data {
		cmp := uni.Cmp(value, otherList.data[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (list *Array[T]) Iter() types.Iterator[T] {
	return newArlIterator[T](list)
}

func (list *Array[T]) Size() int {
	return len(list.data)
}

func (list *Array[T]) Empty() bool {
	return len(list.data) == 0
}

func (list *Array[T]) Clear() {
	list.data = make([]T, 0)
}

func (list *Array[T]) Contains(data T) bool {
	return list.IndexOf(data) != -1
}

func (list *Array[T]) ContainsAll(other types.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !list.Contains(*data) {
			return false
		}
	}
	return true
}

func (list *Array[T]) ContainsAny(other types.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if list.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from List[T] ---
func (list *Array[T]) Get(index int) (*T, error) {
	if list.Empty() {
		return nil, errs.Empty()
	}
	index = list.sanify(index)
	return &list.data[index], nil
}

func (list *Array[T]) Set(index int, data T) {
	index = list.sanify(index)
	list.data[index] = data
}

func (list *Array[T]) Append(data ...T) {
	list.data = append(list.data, data...)
}

func (list *Array[T]) Prepend(data ...T) {
	list.data = append(data, list.data...)
}

func (list *Array[T]) Insert(index int, data T) {
	list.data = append(list.data[:index], append([]T{data}, list.data[index:]...)...)
}

func (list *Array[T]) RemoveFirst(data T) error {
	index := list.IndexOf(data)
	if index == -1 {
		return errs.NotFound()
	}
	list.RemoveAt(index)
	return nil
}

func (list *Array[T]) RemoveAll(data T) error {
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

func (list *Array[T]) RemoveAt(index int) (*T, error) {
	if list.Empty() {
		return nil, errs.Empty()
	}
	index = list.sanify(index)
	data := list.data[index]
	list.data = append(list.data[:index], list.data[index+1:]...)
	return &data, nil
}

func (list *Array[T]) IndexOf(data T) int {
	for index, value := range list.data {
		if uni.Eq(value, data) {
			return index
		}
	}
	return -1
}

func (list *Array[T]) LastIndexOf(data T) int {
	for index := len(list.data) - 1; index >= 0; index-- {
		if uni.Eq(list.data[index], data) {
			return index
		}
	}
	return -1
}

func (list *Array[T]) Slice(start, end int) colls.List[T] {
	return NewArray(list.data[start:end]...)
}

func (list *Array[T]) Sort(comparator types.Comparator[T]) colls.List[T] {
	data := make([]T, len(list.data))
	copy(data, list.data)
	sort.Slice(data, func(i, j int) bool {
		return comparator(data[i], data[j]) < 0
	})
	return NewArray(data...)
}

// create a new list with the data that satisfies the filter function
func (list *Array[T]) Sublist(filter types.Filter[T]) colls.List[T] {
	data := make([]T, 0)
	for _, value := range list.data {
		if filter(value) {
			data = append(data, value)
		}
	}
	return NewArray(data...)
}

// --- Private methods ---
func (list *Array[T]) sanify(index int) int {
	if index < 0 {
		index += len(list.data)
	}
	return index % len(list.data)
}

// --- Iterator struct and constructor ---
type arlIterator[T any] struct {
	list  *Array[T]
	index int
}

func newArlIterator[T any](list *Array[T]) *arlIterator[T] {
	iterator := new(arlIterator[T])
	iterator.list = list
	iterator.index = -1
	return iterator
}

// --- Methods from Iterator[T] ---
func (iterator *arlIterator[T]) Next() (*T, bool) {
	iterator.index++
	if iterator.index >= len(iterator.list.data) {
		return nil, false
	}
	return &iterator.list.data[iterator.index], true
}

func (iterator *arlIterator[T]) Each(f func(T)) {
	for data, ok := iterator.Next(); ok; data, ok = iterator.Next() {
		f(*data)
	}
}
