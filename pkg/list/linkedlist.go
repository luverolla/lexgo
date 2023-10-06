package list

import (
	"fmt"
	"log"
	"sort"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/tau"
)

// List implemented with a doubly linked list
type LkdList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Creates a new list implemented with a doubly linked list
func Lkd[T any](data ...T) *LkdList[T] {
	list := new(LkdList[T])
	list.Append(data...)
	return list
}

// --- Methods from Collection[T] ---
func (list *LkdList[T]) String() string {
	s := "Linked["
	iter := list.Iter()
	next, hasNext := iter.Next()
	for hasNext {
		s += fmt.Sprintf("%v", *next)
		next, hasNext = iter.Next()
		if hasNext {
			s += ","
		}
	}
	s += "]"
	return s
}

func (list *LkdList[T]) Cmp(other any) int {
	otherList, ok := other.(*LkdList[T])
	if !ok {
		return -1
	}
	if list.size != otherList.size {
		return list.size - otherList.size
	}
	iter := list.Iter()
	otherIter := otherList.Iter()
	next, hasNext := iter.Next()
	otherNext, hasOtherNext := otherIter.Next()
	for hasNext && hasOtherNext {
		cmp := tau.Cmp(*next, *otherNext)
		if cmp != 0 {
			return cmp
		}
		next, hasNext = iter.Next()
		otherNext, hasOtherNext = otherIter.Next()
	}
	return 0
}

func (list *LkdList[T]) Iter() tau.Iterator[T] {
	return newLklIter[T](list)
}

func (list *LkdList[T]) Size() int {
	return list.size
}

func (list *LkdList[T]) Empty() bool {
	return list.size == 0
}

func (list *LkdList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *LkdList[T]) Contains(data T) bool {
	return list.IndexOf(data) != -1
}

func (list *LkdList[T]) ContainsAll(other tau.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !list.Contains(*data) {
			return false
		}
	}
	return true
}

func (list *LkdList[T]) ContainsAny(other tau.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if list.Contains(*data) {
			return true
		}
	}
	return false
}

func (list *LkdList[T]) Clone() tau.Collection[T] {
	return list.Slice(0, list.size)
}

// --- Methods from IdxedCollection[T] ---
func (list *LkdList[T]) Get(index int) (*T, error) {
	if list.Empty() {
		return nil, errs.Empty()
	}

	index = list.sanify(index)
	return &list.getNode(index).data, nil
}

func (list *LkdList[T]) Set(index int, data T) {
	index = list.sanify(index)
	list.getNode(index).data = data
}

func (list *LkdList[T]) Insert(index int, data T) {
	index = index % list.size
	if index < 0 {
		index += list.size
	}
	if index == 0 {
		list.Prepend(data)
	} else if index == list.size {
		list.Append(data)
	} else {
		tgt := list.getNode(index)
		newNode := &node[T]{data: data}
		tgt.prev.append(newNode)
		newNode.append(tgt)
	}
	list.size++
}

func (list *LkdList[T]) RemoveAt(index int) (*T, error) {
	if list.Empty() {
		return nil, errs.Empty()
	}
	index = list.sanify(index)
	target := list.getNode(index)
	list.remove(target)
	return &target.data, nil
}

func (list *LkdList[T]) IndexOf(data T) int {
	index := 0
	for node := list.head; node != nil; node = node.next {
		if tau.Eq(node.data, data) {
			return index
		}
		index++
	}
	return -1
}

func (list *LkdList[T]) LastIndexOf(data T) int {
	index := list.size - 1
	for node := list.tail; node != nil; node = node.prev {
		if tau.Eq(node.data, data) {
			return index
		}
		index--
	}
	return -1
}

func (list *LkdList[T]) Swap(i, j int) {
	i = list.sanify(i)
	j = list.sanify(j)
	if i == j {
		return
	}
	nodeI := list.getNode(i)
	nodeJ := list.getNode(j)
	nodeI.data, nodeJ.data = nodeJ.data, nodeI.data
}

func (list *LkdList[T]) Slice(start, end int) tau.IdxedCollection[T] {
	if list.Empty() || start == end {
		return Lkd[T]()
	}
	var actStart = list.sanify(start)
	var actEnd = list.sanify(end)

	if actStart > actEnd {
		actStart, actEnd = actEnd, actStart
	}

	sub := Lkd[T]()
	for i := actStart; i < actEnd; i++ {
		sub.Append(list.getNode(i).data)
	}
	return sub
}

// --- Methods from List[T] ---
func (list *LkdList[T]) Append(data ...T) {
	for _, value := range data {
		newTail := &node[T]{data: value}
		if list.tail == nil {
			list.tail = newTail
			list.head = newTail
		} else {
			list.tail.append(newTail)
			list.tail = newTail
		}
	}
	list.size += len(data)
}

func (list *LkdList[T]) Prepend(data ...T) {
	for _, value := range data {
		newHead := &node[T]{data: value}
		if list.head == nil {
			list.head = newHead
			list.tail = newHead
		} else {
			list.head.prepend(newHead)
			list.head = newHead
		}
	}
	list.size += len(data)
}

func (list *LkdList[T]) RemoveFirst(data T) error {
	index := list.IndexOf(data)
	if index == -1 {
		return errs.NotFound()
	}
	list.RemoveAt(index)
	return nil
}

func (list *LkdList[T]) RemoveAll(data T) error {
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

func (list *LkdList[T]) Sublist(filter tau.Filter[T]) colls.List[T] {
	sub := Lkd[T]()
	for node := list.head; node != nil; node = node.next {
		if filter(node.data) {
			sub.Append(node.data)
		}
	}
	return sub
}

func (list *LkdList[T]) Sort(comparator tau.Comparator[T]) colls.List[T] {
	data := make([]T, list.size)
	for node, i := list.head, 0; node != nil; node, i = node.next, i+1 {
		data[i] = node.data
	}
	sort.Slice(data, func(i, j int) bool {
		return comparator(data[i], data[j]) < 0
	})
	return Lkd(data...)
}

// --- Private Methods ---
func (list *LkdList[T]) sanify(index int) int {
	if index < 0 {
		index += list.size
	}
	return index % list.size
}

func (list *LkdList[T]) getNode(index int) *node[T] {
	if index < 0 || index >= list.size {
		return nil
	}
	if index == 0 {
		return list.head
	}
	if index == list.size-1 {
		return list.tail
	}
	node := list.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

func (list *LkdList[T]) remove(n *node[T]) {
	if n == nil {
		log.Fatal("[Linked] attempted to remove a nil node")
		return
	}
	if n.prev == nil {
		list.head = n.next
		list.head.prev = nil
	} else {
		n.prev.next = n.next
		if n.next != nil {
			n.next.prev = n.prev
		}
	}
	if n.next == nil {
		list.tail = n.prev
		list.tail.next = nil
	} else {
		n.next.prev = n.prev
		if n.prev != nil {
			n.prev.next = n.next
		}
	}
	list.size--
}

// --- Iterator ---
type lklIter[T any] struct {
	list *LkdList[T]
	node *node[T]
}

func newLklIter[T any](list *LkdList[T]) *lklIter[T] {
	return &lklIter[T]{list: list, node: list.head}
}

func (iter *lklIter[T]) Next() (*T, bool) {
	if iter.node == nil {
		return nil, false
	}
	data := iter.node.data
	iter.node = iter.node.next
	return &data, true
}

func (iter *lklIter[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}

// --- Linked Node ---
type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

func (n *node[T]) append(next *node[T]) {
	n.next = next
	next.prev = n
}

func (n *node[T]) prepend(prev *node[T]) {
	n.prev = prev
	prev.next = n
}
