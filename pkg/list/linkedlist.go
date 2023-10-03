package list

import (
	"fmt"
	"log"
	"sort"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/types"
)

type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// --- Constructors ---
func NewLinkedList[T any](data ...T) *LinkedList[T] {
	list := new(LinkedList[T])
	list.Append(data...)
	return list
}

// --- Methods from Collection[T] ---
func (list *LinkedList[T]) String() string {
	s := "LinkedList["
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

func (list *LinkedList[T]) Cmp(other any) int {
	otherList, ok := other.(*LinkedList[T])
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
		cmp := types.Cmp(*next, *otherNext)
		if cmp != 0 {
			return cmp
		}
		next, hasNext = iter.Next()
		otherNext, hasOtherNext = otherIter.Next()
	}
	return 0
}

func (list *LinkedList[T]) Iter() types.Iterator[T] {
	return newLklIterator[T](list)
}

func (list *LinkedList[T]) Size() int {
	return list.size
}

func (list *LinkedList[T]) Empty() bool {
	return list.size == 0
}

func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list *LinkedList[T]) Contains(data T) bool {
	return list.IndexOf(data) != -1
}

func (list *LinkedList[T]) ContainsAll(other types.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !list.Contains(*data) {
			return false
		}
	}
	return true
}

func (list *LinkedList[T]) ContainsAny(other types.Collection[T]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if list.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from List[T] ---
func (list *LinkedList[T]) Append(data ...T) {
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

func (list *LinkedList[T]) Prepend(data ...T) {
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

func (list *LinkedList[T]) Insert(index int, data T) {
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

func (list *LinkedList[T]) RemoveFirst(data T) error {
	index := list.IndexOf(data)
	if index == -1 {
		return errs.NotFound()
	}
	list.RemoveAt(index)
	return nil
}

func (list *LinkedList[T]) RemoveAll(data T) error {
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

func (list *LinkedList[T]) RemoveAt(index int) T {
	target := list.getNode(index)
	list.remove(target)
	return target.data
}

func (list *LinkedList[T]) IndexOf(data T) int {
	index := 0
	for node := list.head; node != nil; node = node.next {
		if types.Eq(node.data, data) {
			return index
		}
		index++
	}
	return -1
}

func (list *LinkedList[T]) LastIndexOf(data T) int {
	index := list.size - 1
	for node := list.tail; node != nil; node = node.prev {
		if types.Eq(node.data, data) {
			return index
		}
		index--
	}
	return -1
}

func (list *LinkedList[T]) Get(index int) T {
	return list.getNode(index).data
}

func (list *LinkedList[T]) Set(index int, data T) {
	index = list.sanify(index)
	list.getNode(index).data = data
}

func (list *LinkedList[T]) Slice(from, to int) List[T] {
	list.sanify(from)
	list.sanify(to)
	if from > to {
		from, to = to, from
	}
	if from == to {
		return NewLinkedList[T]()
	}
	if from == 0 && to == list.size {
		return list
	}
	sub := NewLinkedList[T]()
	for i := from; i < to; i++ {
		sub.Append(list.Get(i))
	}
	return sub
}

func (list *LinkedList[T]) Sublist(filter types.Filter[T]) List[T] {
	sub := NewLinkedList[T]()
	for node := list.head; node != nil; node = node.next {
		if filter(node.data) {
			sub.Append(node.data)
		}
	}
	return sub
}

func (list *LinkedList[T]) Sort(comparator types.Comparator[T]) List[T] {
	data := make([]T, list.size)
	for node, i := list.head, 0; node != nil; node, i = node.next, i+1 {
		data[i] = node.data
	}
	sort.Slice(data, func(i, j int) bool {
		return comparator(data[i], data[j]) < 0
	})
	return NewLinkedList(data...)
}

// --- Private Methods ---
func (list *LinkedList[T]) sanify(index int) int {
	if index < 0 {
		index += list.size
	}
	return index % list.size
}

func (list *LinkedList[T]) getNode(index int) *node[T] {
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

func (list *LinkedList[T]) remove(n *node[T]) {
	if n == nil {
		log.Fatal("[LinkedList] attempted to remove a nil node")
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
type lklIterator[T any] struct {
	list *LinkedList[T]
	node *node[T]
}

func newLklIterator[T any](list *LinkedList[T]) *lklIterator[T] {
	return &lklIterator[T]{list: list, node: list.head}
}

func (iter *lklIterator[T]) Next() (*T, bool) {
	if iter.node == nil {
		return nil, false
	}
	data := iter.node.data
	iter.node = iter.node.next
	return &data, true
}

func (iter *lklIterator[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}

// --- LinkedList Node ---
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
