package priorityqueue

import (
	"GoDS/heap"
)

type pqNode[V any, P any] struct {
	value    V
	priority P
}

type PriorityQueue[V any, P any] struct {
	compare func(a, b P) bool
	root    *heap.Heap[*pqNode[V, P]]
}

func New[V any, P any](compare func(a, b P) bool) *PriorityQueue[V, P] {
	ncompare := func(a, b *pqNode[V, P]) bool {
		return compare(a.priority, b.priority)
	}
	return &PriorityQueue[V, P]{
		root: heap.New[*pqNode[V, P]](ncompare),
	}
}

func (q *PriorityQueue[V, P]) Empty() bool {
	return q.root.Empty()
}

func (q *PriorityQueue[V, P]) Size() int {
	return q.root.Size()
}

func (q *PriorityQueue[V, P]) Top() (V, bool) {
	node, empty := q.root.Top()
	return node.value, empty
}

func (q *PriorityQueue[V, P]) Pop() (V, bool) {
	node, empty := q.root.Pop()
	return node.value, empty
}

func (q *PriorityQueue[V, P]) Push(value V, priority P) {
	q.root.Push(&pqNode[V, P]{
		value:    value,
		priority: priority,
	})
}
