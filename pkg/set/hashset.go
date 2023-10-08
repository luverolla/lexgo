// This package contains implementations of the [tau.Set] interface.
package set

import (
	"fmt"

	"github.com/luverolla/lexgo/pkg/table"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Implementation of a set using a hash table
type HshSet[T any] struct {
	table *table.HshMap[T, any]
}

// Creates a new hash set
func Hsh[T any]() *HshSet[T] {
	return &HshSet[T]{table.Hsh[T, any]()}
}

// --- Methods from Collection[T] ---
func (set *HshSet[T]) String() string {
	s := "HshSet{"
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

func (set *HshSet[T]) Cmp(other any) int {
	otherSet, ok := other.(*HshSet[T])
	if !ok {
		panic(fmt.Sprintf("ERROR: [HshSet.Cmp] %v is not a *HshSet", other))
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

func (set *HshSet[T]) Size() int {
	return set.table.Size()
}

func (set *HshSet[T]) Empty() bool {
	return set.table.Empty()
}

func (set *HshSet[T]) Clear() {
	set.table.Clear()
}

func (set *HshSet[T]) Contains(value T) bool {
	return set.table.HasKey(value)
}

func (set *HshSet[T]) ContainsAll(coll tau.Collection[T]) bool {
	iter := coll.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if !set.table.HasKey(*next) {
			return false
		}
	}
	return true
}

func (set *HshSet[T]) ContainsAny(coll tau.Collection[T]) bool {
	iter := coll.Iter()
	for next, ok := iter.Next(); ok; next, ok = iter.Next() {
		if set.table.HasKey(*next) {
			return true
		}
	}
	return false
}

func (set *HshSet[T]) Clone() tau.Collection[T] {
	return &HshSet[T]{set.table.Clone().(*table.HshMap[T, any])}
}

func (set *HshSet[T]) Iter() tau.Iterator[T] {
	return set.table.Keys()
}

// --- Methods from Set[T] ---
func (set *HshSet[T]) Add(values ...T) {
	for _, value := range values {
		if !set.table.HasKey(value) {
			set.table.Put(value, nil)
		}
	}
}

func (set *HshSet[T]) Remove(value T) error {
	_, err := set.table.Remove(value)
	return err
}
