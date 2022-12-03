package ds

// Set is set data structure that contains int values.
type Set[T comparable] struct {
	vals map[T]bool
}

// NewSet creates new instance of Set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		vals: make(map[T]bool),
	}
}

// Vals returns the values of the set.
func (s *Set[T]) Vals() []T {
	vals := make([]T, s.Len())

	i := 0
	for v := range s.vals {
		vals[i] = v
		i += 1
	}

	return vals
}

// Add adds the value to the set.
func (s *Set[T]) Add(v T) {
	s.vals[v] = true
}

// Has checks if the value exists in the set.
func (s *Set[T]) Has(v T) bool {
	_, exists := s.vals[v]
	return exists
}

// Remove removes the value from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.vals, v)
}

// Len returns the length of the values.
func (s *Set[T]) Len() int {
	return len(s.vals)
}

// SetIntersection returns the intersection of the two sets.
func SetIntersection[T comparable](a, b *Set[T]) *Set[T] {
	s := NewSet[T]()

	for v := range a.vals {
		if b.Has(v) {
			s.Add(v)
		}
	}

	return s
}
