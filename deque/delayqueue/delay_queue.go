package delayqueue

import (
	"GoDS/heap"
	"sync"
	"time"
)

type dqNode[V any] struct {
	value V
	time  time.Time
}

type DelayQueue[V any] struct {
	input  chan *dqNode[V]
	output chan V
	once   sync.Once
	root   *heap.Heap[*dqNode[V]]
}

func New[V any]() *DelayQueue[V] {
	dq := &DelayQueue[V]{
		root: heap.New[*dqNode[V]](func(a, b *dqNode[V]) bool {
			return a.time.Before(b.time)
		}),
		input:  make(chan *dqNode[V], 100),
		output: make(chan V, 100),
	}
	return dq
}

func (q *DelayQueue[V]) Pop() V {
	return <-q.output
}

func (q *DelayQueue[V]) Push(value V, time time.Time) {
	q.input <- &dqNode[V]{
		value: value,
		time:  time,
	}
}

func (q *DelayQueue[V]) Start() {
	q.once.Do(func() {
		go func() {
			for {
				var timegap time.Duration
				dqn, isEmpty := q.root.Top()
				if isEmpty {
					timegap = 0xfffffffffffffff
				} else {
					timegap = dqn.time.Sub(time.Now())
				}
				select {
				case <-time.After(timegap):
					dqn, _ := q.root.Pop()
					q.output <- dqn.value
				case dqn := <-q.input:
					if dqn.time.Sub(time.Now()) < 0 {
						q.output <- dqn.value
					} else {
						q.root.Push(dqn)
					}
				}
			}
		}()
	})
}
