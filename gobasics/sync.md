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

## 互斥锁

- 数据结构

  ```go
  type Mutex struct {
      state int32 // 状态
      sema  uint32 // 控制锁状态的信号量
  }
  ```

- 实现锁的公平性

  - 正常模式
    - 队列中goroutine和活跃的goroutine竞争锁，原队列中的goroutine加入队列首不
  - 饥饿模式
    - 队列头的goroutine执行直接小于1ms，可优先获取锁

- 总结
  - 多个goroutine获取锁是通过原子性实现的，对比和交换（CAS）

## 读写锁

- 数据结构

  ```go
  type RWMutex struct {
      w           Mutex  // held if there are pending writers
      writerSem   uint32 // 用于writer等待读完成排队的信号量
      readerSem   uint32 // 用于reader等待写完成排队的信号量
      readerCount int32  // 读锁的计数器
      readerWait  int32  // 等待读锁释放的数量
  }
  ```

  