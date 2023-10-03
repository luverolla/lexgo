package list_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/list"
)

var data_int = []int{4, 7, 9, 2547347383, 50}
var data_str = []string{"ciao", "", "\r\n", "hey!"}

var lkl_int list.List[int]
var lkl_str list.List[string]

func TestLinkedListAdd(t *testing.T) {
	lkl_int = list.NewLinkedList[int](data_int...)
	lkl_str = list.NewLinkedList[string](data_str...)

	if lkl_int.Size() != len(data_int) {
		t.Errorf("LinkedList[int] size is %d, expected %d", lkl_int.Size(), len(data_int))
	}

	if lkl_str.Size() != len(data_str) {
		t.Errorf("LinkedList[string] size is %d, expected %d", lkl_str.Size(), len(data_str))
	}
}

func TestLinkedListGet(t *testing.T) {
	for i, v := range data_int {
		if lkl_int.Get(i) != v {
			t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i, lkl_int.Get(i), v)
		}
	}

	for i, v := range data_str {
		if lkl_str.Get(i) != v {
			t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i, lkl_str.Get(i), v)
		}
	}
}

func TestLinkedListSet(t *testing.T) {
	for i, v := range data_int {
		lkl_int.Set(i, v+1)
		if lkl_int.Get(i) != v+1 {
			t.Errorf("LinkedList[int] Set(%d) is %d, expected %d", i, lkl_int.Get(i), v+1)
		}
	}

	for i, v := range data_str {
		lkl_str.Set(i, v+"1")
		if lkl_str.Get(i) != v+"1" {
			t.Errorf("LinkedList[string] Set(%d) is %s, expected %s", i, lkl_str.Get(i), v+"1")
		}
	}
}

func TestLinkedListRemove(t *testing.T) {
	idxToRemoveInt := 2
	idxToRemoveStr := 1

	lkl_int_copy := list.NewLinkedList[int](data_int...)
	lkl_str_copy := list.NewLinkedList[string](data_str...)

	lkl_int_copy.RemoveAt(idxToRemoveInt)
	lkl_str_copy.RemoveAt(idxToRemoveStr)

	if lkl_int_copy.Size() != len(data_int)-1 {
		t.Errorf("LinkedList[int] size is %d, expected %d", lkl_int_copy.Size(), len(data_int)-1)
	}

	if lkl_str_copy.Size() != len(data_str)-1 {
		t.Errorf("LinkedList[string] size is %d, expected %d", lkl_str_copy.Size(), len(data_str)-1)
	}

	for i, v := range data_int {
		if i < idxToRemoveInt {
			if lkl_int_copy.Get(i) != v {
				t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i, lkl_int_copy.Get(i), v)
			}
		} else if i > idxToRemoveInt {
			if lkl_int_copy.Get(i-1) != v {
				t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i-1, lkl_int_copy.Get(i-1), v)
			}
		}
	}

	for i, v := range data_str {
		if i < idxToRemoveStr {
			if lkl_str_copy.Get(i) != v {
				t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i, lkl_str_copy.Get(i), v)
			}
		} else if i > idxToRemoveStr {
			if lkl_str_copy.Get(i-1) != v {
				t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i-1, lkl_str_copy.Get(i-1), v)
			}
		}
	}
}

func TestLinkedListRemoveBulk(t *testing.T) {
	lkl_int_copy_1 := list.NewLinkedList[int](data_int...)
	lkl_str_copy_1 := list.NewLinkedList[string](data_str...)

	lkl_int_copy_1.RemoveAll(data_int[0])
	lkl_str_copy_1.RemoveAll(data_str[0])

	if lkl_int_copy_1.Size() != len(data_int)-1 {
		t.Errorf("LinkedList[int] size is %d, expected %d", lkl_int_copy_1.Size(), len(data_int)-1)
	}

	if lkl_str_copy_1.Size() != len(data_str)-1 {
		t.Errorf("LinkedList[string] size is %d, expected %d", lkl_str_copy_1.Size(), len(data_str)-1)
	}

	for i, v := range data_int {
		if i == 0 {
			continue
		}
		if lkl_int_copy_1.Get(i-1) != v {
			t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i-1, lkl_int_copy_1.Get(i-1), v)
		}
	}

	for i, v := range data_str {
		if i == 0 {
			continue
		}
		if lkl_str_copy_1.Get(i-1) != v {
			t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i-1, lkl_str_copy_1.Get(i-1), v)
		}
	}

	lkl_int_copy_1.Clear()
	lkl_str_copy_1.Clear()

	if lkl_int_copy_1.Size() != 0 {
		t.Errorf("LinkedList[int] size is %d, expected %d", lkl_int_copy_1.Size(), 0)
	}

	if lkl_str_copy_1.Size() != 0 {
		t.Errorf("LinkedList[string] size is %d, expected %d", lkl_str_copy_1.Size(), 0)
	}
}

func TestLinkedListContains(t *testing.T) {

	lkl_int_copy_2 := list.NewLinkedList[int](data_int...)
	lkl_str_copy_2 := list.NewLinkedList[string](data_str...)

	t.Logf("data_int: %v", data_int)
	t.Logf("data_str: %v", data_str)

	for _, v := range data_int {
		if !lkl_int_copy_2.Contains(v) {
			t.Errorf("LinkedList[int] does not contain %d", v)
		}
	}

	for _, v := range data_str {
		if !lkl_str_copy_2.Contains(v) {
			t.Errorf("LinkedList[string] does not contain %s", v)
		}
	}
}

func TestLinkedListContainsAll(t *testing.T) {
	lkl_int_copy_3 := list.NewLinkedList[int](data_int...)
	lkl_str_copy_3 := list.NewLinkedList[string](data_str...)

	if !lkl_int_copy_3.ContainsAll(lkl_int_copy_3) {
		t.Errorf("LinkedList[int] does not contain all its elements")
	}

	if !lkl_str_copy_3.ContainsAll(lkl_str_copy_3) {
		t.Errorf("LinkedList[string] does not contain all its elements")
	}

	if !lkl_int_copy_3.ContainsAll(list.NewLinkedList[int](data_int[0])) {
		t.Errorf("LinkedList[int] does not contain all elements of another list")
	}

	if !lkl_str_copy_3.ContainsAll(list.NewLinkedList[string](data_str[0])) {
		t.Errorf("LinkedList[string] does not contain all elements of another list")
	}
}

func TestLinkedListContainsAny(t *testing.T) {
	lkl_int_copy_4 := list.NewLinkedList[int](data_int...)
	lkl_str_copy_4 := list.NewLinkedList[string](data_str...)

	if !lkl_int_copy_4.ContainsAny(lkl_int_copy_4) {
		t.Errorf("LinkedList[int] does not contain any of its elements")
	}

	if !lkl_str_copy_4.ContainsAny(lkl_str_copy_4) {
		t.Errorf("LinkedList[string] does not contain any of its elements")
	}

	if !lkl_int_copy_4.ContainsAny(list.NewLinkedList[int](data_int[0], data_int[0]+1)) {
		t.Errorf("LinkedList[int] does not contain any elements of another list")
	}

	if !lkl_str_copy_4.ContainsAny(list.NewLinkedList[string](data_str[0], "ueeeeeeee")) {
		t.Errorf("LinkedList[string] does not contain any elements of another list")
	}
}
