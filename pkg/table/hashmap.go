package table

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/types"
	"github.com/luverolla/lexgo/pkg/uni"
	"golang.org/x/exp/constraints"
)

// HashMap, implements Map[K, V] with double hashing

type HashMap[K constraints.Ordered, V any] struct {
	inner []hmentry[K, V]
	size  int
}

// --- Constructor ---
func NewHashMap[K constraints.Ordered, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{make([]hmentry[K, V], 0), 0}
}

// --- Methods from Collection[MapEntry[K, V]] ---
func (table *HashMap[K, V]) String() string {
	s := "HashMap["
	for index, value := range table.inner {
		if index != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", value)
	}
	s += "]"
	return s
}

func (table *HashMap[K, V]) Cmp(other any) int {
	otherTable, ok := other.(*HashMap[K, V])
	if !ok {
		return -1
	}
	if table.size != otherTable.size {
		return table.size - otherTable.size
	}
	for index, value := range table.inner {
		cmp := uni.Cmp(value, otherTable.inner[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (table *HashMap[K, V]) Iter() types.Iterator[K] {
	return newHshKeyIter[K](table)
}

func (table *HashMap[K, V]) Size() int {
	return table.size
}

func (table *HashMap[K, V]) Empty() bool {
	return table.size == 0
}

func (table *HashMap[K, V]) Clear() {
	table.inner = make([]hmentry[K, V], 0)
	table.size = 0
}

func (table *HashMap[K, V]) Contains(val K) bool {
	for _, hmentry := range table.inner {
		if hmentry.key == val {
			return true
		}
	}
	return false
}

func (table *HashMap[K, V]) ContainsAll(other types.Collection[K]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !table.Contains(*data) {
			return false
		}
	}
	return true
}

func (table *HashMap[K, V]) ContainsAny(other types.Collection[K]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if table.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from Map[K, V] ---
func (table *HashMap[K, V]) Get(key K) (*V, error) {
	index := table.indexOf(key)
	if index == -1 {
		return nil, errs.NotFound()
	}
	return &table.inner[index].value, nil
}

func (table *HashMap[K, V]) Put(key K, value V) {
	index := table.indexOf(key)
	if index == -1 {
		table.inner = append(table.inner, hmentry[K, V]{key, value})
		table.size++
	} else {
		table.inner[index].value = value
	}
}

func (table *HashMap[K, V]) HasKey(key K) bool {
	return table.Contains(key)
}

func (table *HashMap[K, V]) Remove(key K) (*V, error) {
	index := table.indexOf(key)
	if index == -1 {
		return nil, errs.NotFound()
	}
	value := table.inner[index].value
	table.inner = append(table.inner[:index], table.inner[index+1:]...)
	table.size--
	return &value, nil
}

func (table *HashMap[K, V]) Keys() types.Iterator[K] {
	return newHshKeyIter[K](table)
}

func (table *HashMap[K, V]) Values() types.Iterator[V] {
	return newHshValueIter[K](table)
}

// --- Iterators ---
type hshKeyIter[K constraints.Ordered, V any] struct {
	table *HashMap[K, V]
	index int
}

func newHshKeyIter[K constraints.Ordered, V any](table *HashMap[K, V]) *hshKeyIter[K, V] {
	return &hshKeyIter[K, V]{table, 0}
}

func (iter *hshKeyIter[K, V]) Next() (*K, bool) {
	if iter.index >= iter.table.size {
		return nil, false
	}
	hmentry := iter.table.inner[iter.index]
	iter.index++
	return &hmentry.key, true
}

func (iter *hshKeyIter[K, V]) Each(f func(K)) {
	for _, value := range iter.table.inner {
		f(value.key)
	}
}

type hshValueIter[K constraints.Ordered, V any] struct {
	table *HashMap[K, V]
	index int
}

func newHshValueIter[K constraints.Ordered, V any](table *HashMap[K, V]) *hshValueIter[K, V] {
	return &hshValueIter[K, V]{table, 0}
}

func (iter *hshValueIter[K, V]) Next() (*V, bool) {
	if iter.index >= iter.table.size {
		return nil, false
	}
	hmentry := iter.table.inner[iter.index]
	iter.index++
	return &hmentry.value, true
}

func (iter *hshValueIter[K, V]) Each(f func(V)) {
	for _, value := range iter.table.inner {
		f(value.value)
	}
}

// --- Private methods ---
func (table *HashMap[K, V]) indexOf(key K) int {
	for i := 0; i < table.size; i++ {
		index := table.hash(key, i) % uint(table.size)
		if table.inner[index].key == key {
			return int(index)
		}
	}
	return -1
}

func (table *HashMap[K, V]) hash1(key K) uint {
	return uint(uni.Hash(key) % uint32(table.size))
}

func (table *HashMap[K, V]) hash2(key K) uint {
	PRIME := uint32(7)
	return uint(PRIME - (uni.Hash(key) % PRIME))
}

func (table *HashMap[K, V]) hash(key K, i int) uint {
	return table.hash1(key) + uint(i)*table.hash2(key)
}

// --- Entry ---
type hmentry[K constraints.Ordered, V any] struct {
	key   K
	value V
}

func (entry hmentry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, entry.value)
}

func (entry hmentry[K, V]) Cmp(other any) int {
	otherEntry, ok := other.(hmentry[K, V])
	if !ok {
		return -1
	}
	cmp := uni.Cmp(entry.key, otherEntry.key)
	if cmp != 0 {
		return cmp
	}
	return uni.Cmp(entry.value, otherEntry.value)
}
