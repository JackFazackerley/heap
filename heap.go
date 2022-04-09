package heap

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// ErrEmptyHeap is the error returned by Pop when there are no more results on the heap.
var ErrEmptyHeap = errors.New("empty heap")

type Heap[T compare[X], X constraints.Ordered] []T

func New[T compare[X], X constraints.Ordered]() *Heap[T, X] {
	return &Heap[T, X]{}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T, X]) Push(x T) {
	*h = append(*h, x)
	h.up(h.Len() - 1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap[T, X]) Pop() (T, error) {
	var result T

	n := h.Len()
	if n <= 0 {
		return result, ErrEmptyHeap
	}

	n0 := n - 1
	h.Swap(0, n0)
	h.down(0, n0)

	old := *h
	result = old[n-1]
	*h = old[0 : n-1]

	return result, nil
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T, X]) Remove(i int) (T, error) {
	var result T

	n := h.Len() - 1

	if i < 0 || i > n {
		return result, errors.New("out of range")
	}

	if n != i {
		h.Swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T, X]) Fix(i int) {
	if !h.down(i, h.Len()) {
		h.up(i)
	}
}

func (h *Heap[T, X]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func (h *Heap[T, X]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
