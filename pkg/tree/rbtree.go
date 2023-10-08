package tree

import (
	"fmt"
	"reflect"

	"github.com/luverolla/lexgo/pkg/deque"
	"github.com/luverolla/lexgo/pkg/tau"
)

func Black[T any](node *rbNode[T]) bool {
	return tau.Nil(node) || !node.Red()
}

// Implementation of a Red-Black Tree.
//
// A Red-Black Tree is a binary search tree with one extra bit of storage per
// node: its color, which can be either RED or BLACK. By constraining the node
// colors on any simple path from the root to a leaf, red-black trees ensure
// that no such path is more than twice as long as any other, so that the tree
// is approximately balanced.
//
// The balancing of the tree is not perfect, but it is good enough to allow it
// to guarantee searching in O(log n) time, where n is the total number of
// elements in the tree. The insertion and deletion operations, along with the
// tree rearrangement and recoloring, are also performed in O(log n) time.
type RBTree[T any] struct {
	root *rbNode[T]
	size int
}

// Creates a new empty Red-Black Tree.
func RB[T any]() *RBTree[T] {
	return &RBTree[T]{nil, 0}
}

// --- Methods from tau.Collection[T] ---
func (rb *RBTree[T]) Size() int {
	return rb.size
}

func (rb *RBTree[T]) Empty() bool {
	return rb.size == 0
}

func (rb *RBTree[T]) Clear() {
	rb.root = nil
	rb.size = 0
}

func (rb *RBTree[T]) Contains(val T) bool {
	return !reflect.ValueOf(rb.get(rb.root, val)).IsNil()
}

func (rb *RBTree[T]) ContainsAll(coll tau.Collection[T]) bool {
	it := coll.Iter()
	for next, ok := it.Next(); ok; next, ok = it.Next() {
		if !rb.Contains(*next) {
			return false
		}
	}
	return true
}

func (rb *RBTree[T]) ContainsAny(coll tau.Collection[T]) bool {
	it := coll.Iter()
	for next, ok := it.Next(); ok; next, ok = it.Next() {
		if rb.Contains(*next) {
			return true
		}
	}
	return false
}

func (rb *RBTree[T]) Iter() tau.Iterator[T] {
	return rb.InOrder()
}

// --- Methods from tau.BSTree[T] ---
func (rb *RBTree[T]) Get(val T) tau.BSTreeNode[T] {
	return rb.get(rb.root, val)
}

func (rb *RBTree[T]) Root() tau.BSTreeNode[T] {
	return rb.root
}

func (rb *RBTree[T]) Insert(val T) tau.BSTreeNode[T] {
	rb.insert(rb.root, val)
	return rb.root
}

func (rb *RBTree[T]) Remove(val T) tau.BSTreeNode[T] {
	rb.remove(rb.root, val)
	return rb.root
}

func (rb *RBTree[T]) Min() tau.BSTreeNode[T] {
	return rb.min(rb.root)
}

func (rb *RBTree[T]) Max() tau.BSTreeNode[T] {
	return rb.max(rb.root)
}

func (rb *RBTree[T]) Pred(val T) tau.BSTreeNode[T] {
	return rb.pred(rb.root, val)
}

func (rb *RBTree[T]) Succ(val T) tau.BSTreeNode[T] {
	return rb.succ(rb.root, val)
}

func (rb *RBTree[T]) PreOrder() tau.Iterator[T] {
	return newRBPreOrderIter[T](rb)
}

func (rb *RBTree[T]) InOrder() tau.Iterator[T] {
	return newRBInOrderIter[T](rb)
}

func (rb *RBTree[T]) PostOrder() tau.Iterator[T] {
	return newRBPostOrderIter[T](rb)
}

// --- Private methods ---
func (rb *RBTree[T]) get(root *rbNode[T], val T) *rbNode[T] {
	if tau.Nil(root) {
		return nil
	}
	if tau.Cmp(val, root.val) < 0 {
		return rb.get(root.left, val)
	} else if tau.Cmp(val, root.val) > 0 {
		return rb.get(root.right, val)
	}
	return root
}

func (rb *RBTree[T]) insert(root *rbNode[T], val T) {
	if tau.Nil(root) {
		rb.size++
		rb.root = newRBNode[T](val, BLACK)
		return
	}
	if tau.Cmp(val, root.val) < 0 {
		if tau.Nil(root.left) {
			rb.size++
			root.left = newRBNode[T](val, RED)
			root.left.parent = root
			rb.fixInsert(root.left)
		} else {
			rb.insert(root.left, val)
		}
	} else if tau.Cmp(val, root.val) > 0 {
		if tau.Nil(root.right) {
			rb.size++
			root.right = newRBNode[T](val, RED)
			root.right.parent = root
			rb.fixInsert(root.right)
		} else {
			rb.insert(root.right, val)
		}
	}
}

func (rb *RBTree[T]) remove(root *rbNode[T], val T) {
	nodeToRemove := rb.get(root, val)
	if tau.Nil(nodeToRemove) {
		return
	}

	rb.size--
	if nodeToRemove.left != nil && nodeToRemove.right != nil {
		pred := rb.max(nodeToRemove.left)
		nodeToRemove.val = pred.val
		nodeToRemove = pred
	}

	var child *rbNode[T]
	if nodeToRemove.left != nil {
		child = nodeToRemove.left
	} else {
		child = nodeToRemove.right
	}

	if nodeToRemove.color == BLACK {
		if child != nil {
			nodeToRemove.color = child.color
		}
		rb.fixRemove(nodeToRemove)
	}

	if tau.Nil(nodeToRemove.parent) {
		rb.root = child
	} else if nodeToRemove == nodeToRemove.parent.left {
		nodeToRemove.parent.left = child
	} else {
		nodeToRemove.parent.right = child
	}

	if child != nil {
		child.parent = nodeToRemove.parent
	}

	nodeToRemove.parent = nil
	nodeToRemove.left = nil
	nodeToRemove.right = nil
}

// fix violations of the red-black tree properties after insertion
func (rb *RBTree[T]) fixInsert(node *rbNode[T]) {
	for node != rb.root && !Black(node.parent) {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if !Black(uncle) {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					rb.rotateLeft(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rb.rotateRight(node.parent.parent)
			}
		} else {
			uncle := node.parent.parent.left
			if !Black(uncle) {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					rb.rotateRight(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				rb.rotateLeft(node.parent.parent)
			}
		}
	}
	rb.root.color = BLACK
}

// fix violations of the red-black tree properties after removal
func (rb *RBTree[T]) fixRemove(node *rbNode[T]) {
	for node != rb.root && Black(node) {
		if node == node.parent.left {
			sibling := node.parent.right
			if !Black(sibling) {
				sibling.color = BLACK
				node.parent.color = RED
				rb.rotateLeft(node.parent)
				sibling = node.parent.right
			}
			if Black(sibling.left) && Black(sibling.right) {
				sibling.color = RED
				node = node.parent
			} else {
				if Black(sibling.right) {
					sibling.left.color = BLACK
					sibling.color = RED
					rb.rotateRight(sibling)
					sibling = node.parent.right
				}
				sibling.color = node.parent.color
				node.parent.color = BLACK
				sibling.right.color = BLACK
				rb.rotateLeft(node.parent)
				node = rb.root
			}
		} else {
			sibling := node.parent.left
			if !Black(sibling) {
				sibling.color = BLACK
				node.parent.color = RED
				rb.rotateRight(node.parent)
				sibling = node.parent.left
			}
			if Black(sibling.left) && Black(sibling.right) {
				sibling.color = RED
				node = node.parent
			} else {
				if Black(sibling.left) {
					sibling.right.color = BLACK
					sibling.color = RED
					rb.rotateLeft(sibling)
					sibling = node.parent.left
				}
				sibling.color = node.parent.color
				node.parent.color = BLACK
				sibling.left.color = BLACK
				rb.rotateRight(node.parent)
				node = rb.root
			}
		}
	}
	node.color = BLACK
}

func (tree *RBTree[T]) rotateLeft(root *rbNode[T]) {
	right := root.right
	root.right = right.left
	if right.left != nil {
		right.left.parent = root
	}
	right.parent = root.parent
	if tau.Nil(root.parent) {
		tree.root = right
	} else if root == root.parent.left {
		root.parent.left = right
	} else {
		root.parent.right = right
	}
	right.left = root
	root.parent = right
}

func (tree *RBTree[T]) rotateRight(root *rbNode[T]) {
	left := root.left
	root.left = left.right
	if left.right != nil {
		left.right.parent = root
	}
	left.parent = root.parent
	if tau.Nil(root.parent) {
		tree.root = left
	} else if root == root.parent.left {
		root.parent.left = left
	} else {
		root.parent.right = left
	}
	left.right = root
	root.parent = left
}

func (rb *RBTree[T]) min(root *rbNode[T]) *rbNode[T] {
	if tau.Nil(root) {
		return nil
	}
	if tau.Nil(root.left) {
		return root
	}
	return rb.min(root.left)
}

func (rb *RBTree[T]) max(root *rbNode[T]) *rbNode[T] {
	if tau.Nil(root) {
		return nil
	}
	if tau.Nil(root.right) {
		return root
	}
	return rb.max(root.right)
}

func (rb *RBTree[T]) pred(root *rbNode[T], val T) *rbNode[T] {
	if tau.Nil(root) {
		return nil
	}
	if tau.Cmp(val, root.val) <= 0 {
		return rb.pred(root.left, val)
	}
	right := rb.pred(root.right, val)
	if right != nil {
		return right
	}
	return root
}

func (rb *RBTree[T]) succ(root *rbNode[T], val T) *rbNode[T] {
	if tau.Nil(root) {
		return nil
	}
	if tau.Cmp(val, root.val) >= 0 {
		return rb.succ(root.right, val)
	}
	left := rb.succ(root.left, val)
	if left != nil {
		return left
	}
	return root
}

// --- Debug Helper Methods ---
func (rb *RBTree[T]) PrintfTree() string {
	return rb.printfTree(rb.root, 0)
}

func (rb *RBTree[T]) printfTree(root *rbNode[T], indent int) string {
	if reflect.ValueOf(root).IsNil() {
		return ""
	}
	s := rb.printfTree(root.right, indent+1)
	for i := 0; i < indent; i++ {
		s += "  "
	}
	s += fmt.Sprintf("%v\n", root.val)
	s += rb.printfTree(root.left, indent+1)
	return s
}

// --- Private types ---
type rbColor bool

var (
	RED   rbColor = true
	BLACK rbColor = false
)

type rbNode[T any] struct {
	val    T
	left   *rbNode[T]
	right  *rbNode[T]
	parent *rbNode[T]
	color  rbColor
}

func newRBNode[T any](val T, color rbColor) *rbNode[T] {
	return &rbNode[T]{val, nil, nil, nil, color}
}

func (node *rbNode[T]) Value() T {
	return node.val
}

func (node *rbNode[T]) Left() tau.BSTreeNode[T] {
	return node.left
}

func (node *rbNode[T]) Right() tau.BSTreeNode[T] {
	return node.right
}

func (node *rbNode[T]) Red() bool {
	return node.color == RED
}

func (node *rbNode[T]) sibling() *rbNode[T] {
	if tau.Nil(node.parent) {
		return nil
	}
	if node == node.parent.left {
		return node.parent.right
	}
	return node.parent.left
}

func (node *rbNode[T]) uncle() *rbNode[T] {
	if tau.Nil(node.parent) {
		return nil
	}
	return node.parent.sibling()
}

func (node *rbNode[T]) grandparent() *rbNode[T] {
	if tau.Nil(node.parent) {
		return nil
	}
	return node.parent.parent
}

// --- Iterators ---
type rbPreOrderIter[T any] struct {
	tree  *RBTree[T]
	stack tau.Deque[rbNode[T]]
}

func newRBPreOrderIter[T any](tree *RBTree[T]) *rbPreOrderIter[T] {
	iter := &rbPreOrderIter[T]{tree, deque.Arr[rbNode[T]]()}
	if tree.root != nil {
		iter.stack.PushFront(*tree.root)
	}
	return iter
}

func (iter *rbPreOrderIter[T]) Next() (*T, bool) {
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

func (iter *rbPreOrderIter[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}

type rbInOrderIter[T any] struct {
	tree  *RBTree[T]
	stack tau.Deque[rbNode[T]]
}

func newRBInOrderIter[T any](tree *RBTree[T]) *rbInOrderIter[T] {
	iter := &rbInOrderIter[T]{tree, deque.Arr[rbNode[T]]()}
	node := tree.root
	for node != nil {
		iter.stack.PushFront(*node)
		node = node.left
	}
	return iter
}

func (iter *rbInOrderIter[T]) Next() (*T, bool) {
	if iter.stack.Empty() {
		return nil, false
	}
	node, _ := iter.stack.PopFront()
	if node.right != nil {
		node = node.right
		for node != nil {
			iter.stack.PushFront(*node)
			node = node.left
		}
	}

	if tau.Nil(node) {
		return nil, false
	}

	return &node.val, true
}

func (iter *rbInOrderIter[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}

type rbPostOrderIter[T any] struct {
	tree  *RBTree[T]
	stack tau.Deque[rbNode[T]]
}

func newRBPostOrderIter[T any](tree *RBTree[T]) *rbPostOrderIter[T] {
	iter := &rbPostOrderIter[T]{tree, deque.Arr[rbNode[T]]()}
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

func (iter *rbPostOrderIter[T]) Next() (*T, bool) {
	if iter.stack.Empty() {
		return nil, false
	}
	node, _ := iter.stack.PopFront()
	if !iter.stack.Empty() {
		parent, _ := iter.stack.Front()
		if node == parent.left {
			node = parent.right
			for node != nil {
				iter.stack.PushFront(*node)
				if node.left != nil {
					node = node.left
				} else {
					node = node.right
				}
			}
		}
	}
	return &node.val, true
}

func (iter *rbPostOrderIter[T]) Each(f func(T)) {
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		f(*data)
	}
}
