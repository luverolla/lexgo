package types

import (
	"golang.org/x/exp/constraints"
)

func cmp[T constraints.Ordered](a, b T) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func Eq(a, b any) bool {
	return Cmp(a, b) == 0
}

func Cmp(a, b any) int {
	switch a.(type) {
	case int:
		conva := a.(int)
		convb := b.(int)
		return cmp(conva, convb)
	case int8:
		conva := a.(int8)
		convb := b.(int8)
		return cmp(conva, convb)
	case int16:
		conva := a.(int16)
		convb := b.(int16)
		return cmp(conva, convb)
	case int32:
		conva := a.(int32)
		convb := b.(int32)
		return cmp(conva, convb)
	case int64:
		conva := a.(int64)
		convb := b.(int64)
		return cmp(conva, convb)
	case uint:
		conva := a.(uint)
		convb := b.(uint)
		return cmp(conva, convb)
	case uint8:
		conva := a.(uint8)
		convb := b.(uint8)
		return cmp(conva, convb)
	case uint16:
		conva := a.(uint16)
		convb := b.(uint16)
		return cmp(conva, convb)
	case uint32:
		conva := a.(uint32)
		convb := b.(uint32)
		return cmp(conva, convb)
	case uint64:
		conva := a.(uint64)
		convb := b.(uint64)
		return cmp(conva, convb)
	case float32:
		conva := a.(float32)
		convb := b.(float32)
		return cmp(conva, convb)
	case float64:
		conva := a.(float64)
		convb := b.(float64)
		return cmp(conva, convb)
	case string:
		conva := a.(string)
		convb := b.(string)
		return cmp(conva, convb)
	default:
		conva, oka := a.(Base)
		convb, okb := b.(Base)
		if !oka || !okb {
			panic("[unified.Cmp] CANNOT COMPARE TYPES")
		}
		return conva.Cmp(convb)
	}
}
