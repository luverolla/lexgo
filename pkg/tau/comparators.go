package tau

import "golang.org/x/exp/constraints"

// ASCending order comparator. The smaller is the lesser
func ASCmp[T constraints.Ordered](a, b T) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// DeSCending order comparator. The smaller is the greater
func DSCmp[T constraints.Ordered](a, b T) int {
	return -ASCmp[T](a, b)
}
