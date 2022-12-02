package ds

import "container/heap"

// IntMaxHeap is max heap with int values.
type IntMaxHeap struct {
	vals *intMaxHeap
}

// NewIntMaxHeap creates new instance of IntMaxHeap.
func NewIntMaxHeap() *IntMaxHeap {
	return &IntMaxHeap{
		vals: &intMaxHeap{},
	}
}

// Push pushes new value onto the heap.
func (h *IntMaxHeap) Push(v int) {
	heap.Push(h.vals, v)
}

// Pop removes and returns the maximum value from the heap.
//
// If heap is empty, returns default value of int.
func (h *IntMaxHeap) Pop() int {
	var v int
	if !h.Empty() {
		v = heap.Pop(h.vals).(int)
	}

	return v
}

// Empty checks if the heap is empty.
func (h *IntMaxHeap) Empty() bool {
	return h.vals.Len() == 0
}

// intMaxHeap is the implementation of heap.Interface.
type intMaxHeap []int

// Assert that intMaxHeap implements heap.Interface.
var _ heap.Interface = (*intMaxHeap)(nil)

// Push pushes the element v onto the heap.
func (h *intMaxHeap) Push(v any) {
	*h = append(*h, v.(int))
}

// Pop removes and returns the maximum element from the heap.
func (h *intMaxHeap) Pop() any {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]

	return v
}

// Less reports whether the element with index i
// must sort before the element with index j.
func (h *intMaxHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *intMaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Len is the number of elements in the collection.
func (h *intMaxHeap) Len() int {
	return len(*h)
}
