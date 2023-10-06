package algo

import (
	"github.com/luverolla/lexgo/pkg/colls"
	"github.com/luverolla/lexgo/pkg/tau"
)

// Dijkstra algorithm
// It returns a map where
// the key is the node and the value is the distance from the start node
func Dijkstra[T any](graph colls.Graph[T], start T) (colls.Map[T, int], error) {
	// TODO
	return nil, nil
}

func BellmanFord[T any](graph colls.Graph[T], start, end T) tau.Iterator[T] {
	// TODO
	return nil
}

func FloydWarshall[T any](graph colls.Graph[T], start, end T) tau.Iterator[T] {
	// TODO
	return nil
}
