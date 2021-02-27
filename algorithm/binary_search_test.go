package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearchObj_BinarySearch(t *testing.T) {
	binaryS := NewBSO([]int{1, 2, 3}, 2)
	fmt.Println(binaryS.BinarySearch())
}
