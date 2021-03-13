package algorithm

import (
	"fmt"
	"testing"
)

func TestSortObj_QuicklySort(t *testing.T) {
	qList := new(SortObj)
	qList.List = []int{1, 3, 4, 5}
	qList.QuicklySort()
	fmt.Println(qList.List)
}

func TestBuble(t *testing.T)  {
	ls := []int{9, 3, 2, 7, 4}
	BubbleSort(ls)
	fmt.Println(ls)
}
func TestSelectSort(t *testing.T)  {
	ls := []int{9, 3, 2, 7, 4}
	SelectSort(ls)
	fmt.Println(ls)
}