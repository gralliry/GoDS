package stack

type sNode[V any] struct {
	next  *sNode[V]
	value V
}

type Stack[V any] struct {
	size int
	head sNode[V]
}

func New[V any]() *Stack[V] {
	s := &Stack[V]{head: sNode[V]{}, size: 0}
	s.head.next = nil
	return s
}

func (s *Stack[V]) Empty() bool {
	return s.size == 0 // s.head.next = nil
}

func (s *Stack[V]) Size() int {
	return s.size
}

func (s *Stack[V]) Peek() (V, bool) {
	if s.Empty() {
		return *new(V), false
	}
	return s.head.next.value, true
}

func (s *Stack[V]) Clear() {
	s.head.next = nil
	s.size = 0
}

func (s *Stack[V]) Push(value V) {
	node := &sNode[V]{value: value}
	node.next = s.head.next
	s.head.next = node
	s.size++
}

func (s *Stack[V]) Pop() (V, bool) {
	if s.Empty() {
		return *new(V), false
	}
	node := s.head.next
	s.head.next = node.next
	s.size--
	return node.value, true
}
