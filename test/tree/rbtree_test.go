package tree_test

import (
	"reflect"
	"testing"

	"github.com/luverolla/lexgo/pkg/tau"
	"github.com/luverolla/lexgo/pkg/tree"
)

type RBNode[T any] interface {
	tau.BSTreeNode[T]
	Red() bool
}

func Black[T any](node RBNode[T]) bool {
	return reflect.ValueOf(node).IsNil() || !node.Red()
}

func CheckBlackRootProp(tree *tree.RBTree[int]) bool {
	root := tree.Root().(RBNode[int])
	return Black(root)
}

func CheckRedChildrenProp(node RBNode[int]) bool {
	if Black(node) {
		return true
	}

	return Black(node.Left().(RBNode[int])) && Black(node.Right().(RBNode[int]))
}

func blackHeight(node RBNode[int]) int {
	if reflect.ValueOf(node).IsNil() {
		return 1
	}

	leftHeight := blackHeight(node.Left().(RBNode[int]))
	rightHeight := blackHeight(node.Right().(RBNode[int]))
	maxHeight := max(leftHeight, rightHeight)

	if Black(node) {
		return maxHeight + 1
	}
	return maxHeight
}

func CheckBlackHeightProp(root RBNode[int]) bool {
	if reflect.ValueOf(root).IsNil() {
		return true
	}

	leftHeight := blackHeight(root.Left().(RBNode[int]))
	rightHeight := blackHeight(root.Right().(RBNode[int]))

	if leftHeight != rightHeight {
		return false
	}

	leftCheck := CheckBlackHeightProp(root.Left().(RBNode[int]))
	rightCheck := CheckBlackHeightProp(root.Right().(RBNode[int]))

	return leftCheck && rightCheck
}

func TestRBTreeAdd(t *testing.T) {
	tree := tree.RB[int]()
	if tree.Size() != 0 {
		t.Errorf("RBTree size is %d, expected %d", tree.Size(), 0)
	}

	for i := 0; i < 1000; i++ {
		tree.Insert(i)
		CheckBlackRootProp(tree)
		CheckRedChildrenProp(tree.Root().(RBNode[int]))
		CheckBlackHeightProp(tree.Root().(RBNode[int]))
	}
}

func TestRBTreeRemove(t *testing.T) {
	tree := tree.RB[int]()
	for i := 0; i < 1000; i++ {
		tree.Insert(i)
	}

	for i := 0; i < 1000; i++ {
		tree.Remove(i)
		CheckBlackRootProp(tree)
		CheckRedChildrenProp(tree.Root().(RBNode[int]))
		CheckBlackHeightProp(tree.Root().(RBNode[int]))
	}
}
