package tree

import (
	"github.com/luverolla/lexgo/pkg/deque"
	"github.com/luverolla/lexgo/pkg/types"
)

type AVLTree[T any] struct {
	root *avlNode[T]
	size int
}

// --- Constructor ---
func NewAVLTree[T any]() *AVLTree[T] {
	return &AVLTree[T]{nil, 0}
}

// --- Methods from Collection[T] ---
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
	return t.root.contains(val)
}

func (t *AVLTree[T]) ContainsAll(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !t.Contains(*data) {
			return false
		}
	}
	return true
}

func (t *AVLTree[T]) ContainsAny(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if t.Contains(*data) {
			return true
		}
	}
	return false
}

func (t *AVLTree[T]) Iter() types.Iterator[T] {
	return t.InOrder()
}

// --- Methods from BSTree[T] ---
func (t *AVLTree[T]) Insert(val T) {
	t.root = t.root.insert(val)
	t.size++
}

func (t *AVLTree[T]) Remove(val T) {
	t.root = t.root.remove(val)
	t.size--
}

func (t *AVLTree[T]) Min() T {
	return t.root.min()
}

func (t *AVLTree[T]) Max() T {
	return t.root.max()
}

func (t *AVLTree[T]) Pred(val T) T {
	return t.root.pred(val)
}

func (t *AVLTree[T]) Succ(val T) T {
	return t.root.succ(val)
}

func (t *AVLTree[T]) PreOrder() types.Iterator[T] {
	return newAVLPreOrderIterator(t)
}

func (t *AVLTree[T]) InOrder() types.Iterator[T] {
	return newAVLInOrderIterator(t)
}

func (t *AVLTree[T]) PostOrder() types.Iterator[T] {
	return newAVLPostOrderIterator(t)
}

// --- Node struct and methods ---
type avlNode[T any] struct {
	val   T
	left  *avlNode[T]
	right *avlNode[T]
}

func (n *avlNode[T]) insert(val T) *avlNode[T] {
	if n == nil {
		return &avlNode[T]{val, nil, nil}
	}
	switch {
	case types.Cmp(val, n.val) < 0:
		n.left = n.left.insert(val)
	case types.Cmp(val, n.val) > 0:
		n.right = n.right.insert(val)
	}
	return n.rebalance()
}

func (n *avlNode[T]) remove(val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case types.Cmp(val, n.val) < 0:
		n.left = n.left.remove(val)
	case types.Cmp(val, n.val) > 0:
		n.right = n.right.remove(val)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		n.val = n.right.min()
		n.right = n.right.remove(n.val)
	}
	return n.rebalance()
}

func (n *avlNode[T]) contains(val T) bool {
	if n == nil {
		return false
	}
	switch {
	case types.Cmp(val, n.val) < 0:
		return n.left.contains(val)
	case types.Cmp(val, n.val) > 0:
		return n.right.contains(val)
	}
	return true
}

func (n *avlNode[T]) min() T {
	if n.left == nil {
		return n.val
	}
	return n.left.min()
}

func (n *avlNode[T]) max() T {
	if n.right == nil {
		return n.val
	}
	return n.right.max()
}

func (n *avlNode[T]) pred(val T) T {
	if n == nil {
		return val
	}
	switch {
	case types.Cmp(val, n.val) <= 0:
		return n.left.pred(val)
	default:
		return n.right.pred(val)
	}
}

func (n *avlNode[T]) succ(val T) T {
	if n == nil {
		return val
	}
	switch {
	case types.Cmp(val, n.val) < 0:
		return n.left.succ(val)
	default:
		return n.right.succ(val)
	}
}

func (n *avlNode[T]) height() int {
	if n == nil {
		return 0
	}
	return 1 + max(n.left.height(), n.right.height())
}

func (n *avlNode[T]) balanceFactor() int {
	if n == nil {
		return 0
	}
	return n.left.height() - n.right.height()
}

func (n *avlNode[T]) rotateLeft() *avlNode[T] {
	x := n.right
	n.right = x.left
	x.left = n
	return x
}

func (n *avlNode[T]) rotateRight() *avlNode[T] {
	x := n.left
	n.left = x.right
	x.right = n
	return x
}

func (n *avlNode[T]) rebalance() *avlNode[T] {
	bf := n.balanceFactor()
	switch {
	case bf < -1:
		if n.right.balanceFactor() > 0 {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	case bf > 1:
		if n.left.balanceFactor() < 0 {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// --- Iterators ---
type avlPreOrderIterator[T any] struct {
	tree  *AVLTree[T]
	stack deque.Deque[avlNode[T]]
}

func newAVLPreOrderIterator[T any](tree *AVLTree[T]) *avlPreOrderIterator[T] {
	iter := &avlPreOrderIterator[T]{tree, deque.New[avlNode[T]](deque.ADQ)}
	if tree.root != nil {
		iter.stack.PushFront(*tree.root)
	}
	return iter
}

func (iter *avlPreOrderIterator[T]) Next() (*T, bool) {
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

func (iter *avlPreOrderIterator[T]) Each(f func(T)) {
	for node, ok := iter.Next(); ok; node, ok = iter.Next() {
		f(*node)
	}
}

type avlInOrderIterator[T any] struct {
	tree  *AVLTree[T]
	stack deque.Deque[avlNode[T]]
}

func newAVLInOrderIterator[T any](tree *AVLTree[T]) *avlInOrderIterator[T] {
	iter := &avlInOrderIterator[T]{tree, deque.New[avlNode[T]](deque.ADQ)}
	node := tree.root
	for node != nil {
		iter.stack.PushFront(*node)
		node = node.left
	}
	return iter
}

func (iter *avlInOrderIterator[T]) Next() (*T, bool) {
	if iter.stack.Empty() {
		return nil, false
	}
	node, _ := iter.stack.PopFront()
	if node.right != nil {
		iter.stack.PushFront(*node.right)
	}
	return &node.val, true
}

func (iter *avlInOrderIterator[T]) Each(f func(T)) {
	for node, ok := iter.Next(); ok; node, ok = iter.Next() {
		f(*node)
	}
}

type avlPostOrderIterator[T any] struct {
	tree  *AVLTree[T]
	stack deque.Deque[avlNode[T]]
}

func newAVLPostOrderIterator[T any](tree *AVLTree[T]) *avlPostOrderIterator[T] {
	iter := &avlPostOrderIterator[T]{tree, deque.New[avlNode[T]](deque.ADQ)}
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

func (iter *avlPostOrderIterator[T]) Next() (*T, bool) {
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

func (iter *avlPostOrderIterator[T]) Each(f func(T)) {
	for node, ok := iter.Next(); ok; node, ok = iter.Next() {
		f(*node)
	}
}
