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

- CSP的设计理念：channel
- 在go语音中实现goroutine间通信，分有缓存区和无缓存区
- 分单向和双向模式

- 结构
	```go
      type hchan struct {
       qcount   uint      	// 队列中剩余元素数量
       dataqsiz uint     		// 循环队列的长度（channel的大小）
       buf      unsafe.Pointer // 长度为dataqsiz的底层数组指针，缓存型channel特有
       elemsize uint16		// 元素大小	
       closed   uint32		// 是否关闭
       elemtype *_type 		// 接收、发送的的元素类型
       sendx    uint  		// 已发送元素在循环队列中的索引位置
       recvx    uint  		// 已接收元素在循环队列中的索引位置
       recvq    waitq  		// 接收者sudog等待队列（阻塞等待接收的goroutine）
       sendq    waitq  		// 发送者sudog等待对列（阻塞等待接收的goroutine）
      
       lock mutex				// 互斥锁
      }
  ```
- 发送数据
	```flow
    st=>start: Start
    op=>operation: make(chan T, size)
    op1=>operation: 发送
    op2=>operation: 数据写入buf
    op3=>operation: 取出一个G
    op4=>operation: 数据写入G
    op5=>operation: 唤醒G
    op6=>operation: 将G加入sendq队列
    op7=>operation: 待换醒
    
    cond=>condition: recvq非空？
    cond1=>condition: buf有空位？
    e=>end
    
    st->op->op1->cond
    cond(yes)->op3->op4->op5->e
    cond(no)->cond1
    cond1(yes)->op2->e
    cond1(no)->op6->op7->e
  ```
  - 情况1：recvq有G；直接取出G，写入数据并唤醒。
  - 情况2：recvq没有G，buf有空位；将数据写入buf。
  - 情况3：recvq没有G，buf没有空位；当前goroutine阻塞，并加入sendq队列
  
- 接收数据

    ```flow
    st=>start: Start
    op1=>operation: 接收
    op3=>operation: 取出一个G
    op4=>operation: 从G中读取数据
    op5=>operation: 唤醒G
    op6=>operation: 将G加入sendq队列
    op7=>operation: 待换醒
    op8=>operation: 从buf队首取出数据
    op9=>operation: 取出一个sendG
    op10=>operation: 将sendG数据写入buf队尾
    op11=>operation: 唤醒sendG
    op12=>operation: 从buf中取出数据
    op13=>operation: 将当前的recveG加入recveq
    op14=>operation: recveG待唤醒
    
    cond=>condition: sendq非空？
    cond1=>condition: 有buf？
    cond2=>condition: buf非空？
    e=>end
    
    st->op1->cond
    cond(yes)->cond1
    cond(no)->cond2
    
    cond1(yes)->op8->op9->op10->op11->e
    cond1(no)->op3->op4->op5->e
    
    cond2(yes)->op12->e
    cond2(no)->op13->op14->e
    
    ```
    - 情况1（无缓冲channel，发送G阻塞）：取出sendG，获取sendG的数据并唤醒。
    - 情况2（有缓冲channel，发送G阻塞）：从buf队首取数据，从sendq取出sendG，将sendG的数据写入channel，并唤醒。
    - 情况3（有缓冲channel，缓冲区有数据，无发送G阻塞）:从缓冲区中取出数据。
    - 情况4（无阻塞sendG，缓冲区无数据）：将当前的recveG加入recveq，阻塞当前recveG。

- 思考

    - 这里思考一个问题，那 goroutine1 和 goroutine2 又怎么互相知道自己的数据 ”到“ 了呢？
      - channel结构中的recvq、sendq保存着阻塞等待的goroutine，但goroutine1向环型队列中发送数据时，就会从recvp取出goroutine并唤醒。

- 参考

    - https://mp.weixin.qq.com/s/ZXYpfLNGyej0df2zXqfnHQ
    - https://www.cnblogs.com/-wenli/p/12710361.html
    - https://segmentfault.com/a/1190000019172554

### CSP是一种新的并发编程模型


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


```

```