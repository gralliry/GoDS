package avltree

import (
	"fmt"
	"testing"
)

func TestAvlTree(t *testing.T) {
	tree := NewAVLTree[int](func(i int, i2 int) int {
		return i - i2
	}, Unique)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	fmt.Println(tree.Search(1), tree.Search(2), tree.Search(3))
	tree.Remove(1)
	fmt.Println(tree.Search(1), tree.Search(2), tree.Search(3))
	tree.Remove(2)
	fmt.Println(tree.Search(1), tree.Search(2), tree.Search(3))
}
