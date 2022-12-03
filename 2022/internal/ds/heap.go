package ds

import (
	"container/heap"
	"golang.org/x/exp/constraints"
)

// MaxHeap is max heap with int values.
type MaxHeap[T constraints.Ordered] struct {
	vals *maxHeap[T]
}

// NewMaxHeap creates new instance of MaxHeap.
func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		vals: &maxHeap[T]{},
	}
}

// Push pushes new value onto the heap.
func (h *MaxHeap[T]) Push(v T) {
	heap.Push(h.vals, v)
}

// Pop removes and returns the maximum value from the heap.
//
// If heap is empty, returns default value of int.
func (h *MaxHeap[T]) Pop() T {
	var v T
	if !h.Empty() {
		v = heap.Pop(h.vals).(T)
	}

	return v
}

// Empty checks if the heap is empty.
func (h *MaxHeap[T]) Empty() bool {
	return h.vals.Len() == 0
}

// maxHeap is the implementation of heap.Interface.
type maxHeap[T constraints.Ordered] []T

// Assert that maxHeap implements heap.Interface.
var _ heap.Interface = (*maxHeap[int])(nil)

// Push pushes the element v onto the heap.
func (h *maxHeap[T]) Push(v any) {
	*h = append(*h, v.(T))
}

// Pop removes and returns the maximum element from the heap.
func (h *maxHeap[T]) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]

	return v
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (h *maxHeap[T]) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *maxHeap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Len is the number of elements in the collection.
func (h *maxHeap[T]) Len() int {
	return len(*h)
}
