# sync

### errgroup

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

