package deque

type dequeNode[V any] struct {
	prev  *dequeNode[V]
	next  *dequeNode[V]
	value V
}

type Deque[V any] struct {
	size int
	head dequeNode[V]
	tail dequeNode[V]
}

func NewDeque[V any]() *Deque[V] {
	d := &Deque[V]{size: 0}
	d.head.next = &d.tail
	d.head.prev = nil
	d.tail.prev = &d.head
	d.tail.next = nil
	return d
}

func (d *Deque[V]) Clear() {
	d.head.next = &d.tail
	d.head.prev = nil
	d.tail.prev = &d.head
	d.tail.next = nil
	d.size = 0
}

func (d *Deque[V]) Empty() bool {
	return d.size == 0 // || d.head.next == &d.tail || d.tail.prev == &d.head
}

func (d *Deque[V]) Size() int {
	return d.size
}

func (d *Deque[V]) Front() (value V, ok bool) {
	if d.Empty() {
		return
	}
	return d.head.next.value, true
}

func (d *Deque[V]) Back() (value V, ok bool) {
	if d.Empty() {
		return
	}
	return d.tail.prev.value, true
}

func (d *Deque[V]) PushFront(value V) {
	node := &dequeNode[V]{value: value}
	d.head.next.prev = node
	node.next = d.head.next
	d.head.next = node
	node.prev = &d.head
	d.size++
}

func (d *Deque[V]) PushBack(value V) {
	node := &dequeNode[V]{value: value}
	d.tail.prev.next = node
	node.prev = d.tail.prev
	d.tail.prev = node
	node.next = &d.tail
	d.size++
}

func (d *Deque[V]) PopFront() (value V, ok bool) {
	if d.Empty() {
		return
	}
	node := d.head.next
	d.head.next = node.next
	node.next.prev = &d.head
	d.size--
	return node.value, true
}

func (d *Deque[V]) PopBack() (value V, ok bool) {
	if d.Empty() {
		return
	}
	node := d.tail.prev
	d.tail.prev = node.prev
	node.prev.next = &d.tail
	d.size--
	return node.value, true
}
