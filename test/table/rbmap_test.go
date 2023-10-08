package table_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/table"
)

var rb_keys = []string{"ciao", "becco", "hey!", "castoro", "1245", "mannagg", "gesucrist"}
var rb_vals = []bool{true, false, true, false, true, true, false}

func TestRBMapAdd(t *testing.T) {
	tm := table.RB[string, bool]()

	if tm.Size() != 0 {
		t.Errorf("RBMap size is %d, expected %d", tm.Size(), 0)
	}

	for i, k := range rb_keys {
		tm.Put(k, rb_vals[i])

	}
}

func TestRBMapGet(t *testing.T) {
	tm := table.RB[string, bool]()
	for i, k := range rb_keys {
		tm.Put(k, rb_vals[i])
	}

	for i, k := range rb_keys {
		val, err := tm.Get(k)
		if err != nil {
			t.Errorf("RBMap Get(%s) failed", k)
		}
		if *val != rb_vals[i] {
			t.Errorf("RBMap Get(%s) is %v, expected %v", k, *val, rb_vals[i])
		}
	}
}

func TestRBMapRemove(t *testing.T) {
	rb_copy := table.RB[string, bool]()
	for i, k := range rb_keys {
		rb_copy.Put(k, rb_vals[i])
	}

	for _, k := range rb_keys {
		_, err := rb_copy.Remove(k)
		if err != nil {
			t.Errorf("RBMap Remove(%s) failed", k)
		}
		if rb_copy.HasKey(k) {
			t.Errorf("RBMap still contains key %s", k)
		}
	}

	if rb_copy.Size() != 0 {
		t.Errorf("RBMap size is %d, expected %d", rb_copy.Size(), 0)
	}

	for _, k := range rb_keys {
		if rb_copy.HasKey(k) {
			t.Errorf("RBMap has key %s", k)
		}
	}
}

func TestRBMapIter(t *testing.T) {
	tm := table.RB[string, bool]()
	for i, k := range rb_keys {
		tm.Put(k, rb_vals[i])
	}

	iter := tm.Iter()
	for data, ok := iter.Next(); ok; data, ok = iter.Next() {
		if !tm.HasKey(*data) {
			t.Errorf("RBMap has not key %s", *data)
		}
	}
}
