package ds

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Peek() T {
	var v T
	if s.Len() > 0 {
		v = (*s)[s.Len()-1]
	}

	return v
}

func (s *Stack[T]) Pop() T {
	var v T
	if s.Len() > 0 {
		n := s.Len()
		v = (*s)[n-1]
		*s = (*s)[:n-1]
	}

	return v
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
