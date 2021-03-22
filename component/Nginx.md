# Nginx

### 模块

- 内核模块：加载配置文件
- Handlers模块：处理器模块，用来处理请求
- Filters模块：过滤模块，对处理器模块输出内容进行过滤

### 工作流程

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210321154506.png)

- 请求进入内核，内核加载配置文件，找到对应的location
- 根据location的配置指令，找到对应的Handler去执行请求
- handler的输出结果，经过Filters模块过滤，然后返回给客户端

### 进程模型

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210321155254.png)

- master进程
  - 接收外界信号
  - 管理worker进程
- worker
  - 抢**accept_mutex**，处理请求
- cache 相关的进程
  - 做缓存管理

### 参考

- 模块和工作原理：https://cloud.tencent.com/developer/article/1664470?from=10680
- 进程模型：https://cloud.tencent.com/developer/article/1664471