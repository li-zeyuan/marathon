package algorithm

import (
	"fmt"
	"testing"
	"time"
)

func TestMap_Get_1(t *testing.T) {
	myMap := NewMap()

	fmt.Println(myMap.Get("lizeyuan", time.Second*3))
}

func TestMap_Get_2(t *testing.T) {
	myMap := NewMap()

	myMap.Set("lizeyuan", 1)
	fmt.Println(myMap.Get("lizeyuan", time.Second*3))
}

func TestMap_Get_3(t *testing.T) {
	myMap := NewMap()

	go myMap.Set("lizeyuan", 1)
	fmt.Println(myMap.Get("lizeyuan", time.Second*10))
}