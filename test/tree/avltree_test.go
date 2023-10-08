package tree_test

import (
	"reflect"
	"testing"

	"github.com/luverolla/lexgo/pkg/tau"
	"github.com/luverolla/lexgo/pkg/tree"
)

func Height[T any](root tau.BSTreeNode[T]) int {
	if reflect.ValueOf(root).IsNil() {
		return 0
	}
	return 1 + max(Height(root.Left()), Height(root.Right()))
}

// a binary search tree is balanced when the height of its left and right subtrees differ by at most 1
func IsAVLBalanced[T any](t *tree.AVLTree[T]) bool {
	if reflect.ValueOf(t.Root()).IsNil() {
		return true
	}

	leftHeight := Height(t.Root().Left())
	rightHeight := Height(t.Root().Right())
	diffHeight := leftHeight - rightHeight
	return diffHeight >= -1 && diffHeight <= 1
}

func TestAVLIsBalanced(t *testing.T) {
	avl := tree.AVL[int]()

	for i := 0; i < 100; i++ {
		avl.Insert(i)
		if !IsAVLBalanced(avl) {
			t.Errorf("AVL tree is not balanced after inserting %d\n", i)
		}
	}

	for i := 0; i < 100; i++ {
		avl.Remove(i)
		if !IsAVLBalanced(avl) {
			t.Errorf("AVL tree is not balanced after removing %d\n", i)
		}
	}
}
