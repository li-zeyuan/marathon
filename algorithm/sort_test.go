package algorithm

import (
	"fmt"
	"sync"
	"testing"
)

func TestSortObj_QuicklySort(t *testing.T) {
	qList := new(SortObj)
	qList.List = []int{1, 3, 4, 5}
	qList.QuicklySort()
	fmt.Println(qList.List)
}

func TestDoubleLink_Add2(t *testing.T) {
		threads := []string{"A", "B", "C", "D", "E"}
		n := len(threads)
		// 开启 n 个 goroutine, 每个 goroutine 只输出自身内容, 各自重复 n 遍, 要求所有 goroutine 的输出结果是有序的, 以上面 threads 为例, 要求输出 "ABCDEABCDEABCDEABCDE", threadA 只输出 "AAAAA", threadB 只输出 "BBBBB"
		wg := sync.WaitGroup{}
		c := make(chan string)
		for i := 0; i < n; i++ {
			for _, t := range threads {
				wg.Add(5)
				b := t
				c <- b
				go func() {
					defer wg.Done()
					a := <-c
					fmt.Print(a)
				}()
			}
		}

		wg.Wait()
		close(c)
}