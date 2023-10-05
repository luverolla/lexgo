package list

import (
	"fmt"
	"sort"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/gx"
	"github.com/luverolla/lexgo/pkg/types"
)

type ArrList[T any] struct {
	data []T
}

// --- Constructors ---
func Arr[T any](data ...T) *ArrList[T] {
	list := new(ArrList[T])
	list.data = make([]T, len(data))
	copy(list.data, data)
	return list
}

// --- Methods from Collection[T] ---
func (list *ArrList[T]) String() string {
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

func (list *ArrList[T]) Cmp(other any) int {
	otherList, ok := other.(*ArrList[T])
	if !ok {
		return -1
	}
	if len(list.data) != len(otherList.data) {
		return len(list.data) - len(otherList.data)
	}
	for index, value := range list.data {
		cmp := gx.Cmp(value, otherList.data[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (list *ArrList[T]) Iter() types.Iterator[T] {
	return newArlIter[T](list)
}

func (list *ArrList[T]) Size() int {
	return len(list.data)
}

func (list *ArrList[T]) Empty() bool {
	return len(list.data) == 0
}

func (list *ArrList[T]) Clear() {
	list.data = make([]T, 0)
}

func (list *ArrList[T]) Contains(data T) bool {
	return list.IndexOf(data) != -1
}

func (list *ArrList[T]) ContainsAll(other types.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !list.Contains(*data) {
			return false
		}
	}
	return true
}

func (list *ArrList[T]) ContainsAny(other types.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if list.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from List[T] ---
func (list *ArrList[T]) Get(index int) (*T, error) {
	if list.Empty() {
		return nil, errs.Empty()
	}
	index = list.sanify(index)
	return &list.data[index], nil
}

func (list *ArrList[T]) Set(index int, data T) {
	index = list.sanify(index)
	list.data[index] = data
}

func (list *ArrList[T]) Append(data ...T) {
	list.data = append(list.data, data...)
}

func (list *ArrList[T]) Prepend(data ...T) {
	list.data = append(data, list.data...)
}

func (list *ArrList[T]) Insert(index int, data T) {
	list.data = append(list.data[:index], append([]T{data}, list.data[index:]...)...)
}

func (list *ArrList[T]) RemoveFirst(data T) error {
	index := list.IndexOf(data)
	if index == -1 {
		return errs.NotFound()
	}
	list.RemoveAt(index)
	return nil
}

func (list *ArrList[T]) RemoveAll(data T) error {
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

func (list *ArrList[T]) RemoveAt(index int) (*T, error) {
	if list.Empty() {
		return nil, errs.Empty()
	}
	index = list.sanify(index)
	data := list.data[index]
	list.data = append(list.data[:index], list.data[index+1:]...)
	return &data, nil
}

func (list *ArrList[T]) IndexOf(data T) int {
	for index, value := range list.data {
		if gx.Eq(value, data) {
			return index
		}
	}
	return -1
}

func (list *ArrList[T]) LastIndexOf(data T) int {
	for index := len(list.data) - 1; index >= 0; index-- {
		if gx.Eq(list.data[index], data) {
			return index
		}
	}
	return -1
}

func (list *ArrList[T]) Slice(start, end int) colls.List[T] {
	if start >= end || start == end {
		return Arr[T]()
	}

	var actStart = list.sanify(start)
	var actEnd = list.sanify(end)

	if actStart > actEnd {
		actStart, actEnd = actEnd, actStart
	}

	return Arr(list.data[actStart:actEnd]...)
}

func (list *ArrList[T]) Sort(comparator types.Comparator[T]) colls.List[T] {
	data := make([]T, len(list.data))
	copy(data, list.data)
	sort.Slice(data, func(i, j int) bool {
		return comparator(data[i], data[j]) < 0
	})
	return Arr(data...)
}

// create a new list with the data that satisfies the filter function
func (list *ArrList[T]) Sublist(filter types.Filter[T]) colls.List[T] {
	data := make([]T, 0)
	for _, value := range list.data {
		if filter(value) {
			data = append(data, value)
		}
	}
	return Arr(data...)
}

// --- Private methods ---
func (list *ArrList[T]) sanify(index int) int {
	if index < 0 {
		index += len(list.data)
	}
	return index % len(list.data)
}

// --- Iterator struct and constructor ---
type arlIter[T any] struct {
	list  *ArrList[T]
	index int
}

func newArlIter[T any](list *ArrList[T]) *arlIter[T] {
	iterator := new(arlIter[T])
	iterator.list = list
	iterator.index = -1
	return iterator
}

// --- Methods from Iterator[T] ---
func (iterator *arlIter[T]) Next() (*T, bool) {
	iterator.index++
	if iterator.index >= len(iterator.list.data) {
		return nil, false
	}
	return &iterator.list.data[iterator.index], true
}

func (iterator *arlIter[T]) Each(f func(T)) {
	for data, ok := iterator.Next(); ok; data, ok = iterator.Next() {
		f(*data)
	}
}
