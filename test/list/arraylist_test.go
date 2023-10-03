package list_test

// Make a test file similar to linkedlist_test.go
// But replacing LinkedList with ArrayList

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/list"
)

var data_int_2 = []int{4, 7, 9, 2547347383, 50}
var data_str_2 = []string{"ciao", "", "\r\n", "hey!"}

var arl_int list.List[int]
var arl_str list.List[string]

func TestArrayListAdd(t *testing.T) {
	arl_int = list.NewArrayList[int](data_int_2...)
	arl_str = list.NewArrayList[string](data_str_2...)

	if arl_int.Size() != len(data_int_2) {
		t.Errorf("ArrayList[int] size is %d, expected %d", arl_int.Size(), len(data_int_2))
	}

	if arl_str.Size() != len(data_str_2) {
		t.Errorf("ArrayList[string] size is %d, expected %d", arl_str.Size(), len(data_str_2))
	}
}

func TestArrayListGet(t *testing.T) {
	for i, v := range data_int_2 {
		if arl_int.Get(i) != v {
			t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i, arl_int.Get(i), v)
		}
	}

	for i, v := range data_str_2 {
		if arl_str.Get(i) != v {
			t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i, arl_str.Get(i), v)
		}
	}
}

func TestArrayListSet(t *testing.T) {
	for i, v := range data_int_2 {
		arl_int.Set(i, v+1)
		if arl_int.Get(i) != v+1 {
			t.Errorf("ArrayList[int] Set(%d) is %d, expected %d", i, arl_int.Get(i), v+1)
		}
	}

	for i, v := range data_str_2 {
		arl_str.Set(i, v+"1")
		if arl_str.Get(i) != v+"1" {
			t.Errorf("ArrayList[string] Set(%d) is %s, expected %s", i, arl_str.Get(i), v+"1")
		}
	}
}

func TestArrayListRemove(t *testing.T) {
	arl_int_copy := list.NewArrayList[int](data_int_2...)
	arl_str_copy := list.NewArrayList[string](data_str_2...)

	idxToRemoveInt := 2
	idxToRemoveStr := 1

	arl_int_copy.RemoveAt(idxToRemoveInt)
	arl_str_copy.RemoveAt(idxToRemoveStr)

	if arl_int_copy.Size() != len(data_int_2)-1 {
		t.Errorf("ArrayList[int] size is %d, expected %d", arl_int_copy.Size(), len(data_int_2)-1)
	}

	if arl_str_copy.Size() != len(data_str_2)-1 {
		t.Errorf("ArrayList[string] size is %d, expected %d", arl_str_copy.Size(), len(data_str_2)-1)
	}

	for i, v := range data_int_2 {
		if i < idxToRemoveInt {
			if arl_int_copy.Get(i) != v {
				t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i, arl_int_copy.Get(i), v)
			}
		} else if i > idxToRemoveInt {
			if arl_int_copy.Get(i-1) != v {
				t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i-1, arl_int_copy.Get(i-1), v)
			}
		}
	}

	for i, v := range data_str_2 {
		if i < idxToRemoveStr {
			if arl_str_copy.Get(i) != v {
				t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i, arl_str_copy.Get(i), v)
			}
		} else if i > idxToRemoveStr {
			if arl_str_copy.Get(i-1) != v {
				t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i-1, arl_str_copy.Get(i-1), v)
			}
		}
	}
}

func TestArrayListRemoveBulk(t *testing.T) {

	data_int_2_copy := []int{4, 7, 9, 2547347383, 50}
	data_str_2_copy := []string{"ciao", "", "\r\n", "hey!"}

	arl_int_copy_1 := list.NewArrayList[int](data_int_2_copy...)
	arl_str_copy_1 := list.NewArrayList[string](data_str_2_copy...)

	t.Logf("arl_int_copy_1: %v", arl_int_copy_1)
	t.Logf("arl_str_copy_1: %v", arl_str_copy_1)

	arl_int_copy_1.RemoveAll(4)
	arl_str_copy_1.RemoveAll("")

	t.Logf("arl_int_copy_1 after: %v", arl_int_copy_1)
	t.Logf("arl_str_copy_1 after: %v", arl_str_copy_1)

	if arl_int_copy_1.Size() != len(data_int_2_copy)-1 {
		t.Errorf("ArrayList[int] size is %d, expected %d", arl_int_copy_1.Size(), len(data_int_2_copy)-1)
	}

	if arl_str_copy_1.Size() != len(data_str_2_copy)-1 {
		t.Errorf("ArrayList[string] size is %d, expected %d", arl_str_copy_1.Size(), len(data_str_2_copy)-1)
	}

	for i, v := range data_int_2_copy {
		if v != 4 {
			if arl_int_copy_1.Get(i) != v {
				t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i, arl_int_copy_1.Get(i), v)
			}
		}
	}

	for i, v := range data_str_2_copy {
		if v != "" {
			if arl_str_copy_1.Get(i) != v {
				t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i, arl_str_copy_1.Get(i), v)
			}
		}
	}
}

func TestArrayListContains(t *testing.T) {
	for _, v := range data_int_2 {
		if !arl_int.Contains(v) {
			t.Errorf("ArrayList[int] Contains(%d) is false, expected true", v)
		}
	}

	for _, v := range data_str_2 {
		if !arl_str.Contains(v) {
			t.Errorf("ArrayList[string] Contains(%s) is false, expected true", v)
		}
	}
}

func TestArrayListContainsAll(t *testing.T) {
	if !arl_int.ContainsAll(list.NewArrayList[int](data_int_2...)) {
		t.Errorf("ArrayList[int] ContainsAll(ArrayList[int]) is false, expected true")
	}

	if !arl_str.ContainsAll(list.NewArrayList[string](data_str_2...)) {
		t.Errorf("ArrayList[string] ContainsAll(ArrayList[string]) is false, expected true")
	}
}

func TestArrayListContainsAny(t *testing.T) {
	if !arl_int.ContainsAny(list.NewArrayList[int](data_int_2[0], data_int_2[0]+1)) {
		t.Errorf("ArrayList[int] ContainsAny(ArrayList[int]) is false, expected true")
	}

	if !arl_str.ContainsAny(list.NewArrayList[string](data_str_2[0], "ueeee")) {
		t.Errorf("ArrayList[string] ContainsAny(ArrayList[string]) is false, expected true")
	}
}
