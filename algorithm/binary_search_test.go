package algorithm

import (
	"fmt"
	"testing"
)

func TestBinarySearchObj_BinarySearch(t *testing.T) {
	binaryS := NewBSO([]int{1, 2, 3}, 2)
	fmt.Println(binaryS.BinarySearch())
}

func Test1(t *testing.T) {
	A()
}

func A() {
	p := Person{}
	p.howOld(2)
	p.howOld2("lizeyuan")
}

type Person struct {
	name string
	age int
}

func (p Person) howOld(i int){
	p.age = i
}

func (p *Person) howOld2(s string){
	p.name = s
}