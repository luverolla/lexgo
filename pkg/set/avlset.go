package set

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/table"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Implementation of a set using an AVL tree
type AVLSet[T any] struct {
	table *table.AVLMap[T, any]
}

// Creates a new hash set
func AVL[T any]() *AVLSet[T] {
	return &AVLSet[T]{table.AVL[T, any]()}
}

// --- Methods from Collection[T] ---
func (set *AVLSet[T]) String() string {
	s := "AVLSet{"
	iter := set.Iter()
	first := false
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if first {
			first = false
		} else {
			s += ", "
		}
		s += fmt.Sprintf("%v", *next)
	}
	s += "}"
	return s
}

func (set *AVLSet[T]) Cmp(other any) int {
	otherSet, ok := other.(*AVLSet[T])
	if !ok {
		panic(fmt.Sprintf("ERROR: [AVLSet.Cmp] %v is not a *AVLSet", other))
	}
	if set.Size() != otherSet.Size() {
		return set.Size() - otherSet.Size()
	}
	iter, otherIter := set.Iter(), otherSet.Iter()
	for next, hasNext := iter.Next(); hasNext; next, hasNext = iter.Next() {
		otherNext, _ := otherIter.Next()
		cmp := tau.Cmp(*next, *otherNext)
		if cmp != 0 {
			return cmp
		}
	}
	return 0
}

func (set *AVLSet[T]) Size() int {
	return set.table.Size()
}

func (set *AVLSet[T]) Empty() bool {
	return set.table.Empty()
}

func (set *AVLSet[T]) Clear() {
	set.table.Clear()
}

func (set *AVLSet[T]) Contains(value T) bool {
	return set.table.HasKey(value)
}

func (set *AVLSet[T]) ContainsAll(coll tau.Collection[T]) bool {
	iter := coll.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !set.table.HasKey(*next) {
			return false
		}
	}
	return true
}

func (set *AVLSet[T]) ContainsAny(coll tau.Collection[T]) bool {
	iter := coll.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if set.table.HasKey(*next) {
			return true
		}
	}
	return false
}

func (set *AVLSet[T]) Iter() tau.Iterator[T] {
	return set.table.Keys()
}

func (set *AVLSet[T]) Clone() tau.Collection[T] {
	return &AVLSet[T]{set.table.Clone().(*table.AVLMap[T, any])}
}

// --- Methods from Set[T] ---
func (set *AVLSet[T]) Add(values ...T) {
	for _, value := range values {
		if !set.table.HasKey(value) {
			set.table.Put(value, nil)
		}
	}
}

func (set *AVLSet[T]) Remove(value T) error {
	_, err := set.table.Remove(value)
	return err
}
