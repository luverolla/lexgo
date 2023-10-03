package deque

import "github.com/luverolla/lexgo/pkg/types"

type Deque[T any] interface {
	types.Collection[T]
	PushFront(...T)
	PushBack(...T)
	PopFront() (*T, error)
	PopBack() (*T, error)
	Front() (*T, error)
	Back() (*T, error)
}

type DequeImpl int

const (
	ADQ DequeImpl = iota
	LDQ
)

func New[T any](impl DequeImpl) Deque[T] {
	switch impl {
	case LDQ:
		return nil // TODO: Implement LinkedDeque
	default:
		return NewArrayDeque[T]()
	}
}
