# GMP调度模型

### CSP

- （communicating sequential processes）go推荐的并发模型
- 其他语言一般是通过共享内存实现线程间通信。go推荐通过channel
- “不要以共享内存的方式来通信，相反，要通过通信的来共内存”
- goroutine -> channel -> goroutine

### GM模型
- 缺点：限制GO并发
  - 1、存在全局互斥锁，创建、调度G需要加锁
  - 2、M依赖内存缓存，内存占用高
  - 3、M直接传递G开销大

### GMP模型
- M 的数量默认是 10000，P 的默认数量的 CPU 核数
- go调度器。本质上是把go程序中大量的goroutine分配到少量的线程上执行，利用CPU的多核，实现并发。
  - ![GMP调度模型](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130115056.jpeg)
  - G：goroutine。go协程。
  - P：processor。处理器，go程序创建出来的协程，放到P队列中，供M执行
  - M：thread。内核线程， 是G的执行实体
- 调度过程
  - ![调度过程](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130115135.jpeg)
  - 1、go 关键字创建一个G（goroutine）
  - 2、将G放到P本地队列，或放到全局队列
  - 3、M从P本地队列中取出G执行；若P本地队列为空，从全局队列，或其他P'的本地队列中取。
  - 4、当M在执行G时产生了syscall或阻塞时，M会从本地队列中弹出一个可执行的G'出来执行，并把阻塞的G交给空闲的M服务（或创建一个新的M）
  - 5、M没有可执行G时，就会进入休眠状态，放入休眠M队列
- 缺点
  - 内建函数检测记数器，解决部分抢占式调度，没有解决无内建函数调用"饿死"问题 
  
### 一些调度策略
- https://blog.csdn.net/jigetage/article/details/103350180?utm_source=app&app_version=5.2.0
  
### 参考：
  - https://learnku.com/articles/41728
  - GMP模型演进：https://feiybox.com/2020/03/14/Golang-%E5%8D%8F%E7%A8%8B%E8%B0%83%E5%BA%A6%E5%8E%9F%E7%90%86/
  - Go 面试官：GMP 模型，为什么要有 P？：https://juejin.cn/post/6968311281220583454