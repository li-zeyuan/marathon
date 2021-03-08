package algorithm

import (
	"fmt"
	"sync"
	"time"
)

/*
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母，
最终效果如下：12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ

思路
	- wg做并发同步控制
	- chan用于goroutine间传递信息
*/

func AlternatePrint(n int) {
	wg := new(sync.WaitGroup)
	c := make(chan int, 1)
	for i := 0; i < n; i = i + 2 {
		wg.Add(2)
		c <- i
		go printDigit(i, c, wg) // 传i值进去，而不是从chan中那
		c <- i
		go printAlphabet(i, c, wg)
	}

	wg.Wait()
}

func printDigit(i int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-c
	fmt.Print(i + 1)
	fmt.Print(i + 2)
}

func printAlphabet(i int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-c
	fmt.Print(string(string(i + 65)))
	fmt.Print(string(string(i + 65 + 1)))
}

// ==============================================
/*
写出以下逻辑，要求每秒钟调用一次proc并保证程序不退出
*/
func TickerPanic() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		ticker := time.NewTicker(time.Second)
		for {
			<-ticker.C
			go func() {
				defer func() {
					_ = recover()
				}()

				proc()
			}()
		}

	}()

	select {}
}

func proc() {
	fmt.Println("aa")
	panic("ok")
}
