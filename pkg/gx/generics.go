package gx

import (
	"encoding/binary"
	"hash/fnv"
	"log"
	"math"
	"reflect"

	"golang.org/x/exp/constraints"

	"github.com/luverolla/lexgo/pkg/types"
)

// Checks for equality between two values of generic type
//
// See [Cmp]
func Eq(a, b any) bool {
	return Cmp(a, b) == 0
}

// Get maximum value from a list of values of generic type
//
// See [Cmp]
func Max(vals ...any) any {
	if len(vals) == 0 {
		log.Printf("WARNING: [unified.Max] No values given\r\n")
		return nil
	}
	max := vals[0]
	for _, val := range vals {
		if Cmp(val, max) > 0 {
			max = val
		}
	}
	return max
}

// Get minimum value from a list of values of generic type
//
// See [Cmp]
func Min(vals ...any) any {
	if len(vals) == 0 {
		log.Printf("WARNING: [unified.Min] No values given\r\n")
		return nil
	}
	min := vals[0]
	for _, val := range vals {
		if Cmp(val, min) < 0 {
			min = val
		}
	}
	return min
}

// Tries to compare two values of generic type
// Given a and b the two values (in this order), it returns:
//
//	-1 if a < b
//	 0 if a == b
//	 1 if a > b
//
// It accepts all types under the [constraints.Ordered] interface
// or types that implement the [types.Comparable] interface
// For the latter, the function [types.Comparable.Cmp] must be implemented
//
// If the types of a and b are not comparable, it will panic
func Cmp(a, b any) int {
	switch a.(type) {
	case int:
		a := a.(int)
		b := b.(int)
		return cmp(a, b)
	case int8:
		a := a.(int8)
		b := b.(int8)
		return cmp(a, b)
	case int16:
		a := a.(int16)
		b := b.(int16)
		return cmp(a, b)
	case int32:
		a := a.(int32)
		b := b.(int32)
		return cmp(a, b)
	case int64:
		a := a.(int64)
		b := b.(int64)
		return cmp(a, b)
	case uint:
		a := a.(uint)
		b := b.(uint)
		return cmp(a, b)
	case uint8:
		a := a.(uint8)
		b := b.(uint8)
		return cmp(a, b)
	case uint16:
		a := a.(uint16)
		b := b.(uint16)
		return cmp(a, b)
	case uint32:
		a := a.(uint32)
		b := b.(uint32)
		return cmp(a, b)
	case uint64:
		a := a.(uint64)
		b := b.(uint64)
		return cmp(a, b)
	case float32:
		a := a.(float32)
		b := b.(float32)
		return cmp(a, b)
	case float64:
		a := a.(float64)
		b := b.(float64)
		return cmp(a, b)
	case string:
		a := a.(string)
		b := b.(string)
		return cmp(a, b)
	default:
		a, oka := a.(types.Comparable)
		b, okb := b.(types.Comparable)
		if !oka || !okb {
			log.Fatalf("ERROR: [unified.Cmp] Types %T and %T are not comparable\r\n", a, b)
		}
		return a.Cmp(b)
	}
}

// Hashes a value of generic type
// It accepts all numeric and string types
// It also accepts types that implement the [types.Hashable] interface
// For the latter, the function [types.Hashable.Hash] must be implemented
//
// If the type of given value is not hashable, it will panic
func Hash(v any) uint32 {
	switch val := v.(type) {
	case int, int8, int16, int32, int64:
		s := reflect.ValueOf(val).Int()
		return uint32(s)
	case uint, uint8, uint16, uint32, uint64:
		s := reflect.ValueOf(val).Uint()
		return uint32(s)
	case float32, float64:
		return hashFloat(v.(float32))
	case string:
		return hashString(val)
	default:
		conv, ok := val.(types.Hashable)
		if !ok {
			log.Fatal("ERROR: [unified.Hash] CANNOT HASH TYPE")
		}
		return conv.Hash()
	}
}

// --- Private variables ---
var hashgen = fnv.New32a()
var buf [4]byte

// --- Private functions ---
func cmp[T constraints.Ordered](a, b T) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func hashFloat(v float32) uint32 {
	n := math.Float32bits(v)
	binary.NativeEndian.PutUint32(buf[:], n)
	hashgen.Reset()
	_, err := hashgen.Write(buf[:])
	if err != nil {
		log.Fatal("ERROR: [unified.Hash] HASHGEN.Write() FAILED")
	}
	return hashgen.Sum32()
}

func hashString(v string) uint32 {
	hashgen.Reset()
	_, err := hashgen.Write([]byte(v))
	if err != nil {
		log.Fatal("ERROR: [unified.Hash] HASHGEN.Write() FAILED")
	}
	return hashgen.Sum32()
}
