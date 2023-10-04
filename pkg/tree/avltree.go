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
	cmp  types.Comparator[T]
	size int
}

// --- Constructors ---
func NewAVL[T any]() *AVL[T] {
	return &AVL[T]{nil, nil, 0}
}

func NewAVLCmp[T any](cmp types.Comparator[T]) *AVL[T] {
	return &AVL[T]{nil, cmp, 0}
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
		cmp := t.compare(*next, *otherNext)
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
	return t.contains(t.root, val)
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
	node := t.getNode(t.root, val)
	if node == nil {
		return nil
	}
	return node
}

func (t *AVL[T]) Root() colls.BSTreeNode[T] {
	return t.root
}

func (t *AVL[T]) Insert(val T) colls.BSTreeNode[T] {
	t.root = t.insert(t.root, val)
	t.size++
	return t.root
}

func (t *AVL[T]) Remove(val T) colls.BSTreeNode[T] {
	if t.root == nil {
		return nil
	}
	t.root = t.remove(t.root, val)
	t.size--
	return t.root
}

func (t *AVL[T]) Min() colls.BSTreeNode[T] {
	return t.min(t.root)
}

func (t *AVL[T]) Max() colls.BSTreeNode[T] {
	return t.max(t.root)
}

func (t *AVL[T]) Pred(val T) colls.BSTreeNode[T] {
	return t.pred(t.root, val)
}

func (t *AVL[T]) Succ(val T) colls.BSTreeNode[T] {
	return t.succ(t.root, val)
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

// --- Private Methods ---
func (t *AVL[T]) compare(a, b T) int {
	if t.cmp == nil {
		return uni.Cmp(a, b)
	}
	return t.cmp(a, b)
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

func (t *AVL[T]) getNode(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case t.compare(val, n.val) < 0:
		return t.getNode(n.left, val)
	case t.compare(val, n.val) > 0:
		return t.getNode(n.right, val)
	}
	return n
}

func (t *AVL[T]) insert(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return &avlNode[T]{val, nil, nil}
	}
	switch {
	case t.compare(val, n.val) < 0:
		n.left = t.insert(n.left, val)
	case t.compare(val, n.val) > 0:
		n.right = t.insert(n.right, val)
	}
	return t.rebalance(n)
}

func (t *AVL[T]) remove(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case t.compare(val, n.val) < 0:
		n.left = t.remove(n.left, val)
	case t.compare(val, n.val) > 0:
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

func (t *AVL[T]) contains(n *avlNode[T], val T) bool {
	return t.getNode(n, val) != nil
}

func (t *AVL[T]) min(n *avlNode[T]) *avlNode[T] {
	if n.left == nil {
		return n
	}
	return t.min(n.left)
}

func (t *AVL[T]) max(n *avlNode[T]) *avlNode[T] {
	if n.right == nil {
		return n
	}
	return t.max(n.right)
}

func (t *AVL[T]) pred(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case t.compare(val, n.val) < 0:
		return t.pred(n.left, val)
	case t.compare(val, n.val) > 0:
		return t.pred(n.right, val)
	default:
		if n.left != nil {
			return t.max(n.left)
		}
	}
	return n
}

func (t *AVL[T]) succ(n *avlNode[T], val T) *avlNode[T] {
	if n == nil {
		return nil
	}
	switch {
	case t.compare(val, n.val) < 0:
		return t.succ(n.left, val)
	case t.compare(val, n.val) > 0:
		return t.succ(n.right, val)
	default:
		if n.right != nil {
			return t.min(n.right)
		}
	}
	return n
}

func (t *AVL[T]) height(n *avlNode[T]) int {
	if n == nil {
		return 0
	}
	return 1 + max(t.height(n.left), t.height(n.right))
}

func (t *AVL[T]) balanceFactor(n *avlNode[T]) int {
	if n == nil {
		return 0
	}
	return t.height(n.left) - t.height(n.right)
}

func (t *AVL[T]) rotateLeft(n *avlNode[T]) *avlNode[T] {
	x := n.right
	n.right = x.left
	x.left = n
	return x
}

func (t *AVL[T]) rotateRight(n *avlNode[T]) *avlNode[T] {
	x := n.left
	n.left = x.right
	x.right = n
	return x
}

func (t *AVL[T]) rebalance(n *avlNode[T]) *avlNode[T] {
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
