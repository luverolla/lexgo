package table

import (
	"fmt"
	"log"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/gx"
	"github.com/luverolla/lexgo/pkg/tree"
	"github.com/luverolla/lexgo/pkg/types"
	"golang.org/x/exp/constraints"
)

type AVLMap[K constraints.Ordered, V any] struct {
	tree *tree.AVLTree[avlEntry[K, V]]
}

// --- Constructor ---
func AVL[K constraints.Ordered, V any]() *AVLMap[K, V] {
	return &AVLMap[K, V]{tree.AVL[avlEntry[K, V]]()}
}

// --- Methods from Collection[MapEntry[K, V]] ---
func (table *AVLMap[K, V]) String() string {
	return table.tree.String()
}

func (table *AVLMap[K, V]) Cmp(other any) int {
	return table.tree.Cmp(other)
}

func (table *AVLMap[K, V]) Iter() types.Iterator[K] {
	return newAvlKeyIter[K](table)
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

func (table *AVLMap[K, V]) ContainsAll(c types.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if !table.Contains(*next) {
			return false
		}
	}
	return true
}

func (table *AVLMap[K, V]) ContainsAny(c types.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if table.Contains(*next) {
			return true
		}
	}
	return false
}

// --- Methods from Map[K, V] ---
func (table *AVLMap[K, V]) Get(key K) (*V, error) {
	if table.Empty() {
		return nil, errs.Empty()
	}
	entry := avlEntry[K, V]{key, nil}
	node := table.tree.Get(entry)
	if node == nil {
		return nil, errs.NotFound()
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
	if node == nil {
		return nil, errs.NotFound()
	}
	table.tree.Remove(entry)
	return node.Value().value, nil
}

func (table *AVLMap[K, V]) Keys() types.Iterator[K] {
	return table.Iter()
}

func (table *AVLMap[K, V]) Values() types.Iterator[V] {
	return newAvlValueIter[K](table)
}

// --- Iterator ---
type avlKeyIter[K constraints.Ordered, V any] struct {
	inner types.Iterator[avlEntry[K, V]]
}

func newAvlKeyIter[K constraints.Ordered, V any](table *AVLMap[K, V]) *avlKeyIter[K, V] {
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

type avlValueIter[K constraints.Ordered, V any] struct {
	inner types.Iterator[avlEntry[K, V]]
}

func newAvlValueIter[K constraints.Ordered, V any](table *AVLMap[K, V]) *avlValueIter[K, V] {
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
type avlEntry[K constraints.Ordered, V any] struct {
	key   K
	value *V
}

func (entry avlEntry[K, V]) Cmp(other any) int {
	oth, ok := other.(avlEntry[K, V])
	if !ok {
		log.Fatalf("ERROR: [table.AVLTreeMap] right hand side of comparison is not an AVL entry")
	}
	return gx.Cmp(entry.key, oth.key)
}

func (entry avlEntry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, *entry.value)
}
