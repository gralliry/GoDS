package delayqueue

import (
	"github.com/gralliry/gods/heap"
	"math"
	"sync"
	"time"
)

type delayqueueNode[V any] struct {
	value      V
	expireTime time.Time
}

type DelayQueue[V any] struct {
	input     chan *delayqueueNode[V]
	output    chan V
	onceStart sync.Once
	onceClose sync.Once
	root      *heap.Heap[*delayqueueNode[V]]
}

func NewDelayQueue[V any]() *DelayQueue[V] {
	dq := &DelayQueue[V]{
		root: heap.NewHeap[*delayqueueNode[V]](func(a, b *delayqueueNode[V]) bool {
			return a.expireTime.Before(b.expireTime)
		}),
		input:  make(chan *delayqueueNode[V], 100),
		output: make(chan V, 100),
	}
	return dq
}

func (q *DelayQueue[V]) Pop() (value V, ok bool) {
	value, ok = <-q.output
	return value, ok
}

func (q *DelayQueue[V]) Push(value V, exTime time.Time) (ok bool) {
	select {
	case q.input <- &delayqueueNode[V]{
		value:      value,
		expireTime: exTime,
	}:
		return true
	default:
		return false
	}
}

func (q *DelayQueue[V]) Start() {
	q.onceStart.Do(func() {
		go func() {
			var (
				timegap   time.Duration
				dqn, node *delayqueueNode[V]
				ok        bool
			)
			for {
				if dqn, ok = q.root.Top(); ok {
					timegap = dqn.expireTime.Sub(time.Now())
					if timegap < 0 {
						q.output <- dqn.value
						q.root.Pop()
						continue
					}
				} else {
					timegap = math.MaxInt64
				}
				select {
				case <-time.After(timegap):
					if node, ok = q.root.Pop(); ok {
						q.output <- node.value
					}
				case node, ok = <-q.input:
					if !ok {
						break
					}
					if node.expireTime.Before(time.Now()) {
						q.output <- node.value
					} else {
						q.root.Push(node)
					}
				}
			}
		}()
	})
}

func (q *DelayQueue[V]) Close() {
	q.onceClose.Do(func() {
		close(q.input)
		close(q.output)
	})
}
