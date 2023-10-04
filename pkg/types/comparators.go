package types

import "golang.org/x/exp/constraints"

func AsCmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func DsCmp[T constraints.Ordered](a, b T) int {
	return -AsCmp[T](a, b)
}
