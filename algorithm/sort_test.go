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
