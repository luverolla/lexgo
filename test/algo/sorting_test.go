package algo_test

import (
	"fmt"
	"testing"

	"github.com/luverolla/lexgo/pkg/algo"
	"github.com/luverolla/lexgo/pkg/list"
	"github.com/luverolla/lexgo/pkg/tau"
)

var startData = list.Lkd([]int{90, 47, 2, 145, 30, 750, 99, 25, 76}...)
var sortedData = list.Lkd([]int{2, 25, 30, 47, 76, 90, 99, 145, 750}...)

func TestQuickSort(t *testing.T) {
	sorted := algo.QuickSort(startData, tau.ASCmp)

	if sorted.Size() != startData.Size() {
		t.Errorf("QuickSort() = %v, want %v", sorted.Size(), startData.Size())
	}

	if !tau.Eq(sorted, sortedData) {
		t.Errorf("QuickSort() = %v, want %v", sorted, sortedData)
	}

}

func TestMergeSort(t *testing.T) {
	sorted := algo.MergeSort(startData, tau.ASCmp)

	if sorted.Size() != startData.Size() {
		t.Errorf("MergeSort() = %v, want %v", sorted.Size(), startData.Size())
	}

	if !tau.Eq(sorted, sortedData) {
		t.Errorf("MergeSort() = %v, want %v", sorted, sortedData)
	}
}

func TestBubbleSort(t *testing.T) {
	sorted := algo.BubbleSort(startData, tau.ASCmp)

	if sorted.Size() != startData.Size() {
		t.Errorf("BubbleSort() = %v, want %v", sorted.Size(), startData.Size())
	}

	if !tau.Eq(sorted, sortedData) {
		t.Errorf("BubbleSort() = %v, want %v", sorted, sortedData)
	}
}

func TestInsertionSort(t *testing.T) {
	sorted := algo.InsertionSort(startData, tau.ASCmp)

	if sorted.Size() != startData.Size() {
		t.Errorf("InsertionSort() = %v, want %v", sorted.Size(), startData.Size())
	}

	if !tau.Eq(sorted, sortedData) {
		t.Errorf("InsertionSort() = %v, want %v", sorted, sortedData)
	}
}

func TestSelectionSort(t *testing.T) {
	sorted := algo.SelectionSort(startData, tau.ASCmp)

	if sorted.Size() != startData.Size() {
		t.Errorf("SelectionSort() = %v, want %v", sorted.Size(), startData.Size())
	}

	if !tau.Eq(sorted, sortedData) {
		t.Errorf("SelectionSort() = %v, want %v", sorted, sortedData)
	}
}

func TestHeapSort(t *testing.T) {
	sorted := algo.HeapSort(startData, tau.ASCmp)

	if sorted.Size() != startData.Size() {
		t.Errorf("HeapSort() = %v, want %v", sorted.Size(), startData.Size())
	}

	if !tau.Eq(sorted, sortedData) {
		t.Errorf("HeapSort() = %v, want %v", sorted, sortedData)
	}
}

// --- Examples ---
func ExampleQuickSort() {
	// Sorts a [tau.List] of integers using the QuickSort algorithm
	// and the ascending order comparator
	list := list.Arr[int](1, 5, 2, 3, 4)
	sorted := algo.QuickSort(list, tau.ASCmp)
	fmt.Printf("%v", sorted)
}
