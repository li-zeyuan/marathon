package algorithm

import (
	"fmt"
	"sync"
	"testing"

	"github.com/bmizerany/assert"
	"time"
)

func TestGetSingleton(t *testing.T) {
	s1 := GetSingleton()
	s2 := GetSingleton()

	if s1 != s2 {
		t.Fatal("not equal")
	}

	assert.Equal(t, s1, s2)
}

func TestDecorator(t *testing.T) {
	coolFunc(myFunc)
}

func TestGetIota(t *testing.T) {
	GetIota()
}

func TestFac(t *testing.T) {
	fmt.Println(Sum(4))
}

func TestPriFib(t *testing.T) {
	PriFib(5)
}

func TestChanel(t *testing.T) {
	ch := make(chan int)

	close(ch)
	ch <- 1

	fmt.Println("iii")
}

func TestIsPalindrome(t *testing.T)  {
	fmt.Println(isPalindrome(123121))
}

func TestLargestNumber(t *testing.T)  {
	str := largestNumber([]int{3,30,34,5,9})
	fmt.Println(str)
}

func TestRW(t *testing.T)  {
	rw := sync.RWMutex{}
	wg := sync.WaitGroup{}

	rw.Lock()

	go func() {
		wg.Add(1)
		defer rw.RUnlock()
		defer wg.Done()
		fmt.Println("111")
		rw.RLock()
		fmt.Println("222")
	}()

	time.Sleep(time.Second)
	rw.Unlock()
	wg.Wait()

	fmt.Println()
}