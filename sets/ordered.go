package sets

// New creates new `UnOrdered` set data structure given list of items
func New[T comparable](v ...T) *UnOrdered[T] {
	set := NewWithCap[T](len(v))
	set.Add(v...)
	return set
}

// New creates new `UnOrdered` set data structure with pre-defined cap
func NewWithCap[T comparable](cap int) *UnOrdered[T] {
	return &UnOrdered[T]{data: make(map[T]struct{}, cap)}
}

// UnOrdered implements un-orderd set data struct
type UnOrdered[T comparable] struct {
	data map[T]struct{}
}

func (s *UnOrdered[T]) Add(items ...T) *UnOrdered[T] {
	for _, i := range items {
		s.data[i] = struct{}{}
	}
	return s
}

func (s *UnOrdered[T]) Length() int {
	return len(s.data)
}

func (s *UnOrdered[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *UnOrdered[T]) Pop(v T) *UnOrdered[T] {
	delete(s.data, v)
	return s
}

func (s *UnOrdered[T]) Merge(next *UnOrdered[T]) *UnOrdered[T] {
	for k := range next.data {
		s.Add(k)
	}
	return s
}

func (s *UnOrdered[T]) Copy() *UnOrdered[T] {
	return New(s.Values()...)
}

func (s *UnOrdered[T]) Values() []T {
	out := make([]T, 0, s.Length())
	for k := range s.data {
		out = append(out, k)
	}
	return out
}
