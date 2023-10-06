// This package contains implementation for the interface [tau.BSTree]
package tree

import (
	"fmt"
	"log"

	"github.com/luverolla/lexgo/pkg/deque"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Binary search tree implemented with an AVL tree
type AVLTree[T any] struct {
	root *avlNode[T]
	size int
}

// Creates a new binary search tree implemented with an AVL tree
func AVL[T any]() *AVLTree[T] {
	return &AVLTree[T]{nil, 0}
}

// --- Methods from Collection[T] ---
func (t *AVLTree[T]) String() string {
	s := "AVLTree["
	iter := t.Iter()
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

func (t *AVLTree[T]) Cmp(other any) int {
	otherTree, ok := other.(*AVLTree[T])
	if !ok {
		log.Fatal("ERROR: [tree.AVL] right hand side of comparison is not an AVL tree")
	}
	if t.size != otherTree.size {
		return t.size - otherTree.size
	}
	iter := t.Iter()
	otherIter := otherTree.Iter()
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

func (t *AVLTree[T]) Size() int {
	return t.size
}

func (t *AVLTree[T]) Empty() bool {
	return t.size == 0
}

func (t *AVLTree[T]) Clear() {
	t.root = nil
	t.size = 0
}

func (t *AVLTree[T]) Contains(val T) bool {
	return t.contains(t.root, val)
}

func (t *AVLTree[T]) ContainsAll(c tau.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !t.Contains(*data) {
			return false
		}
	}
	return true
}

func (t *AVLTree[T]) ContainsAny(c tau.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if t.Contains(*data) {
			return true
		}
	}
	return false
}

func (t *AVLTree[T]) Iter() tau.Iterator[T] {
	return t.InOrder()
}

// --- Methods from BSTree[T] ---
func (t *AVLTree[T]) Get(val T) tau.BSTreeNode[T] {
	node := t.getNode(t.root, val)
	if node == nil {
		return nil
	}
	return node
}

func (t *AVLTree[T]) Root() tau.BSTreeNode[T] {
	return t.root
}

func (t *AVLTree[T]) Insert(val T) tau.BSTreeNode[T] {
	t.root = t.insert(t.root, val)
	t.size++
	return t.root
}

func (t *AVLTree[T]) Remove(val T) tau.BSTreeNode[T] {
	if t.root == nil {
		return nil
	}
	t.root = t.remove(t.root, val)
	t.size--
	return t.root
}

func (t *AVLTree[T]) Min() tau.BSTreeNode[T] {
	return t.min(t.root)
}

func (t *AVLTree[T]) Max() tau.BSTreeNode[T] {
	return t.max(t.root)
}

func (t *AVLTree[T]) Pred(val T) tau.BSTreeNode[T] {
	return t.pred(t.root, val)
}

func (t *AVLTree[T]) Succ(val T) tau.BSTreeNode[T] {
	return t.succ(t.root, val)
}

func (t *AVLTree[T]) PreOrder() tau.Iterator[T] {
	return newAVLPreOrderIter(t)
}

func (t *AVLTree[T]) InOrder() tau.Iterator[T] {
	return newAVLInOrderIter(t)
}

func (t *AVLTree[T]) PostOrder() tau.Iterator[T] {
	return newAVLPostOrderIter(t)
}

// --- Node struct and methods ---
type avlNode[T any] struct {
	val   T
	left  *avlNode[T]
	right *avlNode[T]
}

func (n *avlNode[T]) Value() T {
	return n.val
}

func (n *avlNode[T]) Left() tau.BSTreeNode[T] {
	return n.left
}

func (n *avlNode[T]) Right() tau.BSTreeNode[T] {
	return n.right
}

func (t *AVLTree[T]) getNode(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case tau.Cmp(val, n.val) < 0:
		return t.getNode(n.left, val)
	case tau.Cmp(val, n.val) > 0:
		return t.getNode(n.right, val)
	}
	return n
}

func (t *AVLTree[T]) insert(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return &avlNode[T]{val, nil, nil}
	}
	switch {
	case tau.Cmp(val, n.val) < 0:
		n.left = t.insert(n.left, val)
	case tau.Cmp(val, n.val) > 0:
		n.right = t.insert(n.right, val)
	}
	return t.rebalance(n)
}

func (t *AVLTree[T]) remove(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case tau.Cmp(val, n.val) < 0:
		n.left = t.remove(n.left, val)
	case tau.Cmp(val, n.val) > 0:
		n.right = t.remove(n.right, val)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		n.val = t.min(n.right).val
		n.right = t.remove(n.right, n.val)
	}
	return t.rebalance(n)
}

func (t *AVLTree[T]) contains(n *avlNode[T], val T) bool {
	return t.getNode(n, val) != nil
}

func (t *AVLTree[T]) min(n *avlNode[T]) *avlNode[T] {
	if n.left == nil {
		return n
	}
	return t.min(n.left)
}

func (t *AVLTree[T]) max(n *avlNode[T]) *avlNode[T] {
	if n.right == nil {
		return n
	}
	return t.max(n.right)
}

func (t *AVLTree[T]) pred(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case tau.Cmp(val, n.val) < 0:
		return t.pred(n.left, val)
	case tau.Cmp(val, n.val) > 0:
		return t.pred(n.right, val)
	default:
		if n.left != nil {
			return t.max(n.left)
		}
	}
	return n
}

func (t *AVLTree[T]) succ(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case tau.Cmp(val, n.val) < 0:
		return t.succ(n.left, val)
	case tau.Cmp(val, n.val) > 0:
		return t.succ(n.right, val)
	default:
		if n.right != nil {
			return t.min(n.right)
		}
	}
	return n
}

func (t *AVLTree[T]) height(n *avlNode[T]) int {
	if n == nil {
		return 0
	}
	return 1 + max(t.height(n.left), t.height(n.right))
}

func (t *AVLTree[T]) balanceFactor(n *avlNode[T]) int {
	if n == nil {
		return 0
	}
	return t.height(n.left) - t.height(n.right)
}

func (t *AVLTree[T]) rotateLeft(n *avlNode[T]) *avlNode[T] {
	x := n.right
	n.right = x.left
	x.left = n
	return x
}

func (t *AVLTree[T]) rotateRight(n *avlNode[T]) *avlNode[T] {
	x := n.left
	n.left = x.right
	x.right = n
	return x
}

func (t *AVLTree[T]) rebalance(n *avlNode[T]) *avlNode[T] {
	bf := t.balanceFactor(n)
	switch {
	case bf < -1:
		if t.balanceFactor(n.right) > 0 {
			n.right = t.rotateRight(n.right)
		}
		return t.rotateLeft(n)
	case bf > 1:
		if t.balanceFactor(n.left) < 0 {
			n.left = t.rotateLeft(n.left)
		}
		return t.rotateRight(n)
	}
	return n
}

// --- Iterator ---
type avlPreOrderIter[T any] struct {
	tree  *AVLTree[T]
	stack tau.Deque[avlNode[T]]
}

func newAVLPreOrderIter[T any](tree *AVLTree[T]) *avlPreOrderIter[T] {
	iter := &avlPreOrderIter[T]{tree, deque.Arr[avlNode[T]]()}
	if tree.root != nil {
		iter.stack.PushFront(*tree.root)
	}
	return iter
}

func (iter *avlPreOrderIter[T]) Next() (*T, bool) {
	if iter.stack.Empty() {
		return nil, false
	}
	node, _ := iter.stack.PopFront()
	if node.right != nil {
		iter.stack.PushFront(*node.right)
	}
	if node.left != nil {
		iter.stack.PushFront(*node.left)
	}
	return &node.val, true
}

func (iter *avlPreOrderIter[T]) Each(f func(T)) {
	for node, ok := iter.Next(); ok; node, ok = iter.Next() {
		f(*node)
	}
}

type avlInOrderIter[T any] struct {
	tree  *AVLTree[T]
	stack tau.Deque[avlNode[T]]
}

func newAVLInOrderIter[T any](tree *AVLTree[T]) *avlInOrderIter[T] {
	iter := &avlInOrderIter[T]{tree, deque.Arr[avlNode[T]]()}
	node := tree.root
	for node != nil {
		iter.stack.PushFront(*node)
		node = node.left
	}
	return iter
}

func (iter *avlInOrderIter[T]) Next() (*T, bool) {
	if iter.stack.Empty() {
		return nil, false
	}
	node, _ := iter.stack.PopFront()
	if node.right != nil {
		iter.stack.PushFront(*node.right)
	}
	return &node.val, true
}

func (iter *avlInOrderIter[T]) Each(f func(T)) {
	for node, ok := iter.Next(); ok; node, ok = iter.Next() {
		f(*node)
	}
}

type avlPostOrderIter[T any] struct {
	tree  *AVLTree[T]
	stack tau.Deque[avlNode[T]]
}

func newAVLPostOrderIter[T any](tree *AVLTree[T]) *avlPostOrderIter[T] {
	iter := &avlPostOrderIter[T]{tree, deque.Arr[avlNode[T]]()}
	node := tree.root
	for node != nil {
		iter.stack.PushFront(*node)
		if node.left != nil {
			node = node.left
		} else {
			node = node.right
		}
	}
	return iter
}

func (iter *avlPostOrderIter[T]) Next() (*T, bool) {
	if iter.stack.Empty() {
		return nil, false
	}
	node, _ := iter.stack.PopFront()
	if !iter.stack.Empty() {
		parent, _ := iter.stack.Front()
		if node == parent.left && parent.right != nil {
			iter.stack.PushFront(*parent.right)
			node = parent.right
			for node.left != nil {
				iter.stack.PushFront(*node.left)
				node = node.left
			}
		}
	}
	return &node.val, true
}

func (iter *avlPostOrderIter[T]) Each(f func(T)) {
	for node, ok := iter.Next(); ok; node, ok = iter.Next() {
		f(*node)
	}
}
