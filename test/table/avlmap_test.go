package table_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/table"
)

var tm_keys = []string{"ciao", "becco", "hey!", "castoro"}
var tm_vals = []bool{true, false, true, false}

func TestAVLMapAdd(t *testing.T) {
	tm := table.AVL[string, bool]()

	if tm.Size() != 0 {
		t.Errorf("AVLMap size is %d, expected %d", tm.Size(), 0)
	}

	for i, k := range tm_keys {
		tm.Put(k, tm_vals[i])
	}

	if tm.Size() != 4 {
		t.Errorf("AVLMap size is %d, expected %d", tm.Size(), 4)
	}

	if !tm.Contains("ciao") {
		t.Errorf("AVLMap has not key %s", "ciao")
	}

	if !tm.Contains("becco") {
		t.Errorf("AVLMap has not key %s", "becco")
	}

	if !tm.Contains("hey!") {
		t.Errorf("AVLMap has not key %s", "hey!")
	}

	if !tm.Contains("castoro") {
		t.Errorf("AVLMap has not key %s", "castoro")
	}
}

func TestAVLMapGet(t *testing.T) {
	tm := table.AVL[string, bool]()
	for i, k := range tm_keys {
		tm.Put(k, tm_vals[i])
	}

	for i, k := range tm_keys {
		val, _ := tm.Get(k)
		if *val != tm_vals[i] {
			t.Errorf("AVLMap Get(%s) is %v, expected %v", k, *val, tm_vals[i])
		}
	}
}

func TestAVLMapRemove(t *testing.T) {
	tm_copy := table.AVL[string, bool]()
	for i, k := range tm_keys {
		tm_copy.Put(k, tm_vals[i])
	}

	for _, k := range tm_keys {
		_, err := tm_copy.Remove(k)
		if err != nil {
			t.Errorf("AVLMap Remove(%s) failed", k)
		}
	}

	if tm_copy.Size() != 0 {
		t.Errorf("AVLMap size is %d, expected %d", tm_copy.Size(), 0)
	}

	for _, k := range tm_keys {
		if tm_copy.HasKey(k) {
			t.Errorf("AVLMap has key %s", k)
		}
	}
}

func TestAVLMapIter(t *testing.T) {
	tm := table.AVL[string, bool]()
	for i, k := range tm_keys {
		tm.Put(k, tm_vals[i])
	}

	iter := tm.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !tm.HasKey(*data) {
			t.Errorf("AVLMap has not key %s", *data)
		}
	}
}
