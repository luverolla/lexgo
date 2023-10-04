package table

import (
	"fmt"
	"log"

	"github.com/luverolla/lexgo/pkg/errs"
	"github.com/luverolla/lexgo/pkg/tree"
	"github.com/luverolla/lexgo/pkg/types"
	"github.com/luverolla/lexgo/pkg/uni"
	"golang.org/x/exp/constraints"
)

type AVLTreeMap[K constraints.Ordered, V any] struct {
	tree *tree.AVL[avlEntry[K, V]]
}

// --- Constructor ---
func NewAVLTreeMap[K constraints.Ordered, V any]() *AVLTreeMap[K, V] {
	return &AVLTreeMap[K, V]{tree.NewAVL[avlEntry[K, V]]()}
}

// --- Methods from Collection[MapEntry[K, V]] ---
func (table *AVLTreeMap[K, V]) String() string {
	return table.tree.String()
}

func (table *AVLTreeMap[K, V]) Cmp(other any) int {
	return table.tree.Cmp(other)
}

func (table *AVLTreeMap[K, V]) Iter() types.Iterator[K] {
	return newAvlKeyIter[K](table)
}

func (table *AVLTreeMap[K, V]) Size() int {
	return table.tree.Size()
}

func (table *AVLTreeMap[K, V]) Empty() bool {
	return table.tree.Empty()
}

func (table *AVLTreeMap[K, V]) Clear() {
	table.tree.Clear()
}

func (table *AVLTreeMap[K, V]) Contains(val K) bool {
	return table.tree.Contains(avlEntry[K, V]{val, nil})
}

func (table *AVLTreeMap[K, V]) ContainsAll(c types.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if !table.Contains(*next) {
			return false
		}
	}
	return true
}

func (table *AVLTreeMap[K, V]) ContainsAny(c types.Collection[K]) bool {
	iter := c.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		if table.Contains(*next) {
			return true
		}
	}
	return false
}

// --- Methods from Map[K, V] ---
func (table *AVLTreeMap[K, V]) Get(key K) (*V, error) {
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

func (table *AVLTreeMap[K, V]) HasKey(key K) bool {
	return table.Contains(key)
}

func (table *AVLTreeMap[K, V]) Put(key K, value V) {
	entry := avlEntry[K, V]{key, &value}
	table.tree.Insert(entry)
}

func (table *AVLTreeMap[K, V]) Remove(key K) (*V, error) {
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

func (table *AVLTreeMap[K, V]) Keys() types.Iterator[K] {
	return table.Iter()
}

func (table *AVLTreeMap[K, V]) Values() types.Iterator[V] {
	return newAvlValueIter[K](table)
}

// --- Iterator ---
type avlKeyIter[K constraints.Ordered, V any] struct {
	inner types.Iterator[avlEntry[K, V]]
}

func newAvlKeyIter[K constraints.Ordered, V any](table *AVLTreeMap[K, V]) *avlKeyIter[K, V] {
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

func newAvlValueIter[K constraints.Ordered, V any](table *AVLTreeMap[K, V]) *avlValueIter[K, V] {
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
	return uni.Cmp(entry.key, oth.key)
}

func (entry avlEntry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, *entry.value)
}
