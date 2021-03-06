package algorithm

import (
	"testing"

	"fmt"
)

func TestReversedSlice(t *testing.T) {
	strList := []string{"A", "B", "C", "D"}
	sList := ReversedSlice(strList)
	fmt.Println(sList)
}
