package list_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/list"
)

var data_int = []int{4, 7, 9, 2547347383, 50, 78, 9, 77, 32, 9}
var data_str = []string{"ciao", "", "becco", "\r\n", "hey!", "", "castoro"}

var lkl_int colls.List[int]
var lkl_str colls.List[string]

func TestLinkedListAdd(t *testing.T) {
	lkl_int = list.NewLinked[int](data_int...)
	lkl_str = list.NewLinked[string](data_str...)

	if lkl_int.Size() != len(data_int) {
		t.Errorf("LinkedList[int] size is %d, expected %d", lkl_int.Size(), len(data_int))
	}

	if lkl_str.Size() != len(data_str) {
		t.Errorf("LinkedList[string] size is %d, expected %d", lkl_str.Size(), len(data_str))
	}
}

func TestLinkedListGet(t *testing.T) {
	for i, v := range data_int {
		val, _ := lkl_int.Get(i)
		if *val != v {
			t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i, *val, v)
		}
	}

	for i, v := range data_str {
		val, _ := lkl_str.Get(i)
		if *val != v {
			t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i, *val, v)
		}
	}
}

func TestLinkedListSet(t *testing.T) {
	for i, v := range data_int {
		lkl_int.Set(i, v+1)
		val, _ := lkl_int.Get(i)
		if *val != v+1 {
			t.Errorf("LinkedList[int] Set(%d) is %d, expected %d", i, *val, v+1)
		}
	}

	for i, v := range data_str {
		lkl_str.Set(i, v+"1")
		val, _ := lkl_str.Get(i)
		if *val != v+"1" {
			t.Errorf("LinkedList[string] Set(%d) is %s, expected %s", i, *val, v+"1")
		}
	}
}

func TestLinkedListRemove(t *testing.T) {
	lkl_int_copy := list.NewLinked[int](data_int...)
	lkl_str_copy := list.NewLinked[string](data_str...)

	idxToRemoveInt := 2
	idxToRemoveStr := 1

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
			val, _ := lkl_int_copy.Get(i)
			if *val != v {
				t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i, *val, v)
			}
		} else if i > idxToRemoveInt {
			val, _ := lkl_int_copy.Get(i - 1)
			if *val != v {
				t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", i-1, *val, v)
			}
		}
	}

	for i, v := range data_str {
		if i < idxToRemoveStr {
			val, _ := lkl_str_copy.Get(i)
			if *val != v {
				t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i, *val, v)
			}
		} else if i > idxToRemoveStr {
			val, _ := lkl_str_copy.Get(i - 1)
			if *val != v {
				t.Errorf("LinkedList[string] Get(%d) is %s, expected %s", i-1, *val, v)
			}
		}
	}
}

func TestLinkedListRemoveBulk(t *testing.T) {
	lkl_int_copy := list.NewLinked[int](data_int...)
	lkl_str_copy := list.NewLinked[string](data_str...)

	intToRemove := 9
	strToRemove := ""

	expRemInt := 3
	expRemStr := 2

	lkl_int_copy.RemoveAll(intToRemove)
	lkl_str_copy.RemoveAll(strToRemove)

	if lkl_int_copy.Size() != len(data_int)-expRemInt {
		t.Errorf("LinkedList[int] size is %d, expected %d", lkl_int_copy.Size(), len(data_int)-1)
	}

	if lkl_str_copy.Size() != len(data_str)-expRemStr {
		t.Errorf("LinkedList[string] size is %d, expected %d", lkl_str_copy.Size(), len(data_str)-1)
	}

	i := 0
	j := 0
	for i < len(data_int) {
		if data_int[i] == intToRemove {
			i++
			continue
		}
		val, _ := lkl_int_copy.Get(j)
		if *val != data_int[i] {
			t.Errorf("LinkedList[int] Get(%d) is %d, expected %d", j, *val, data_int[i])
		}
		i++
		j++
	}

	i = 0
	j = 0
	for i < len(data_str) {
		if data_str[i] == strToRemove {
			i++
			continue
		}
		val, _ := lkl_str_copy.Get(j)
		if *val != data_str[i] {
			t.Errorf("LinkedList[int] Get(%d) is %s, expected %s", j, *val, data_str[i])
		}
		i++
		j++
	}

}

func TestLinkedListContains(t *testing.T) {
	lkl_int_copy := list.NewLinked[int](data_int...)
	lkl_str_copy := list.NewLinked[string](data_str...)

	for _, v := range data_int {
		if !lkl_int_copy.Contains(v) {
			t.Errorf("LinkedList[int] Contains(%d) is false, expected true", v)
		}
	}

	for _, v := range data_str {
		if !lkl_str_copy.Contains(v) {
			t.Errorf("LinkedList[string] Contains(%s) is false, expected true", v)
		}
	}
}

func TestLinkedListContainsAll(t *testing.T) {

	lkl_int_copy := list.NewLinked[int](data_int...)
	lkl_str_copy := list.NewLinked[string](data_str...)

	if !lkl_int_copy.ContainsAll(list.NewLinked[int](data_int...)) {
		t.Errorf("LinkedList[int] ContainsAll(LinkedList[int]) is false, expected true")
	}

	if !lkl_str_copy.ContainsAll(list.NewLinked[string](data_str...)) {
		t.Errorf("LinkedList[string] ContainsAll(LinkedList[string]) is false, expected true")
	}
}

func TestLinkedListContainsAny(t *testing.T) {

	test_data_int := []int{1000, 10, 0, 0, 9, 0, 1000}
	test_data_str := []string{"alpha", "beta", "becco", "gamma", "lambda"}

	sub_int := list.NewLinked[int](test_data_int...)
	sub_str := list.NewLinked[string](test_data_str...)

	lkl_int_copy := list.NewLinked[int](data_int...)
	lkl_str_copy := list.NewLinked[string](data_str...)

	if !lkl_int_copy.ContainsAny(sub_int) {
		t.Errorf("LinkedList[int] ContainsAny(LinkedList[int]) is false, expected true")
	}

	if !lkl_str_copy.ContainsAny(sub_str) {
		t.Errorf("LinkedList[string] ContainsAny(LinkedList[string]) is false, expected true")
	}
}
