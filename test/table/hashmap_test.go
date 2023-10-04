package table_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/table"
)

var hm_keys = []string{"ciao", "becco", "hey!", "castoro"}
var hm_vals = []bool{true, false, true, false}

func TestHashMapAdd(t *testing.T) {
	hm := table.Hsh[string, bool]()
	if hm.Size() != 0 {
		t.Errorf("HashMap size is %d, expected %d", hm.Size(), 0)
	}

	for i, k := range hm_keys {
		hm.Put(k, hm_vals[i])
	}

	if hm.Size() != 4 {
		t.Errorf("HashMap size is %d, expected %d", hm.Size(), 4)
	}

	if !hm.Contains("ciao") {
		t.Errorf("HashMap has not key %s", "ciao")
	}

	if !hm.Contains("becco") {
		t.Errorf("HashMap has not key %s", "becco")
	}

	if !hm.Contains("hey!") {
		t.Errorf("HashMap has not key %s", "hey!")
	}

	if !hm.Contains("castoro") {
		t.Errorf("HashMap has not key %s", "castoro")
	}
}

func TestHashMapGet(t *testing.T) {
	hm := table.Hsh[string, bool]()
	for i, k := range hm_keys {
		hm.Put(k, hm_vals[i])
	}

	for i, k := range hm_keys {
		val, _ := hm.Get(k)
		if *val != hm_vals[i] {
			t.Errorf("HashMap Get(%s) is %v, expected %v", k, *val, hm_vals[i])
		}
	}
}

func TestHashMapRemove(t *testing.T) {
	hm_copy := table.Hsh[string, bool]()
	for i, k := range hm_keys {
		hm_copy.Put(k, hm_vals[i])
	}

	for _, k := range hm_keys {
		_, err := hm_copy.Remove(k)
		if err != nil {
			t.Errorf("HashMap Remove(%s) failed", k)
		}
	}

	if hm_copy.Size() != 0 {
		t.Errorf("HashMap size is %d, expected %d", hm_copy.Size(), 0)
	}

	for _, k := range hm_keys {
		if hm_copy.HasKey(k) {
			t.Errorf("HashMap has key %s", k)
		}
	}
}

func TestHashMapIter(t *testing.T) {
	hm := table.Hsh[string, bool]()
	for i, k := range hm_keys {
		hm.Put(k, hm_vals[i])
	}

	iter := hm.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !hm.HasKey(*data) {
			t.Errorf("HashMap has not key %s", *data)
		}
	}
}
