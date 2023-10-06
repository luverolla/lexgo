package algo

import "github.com/luverolla/lexgo/pkg/tau"

func QuickSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	return quickSort(coll, cmp, 0, coll.Size()-1)
}

func MergeSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	return mergeSort(coll, cmp, 0, coll.Size()-1)
}

func BubbleSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	return bubbleSort(coll, cmp)
}

func InsertionSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	return insertionSort(coll, cmp)
}

func SelectionSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	return selectionSort(coll, cmp)
}

func HeapSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	return heapSort(coll, cmp)
}

func heapSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	copy := coll.Clone().(tau.IdxedCollection[T])
	for i := copy.Size()/2 - 1; i >= 0; i-- {
		heapify(copy, cmp, copy.Size(), i)
	}
	for i := copy.Size() - 1; i >= 0; i-- {
		copy.Swap(0, i)
		heapify(copy, cmp, i, 0)
	}
	return copy
}

func heapify[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T], size, root int) {
	largest := root
	left := 2*root + 1
	right := 2*root + 2
	if left < size {
		largestVal, _ := coll.Get(largest)
		leftVal, _ := coll.Get(left)
		if cmp(*leftVal, *largestVal) > 0 {
			largest = left
		}
	}
	if right < size {
		largestVal, _ := coll.Get(largest)
		rightVal, _ := coll.Get(right)
		if cmp(*rightVal, *largestVal) > 0 {
			largest = right
		}
	}
	if largest != root {
		coll.Swap(root, largest)
		heapify(coll, cmp, size, largest)
	}
}

func bubbleSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	copy := coll.Clone().(tau.IdxedCollection[T])
	for i := 0; i < copy.Size()-1; i++ {
		for j := 0; j < copy.Size()-i-1; j++ {
			curr, _ := copy.Get(j)
			next, _ := copy.Get(j + 1)
			if cmp(*curr, *next) > 0 {
				copy.Swap(j, j+1)
			}
		}
	}
	return copy
}

func insertionSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	copy := coll.Clone().(tau.IdxedCollection[T])
	for i := 1; i < copy.Size(); i++ {
		curr, _ := copy.Get(i)
		j := i - 1
		val, _ := copy.Get(j)
		for j >= 0 && cmp(*curr, *val) < 0 {
			val, _ = copy.Get(j)
			copy.Set(j+1, *val)
			j--
		}
		copy.Set(j+1, *curr)
	}
	return copy
}

func selectionSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T]) tau.IdxedCollection[T] {
	copy := coll.Clone().(tau.IdxedCollection[T])
	for i := 0; i < copy.Size()-1; i++ {
		min := i
		for j := i + 1; j < copy.Size(); j++ {
			valJ, _ := copy.Get(j)
			valMin, _ := copy.Get(min)
			if cmp(*valJ, *valMin) < 0 {
				min = j
			}
		}
		copy.Swap(i, min)
	}
	return copy
}

func mergeSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T], start, end int) tau.IdxedCollection[T] {
	copy := coll.Clone().(tau.IdxedCollection[T])
	if start < end {
		mid := (start + end) / 2
		mergeSort(copy, cmp, start, mid)
		mergeSort(copy, cmp, mid+1, end)
		merge(copy, cmp, start, mid, end)
	}
	return copy
}

func merge[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T], start, mid, end int) {
	leftSize := mid - start + 1
	rightSize := end - mid
	left := make([]T, leftSize)
	right := make([]T, rightSize)
	for i := 0; i < leftSize; i++ {
		val, _ := coll.Get(start + i)
		left[i] = *val
	}
	for i := 0; i < rightSize; i++ {
		val, _ := coll.Get(mid + 1 + i)
		right[i] = *val
	}
	i, j, k := 0, 0, start
	for i < leftSize && j < rightSize {
		if cmp(left[i], right[j]) <= 0 {
			coll.Set(k, left[i])
			i++
		} else {
			coll.Set(k, right[j])
			j++
		}
		k++
	}
	for i < leftSize {
		coll.Set(k, left[i])
		i++
		k++
	}
	for j < rightSize {
		coll.Set(k, right[j])
		j++
		k++
	}
}

func quickSort[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T], start, end int) tau.IdxedCollection[T] {
	copy := coll.Clone().(tau.IdxedCollection[T])
	if start < end {
		pivot := qsPartition(copy, cmp, start, end)
		quickSort(copy, cmp, start, pivot-1)
		quickSort(copy, cmp, pivot+1, end)
	}
	return copy
}

func qsPartition[T any](coll tau.IdxedCollection[T], cmp tau.Comparator[T], start, end int) int {
	pivot, _ := coll.Get(end)
	i := start - 1
	for j := start; j < end; j++ {
		curr, _ := coll.Get(j)
		if cmp(*curr, *pivot) < 0 {
			i++
			coll.Swap(i, j)
		}
	}
	coll.Swap(i+1, end)
	return i + 1
}
