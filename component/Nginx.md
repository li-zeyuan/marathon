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
##### 启动
- ![v2-916b9832017683d98c248fde1717ac91_r](https://raw.githubusercontent.com/li-zeyuan/access/master/img/v2-916b9832017683d98c248fde1717ac91_r.jpeg)
- 启动进程启动后，fork出master进程后结束
- master进程交给init进程接管
- master进程worker出worker进程
##### worker进程工作流程

- ![worker](https://raw.githubusercontent.com/li-zeyuan/access/master/img/worker.png)


### I/O模型
- 通过epoll实现IO多路复用
- 过程：
  - 1、当一个请求accept时，worker调用epoll_ctl向epoll注册socket和回调事件
  - 2、继续监听\处理其他请求
  - 3、回调事件被触发，worker处理对应的socket

### 信号管理

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/1468231-20190604222852999-553607453.png)

### 惊群问题
- what：当一个请求accept时，多个worker进程被唤醒去争夺处理权，请求被其中一个worker成功处理后，其他的worker又进入休眠，是一种资源浪费的现象
- why: 因为多个worker监听同一个端口
- resolve：nginx设置了一个accept_mutex锁

### 不停机更新配置

- 修改nginx.conf的配置
- master节点会根据新配置fork出新的worker进程
- 新的请求由新的worker进程处理
- 等老的worker进程处理请求完成后会被kill

### 为什么高效

- 多进程模型；采用多个worker 进程实现对 多cpu 的利用
- 异步非阻塞；通过epoll实现IO多路复用机制
- 事件模型；worker进程在处理request时，遇到阻塞，则向epoll注册一个事件，然后继续处理其他request。直到事件被触发，worker继续处理该request

### 如何实现高可用

![v2-ec3208d1ea659d126fe2a008ec5ae927_r](https://raw.githubusercontent.com/li-zeyuan/access/master/img/v2-ec3208d1ea659d126fe2a008ec5ae927_r.jpeg)

- Keepalived + 双机热备
- 请求过来先落到keepalived，keepalived有虚拟ip（vip）
- keepalived通过心跳监控nginx的健康状态，然后做故障转移
- 参考：https://mp.weixin.qq.com/s?__biz=MzIwMzY1OTU1NQ==&mid=2247508995&idx=2&sn=9afa90512c951783982cec79a95ce6b1&chksm=96cee44fa1b96d59cd137d0f8b53cd55ddb814c4d7a45099f4290127d9659100006ba8a38d4c&scene=27#wechat_redirect

### 限流

参考：https://www.cnblogs.com/Nicholas0707/p/12093173.html

##### 漏桶算法

<img src="https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-09-16_17-35-23.png" alt="Snipaste_2021-09-16_17-35-23" style="zoom: 33%;" />

- 请求从上方不定速进入
- 请求从下方匀速流出
- 超出桶容量，请求会被丢弃

##### 令牌桶算法

<img src="https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-09-16_18-06-42.png" alt="Snipaste_2021-09-16_17-35-23" style="zoom: 33%;" />

- 令牌匀速生成，并放入令牌桶中
- 请求到达时，获取令牌成功才能正常被处理
- 获取令牌失败的请求将被缓存

##### 区别

- 令牌桶比漏桶多一个队列，用来缓存请求
- 令牌桶（桶中令牌为最大值）允许突发流处理，令牌桶则限制请求的速度不超过设定的阀值

- Nginx请求限速模块采用的漏桶算法 + 令牌桶算法

##### 限制并发数：ngx_http_limit_conn_module

- 限制客户端的IP并发连接数

##### 限制请求速率：ngx_http_limit_req_module

- 限制请求处理速率

### 参考

- nginx 多进程 + io多路复用 实现高并发：https://zhuanlan.zhihu.com/p/346243441
- nginx快速入门之基本原理篇：https://zhuanlan.zhihu.com/p/31196264
- 模块和工作原理：https://cloud.tencent.com/developer/article/1664470?from=10680
- 进程模型：https://cloud.tencent.com/developer/article/1664471
- 7层网络以及5种Linux IO模型以及相应IO基础：https://www.cnblogs.com/jing99/p/11984966.html)https://www.cnblogs.com/jing99/p/11984966.html