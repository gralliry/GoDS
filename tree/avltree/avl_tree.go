// Package avltree https://www.youtube.com/results?search_query=avl+tree
package avltree

type Mode int

const (
	Multiple Mode = iota
	Unique
)

type aNode[V any] struct {
	v V
	l *aNode[V]
	r *aNode[V]
	h int
}

type AVLTree[V any] struct {
	mode    Mode
	compare func(V, V) int
	root    *aNode[V]
}

func NewAVLTree[V any](compare func(V, V) int, mode Mode) *AVLTree[V] {
	return &AVLTree[V]{compare: compare, root: nil, mode: mode}
}

func (t *AVLTree[V]) Insert(v V) (ok bool) {
	t.root, ok = t.insert(t.root, v)
	return ok
}

func (t *AVLTree[V]) Remove(v V) {
	t.root = t.delete(t.root, v)
}

func (t *AVLTree[V]) Search(v V) bool {
	return t.search(t.root, v)
}

func (t *AVLTree[V]) search(node *aNode[V], v V) bool {
	if node == nil {
		return false
	}
	compareRes := t.compare(v, node.v)
	if compareRes < 0 {
		return t.search(node.l, v)
	} else if compareRes > 0 {
		return t.search(node.r, v)
	} else {
		return true
	}
}

func (t *AVLTree[V]) insert(node *aNode[V], v V) (adjustedNode *aNode[V], ok bool) {
	if node == nil {
		return &aNode[V]{v: v, h: 1, l: nil, r: nil}, true
	}
	compareRes := t.compare(v, node.v)
	if compareRes < 0 {
		adjustedNode, ok = t.insert(node.l, v)
		node.l = adjustedNode
		node = t.adjust(node)
		return node, ok
	} else if compareRes > 0 {
		adjustedNode, ok = t.insert(node.l, v)
		node.r = adjustedNode
		node = t.adjust(node)
		return node, ok
	} else if t.mode == Unique {
		return node, false
	} else if t.mode == Multiple {
		adjustedNode, ok = t.insert(node.l, v)
		node.r = adjustedNode
		node = t.adjust(node)
		return node, ok
	} else {
		// log.Fatal("Unknown Mode")
		return node, false
	}
}

func (t *AVLTree[V]) delete(node *aNode[V], v V) *aNode[V] {
	if node == nil {
		return nil
	}
	compareRes := t.compare(v, node.v)
	if compareRes < 0 {
		node.l = t.delete(node.l, v)
		return node
	} else if compareRes > 0 {
		node.r = t.delete(node.r, v)
		return node
	} else if node.l != nil && node.r != nil {
		// find the max node in l subtree
		node.l = t.getMaxNode(node.l, node.r)
		return node
	} else if node.l == nil {
		return node.r
	} else if node.r == nil {
		return node.l
	} else {
		return nil
	}
}

func (t *AVLTree[V]) adjust(node *aNode[V]) *aNode[V] {
	if node == nil {
		return node
	}
	// blance factor
	bF := t.blanceFactor(node)
	// adjust
	if bF > 1 {
		son := node.l
		sBF := t.blanceFactor(son)
		if sBF >= 0 {
			// LL
			node.l = son.r
			son.r = node
			// adjust height
			node.h -= 2
			return son
		} else {
			// LR
			grandson := son.r
			son.r = grandson.l
			grandson.l = son
			node.l = grandson.r
			grandson.r = node
			// adjust height
			node.h -= 2
			son.h -= 1
			grandson.h += 1
			return grandson
		}
	} else if bF < -1 {
		son := node.r
		sBF := t.blanceFactor(son)
		if sBF >= 0 {
			// RL
			grandson := son.l
			son.l = grandson.r
			grandson.r = son
			node.r = grandson.l
			grandson.l = node
			// adjust height
			node.h -= 2
			son.h -= 1
			grandson.h += 1
			return grandson
		} else {
			// RR
			node.r = son.l
			son.l = node
			// adjust height
			node.h -= 2
			return son
		}
	} else {
		return node
	}
}

func (t *AVLTree[V]) blanceFactor(node *aNode[V]) int {
	if node == nil {
		return 0
	}
	lH, rH := 0, 0
	if node.l != nil {
		lH = node.l.h
	}
	if node.r != nil {
		rH = node.r.h
	}
	return lH - rH
}

func (t *AVLTree[V]) getMaxNode(node *aNode[V], subnode *aNode[V]) *aNode[V] {
	if node == nil {
		return nil
	}
	if node.r != nil {
		node.r = t.getMaxNode(node.r, subnode)
		return node
	} else {
		node.r = subnode
		return node
	}
}
