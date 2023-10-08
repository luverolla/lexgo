package set

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/table"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Implementation of a set using an RB tree
type RBSet[T any] struct {
	table *table.RBMap[T, any]
}

// Creates a new hash set
func RB[T any]() *RBSet[T] {
	return &RBSet[T]{table.RB[T, any]()}
}

// --- Methods from Collection[T] ---
func (set *RBSet[T]) String() string {
	s := "RBSet{"
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

func (set *RBSet[T]) Cmp(other any) int {
	otherSet, ok := other.(*RBSet[T])
	if !ok {
		panic(fmt.Sprintf("ERROR: [RBSet.Cmp] %v is not a *RBSet", other))
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

func (set *RBSet[T]) Size() int {
	return set.table.Size()
}

func (set *RBSet[T]) Empty() bool {
	return set.table.Empty()
}

func (set *RBSet[T]) Clear() {
	set.table.Clear()
}

func (set *RBSet[T]) Contains(value T) bool {
	return set.table.HasKey(value)
}

func (set *RBSet[T]) ContainsAll(coll tau.Collection[T]) bool {
	iter := coll.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !set.table.HasKey(*next) {
			return false
		}
	}
	return true
}

func (set *RBSet[T]) ContainsAny(coll tau.Collection[T]) bool {
	iter := coll.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if set.table.HasKey(*next) {
			return true
		}
	}
	return false
}

func (set *RBSet[T]) Iter() tau.Iterator[T] {
	return set.table.Keys()
}

func (set *RBSet[T]) Clone() tau.Collection[T] {
	return &RBSet[T]{set.table.Clone().(*table.RBMap[T, any])}
}

// --- Methods from Set[T] ---
func (set *RBSet[T]) Add(values ...T) {
	for _, value := range values {
		if !set.table.HasKey(value) {
			set.table.Put(value, nil)
		}
	}
}

func (set *RBSet[T]) Remove(value T) error {
	_, err := set.table.Remove(value)
	return err
}
