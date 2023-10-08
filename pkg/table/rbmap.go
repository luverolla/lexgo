// This package contains implementation for the interface [tau.Map]
package table

import (
	"fmt"
	"log"
	"reflect"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/tau"
	"github.com/luverolla/lexgo/pkg/tree"
)

// Sorted map implemented with an RB tree
type RBMap[K any, V any] struct {
	tree *tree.RBTree[rbEntry[K, V]]
}

// Creates a new map implemented with an RB tree
func RB[K any, V any]() *RBMap[K, V] {
	return &RBMap[K, V]{tree.RB[rbEntry[K, V]]()}
}

// --- Methods from Collection[MapEntry[K, V]] ---
func (table *RBMap[K, V]) String() string {
	s := "RBMap["
	iter := table.tree.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		s += fmt.Sprintf("%v", *next)
		if hasNext {
			s += ","
		}
	}
	s += "]"
	return s
}

func (table *RBMap[K, V]) Cmp(other any) int {
	rbOther, ok := other.(*RBMap[K, V])
	if !ok {
		panic(fmt.Sprintf("ERROR: [RBMap.Cmp] %v is not a *RBMap", other))
	}

	if table.Size() != rbOther.Size() {
		return table.Size() - rbOther.Size()
	}

	iter := table.Iter()
	otherIter := rbOther.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		otherNext, _ := otherIter.Next()
		cmp := tau.Cmp(*next, *otherNext)
		if cmp != 0 {
			return cmp
		}
	}

	return 0
}

func (table *RBMap[K, V]) Iter() tau.Iterator[K] {
	return newRBKeyIter[K](table)
}

func (table *RBMap[K, V]) Size() int {
	return table.tree.Size()
}

func (table *RBMap[K, V]) Empty() bool {
	return table.tree.Empty()
}

func (table *RBMap[K, V]) Clear() {
	table.tree.Clear()
}

func (table *RBMap[K, V]) Contains(val K) bool {
	return table.tree.Contains(rbEntry[K, V]{val, nil})
}

func (table *RBMap[K, V]) ContainsAll(c tau.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if !table.Contains(*next) {
			return false
		}
	}
	return true
}

func (table *RBMap[K, V]) ContainsAny(c tau.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if table.Contains(*next) {
			return true
		}
	}
	return false
}

func (table *RBMap[K, V]) Clone() tau.Collection[K] {
	clone := RB[K, V]()
	iter := table.tree.InOrder()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		clone.Put(next.key, *next.value)
	}
	return clone
}

// --- Debug ---
func (table *RBMap[K, V]) Tree() *tree.RBTree[rbEntry[K, V]] {
	return table.tree
}

// --- Methods from Map[K, V] ---
func (table *RBMap[K, V]) Get(key K) (*V, error) {
	if table.Empty() {
		return nil, errs.Empty()
	}
	entry := rbEntry[K, V]{key, nil}
	node := table.tree.Get(entry)
	if reflect.ValueOf(node).IsNil() {
		return nil, errs.NotFound(key)
	}
	return node.Value().value, nil
}

func (table *RBMap[K, V]) HasKey(key K) bool {
	return table.Contains(key)
}

func (table *RBMap[K, V]) Put(key K, value V) {
	entry := rbEntry[K, V]{key, &value}
	table.tree.Insert(entry)
}

func (table *RBMap[K, V]) Remove(key K) (*V, error) {
	if table.Empty() {
		return nil, errs.Empty()
	}
	entry := rbEntry[K, V]{key, nil}
	node := table.tree.Get(entry)

	if node.(any) == nil {
		return nil, errs.NotFound(key)
	}

	table.tree.Remove(entry)
	return node.Value().value, nil
}

func (table *RBMap[K, V]) Keys() tau.Iterator[K] {
	return table.Iter()
}

func (table *RBMap[K, V]) Values() tau.Iterator[V] {
	return newRBValueIter[K](table)
}

// --- Iterator ---
type rbKeyIter[K any, V any] struct {
	inner tau.Iterator[rbEntry[K, V]]
}

func newRBKeyIter[K any, V any](table *RBMap[K, V]) *rbKeyIter[K, V] {
	return &rbKeyIter[K, V]{table.tree.InOrder()}
}

func (iter *rbKeyIter[K, V]) Next() (*K, bool) {
	next, hasNext := iter.inner.Next()
	if !hasNext {
		return nil, false
	}
	return &next.key, true
}

func (iter *rbKeyIter[K, V]) Each(f func(K)) {
	iter.inner.Each(func(entry rbEntry[K, V]) {
		f(entry.key)
	})
}

type rbValueIter[K any, V any] struct {
	inner tau.Iterator[rbEntry[K, V]]
}

func newRBValueIter[K any, V any](table *RBMap[K, V]) *rbValueIter[K, V] {
	return &rbValueIter[K, V]{table.tree.InOrder()}
}

func (iter *rbValueIter[K, V]) Next() (*V, bool) {
	next, hasNext := iter.inner.Next()
	if !hasNext {
		return nil, false
	}
	return next.value, true
}

func (iter *rbValueIter[K, V]) Each(f func(V)) {
	iter.inner.Each(func(entry rbEntry[K, V]) {
		f(*entry.value)
	})
}

// --- Entry ---
type rbEntry[K any, V any] struct {
	key   K
	value *V
}

func (entry rbEntry[K, V]) Cmp(other any) int {
	oth, ok := other.(rbEntry[K, V])
	if !ok {
		log.Fatalf("ERROR: [table.RBTreeMap] right hand side of comparison is not an RB entry")
	}
	return tau.Cmp(entry.key, oth.key)
}

func (entry rbEntry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, entry.value)
}
