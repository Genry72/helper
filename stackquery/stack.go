package stackquery

// Stack - стек (LIFO) (Last In First Out) - последним пришел, первым ушел.
type Stack[V any] struct {
	stack []V
}

func NewStack[V any](count int) *Stack[V] {
	return &Stack[V]{
		stack: make([]V, 0, count),
	}
}

func NewStackFromSlice[V any](sl []V) *Stack[V] {
	return &Stack[V]{
		stack: sl,
	}
}

func (s *Stack[V]) Len() int {
	return len(s.stack)
}

func (s *Stack[V]) Push(v V) {
	s.stack = append(s.stack, v)
}

func (s *Stack[V]) Pop() V {
	if s.Len() == 0 {
		panic("empty stack")
	}
	v := s.stack[s.Len()-1]
	s.stack = s.stack[:s.Len()-1]
	return v
}
