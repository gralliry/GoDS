package blockqueue

import (
	"sync"
)

type BlockQueue[V any] struct {
	once  sync.Once
	queue chan V
}

func NewBlockQueue[V any](cap int) *BlockQueue[V] {
	return &BlockQueue[V]{
		queue: make(chan V, cap),
	}
}

func (bq *BlockQueue[V]) Size() int {
	return len(bq.queue)
}

func (bq *BlockQueue[V]) Push(value V) bool {
	select {
	case bq.queue <- value:
		return true
	default:
		return false
	}
}

func (bq *BlockQueue[V]) Pop() (V, bool) {
	value, ok := <-bq.queue
	return value, ok
}

func (bq *BlockQueue[V]) Close() {
	bq.once.Do(func() {
		close(bq.queue)
	})
}
