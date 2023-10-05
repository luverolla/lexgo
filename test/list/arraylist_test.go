package list_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/list"
	"github.com/luverolla/lexgo/pkg/types"
)

var data_int_2 = []int{4, 7, 9, 2547347383, 50, 78, 9, 77, 32, 9}
var data_str_2 = []string{"ciao", "", "becco", "\r\n", "hey!", "", "castoro"}

var arl_int colls.List[int]
var arl_str colls.List[string]

func TestArrayListAdd(t *testing.T) {
	arl_int = list.Arr[int](data_int_2...)
	arl_str = list.Arr[string](data_str_2...)

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
	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

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
	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

	intToRemove := 9
	strToRemove := ""

	expRemInt := 3
	expRemStr := 2

	arl_int_copy.RemoveAll(intToRemove)
	arl_str_copy.RemoveAll(strToRemove)

	if arl_int_copy.Size() != len(data_int_2)-expRemInt {
		t.Errorf("ArrayList[int] size is %d, expected %d", arl_int_copy.Size(), len(data_int_2)-1)
	}

	if arl_str_copy.Size() != len(data_str_2)-expRemStr {
		t.Errorf("ArrayList[string] size is %d, expected %d", arl_str_copy.Size(), len(data_str_2)-1)
	}

	i := 0
	j := 0
	for i < len(data_int_2) {
		if data_int_2[i] == intToRemove {
			i++
			continue
		}
		val, _ := arl_int_copy.Get(j)
		if *val != data_int_2[i] {
			t.Errorf("ArrayList[int] Get(%d) is %d, expected %d", j, *val, data_int_2[i])
		}
		i++
		j++
	}

	i = 0
	j = 0
	for i < len(data_str_2) {
		if data_str_2[i] == strToRemove {
			i++
			continue
		}
		val, _ := arl_str_copy.Get(j)
		if *val != data_str_2[i] {
			t.Errorf("ArrayList[int] Get(%d) is %s, expected %s", j, *val, data_str_2[i])
		}
		i++
		j++
	}

}

func TestArrayListContains(t *testing.T) {
	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

	for _, v := range data_int_2 {
		if !arl_int_copy.Contains(v) {
			t.Errorf("ArrayList[int] Contains(%d) is false, expected true", v)
		}
	}

	for _, v := range data_str_2 {
		if !arl_str_copy.Contains(v) {
			t.Errorf("ArrayList[string] Contains(%s) is false, expected true", v)
		}
	}
}

func TestArrayListContainsAll(t *testing.T) {

	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

	if !arl_int_copy.ContainsAll(list.Arr[int](data_int_2...)) {
		t.Errorf("ArrayList[int] ContainsAll(ArrayList[int]) is false, expected true")
	}

	if !arl_str_copy.ContainsAll(list.Arr[string](data_str_2...)) {
		t.Errorf("ArrayList[string] ContainsAll(ArrayList[string]) is false, expected true")
	}
}

func TestArrayListContainsAny(t *testing.T) {

	test_data_int := []int{1000, 10, 0, 0, 9, 0, 1000}
	test_data_str := []string{"alpha", "beta", "becco", "gamma", "lambda"}

	sub_int := list.Arr[int](test_data_int...)
	sub_str := list.Arr[string](test_data_str...)

	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

	if !arl_int_copy.ContainsAny(sub_int) {
		t.Errorf("ArrayList[int] ContainsAny(ArrayList[int]) is false, expected true")
	}

	if !arl_str_copy.ContainsAny(sub_str) {
		t.Errorf("ArrayList[string] ContainsAny(ArrayList[string]) is false, expected true")
	}
}

func TestArrayListSlice(t *testing.T) {
	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

	start := 2
	end := 6

	sub_int := arl_int_copy.Slice(start, end)
	sub_str := arl_str_copy.Slice(start, end)

	if sub_int.Size() != end-start {
		t.Errorf("ArrayList[int] Slice(%d, %d) size is %d, expected %d", start, end, sub_int.Size(), end-start)
	}

	if sub_str.Size() != end-start {
		t.Errorf("ArrayList[string] Slice(%d, %d) size is %d, expected %d", start, end, sub_str.Size(), end-start)
	}

	for i := start; i < end; i++ {
		val, _ := sub_int.Get(i - start)
		if *val != data_int_2[i] {
			t.Errorf("ArrayList[int] Slice(%d, %d) Get(%d) is %d, expected %d", start, end, i-start, *val, data_int_2[i])
		}
	}

	for i := start; i < end; i++ {
		val, _ := sub_str.Get(i - start)
		if *val != data_str_2[i] {
			t.Errorf("ArrayList[string] Slice(%d, %d) Get(%d) is %s, expected %s", start, end, i-start, *val, data_str_2[i])
		}
	}
}

func TestArrayListSublist(t *testing.T) {
	arl_int_copy := list.Arr[int](data_int_2...)
	arl_str_copy := list.Arr[string](data_str_2...)

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
		t.Errorf("ArrayList[int] Sublist() size is %d, expected %d", sub_int.Size(), len(valuesInt))
	}

	if sub_str.Size() != len(valuesStr) {
		t.Errorf("ArrayList[string] Sublist() size is %d, expected %d", sub_str.Size(), len(valuesStr))
	}

	for i, v := range valuesInt {
		val, _ := sub_int.Get(i)
		if *val != v {
			t.Errorf("ArrayList[int] Sublist() Get(%d) is %d, expected %d", i, *val, v)
		}
	}

	for i, v := range valuesStr {
		val, _ := sub_str.Get(i)
		if *val != v {
			t.Errorf("ArrayList[string] Sublist() Get(%d) is %s, expected %s", i, *val, v)
		}
	}
}
