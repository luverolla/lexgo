package types_test

import (
	"testing"

	"github.com/luverolla/lexgo/pkg/types"
)

func TestUnifiedDo(t *testing.T) {

	a := int(1)
	b := int(2)

	a8 := int8(a)
	b8 := int8(b)

	if !types.Eq(a, a) {
		t.Errorf("Eq(%v, %v) = false", a, a)
	}

	if types.Eq(a, b) {
		t.Errorf("Eq(%v, %v) = true", a, b)
	}

	if types.Eq(a8, b8) {
		t.Errorf("Eq(%v, %v) = true", a8, b8)
	}
}
