package algorithm

import (
	"fmt"
	"testing"
)

func TestBinaryTree_Search(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(4)

	fmt.Println(tree.Search(4))
}

func TestBinaryTree_MiddleShow(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(8)

	tree.MiddleShow()
}

func TestNewBinaryTree(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(8)

	tree.FirstShow()
}

func TestBinaryTree_LevelShow(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(3)

	tree.LevelShow()
}
