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
	fmt.Print(fmt.Sprint(i + 65))
	fmt.Print(fmt.Sprint(i + 65 + 1))
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

// ============================================================
/*
为 sync.WaitGroup 中Wait函数支持 WaitTimeout 功能.
思路
1、一个goroutine去处理timeout，一个goroutine去处理wg.Wait()
2、chan去阻塞，获取优先返回的goroutine
*/
func WaitTimeoutDo() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}

	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	// 要求手写代码
	// 要求sync.WaitGroup支持timeout功能
	// 如果timeout到了超时时间返回true
	// 如果WaitGroup自然结束返回false

	c := make(chan bool)

	go func() {
		ticker := time.NewTicker(timeout)
		<-ticker.C
		c <- true
	}()

	go func() {
		wg.Wait()
		c <- false
	}()

	return <-c
}
