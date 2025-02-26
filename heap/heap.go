package heap

type Heap[V any] struct {
	compare func(a, b V) bool
	data    []V
}

func NewHeap[V any](compare func(a, b V) bool) *Heap[V] {
	return &Heap[V]{
		compare: compare,
		data:    make([]V, 1, 100),
	}
}

func (h *Heap[V]) Size() int {
	return len(h.data) - 1
}

func (h *Heap[V]) Empty() bool {
	return len(h.data) == 1
}

func (h *Heap[V]) Push(value V) {
	h.data = append(h.data, value)
	h.up(len(h.data) - 1)
}

func (h *Heap[V]) Pop() (value V, ok bool) {
	size := len(h.data) - 1
	if size == 0 {
		return h.data[0], false
	}
	h.data[0] = h.data[1]
	h.data[1] = h.data[size]
	h.data = h.data[:size]
	h.down(1, size-1)
	return h.data[0], true
}

func (h *Heap[V]) Top() (value V, ok bool) {
	if len(h.data) == 1 {
		return h.data[0], false
	}
	return h.data[1], true
}

func (h *Heap[V]) up(index int) {
	parent := index / 2
	if parent == 0 {
		return
	}
	if h.compare(h.data[index], h.data[parent]) {
		// The Node is lighter than Parent Node, rises
		h.swap(index, parent)
		h.up(parent)
	}
}

func (h *Heap[V]) down(index int, maxindex int) {
	lchild, rchild := index*2, index*2+1
	if lchild > maxindex {
		// The heap is a complete binary tree
		// If the left child does not exist, then the right child must not exist either.
		// So the node has no children
		return
	} else if rchild > maxindex {
		// This Node has only Left child
		if !h.compare(h.data[index], h.data[lchild]) {
			// The Node is heavier than L
			h.swap(index, lchild)
			h.down(lchild, maxindex)
		}
		return
	}
	isLheavier := h.compare(h.data[index], h.data[lchild])
	isRheavier := h.compare(h.data[index], h.data[rchild])
	if !isLheavier && !isRheavier {
		// heavier than L and R, adjust to the lightest child
		if h.compare(h.data[lchild], h.data[rchild]) {
			// L is lighter
			h.swap(index, lchild)
			h.down(lchild, maxindex)
		} else {
			// R is lighter
			h.swap(index, rchild)
			h.down(rchild, maxindex)
		}
		return
	} else if isLheavier && !isRheavier {
		// lighter than L, but heavier than R, adjust to R
		h.swap(index, rchild)
		h.down(rchild, maxindex)
	} else if !isLheavier && isRheavier {
		// lighter than R, but heavier than L, adjust to L
		h.swap(index, lchild)
		h.down(lchild, maxindex)
	}
	// This Node is heavier than L and R, No need to adjust
}

func (h *Heap[V]) swap(a, b int) {
	h.data[a], h.data[b] = h.data[b], h.data[a]
}
