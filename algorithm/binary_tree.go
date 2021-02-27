package algorithm

import "fmt"

type TNode struct {
	Value     int
	leftNode  *TNode
	rightNode *TNode
}

func NewTNode(data int) *TNode {
	return &TNode{
		Value:     data,
		leftNode:  nil,
		rightNode: nil,
	}
}

func (n *TNode) insert(data int) {
	if data < n.Value {
		if n.leftNode != nil {
			n.leftNode.insert(data)
		} else {
			n.leftNode = NewTNode(data)
		}
	} else {
		if n.rightNode != nil {
			n.rightNode.insert(data)
		} else {
			n.rightNode = NewTNode(data)
		}
	}
}

func (n *TNode) middleShow() {
	if n.leftNode != nil {
		n.leftNode.middleShow()
	}

	fmt.Println(n.Value)
	if n.rightNode != nil {
		n.rightNode.middleShow()
	}
}

func (n *TNode) firstShow() {
	fmt.Println(n.Value)
	if n.leftNode != nil {
		n.leftNode.firstShow()
	}

	if n.rightNode != nil {
		n.rightNode.firstShow()
	}
}

func (n *TNode) search(data int) *TNode {
	if n.Value == data {
		return n
	} else if data < n.Value {
		if n.leftNode != nil {
			return n.leftNode.search(data)
		}
	} else {
		if n.rightNode != nil {
			return n.rightNode.search(data)
		}
	}

	return nil
}

func (n *TNode) reverse() {
	leftNode := n.leftNode
	n.leftNode = n.rightNode
	n.rightNode = leftNode

	if n.leftNode != nil {
		n.leftNode.reverse()
	}

	if n.rightNode != nil {
		n.rightNode.reverse()
	}
}

// =======================================

type BinaryTree struct {
	len  int
	root *TNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (b *BinaryTree) Insert(data int) {
	b.len++
	if b.root == nil {
		b.root = NewTNode(data)
	} else {
		b.root.insert(data)
	}
}

func (b *BinaryTree) Search(data int) *TNode {
	if b.root != nil {
		return b.root.search(data)
	}

	return nil
}

/*
中序
1、递归实现
2、非递归实现，借助栈保存右节点
*/
func (b *BinaryTree) MiddleShow() {
	if b.root != nil {
		b.root.middleShow()
	}
}

// 先序
func (b *BinaryTree) FirstShow() {
	if b.root != nil {
		b.root.firstShow()
	}
}

/*
逐层遍历
用一个队列保存节点，先进先出
*/
func (b *BinaryTree) LevelShow() {
	queue := make([]*TNode, 0)
	result := make([]int, 0)
	queue = append(queue, b.root)

	for len(queue) > 0 {
		result = append(result, queue[0].Value)

		if queue[0].leftNode != nil {
			queue = append(queue, queue[0].leftNode)
		}
		if queue[0].rightNode != nil {
			queue = append(queue, queue[0].rightNode)
		}

		queue = queue[1:]
	}

	fmt.Println(result)
}

// 翻转二叉树
func (b *BinaryTree) Reverse() {
	if b.root != nil {
		b.root.reverse()
	}
}

/*
参考
- https://learnku.com/articles/48638
*/
