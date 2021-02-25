package algorithm

import (
	"fmt"
	"testing"
)

func TestDoubleLink_Add(t *testing.T) {
	doubleLink := NewDoubleLink(1)
	doubleLink.Add(2)
	fmt.Println(doubleLink)
}

func TestDoubleLink_Scan(t *testing.T) {
	doubleLink := NewDoubleLink(1)
	doubleLink.Add(2)
	doubleLink.Scan()
}

func TestDoubleLink_Reverse(t *testing.T) {
	doubleLink := NewDoubleLink(1)
	doubleLink.Add(2)
	doubleLink.Reverse()

	doubleLink.Scan()
}

func TestDoubleLink_Length(t *testing.T) {
	doubleLink := NewDoubleLink(2)
	doubleLink.Add(1)
	fmt.Println(doubleLink.Length())
}

func TestDoubleLink_Append(t *testing.T) {
	doubleLink := NewDoubleLink(1)
	doubleLink.Append(2)

	doubleLink.Scan()
}

func TestDoubleLink_Insert(t *testing.T) {
	doubleLink := NewDoubleLink(1)
	doubleLink.Append(3)

	doubleLink.Insert(2, 2)
	doubleLink.Scan()
}

func TestDoubleLink_Del(t *testing.T) {
	doubleLink := NewDoubleLink(1)
	doubleLink.Append(2)
	doubleLink.Append(3)

	doubleLink.Del(1)
	doubleLink.Scan()
}
