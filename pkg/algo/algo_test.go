package algo

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/list"
	"github.com/luverolla/lexgo/pkg/tau"
)

func ExampleMergeSort() {
	list := list.Arr(1, 5, 2, 3, 4)
	sorted := MergeSort(list, tau.ASCmp)
	fmt.Printf("%v", sorted)
	// Output: ArrList[1,2,3,4,5]
}

func ExampleInsertionSort() {
	list := list.Arr(1, 5, 2, 3, 4)
	sorted := InsertionSort(list, tau.ASCmp)
	fmt.Printf("%v", sorted)
	// Output: ArrList[1,2,3,4,5]
}

func ExampleBubbleSort() {
	list := list.Arr(1, 5, 2, 3, 4)
	sorted := BubbleSort(list, tau.ASCmp)
	fmt.Printf("%v", sorted)
	// Output: ArrList[1,2,3,4,5]
}

func ExampleSelectionSort() {
	list := list.Arr(1, 5, 2, 3, 4)
	sorted := SelectionSort(list, tau.ASCmp)
	fmt.Printf("%v", sorted)
	// Output: ArrList[1,2,3,4,5]
}

func ExampleHeapSort() {
	list := list.Arr(1, 5, 2, 3, 4)
	sorted := HeapSort(list, tau.ASCmp)
	fmt.Printf("%v", sorted)
	// Output: ArrList[1,2,3,4,5]
}
