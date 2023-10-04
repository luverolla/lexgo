package types

import "fmt"

type Hashable interface {
	Hash() uint
}

type Comparable interface {
	Cmp(any) int
}

type Base interface {
	fmt.Stringer
	Hashable
	Comparable
}