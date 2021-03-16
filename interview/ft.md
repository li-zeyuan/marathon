# 前面的笔试题

可以查看保存的Word文档 
 然后流程是每题进行解析

（中间defer的问题是答错了的，）

1. 开始问一些项目相关的东西（业务，团队文档，工作）
2. 有听说过或者用过errgroup吗？  
    当时答：没有
3. 如果我有一堆goroutine，并发执行一个同样的逻辑，要获取第一个错误，有什么办法 
    当时答： 主goroutine 用select，一个是接收子gorotine中的返回值，channel size为1，然后返回

更好的答案： errgroup里面的实现，通过context来继续处理，用了sync.Once,然后触发所有的goroutine的context.Cancel，然后返回第一个错误

1. 如果我有一堆goroutine，并发执行一个同样的逻辑，要获取所有错误，有什么办法 
    当时答： 主goroutine 用select，一个是接收子gorotine中的返回值，channel size为goroutine数，然后返回

更好的答案：需要基于errgroup的代码上面修改

 

```
type Group struct {
    ...
    errs   []error   // 
    returnAll bool 
}
func (g *Group) Go(f func() error) {
     ...
    if err := f(); err != nil {
           if !g.returnAll{
            g.errOnce.Do(func() {
                g.err = err
                if g.cancel != nil {
                    g.cancel()
                }
            })
        } else{
            g.errs = append(g.errs, err)
       }
    }
}
```

通过这样就能解决返回全部的问题。

1. 算法题： leetcode 179（原题） 
    解题思路：排序 
    主要特殊的点：

 

```
func(s strCompare) Less(i,j int) bool {
    o12 := s[i] + s[j]
    o21 := s[j] + s[i]
    return o12 > o21 
}
```

1. 看你做过6.824, 能解释一下Raft 算法吗？ 
    当时答：  基本简单解释
2. 数学题： 甲乙两人抛硬币，甲先抛，请问乙的获胜概率是多少 
    当时答： 推导了很久都没算出来，最后是面试官给出提示来进行 
    答案： 1/4 的等比数列和

# 二面

1 
 ![img](C:\Users\XJY\Desktop\learnproject\src\readygo\interview\index_files\edee65af-a1d5-4ee5-b339-2fb558ddc7be.png) 
 当时处理： 
 这个不知道什么是字节序，然后面试官更换了题目

后面回溯可以用下面方法进行判断：

 

```
package main
import (
    "encoding/binary"
    "fmt"
    "unsafe"
)
const INT_SIZE int = int(unsafe.Sizeof(0))
//判断我们系统中的字节序类型
func systemEdian() {
    var i int = 0x1
    bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
    if bs[0] == 0 {
        fmt.Println("system edian is little endian")
    } else {
        fmt.Println("system edian is big endian")
    }
}
func testBigEndian() {
    // 0000 0000 0000 0000   0000 0001 1111 1111
    var testInt int32 = 256
    fmt.Printf("%d use big endian: \n", testInt)
    var testBytes []byte = make([]byte, 4)
    binary.BigEndian.PutUint32(testBytes, uint32(testInt))
    fmt.Println("int32 to bytes:", testBytes)
    convInt := binary.BigEndian.Uint32(testBytes)
    fmt.Printf("bytes to int32: %d\n\n", convInt)
}
func testLittleEndian() {
    // 0000 0000 0000 0000   0000 0001 1111 1111
    var testInt int32 = 256
    fmt.Printf("%d use little endian: \n", testInt)
    var testBytes []byte = make([]byte, 4)
    binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
    fmt.Println("int32 to bytes:", testBytes)
    convInt := binary.LittleEndian.Uint32(testBytes)
    fmt.Printf("bytes to int32: %d\n\n", convInt)
}
func main() {
    systemEdian()
    fmt.Println("")
    testBigEndian()
    testLittleEndian()
}
```

参考资料： 
 https://www.cnblogs.com/-wenli/p/12323809.html

\2. 
 ![img](C:\Users\XJY\Desktop\learnproject\src\readygo\interview\index_files\36bd0c2d-936c-4226-a733-fcf2e7d90293.png)

当时回答： 性能表现、原理，场景设计

1. ![img](C:\Users\XJY\Desktop\learnproject\src\readygo\interview\index_files\e2cb64cc-1ffe-4eb8-a91e-3081008189be.png) 
    架构设计，数据结构设计能力,这种设计方案占用的元数据大小

用户ID int64, TS int64,  Operator uint(byte=1) 
 当时回答：  
 问题一： 
 使用了每个用户维护一个堆（堆中元素为365个，每天一个） 
 里面的结构如下

 

```
type Item struct {
    TS int64
    Data &[]uint
}
```

1. Tcp协议中Seq字段（非Flag字段的用处）

当时回答： 1. 三次握手，四次挥手  2. TCP拥塞控制协议(解释了一下拥塞协议）

![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210316134213.png)

 