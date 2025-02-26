package queue

type queueNode[V any] struct {
	value V
	next  *queueNode[V]
}

type Queue[V any] struct {
	head *queueNode[V]
	tail *queueNode[V]
	size int
}

func NewQueue[V any]() *Queue[V] {
	sentinel := &queueNode[V]{}
	return &Queue[V]{
		head: sentinel,
		tail: sentinel,
		size: 0,
	}
}

func (q *Queue[V]) Empty() bool {
	return q.size == 0
}

func (q *Queue[V]) Size() int {
	return q.size
}

func (q *Queue[V]) Peek() (V, bool) {
	if q.Empty() {
		return *new(V), false
	}
	return q.head.next.value, true
}

func (q *Queue[V]) Push(value V) {
	node := &queueNode[V]{value: value}
	q.tail.next = node
	q.tail = node
	q.size++
}

func (q *Queue[V]) Pop() (value V, ok bool) {
	if q.Empty() {
		return
	}
	node := q.head.next
	q.head.next = node.next
	if node == q.tail {
		q.tail = q.head
	}
	q.size--
	return node.value, true
}
