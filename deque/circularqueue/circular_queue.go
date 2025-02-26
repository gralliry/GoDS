package circularqueue

type CircularQueue[V any] struct {
	front int
	rear  int
	cap   int
	data  []V
}

func NewCircularQueue[V any](capacity int) *CircularQueue[V] {
	return &CircularQueue[V]{
		front: 0,
		rear:  0,
		cap:   capacity,
		data:  make([]V, capacity),
	}
}

func (cq *CircularQueue[V]) Enqueue(v V) bool {
	if cq.Full() {
		return false
	}
	cq.data[cq.rear] = v
	cq.rear = (cq.rear + 1) % cq.cap
	return true
}

func (cq *CircularQueue[V]) Dequeue() (V, bool) {
	if cq.Empty() {
		var zero V
		return zero, false
	}
	v := cq.data[cq.front]
	cq.front = (cq.front + 1) % cq.cap
	return v, true
}

func (cq *CircularQueue[V]) Empty() bool {
	return cq.front == cq.rear
}

func (cq *CircularQueue[V]) Full() bool {
	return (cq.rear+1)%cq.cap == cq.front
}

func (cq *CircularQueue[V]) Size() int {
	return (cq.rear - cq.front + cq.cap) % cq.cap
}
