package algorithm

import (
	"fmt"
	"sync"
	"time"
)

// 单例
var (
	single *Singleton
	sOne   sync.Once
)

type Singleton struct{}

func GetSingleton() *Singleton {
	sOne.Do(func() {
		single = new(Singleton)
	})

	return single
}

// 装饰器
func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func coolFunc(a func()) { // go 通过函数传参实现装饰器
	fmt.Printf("Starting function execution: %s\n", time.Now())
	a()
	fmt.Printf("End of function execution: %s\n", time.Now())
}

// 闭包
