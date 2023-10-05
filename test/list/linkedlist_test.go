package list_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/list"
	"github.com/luverolla/lexgo/pkg/types"
)

var data_int = []int{4, 7, 9, 2547347383, 50, 78, 9, 77, 32, 9}
var data_str = []string{"ciao", "", "becco", "\r\n", "hey!", "", "castoro"}

var lkl_int colls.List[int]
var lkl_str colls.List[string]

func TestLinkedListAdd(t *testing.T) {
	lkl_int = list.Lkd[int](data_int...)
	lkl_str = list.Lkd[string](data_str...)

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
	lkl_int_copy := list.Lkd[int](data_int...)
	lkl_str_copy := list.Lkd[string](data_str...)

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
	lkl_int_copy := list.Lkd[int](data_int...)
	lkl_str_copy := list.Lkd[string](data_str...)

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
	lkl_int_copy := list.Lkd[int](data_int...)
	lkl_str_copy := list.Lkd[string](data_str...)

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

	lkl_int_copy := list.Lkd[int](data_int...)
	lkl_str_copy := list.Lkd[string](data_str...)

	if !lkl_int_copy.ContainsAll(list.Lkd[int](data_int...)) {
		t.Errorf("LinkedList[int] ContainsAll(LinkedList[int]) is false, expected true")
	}

	if !lkl_str_copy.ContainsAll(list.Lkd[string](data_str...)) {
		t.Errorf("LinkedList[string] ContainsAll(LinkedList[string]) is false, expected true")
	}
}

func TestLinkedListContainsAny(t *testing.T) {

	test_data_int := []int{1000, 10, 0, 0, 9, 0, 1000}
	test_data_str := []string{"alpha", "beta", "becco", "gamma", "lambda"}

	sub_int := list.Lkd[int](test_data_int...)
	sub_str := list.Lkd[string](test_data_str...)

	lkl_int_copy := list.Lkd[int](data_int...)
	lkl_str_copy := list.Lkd[string](data_str...)

	if !lkl_int_copy.ContainsAny(sub_int) {
		t.Errorf("LinkedList[int] ContainsAny(LinkedList[int]) is false, expected true")
	}

	if !lkl_str_copy.ContainsAny(sub_str) {
		t.Errorf("LinkedList[string] ContainsAny(LinkedList[string]) is false, expected true")
	}
}

func TestLinkedListSlice(t *testing.T) {
	arl_int_copy := list.Lkd[int](data_int_2...)
	arl_str_copy := list.Lkd[string](data_str_2...)

	start := 2
	end := 6

	sub_int := arl_int_copy.Slice(start, end)
	sub_str := arl_str_copy.Slice(start, end)

	if sub_int.Size() != end-start {
		t.Errorf("LinkedList[int] Slice(%d, %d) size is %d, expected %d", start, end, sub_int.Size(), end-start)
	}

	if sub_str.Size() != end-start {
		t.Errorf("LinkedList[string] Slice(%d, %d) size is %d, expected %d", start, end, sub_str.Size(), end-start)
	}

	for i := start; i < end; i++ {
		val, _ := sub_int.Get(i - start)
		if *val != data_int_2[i] {
			t.Errorf("LinkedList[int] Slice(%d, %d) Get(%d) is %d, expected %d", start, end, i-start, *val, data_int_2[i])
		}
	}

	for i := start; i < end; i++ {
		val, _ := sub_str.Get(i - start)
		if *val != data_str_2[i] {
			t.Errorf("LinkedList[string] Slice(%d, %d) Get(%d) is %s, expected %s", start, end, i-start, *val, data_str_2[i])
		}
	}
}

func TestLinkedListSublist(t *testing.T) {
	arl_int_copy := list.Lkd[int](data_int_2...)
	arl_str_copy := list.Lkd[string](data_str_2...)

	var evenIntFilter types.Filter[int] = func(v int, args ...any) bool {
		return v%2 == 0
	}

	var emptyStrFilter types.Filter[string] = func(s string, args ...any) bool {
		return s == ""
	}

	valuesInt := make([]int, 0)
	valuesStr := make([]string, 0)

	for _, v := range data_int_2 {
		if evenIntFilter(v) {
			valuesInt = append(valuesInt, v)
		}
	}

	for _, v := range data_str_2 {
		if emptyStrFilter(v) {
			valuesStr = append(valuesStr, v)
		}
	}

	sub_int := arl_int_copy.Sublist(evenIntFilter)
	sub_str := arl_str_copy.Sublist(emptyStrFilter)

	if sub_int.Size() != len(valuesInt) {
		t.Errorf("LinkedList[int] Sublist() size is %d, expected %d", sub_int.Size(), len(valuesInt))
	}

	if sub_str.Size() != len(valuesStr) {
		t.Errorf("LinkedList[string] Sublist() size is %d, expected %d", sub_str.Size(), len(valuesStr))
	}

	for i, v := range valuesInt {
		val, _ := sub_int.Get(i)
		if *val != v {
			t.Errorf("LinkedList[int] Sublist() Get(%d) is %d, expected %d", i, *val, v)
		}
	}

	for i, v := range valuesStr {
		val, _ := sub_str.Get(i)
		if *val != v {
			t.Errorf("LinkedList[string] Sublist() Get(%d) is %s, expected %s", i, *val, v)
		}
	}
}
