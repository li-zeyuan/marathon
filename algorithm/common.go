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

// ===========================

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

// ==================================

// iota :一个const中，iota初始值为0（第一行为0），可跳过，可占位
const (
	a = iota
	b = iota
)
const (
	name = "menglu" // 占位
	c    = iota
	d    = iota
)

func GetIota() {
	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 1
	fmt.Println(d) // 2
}

/*
求阶乘和
1+2！+3！+4！
*/
// 求一个数阶乘
func Fac(i int) int {
	if i == 1 {
		return i
	}

	return i * Fac(i-1)
}

func Sum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += Fac(i)
	}

	return sum
}

/*
斐波纳契数列，又称黄金分割数列，指的是这样一个数列：1、1、2、3、5、8、13、21、……
在数学上，斐波纳契数列以如下被以递归的方法定义：F0=0，F1=1，Fn=F(n-1)+F(n-2)（n>=2，n∈N*）
*/
func PriFib(num int) {
	ch := make(chan int)

	go func(ch chan int, n int) {
		pre, cur := 0, 1
		for i := 0; i < n; i++ {
			ch <- cur
			pre, cur = cur, pre+cur
		}

		close(ch)
	}(ch, num)

	for i := range ch {
		fmt.Print(i)
	}
}

/*
回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

输入：x = 121
输出：true
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
输入：x = -101
输出：false

思路：反转后半数字，和前一半数字做比较
1、负数、和最后一位为0的数，不是回文数
2、x重复/10，为取末位数做准备
3、result 上一次的值 * 10，再加上x值的末位
4、若该回文数的长度为奇数，后半数字得 /10 去掉一位再和前半数字做比较
*/
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	result := 0
	for x > result {
		result = result*10 + x%10 // result 上一次的值 * 10，再加上x值的末位
		x = x / 10                // x重复/10，为取末位数做准备
	}

	return result == x || x == result/10
}
