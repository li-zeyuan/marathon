package algorithm

import (
	"fmt"
	"testing"
)

func TestLList_Add(t *testing.T) {
	lList := NewList()
	lList.Add(1)
	lList.Add(2)

	lList.Scan()
}

func TestLList_Append(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	lList.Scan()
}

func TestLList_Length(t *testing.T) {
	lList := NewList()

	lList.Append(1)
	lList.Append(2)
	fmt.Println(lList.Length())
}
