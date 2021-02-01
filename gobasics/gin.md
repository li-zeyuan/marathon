# gin

- 是一个轻量级web框架
- 基于http包

### 中间件

- 需要返回gin.HandlerFunc函数，通next一次来执行
- 项目中用来：token认证，接口权限认证、限流、日志记录
  - 限流：redis记录用户单位时间内的访问次数，超频则拦截

### 路由

- RouterGroup 路由组

### content

- 保存请求上下文信息
- 实现链路追踪

### 单元测试

- 提供接口单元测试

### 参考

- go语音中国网：https://studygolang.com/articles/32543?fr=sidebar