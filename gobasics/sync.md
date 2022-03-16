# sync

## errgroup

- 使用

  - ```go
    package main
    
    import (
        "context"
        "fmt"
        "time"
    
        "golang.org/x/sync/errgroup"
    )
    
    func main() {
        group, _ := errgroup.WithContext(context.Background())
        for i := 0; i < 5; i++ {
            index := i
            group.Go(func() error {
                if index%2 == 0 {
                    return fmt.Errorf("something has failed on grouting:%d", index)
                }
                return nil
            })
        }
        if err := group.Wait(); err != nil {
            fmt.Println(err)
        }
    }
    ```

  

### 结构

- ```go
  type Group struct {
    cancel  func()             //context cancel()
      wg      sync.WaitGroup     // 线程同步    
      errOnce sync.Once          //只会传递第一个出现错的协程的 error
      err     error              //传递子协程错误
  }
  ```

### 小结

- 首先传递context初始化一个errgroup
- 子协程监听ctx.Done( )获取获取撤销信号
- 只能记录最先出错的协程
- 守护线程，主协程等待子协程退出后才结束

## Mutex
### 数据结构
- ```go
  type Mutex struct {
	    state int32
	    sema  uint32
    }
  ```
- Mutex.state
  - ![Snipaste_2022-02-17_14-59-04](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-02-17_14-59-04.png)
  - waiter_num：等待goroutine
  - starving：是否处于饥饿状态
  - woken：是否有goroutine正在加锁
  - locked：是否有goroutine持有该锁
- Mutex.sema
  - 信号量，用于唤醒阻塞的goroutine
### 锁的两种模式

```
引入饥饿模式，保证锁公平性
公平性：goroutine获取锁的顺序与请求锁的顺序一致
```

- 正常模式
  - 唤醒阻塞队列队头的goroutine，该goroutine与新请求锁的goroutine共同竞争锁，新请求的goroutine大概率会获得锁，因为占有cpu时间片
  
- 饥饿模式
  - 唤醒阻塞队列队头的goroutine直接获得锁
  - 新请求锁goroutine不参与锁竞争
  
  ```
  饥饿模式的触发条件：
  	1、有一个goroutine获取锁的时间超过1ms
  饥饿模式的取消条件：
  	1、获取到锁的goroutine为阻塞队列的最后一个时，恢复正常模式
  	2、获取到锁的goroutine等待时间小于1ms，恢复正常模式
  ```

### 注意点

- 不可重入，重复Mutex.Lock会panic
- 先调用Mutex.Unlock会panic
- 同一把锁，一个goroutine Mutex.Lock，另一个goroutine可以Mutex.Unlock

## Map

- 数据结构

  - ```go
    type Map struct {
    	mu Mutex // 加锁
    	read atomic.Value // readOnly，只读
    	dirty map[interface{}]*entry // 包含新写入的数据，misses计数到伐值则拷贝到read
    	misses int // 读失败计数
    }
    ```

- 大体流程

  - ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210319141225.png)

- 总结
  - 线程安全
  - 读写分离，降低锁的时间来提高效率
  - misses增加到dirty的长度时，将dict拷贝给read

- 参考
  - https://juejin.cn/post/6844903895227957262
  - https://www.cnblogs.com/qcrao-2018/p/12833787.html

## 读写锁

##### 数据结构

```go
type RWMutex struct {
    w           Mutex  // 保证只会有一个写锁加锁成功
    writerSem   uint32 // 用于writer等待读完成排队的信号量
    readerSem   uint32 // 用于reader等待写完成排队的信号量
    readerCount int32  // 读操作goroutine数量
    readerWait  int32  // 阻塞写操作goroutine的读操作goroutine数量
}
```

##### 加读锁

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210323101222.png)
- readerCount > 0，说明存在读锁，则加锁成功
- readerCount < 0，说明存在写锁，则阻塞，等待readerSem信号量唤醒

##### 释放读锁

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210323102653.png)
- readerCount - 1，若readerCount<0，说明有写锁等待
- readerWait - 1，若readerWait == 0，说明最后一个解读锁了，则唤起写锁信号量（释放全部读锁后，唤醒写锁）

##### 加写锁

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210323103907.png)
- m.lock保证读锁间互斥
- readerCount - rwmutexMaxReaders ，阻塞后面的读锁再加锁
- readerWait>0，说明存在读锁，则阻塞，等待写锁信号量唤醒

##### 释放写锁

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210323105015.png)
- readerCount + rwmutexMaxReaders，readerCount复位
- 读锁信号量唤醒所有读锁

##### 总结

- 写锁通过递减rwmutexMaxReaders常量，使readerCount < 0，实现对读锁的抢占
- atomic.AddInt32操作是通过LOCK来进行CPU总线加锁的
- m.lock保证写锁之间的公平
- 先入先出（FIFO）的原则进行加锁，实现公平读写锁，解决线程饥饿问题

##### 参考

- https://cloud.tencent.com/developer/article/1557629
- https://www.techclone.cn/post/tech/go/go-rwlock/#%E8%AF%BB%E5%86%99%E9%94%81%E5%BC%95%E5%85%A5
- Golang 读写锁设计：https://segmentfault.com/a/1190000040406605

