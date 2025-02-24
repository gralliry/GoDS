package avltree

type aNode[V any] struct {
	left   *aNode[V]
	right  *aNode[V]
	height int
}

type AVLTree[V any] struct {
	compare func(V, V) bool
	root    *aNode[V]
}

func New[V any]() *AVLTree[V] {
	return &AVLTree[V]{}
}

func (t *AVLTree[V]) adjust(node *aNode[V]) *aNode[V] {
	if node == nil {
		return node
	}
	// blance factor
	bF := t.blanceFactor(node)
	// adjust
	if bF > 1 {
		son := node.left
		sBF := t.blanceFactor(son)
		if sBF >= 0 {
			// LL
			node.left = son.right
			son.right = node
			return son
		} else {
			// LR
			grandson := son.right
			son.right = grandson.left
			grandson.left = son
			node.left = grandson.right
			grandson.right = node
			return grandson
		}
	} else if bF < -1 {
		son := node.right
		sBF := t.blanceFactor(son)
		if sBF >= 0 {
			// RL
			grandson := son.left
			son.left = grandson.right
			grandson.right = son
			node.right = grandson.left
			grandson.left = node
			return grandson
		} else {
			// RR
			node.right = son.left
			son.left = node
			return son
		}
	} else {
		return node
	}
}

func (t *AVLTree[V]) height(node *aNode[V]) int {
	if node == nil {
		return 0
	}
	lH, rH := 0, 0
	if node.left != nil {
		lH = node.left.height
	}
	if node.right != nil {
		rH = node.right.height
	}
	return max(lH, rH) + 1
}

func (t *AVLTree[V]) blanceFactor(node *aNode[V]) int {
	if node == nil {
		return 0
	}
	lH, rH := 0, 0
	if node.left != nil {
		lH = node.left.height
	}
	if node.right != nil {
		rH = node.right.height
	}
	return lH - rH
}
