package heap

import "testing"

func TestHeap(t *testing.T) {
	heap := NewHeap[int](func(a, b int) bool {
		return a < b
	})
	heap.Push(1)
	heap.Push(2)
	heap.Push(3)
	heap.Push(4)
	heap.Pop()
	heap.Pop()
	for i := 1; i <= len(heap.data)-1; i++ {
		println(heap.data[i])
	}
}
