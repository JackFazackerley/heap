package heap

import (
	"golang.org/x/exp/constraints"
)

type compare[T constraints.Ordered] interface {
	Compare() T
}

func (h Heap[T, X]) Len() int           { return len(h) }
func (h Heap[T, X]) Less(i, j int) bool { return h[i].Compare() < h[j].Compare() }
func (h Heap[T, X]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
