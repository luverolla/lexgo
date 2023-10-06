package table

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Unsorted map implemented with a hash table
type HshMap[K any, V any] struct {
	inner []hshEntry[K, V]
	size  int
}

// --- Constructor ---
func Hsh[K any, V any]() *HshMap[K, V] {
	return &HshMap[K, V]{make([]hshEntry[K, V], 0), 0}
}

// --- Methods from Collection[MapEntry[K, V]] ---
func (table *HshMap[K, V]) String() string {
	s := "HshMap["
	for index, value := range table.inner {
		if index != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", value)
	}
	s += "]"
	return s
}

func (table *HshMap[K, V]) Cmp(other any) int {
	otherTable, ok := other.(*HshMap[K, V])
	if !ok {
		return -1
	}
	if table.size != otherTable.size {
		return table.size - otherTable.size
	}
	for index, value := range table.inner {
		cmp := tau.Cmp(value, otherTable.inner[index])
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (table *HshMap[K, V]) Iter() tau.Iterator[K] {
	return newHshKeyIter[K](table)
}

func (table *HshMap[K, V]) Size() int {
	return table.size
}

func (table *HshMap[K, V]) Empty() bool {
	return table.size == 0
}

func (table *HshMap[K, V]) Clear() {
	table.inner = make([]hshEntry[K, V], 0)
	table.size = 0
}

func (table *HshMap[K, V]) Contains(val K) bool {
	for _, hshEntry := range table.inner {
		if tau.Eq(hshEntry.key, val) {
			return true
		}
	}
	return false
}

func (table *HshMap[K, V]) ContainsAll(other tau.Collection[K]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !table.Contains(*data) {
			return false
		}
	}
	return true
}

func (table *HshMap[K, V]) ContainsAny(other tau.Collection[K]) bool {
	iter := other.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if table.Contains(*data) {
			return true
		}
	}
	return false
}

// --- Methods from Map[K, V] ---
func (table *HshMap[K, V]) Get(key K) (*V, error) {
	index := table.indexOf(key)
	if index == -1 {
		return nil, errs.NotFound()
	}
	return &table.inner[index].value, nil
}

func (table *HshMap[K, V]) Put(key K, value V) {
	index := table.indexOf(key)
	if index == -1 {
		table.inner = append(table.inner, hshEntry[K, V]{key, value})
		table.size++
	} else {
		table.inner[index].value = value
	}
}

func (table *HshMap[K, V]) HasKey(key K) bool {
	return table.Contains(key)
}

func (table *HshMap[K, V]) Remove(key K) (*V, error) {
	index := table.indexOf(key)
	if index == -1 {
		return nil, errs.NotFound()
	}
	value := table.inner[index].value
	table.inner = append(table.inner[:index], table.inner[index+1:]...)
	table.size--
	return &value, nil
}

func (table *HshMap[K, V]) Keys() tau.Iterator[K] {
	return newHshKeyIter[K](table)
}

func (table *HshMap[K, V]) Values() tau.Iterator[V] {
	return newHshValueIter[K](table)
}

// --- Iterators ---
type hshKeyIter[K any, V any] struct {
	table *HshMap[K, V]
	index int
}

func newHshKeyIter[K any, V any](table *HshMap[K, V]) *hshKeyIter[K, V] {
	return &hshKeyIter[K, V]{table, 0}
}

func (iter *hshKeyIter[K, V]) Next() (*K, bool) {
	if iter.index >= iter.table.size {
		return nil, false
	}
	hshEntry := iter.table.inner[iter.index]
	iter.index++
	return &hshEntry.key, true
}

func (iter *hshKeyIter[K, V]) Each(f func(K)) {
	for _, value := range iter.table.inner {
		f(value.key)
	}
}

type hshValueIter[K any, V any] struct {
	table *HshMap[K, V]
	index int
}

func newHshValueIter[K any, V any](table *HshMap[K, V]) *hshValueIter[K, V] {
	return &hshValueIter[K, V]{table, 0}
}

func (iter *hshValueIter[K, V]) Next() (*V, bool) {
	if iter.index >= iter.table.size {
		return nil, false
	}
	hshEntry := iter.table.inner[iter.index]
	iter.index++
	return &hshEntry.value, true
}

func (iter *hshValueIter[K, V]) Each(f func(V)) {
	for _, value := range iter.table.inner {
		f(value.value)
	}
}

// --- Private methods ---
func (table *HshMap[K, V]) indexOf(key K) int {
	for i := 0; i < table.size; i++ {
		index := table.hash(key, i) % uint(table.size)
		if tau.Eq(table.inner[index].key, key) {
			return int(index)
		}
	}
	return -1
}

func (table *HshMap[K, V]) hash1(key K) uint {
	return uint(tau.Hash(key) % uint32(table.size))
}

func (table *HshMap[K, V]) hash2(key K) uint {
	PRIME := uint32(7)
	return uint(PRIME - (tau.Hash(key) % PRIME))
}

func (table *HshMap[K, V]) hash(key K, i int) uint {
	return table.hash1(key) + uint(i)*table.hash2(key)
}

// --- Entry ---
type hshEntry[K any, V any] struct {
	key   K
	value V
}

func (entry hshEntry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, entry.value)
}

func (entry hshEntry[K, V]) Cmp(other any) int {
	otherEntry, ok := other.(hshEntry[K, V])
	if !ok {
		return -1
	}
	cmp := tau.Cmp(entry.key, otherEntry.key)
	if cmp != 0 {
		return cmp
	}
	return tau.Cmp(entry.value, otherEntry.value)
}
