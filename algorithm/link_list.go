package algorithm

import "fmt"

// 单向链表

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
	if l.header == nil {
		l.header = NewNode(data)
	} else {
		n := NewNode(data)
		n.Next = l.header
		l.header = n
	}
}

func (l *LList) Append(data interface{}) {
	if l.header == nil {
		l.header = NewNode(data)
	} else {
		curNode := l.header
		for curNode.Next != nil {
			curNode = curNode.Next
		}

		curNode.Next = NewNode(data)
	}
}

// todo
func (l *LList) Insert(i int, data interface{}) {
	//if i <= 0 {
	//	l.Add(data)
	//}
	//if i >= l.Length() {
	//	l.Append(data)
	//}
	//
	//curNode := l.header
	//for j := 0 ; j > i; j ++ {
	//
	//}
}

func (l *LList) Length() int {
	curNode := l.header

	length := 0
	for curNode != nil {
		length++
		curNode = curNode.Next
	}

	return length
}

func (l *LList) Scan() {
	curNode := l.header

	for curNode != nil {
		fmt.Println(curNode.data)
		curNode = curNode.Next
	}
}
