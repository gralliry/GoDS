package array

type Array[T any] struct {
	slice []T
}

func New[T any]() *Array[T] {
	return &Array[T]{slice: make([]T, 0)}
}

func Of[T any](slice []T) *Array[T] {
	ary := &Array[T]{slice: make([]T, len(slice))}
	copy(ary.slice, slice)
	return ary
}

func (s *Array[T]) Get(index int) T {
	return s.slice[index]
}

func (s *Array[T]) Set(index int, value T) {
	s.slice[index] = value
}

func (s *Array[T]) Append(value T) {
	s.slice = append(s.slice, value)
}

func (s *Array[T]) Insert(index int, value T) {
	s.slice = append(s.slice[:index], *new(T))
	copy(s.slice[index+1:], s.slice[index:])
	s.slice[index] = value
}

func (s *Array[T]) Remove(index int) T {
	value := s.slice[index]
	copy(s.slice[index:], s.slice[index+1:])
	s.slice = s.slice[:len(s.slice)-1]
	return value
}

func (s *Array[T]) Extend(slice []T) {
	s.slice = append(s.slice, slice...)
}

func (s *Array[T]) Len() int {
	return len(s.slice)
}

func (s *Array[T]) Clear() {
	s.slice = make([]T, 0)
}

func (s *Array[T]) Range(f func(index int, value T) bool) {
	for i, v := range s.slice {
		if !f(i, v) {
			break
		}
	}
}

func (s *Array[T]) Filter(filter func(T) bool) {
	size := 0
	for i, elem := range s.slice {
		if filter(elem) {
			s.slice[size] = s.slice[i]
			size++
		}
	}
	s.slice = s.slice[:size]
}

func (s *Array[T]) Copy() *Array[T] {
	return Of[T](s.slice)
}

func (s *Array[T]) Slice() []T {
	slice := make([]T, len(s.slice))
	copy(slice, s.slice)
	return slice
}
