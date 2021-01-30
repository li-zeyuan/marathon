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

