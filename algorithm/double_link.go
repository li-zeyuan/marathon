package algorithm

import (
	"fmt"
)

type DoubleNode struct {
	Data interface{}
	Next *DoubleNode
	Pre  *DoubleNode
}

func NewDNode(data interface{}) *DoubleNode {
	return &DoubleNode{
		Data: data,
	}
}

type DoubleLink struct {
	Head *DoubleNode
}

func NewDoubleLink(data interface{}) DoubleLink {
	node := NewDNode(data)
	return DoubleLink{
		Head: node,
	}
}

func (d *DoubleLink) Add(data interface{}) {
	node := NewDNode(data)
	d.Head.Pre = node

	node.Next = d.Head
	d.Head = node
}

func (d *DoubleLink) Append(data interface{}) {
	node := NewDNode(data)

	curNode := d.Head
	for curNode.Next != nil {
		curNode = curNode.Next
	}

	curNode.Next = node
	node.Pre = curNode
}

func (d *DoubleLink) Insert(i int, data interface{}) {
	if i == 0 {
		d.Add(data)
		return
	}
	if i >= d.Length() {
		d.Append(data)
		return
	}

	curNode := d.Head
	for j := 0; j < d.Length(); j++ {
		if j == i {
			node := NewDNode(data)
			node.Next = curNode
			curNode.Pre.Next = node
			node.Pre = curNode.Pre
			curNode.Pre = node
			break
		}

		curNode = curNode.Next
	}
}

func (d *DoubleLink) Del(i int) {
	if i < 0 || i > d.Length()-1 {
		return
	}

	if i == 0 {
		h := d.Head
		d.Head = d.Head.Next
		d.Head.Pre = nil
		h.Next = nil
		return
	}

	curNode := d.Head
	for i < d.Length() {
		if i == 0 {
			curNode.Pre.Next = curNode.Next
			curNode.Pre = curNode.Next
			break
		}

		curNode = curNode.Next
		i--
	}
}

func (d *DoubleLink) Length() int {
	i := 0
	curNode := d.Head
	for curNode != nil {
		i++
		curNode = curNode.Next
	}

	return i
}

func (d *DoubleLink) Scan() {
	curNode := d.Head
	for curNode != nil {
		fmt.Print(curNode.Data, " ")
		curNode = curNode.Next
	}

	fmt.Println()
}

func (d *DoubleLink) Reverse() {
	pre := new(DoubleNode)
	curNode := d.Head
	for curNode != nil {
		next := curNode.Next
		curNode.Pre = curNode.Next
		if pre.Data == nil {
			pre = nil
		}
		curNode.Next = pre

		pre = curNode
		curNode = next
	}

	d.Head = pre
}

/*
参考
https://studygolang.com/articles/18042
*/
