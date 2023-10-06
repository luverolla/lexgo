package types_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/tau"
)

func TestUnifiedDo(t *testing.T) {

	a := int(1)
	b := int(2)

	a8 := int8(a)
	b8 := int8(b)

	if !tau.Eq(a, a) {
		t.Errorf("Eq(%v, %v) = false", a, a)
	}

	if tau.Eq(a, b) {
		t.Errorf("Eq(%v, %v) = true", a, b)
	}

	if tau.Eq(a8, b8) {
		t.Errorf("Eq(%v, %v) = true", a8, b8)
	}
}
