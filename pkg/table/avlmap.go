// This package contains implementation for the interface [tau.Map]
package table

import (
	"fmt"
	"log"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/tau"
	"github.com/luverolla/lexgo/pkg/tree"
)

// Sorted map implemented with an AVL tree
type AVLMap[K any, V any] struct {
	tree *tree.AVLTree[avlEntry[K, V]]
}

// Creates a new map implemented with an AVL tree
func AVL[K any, V any]() *AVLMap[K, V] {
	return &AVLMap[K, V]{tree.AVL[avlEntry[K, V]]()}
}

// --- Methods from Collection[MapEntry[K, V]] ---
func (table *AVLMap[K, V]) String() string {
	s := "AVLMap{"
	iter := table.tree.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		s += fmt.Sprintf("%v", *next)
		if hasNext {
			s += ","
		}
	}
	s += "}"
	return s
}

func (table *AVLMap[K, V]) Cmp(other any) int {
	oth, ok := other.(*AVLMap[K, V])
	if !ok {
		panic(fmt.Sprintf("ERROR: [AVLMap.Cmp] %v is not a *AVLMap", other))
	}

	if table.Size() != oth.Size() {
		return table.Size() - oth.Size()
	}

	iter, otherIter := table.tree.Iter(), oth.tree.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		otherNext, _ := otherIter.Next()
		cmp := tau.Cmp(*next, *otherNext)
		if cmp != 0 {
			return cmp
		}
	}

	return 0
}

func (table *AVLMap[K, V]) Iter() tau.Iterator[K] {
	return newAVLKeyIter[K](table)
}

func (table *AVLMap[K, V]) Size() int {
	return table.tree.Size()
}

func (table *AVLMap[K, V]) Empty() bool {
	return table.tree.Empty()
}

func (table *AVLMap[K, V]) Clear() {
	table.tree.Clear()
}

func (table *AVLMap[K, V]) Contains(val K) bool {
	return table.tree.Contains(avlEntry[K, V]{val, nil})
}

func (table *AVLMap[K, V]) ContainsAll(c tau.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if !table.Contains(*next) {
			return false
		}
	}
	return true
}

func (table *AVLMap[K, V]) ContainsAny(c tau.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if table.Contains(*next) {
			return true
		}
	}
	return false
}

func (table *AVLMap[K, V]) Clone() tau.Collection[K] {
	clone := AVL[K, V]()
	iter := table.tree.InOrder()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		clone.Put(next.key, *next.value)
	}
	return clone
}

// --- Methods from Map[K, V] ---
func (table *AVLMap[K, V]) Get(key K) (*V, error) {
	if table.Empty() {
		return nil, errs.Empty()
	}
	entry := avlEntry[K, V]{key, nil}
	node := table.tree.Get(entry)
	if tau.Nil(node) {
		return nil, errs.NotFound(key)
	}
	return node.Value().value, nil
}

func (table *AVLMap[K, V]) HasKey(key K) bool {
	return table.Contains(key)
}

func (table *AVLMap[K, V]) Put(key K, value V) {
	entry := avlEntry[K, V]{key, &value}
	table.tree.Insert(entry)
}

func (table *AVLMap[K, V]) Remove(key K) (*V, error) {
	if table.Empty() {
		return nil, errs.Empty()
	}
	entry := avlEntry[K, V]{key, nil}
	node := table.tree.Get(entry)
	if tau.Nil(node) {
		return nil, errs.NotFound(key)
	}
	table.tree.Remove(entry)
	return node.Value().value, nil
}

func (table *AVLMap[K, V]) Keys() tau.Iterator[K] {
	return table.Iter()
}

func (table *AVLMap[K, V]) Values() tau.Iterator[V] {
	return newAVLValueIter[K](table)
}

// --- Iterator ---
type avlKeyIter[K any, V any] struct {
	inner tau.Iterator[avlEntry[K, V]]
}

func newAVLKeyIter[K any, V any](table *AVLMap[K, V]) *avlKeyIter[K, V] {
	return &avlKeyIter[K, V]{table.tree.InOrder()}
}

func (iter *avlKeyIter[K, V]) Next() (*K, bool) {
	next, hasNext := iter.inner.Next()
	if !hasNext {
		return nil, false
	}
	return &next.key, true
}

func (iter *avlKeyIter[K, V]) Each(f func(K)) {
	iter.inner.Each(func(entry avlEntry[K, V]) {
		f(entry.key)
	})
}

type avlValueIter[K any, V any] struct {
	inner tau.Iterator[avlEntry[K, V]]
}

func newAVLValueIter[K any, V any](table *AVLMap[K, V]) *avlValueIter[K, V] {
	return &avlValueIter[K, V]{table.tree.InOrder()}
}

func (iter *avlValueIter[K, V]) Next() (*V, bool) {
	next, hasNext := iter.inner.Next()
	if !hasNext {
		return nil, false
	}
	return next.value, true
}

func (iter *avlValueIter[K, V]) Each(f func(V)) {
	iter.inner.Each(func(entry avlEntry[K, V]) {
		f(*entry.value)
	})
}

// --- Entry ---
type avlEntry[K any, V any] struct {
	key   K
	value *V
}

func (entry avlEntry[K, V]) Cmp(other any) int {
	oth, ok := other.(avlEntry[K, V])
	if !ok {
		log.Fatalf("ERROR: [table.AVLTreeMap] right hand side of comparison is not an AVL entry")
	}
	return tau.Cmp(entry.key, oth.key)
}

func (entry avlEntry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, *entry.value)
}
