# context

### context包

​	![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130151045.png)

- 在goroutine中传递上下文信息、信号控制、公共参数等

### context数据结构

```go

type Context interface {
    Deadline() (deadline time.Time, ok bool)		    // 获取当前context的截止时间
    Done() <-chan struct{}							  // 识别channel是否被关闭
    Err() error										 // 获取context被关闭的原因
    Value(key interface{}) interface{}				   // 获取当前context中所存储的value
}
```

### 使用场景

- 每个请求，都会有一个goroutine去处理，这个goroutine又会创建额外的goroutine去访问数据库、RPC服务等。这些goroutine同属于一个请求，可以同context包去传递request_id，请求goroutine超时取消后，通知所有的goroutine也马上退出并释放。

### context的继承

- WithCancel：创建一个可以取消的Context
- WithDeadline：创建一个到截止日期就取消的Context
- WithTimeout：创建一个超时自动取消的Context
- WithValue：在Context中设置键值对

### cancelCtx 结构

```go
type cancelCtx struct {
 Context

 mu       sync.Mutex            // protects following fields
 done     chan struct{}         // created lazily, closed by first cancel call
 children map[canceler]struct{} // set to nil by the first cancel call
 err      error                 // set to non-nil by the first cancel call
}
```

- mu：并发安全，加互斥锁进行操作
- done：context取消会关闭
- children：包含context对应的子集，关闭通知所有的子集context
- err：报错信息

### 在项目中使用

- 自定义mapCtx，实现了context所定义的方法

- ```go
type mapCtx struct {
  	Keys map[string]interface{}
  }
  
  func NewContext() *mapCtx {
	return &mapCtx{Keys: make(map[string]interface{})}
	}
	
	func (*mapCtx) Deadline() (deadline time.Time, ok bool) {
		return
	}
	
	func (*mapCtx) Done() <-chan struct{} {
		return nil
	}
	
	func (*mapCtx) Err() error {
		return nil
	}
	```
	
- 使用

  - ```go
    func NewInfra(requestID string) *Infra {
    	infra := new(Infra)
        ...
    	infra.Context = middlecontext.NewContext()
    	return infra
    }
    ```

  - new出的mapCtx赋值给infra

  - 业务中就是传递infra

  - 用于request_id请求链路追踪、传递权限值等

### 参考

- 煎鱼一文吃透 Go 语言解密之上下文 context：https://mp.weixin.qq.com/s/A03G3_kCvVFN3TxB-92GVw

