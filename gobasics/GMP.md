# GMP调度模型

### CSP

- （communicating sequential processes）go推荐的并发模型
- 其他语言一般是通过共享内存实现线程间通信。go推荐通过channel
- “不要以共享内存的方式来通信，相反，要通过通信的来共内存”
- goroutine -> channel -> goroutine


### GMP

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
- 参考：
  - https://learnku.com/articles/41728