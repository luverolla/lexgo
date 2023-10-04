package tree

import (
	"fmt"
	"log"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/deque"
	"github.com/luverolla/lexgo/pkg/types"
	"github.com/luverolla/lexgo/pkg/uni"
)

type AVL[T any] struct {
	root *avlNode[T]
	size int
}

// --- Constructor ---
func NewAVL[T any]() *AVL[T] {
	return &AVL[T]{nil, 0}
}

// --- Methods from Collection[T] ---
func (t *AVL[T]) String() string {
	s := "AVL["
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

func (t *AVL[T]) Cmp(other any) int {
	otherTree, ok := other.(*AVL[T])
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
		cmp := uni.Cmp(*next, *otherNext)
		if cmp != 0 {
			return cmp
		}
		next, hasNext = iter.Next()
		otherNext, hasOtherNext = otherIter.Next()
	}
	return 0
}

func (t *AVL[T]) Size() int {
	return t.size
}

func (t *AVL[T]) Empty() bool {
	return t.size == 0
}

func (t *AVL[T]) Clear() {
	t.root = nil
	t.size = 0
}

func (t *AVL[T]) Contains(val T) bool {
	return t.root.contains(val)
}

func (t *AVL[T]) ContainsAll(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !t.Contains(*data) {
			return false
		}
	}
	return true
}

func (t *AVL[T]) ContainsAny(c types.Collection[T]) bool {
	iter := c.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if t.Contains(*data) {
			return true
		}
	}
	return false
}

func (t *AVL[T]) Iter() types.Iterator[T] {
	return t.InOrder()
}

// --- Methods from BSTree[T] ---
func (t *AVL[T]) Get(val T) colls.BSTreeNode[T] {
	node := t.root.getNode(val)
	if node == nil {
		return nil
	}
	return node
}

func (t *AVL[T]) Root() colls.BSTreeNode[T] {
	return t.root
}

func (t *AVL[T]) Insert(val T) colls.BSTreeNode[T] {
	t.root = t.root.insert(val)
	t.size++
	return t.root
}

func (t *AVL[T]) Remove(val T) colls.BSTreeNode[T] {
	t.root = t.root.remove(val)
	if t.root == nil {
		return nil
	}
	t.size--
	return t.root
}

func (t *AVL[T]) Min() colls.BSTreeNode[T] {
	return t.root.min()
}

func (t *AVL[T]) Max() colls.BSTreeNode[T] {
	return t.root.max()
}

func (t *AVL[T]) Pred(val T) colls.BSTreeNode[T] {
	return t.root.pred(val)
}

func (t *AVL[T]) Succ(val T) colls.BSTreeNode[T] {
	return t.root.succ(val)
}

func (t *AVL[T]) PreOrder() types.Iterator[T] {
	return newAVLPreOrderIterator(t)
}

func (t *AVL[T]) InOrder() types.Iterator[T] {
	return newAVLInOrderIterator(t)
}

func (t *AVL[T]) PostOrder() types.Iterator[T] {
	return newAVLPostOrderIterator(t)
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

func (n *avlNode[T]) Left() colls.BSTreeNode[T] {
	return n.left
}

func (n *avlNode[T]) Right() colls.BSTreeNode[T] {
	return n.right
}

func (n *avlNode[T]) getNode(val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case uni.Cmp(val, n.val) < 0:
		return n.left.getNode(val)
	case uni.Cmp(val, n.val) > 0:
		return n.right.getNode(val)
	}
	return n
}

func (n *avlNode[T]) insert(val T) *avlNode[T] {
	if n == nil {
		return &avlNode[T]{val, nil, nil}
	}
	switch {
	case uni.Cmp(val, n.val) < 0:
		n.left = n.left.insert(val)
	case uni.Cmp(val, n.val) > 0:
		n.right = n.right.insert(val)
	}
	return n.rebalance()
}

func (n *avlNode[T]) remove(val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case uni.Cmp(val, n.val) < 0:
		n.left = n.left.remove(val)
	case uni.Cmp(val, n.val) > 0:
		n.right = n.right.remove(val)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		n.val = n.right.min().val
		n.right = n.right.remove(n.val)
	}
	return n.rebalance()
}

func (n *avlNode[T]) contains(val T) bool {
	return n.getNode(val) != nil
}

func (n *avlNode[T]) min() *avlNode[T] {
	if n.left == nil {
		return n
	}
	return n.left.min()
}

func (n *avlNode[T]) max() *avlNode[T] {
	if n.right == nil {
		return n
	}
	return n.right.max()
}

func (n *avlNode[T]) pred(val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case uni.Cmp(val, n.val) < 0:
		return n.left.pred(val)
	case uni.Cmp(val, n.val) > 0:
		return n.right.pred(val)
	default:
		if n.left != nil {
			return n.left.max()
		}
	}
	return n
}

func (n *avlNode[T]) succ(val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case uni.Cmp(val, n.val) < 0:
		return n.left.succ(val)
	case uni.Cmp(val, n.val) > 0:
		return n.right.succ(val)
	default:
		if n.right != nil {
			return n.right.min()
		}
	}
	return n
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
	tree  *AVL[T]
	stack colls.Deque[avlNode[T]]
}

func newAVLPreOrderIterator[T any](tree *AVL[T]) *avlPreOrderIterator[T] {
	iter := &avlPreOrderIterator[T]{tree, deque.NewArray[avlNode[T]]()}
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
	tree  *AVL[T]
	stack colls.Deque[avlNode[T]]
}

func newAVLInOrderIterator[T any](tree *AVL[T]) *avlInOrderIterator[T] {
	iter := &avlInOrderIterator[T]{tree, deque.NewArray[avlNode[T]]()}
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
	tree  *AVL[T]
	stack colls.Deque[avlNode[T]]
}

func newAVLPostOrderIterator[T any](tree *AVL[T]) *avlPostOrderIterator[T] {
	iter := &avlPostOrderIterator[T]{tree, deque.NewArray[avlNode[T]]()}
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
