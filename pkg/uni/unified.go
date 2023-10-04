package uni

import (
	"encoding/binary"
	"hash/fnv"
	"log"
	"math"

	"golang.org/x/exp/constraints"

	"github.com/luverolla/lexgo/pkg/types"
)

func Eq(a, b any) bool {
	return Cmp(a, b) == 0
}

func Cmp(a, b any) int {
	switch a.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
		conva := a.(int)
		convb := b.(int)
		return cmp(conva, convb)
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
		conva := a.(uint)
		convb := b.(uint)
		return cmp(conva, convb)
	case float32:
	case float64:
		conva := a.(float64)
		convb := b.(float64)
		return cmp(conva, convb)
	case string:
		conva := a.(string)
		convb := b.(string)
		return cmp(conva, convb)
	default:
		conva, oka := a.(types.Comparable)
		convb, okb := b.(types.Comparable)
		if !oka || !okb {
			log.Fatal("ERROR: [unified.Cmp] CANNOT COMPARE TYPES")
		}
		return conva.Cmp(convb)
	}
	return -1
}

func Hash(v any) uint {
	switch val := v.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	case uint:
	case uint8:
	case uint16:
	case uint32:
	case uint64:
		return uint(val)
	case float32:
	case float64:
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
	return 0
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

func hashFloat(v float32) uint {
	n := math.Float32bits(v)
	binary.NativeEndian.PutUint32(buf[:], n)
	hashgen.Reset()
	_, err := hashgen.Write(buf[:])
	if err != nil {
		log.Fatal("ERROR: [unified.Hash] HASHGEN.Write() FAILED")
	}
	return uint(hashgen.Sum32())
}

func hashString(v string) uint {
	hashgen.Reset()
	_, err := hashgen.Write([]byte(v))
	if err != nil {
		log.Fatal("ERROR: [unified.Hash] HASHGEN.Write() FAILED")
	}
	return uint(hashgen.Sum32())
}
