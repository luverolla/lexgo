package list_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/list"
)

var data_int_2 = []int{4, 7, 9, 2547347383, 50}
var data_str_2 = []string{"ciao", "", "\r\n", "hey!"}

var arl_int colls.List[int]
var arl_str colls.List[string]

func TestArrayListAdd(t *testing.T) {
	arl_int = list.NewArray[int](data_int_2...)
	arl_str = list.NewArray[string](data_str_2...)

	if arl_int.Size() != len(data_int_2) {
		t.Errorf("ArrayList[int] size is %d, expected %d", arl_int.Size(), len(data_int_2))
	}

	if arl_str.Size() != len(data_str_2) {
		t.Errorf("ArrayList[string] size is %d, expected %d", arl_str.Size(), len(data_str_2))
	}
}

func TestArrayListGet(t *testing.T) {
	for i, v := range data_int_2 {
		val, _ := arl_int.Get(i)
		if *val != v {
			t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i, *val, v)
		}
	}

	for i, v := range data_str_2 {
		val, _ := arl_str.Get(i)
		if *val != v {
			t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i, *val, v)
		}
	}
}

func TestArrayListSet(t *testing.T) {
	for i, v := range data_int_2 {
		arl_int.Set(i, v+1)
		val, _ := arl_int.Get(i)
		if *val != v+1 {
			t.Errorf("ArrayList[int] Set(%d) is %d, expected %d", i, *val, v+1)
		}
	}

	for i, v := range data_str_2 {
		arl_str.Set(i, v+"1")
		val, _ := arl_str.Get(i)
		if *val != v+"1" {
			t.Errorf("ArrayList[string] Set(%d) is %s, expected %s", i, *val, v+"1")
		}
	}
}

func TestArrayListRemove(t *testing.T) {
	arl_int_copy := list.NewArray[int](data_int_2...)
	arl_str_copy := list.NewArray[string](data_str_2...)

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
			val, _ := arl_int_copy.Get(i)
			if *val != v {
				t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i, *val, v)
			}
		} else if i > idxToRemoveInt {
			val, _ := arl_int_copy.Get(i - 1)
			if *val != v {
				t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i-1, *val, v)
			}
		}
	}

	for i, v := range data_str_2 {
		if i < idxToRemoveStr {
			val, _ := arl_str_copy.Get(i)
			if *val != v {
				t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i, *val, v)
			}
		} else if i > idxToRemoveStr {
			val, _ := arl_str_copy.Get(i - 1)
			if *val != v {
				t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i-1, *val, v)
			}
		}
	}
}

func TestArrayListRemoveBulk(t *testing.T) {

	data_int_2_copy := []int{4, 7, 9, 2547347383, 50}
	data_str_2_copy := []string{"ciao", "", "\r\n", "hey!"}

	arl_int_copy_1 := list.NewArray[int](data_int_2_copy...)
	arl_str_copy_1 := list.NewArray[string](data_str_2_copy...)

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
			val, _ := arl_int_copy_1.Get(i)
			if *val != v {
				t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", i, *val, v)
			}
		}
	}

	for i, v := range data_str_2_copy {
		if v != "" {
			val, _ := arl_str_copy_1.Get(i)
			if *val != v {
				t.Errorf("ArrayList[string] Get(%d) is %s, expected %s", i, *val, v)
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
	if !arl_int.ContainsAll(list.NewArray[int](data_int_2...)) {
		t.Errorf("ArrayList[int] ContainsAll(ArrayList[int]) is false, expected true")
	}

	if !arl_str.ContainsAll(list.NewArray[string](data_str_2...)) {
		t.Errorf("ArrayList[string] ContainsAll(ArrayList[string]) is false, expected true")
	}
}

func TestArrayListContainsAny(t *testing.T) {
	if !arl_int.ContainsAny(list.NewArray[int](data_int_2[0], data_int_2[0]+1)) {
		t.Errorf("ArrayList[int] ContainsAny(ArrayList[int]) is false, expected true")
	}

	if !arl_str.ContainsAny(list.NewArray[string](data_str_2[0], "ueeee")) {
		t.Errorf("ArrayList[string] ContainsAny(ArrayList[string]) is false, expected true")
	}
}
