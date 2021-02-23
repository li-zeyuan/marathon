package algorithm

import "fmt"

/*
单向链表
h->1->2->nil
 */

type Node struct {
	data interface{}
	Next *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
	}
}

type LList struct {
	header *Node
}

func NewList() *LList {
	return &LList{}
}

// 表头增加节点
func (l *LList) Add(data interface{}) {
	n := NewNode(data)
	if l.header == nil {
		l.header = n
		n.Next = new(Node)
	} else {
		n.Next = l.header
		l.header = n
		n.Next = new(Node)
	}
}

func (l *LList) Append(data interface{}) {
	n := NewNode(data)
	if l.header == nil {
		l.header = n
		n.Next = new(Node)
	} else {
		curNode := l.header
		for curNode.Next != nil {
			curNode = curNode.Next
		}

		curNode.data = data
		curNode.Next = new(Node)
	}
}

func (l *LList) Insert(i int, data interface{}) {
	if i <= 0 {
		l.Add(data)
		return
	}
	if i >= l.Length() {
		l.Append(data)
		return
	}

	curNode := l.header
	curIdx := 0
	for curNode != nil {
		if curIdx == i {
			n := NewNode(data)
			n.Next = curNode.Next
			curNode.Next = n
			break
		}
		curIdx ++
	}
}

// 翻转链表：遍历链表，把当前的node.next指向前一个节点
func (l *LList)Revert() {
	preNode := new(Node)

	curNode := l.header
	lastNode := new(Node)
	for curNode != nil {
		lastNode = curNode

		next := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = next
	}

	l.header = lastNode
}

func (l *LList)Delete(i int) {
 	if i + 1 > l.Length() {
		return
	}

	if i == 0 {
		l.header = l.header.Next
		return
	}

	curNode := l.header
	lastNode := l.header
	for curNode.Next != nil{
		if i == 0 {
			lastNode.Next = curNode.Next
			break
		}

		i --
		lastNode = curNode
		curNode = curNode.Next
	}
}

func (l *LList) Length() int {
	curNode := l.header

	length := 0
	for curNode.Next != nil {
		length++
		curNode = curNode.Next
	}

	return length
}

func (l *LList) Scan() {
	curNode := l.header

	for curNode != nil {
		fmt.Print(curNode.data, " ")
		curNode = curNode.Next
	}
	fmt.Println()
}
