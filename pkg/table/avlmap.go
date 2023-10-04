package table

import (
	"fmt"

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

func (table *AVLTreeMap[K, V]) Iter() types.Iterator[avlEntry[K, V]] {
	return table.tree.Iter()
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

func (table *AVLTreeMap[K, V]) ContainsAll(c types.Collection[avlEntry[K, V]]) bool {
	return table.tree.ContainsAll(c)
}

func (table *AVLTreeMap[K, V]) ContainsAny(c types.Collection[avlEntry[K, V]]) bool {
	return table.tree.ContainsAny(c)
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
	table.tree.Remove(entry)
	if node == nil {
		return nil, errs.NotFound()
	}
	return node.Value().value, nil
}

// --- Entry ---
type avlEntry[K constraints.Ordered, V any] struct {
	key   K
	value *V
}

func (entry *avlEntry[K, V]) Cmp(other any) int {
	return uni.Cmp(entry.key, other)
}

func (entry *avlEntry[K, V]) String() string {
	return fmt.Sprintf("(%v: %v)", entry.key, *entry.value)
}
