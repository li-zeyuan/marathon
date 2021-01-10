# go基础

### map

- hash冲突：map的底层数据结构是数组，当向map中存储一个kv时，通过hash计算得出这个kv应该存储在底层数组的哪个下标，如果在始之前该数组下标已经存在kv（前后两个kv的hash值一样），这时就产生了冲突。

- hash冲突解决：

  - 开放定址法：当存储kv产生hash冲突时，就从数组冲突下标往后查找，找到一个空值下标就存储该kv。
  - 拉链法：当产生hash冲突时，就在冲突的下标形成一个链表，通过指针相连接。

- go map的实现原理：

  - 底层是一个bucket数组，每个bucket可以存储8个kv，当超过8个后，会产生一个新的bucket，并通过overflow指针指向新bucket。tophash通常包含该bucket中每个键的hash值的高八位。

    ```go
    // bucket的结构
    type bmap{
    	//tophash通常包含该bucket中每个键的hash值的高八位
    	tophash [bucketCnt]uint8
        overflow *[]*bmap
    }
    ```

    

- kv存储的过程：当往map中存储kv时，对k进行hash，定位到底层数组的下标（bucket），k的hash值高8位和bucket的tophash对比，判断k是否已经存在。将kv存储到该bucket中，若bucket满了，新建一个新的bucket，并用overflow指向新的bucket。

- 参考

  - https://learnku.com/articles/35019
  - https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/

### slice与array 

- 相同点
  - len()获取长度，通过下表获取
  - 分配一块连续的内存空间

- array

  - 值类型
  - var，:= 创建，不可用make（运行时）、**append**、**copy**
  - 创建后长度、容量不可改变

- slice  

  ```go
  // slice 源代码表示
  type SliceHeader struct {
  	Data uintptr // 指向底层数组
  	Len  int
  	Cap  int
  }
  ```

  - 引用类型
  - var，:= 、make创建
  - 可扩充，1024个元素内2倍增长，往后1/4增长
  - 底层实现是指向一个array，容量为底层array的大小

- 使用场景

  - 一般使用slice
  - 不确定len用 `slice`，确定大小使用`array`。

### channel

- CSP的设计理念：

```go
  type hchan struct {
      qcount   uint           // total data in the queue
      dataqsiz uint           // size of the circular queue
      buf      unsafe.Pointer // points to an array of dataqsiz elements
      elemsize uint16
      closed   uint32
      elemtype *_type // element type
      sendx    uint   // send index
      recvx    uint   // receive index
      recvq    waitq  // list of recv waiters
      sendq    waitq  // list of send waiters
  
      // lock protects all fields in hchan, as well as several
      // fields in sudogs blocked on this channel.
      //
      // Do not change another G's status while holding this lock
      // (in particular, do not ready a G), as this can deadlock
      // with stack shrinking.
      lock mutex
  }
```

- 分有缓冲区、无缓冲区两种
- 通过lock mutex实现线程安全

### GMP

- go调度器。本质上是把go程序中大量的goroutine分配到少量的线程上执行，利用CPU的多核，实现并发。
  - ![](.\goBasics.assets\GMP调度模型.jpeg)
  - G：goroutine。go协程。
  - P：processor。处理器，go程序创建出来的协程，放到P队列中，供M执行
  - M：thread。内核线程， 是G的执行实体
- 调度过程
  - ![](.\goBasics.assets\调度过程.jpeg)
  - 1、go 关键字创建一个G（goroutine）
  - 2、将G放到P本地队列，或放到全局队列
  - 3、M从P本地队列中取出G执行；若P本地队列为空，从全局队列，或其他P'的本地队列中取。
  - 4、当M在执行G时产生了syscall或阻塞时，M会从本地队列中弹出一个可执行的G'出来执行，并把阻塞的G交给空闲的M服务（或创建一个新的M）
  - 5、M没有可执行G时，就会进入休眠状态，放入休眠M队列
- 参考：
  - https://learnku.com/articles/41728

### 深拷贝与浅拷贝

- 深拷贝：开辟新的内存空间，新旧对象不共享内存
  - 值类型数据赋值：array，struct...
  - 内建函数copy( )
- 浅拷贝：复制了指向对象的引用，并没有开辟新的内存地址，新旧对象指向同一个内存地址
  - 引用类型赋值：指针、slice、map...

### oop

### gc

### 逃逸分析
- 定义：go的内存分配由编译器完成，通过逃逸分析，决定内存分配是在栈上还是在堆上。若变量的生命周期是完全可知，则分配到栈上，否则分配到堆（逃逸）。
 - 编译器尽可能地内存分配到栈，几种内存分配到堆得情况（逃逸）
   	- 变量类型不确定
      	- 函数内暴露给外部的指针
      	- 变量所占内存较大
              	- 变量的大小不确定

- 逃逸分析的作用：写出更好的程序，使内存尽可能地分配到栈，减小gc压力，减少内存分配开销。
- 参考
  - 逃逸分析：https://mp.weixin.qq.com/s/xhBVv6JEPY8R3kCJlbirYw
  - 堆：https://www.jianshu.com/p/6b526aa481b1

### 内存分析

### gin框架

### Nginx

### TCP与UDP

