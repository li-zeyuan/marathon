# Nginx

### 进程模型

- ![Snipaste_2021-08-03_11-26-43](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-08-03_11-26-43.png)

- master进程：

  - 接收来自外界信号
  - 向worker进程分发信号
  - 监控worker进程运行状态

- worker进程：

  - 连接accept后，读取请求、解析请求、处理请求

  - 独立进程 ，一个请求只能被一个worker处理

- proxy cache

  - 缓存静态资源

### 工作流程

- 


### I/O模型

### 信号管理

### 惊群问题

### 如何不停机更新配置？

### 为什么高效

- Nginx master-worker进程机制。
- IO多路复用机制。
- Accept锁及REUSEPORT机制。
- sendfile零拷贝机制

### 参考
- nginx快速入门之基本原理篇：https://zhuanlan.zhihu.com/p/31196264
- 模块和工作原理：https://cloud.tencent.com/developer/article/1664470?from=10680
- 进程模型：https://cloud.tencent.com/developer/article/1664471
- 7层网络以及5种Linux IO模型以及相应IO基础：https://www.cnblogs.com/jing99/p/11984966.html)https://www.cnblogs.com/jing99/p/11984966.html