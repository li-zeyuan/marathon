# TCP与UDP

- TCP

  - 面向连接
  - 可靠交付、流量控制、拥塞控制、全双工
  - 字节流传输
  - 只能一对一
- UDP

  - 面向报文，不拆包、组包
  - 可一对一、一对多、多对多

### TCP三次握手四次挥手
​	![三次握手](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130134349.png)

- 1、client向server发送一个syn
  - 2、server收到syn包，响应一个ack；同时发送一个请求报文syn
  - 3、client收到syn后，响应确认报文ack
- TCP四次挥手
  ​	![四次挥手](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130134505.png)

  - 1、client向server发送FIN报文，请求关闭连接
  - 2、server回复确认报文ACK
  - 3、server向client发送FIN报文
  - 4、client回复确认报文ACK
- 为什么建立连接是三次握手，而关闭连接却是四次挥手呢？

  - 三次握手：在第二次握手时，server回复ACK的同时，也发送了SYN
  - 四次挥手：TCP是全双工，需要双方都关闭才断开连接。client发送FIN报文，代表client没有发送数据，可以关闭；server发送FIN报文，代表server没有发送数据，可以关闭。需要对client、server的FIN报文分别确认。
- TCP 短连接和长连接

  - 短连接：client发送消息，server响应。完成一次读写操作后，client就发起关闭请求
  - 长连接：在连接建立完成后，client、server会一直复用这个连接，直到该连接关闭

### TCP粘包、拆包

- 拆包
  - 当应用层发送过来的数据大于缓冲区剩余空间时，将发生拆包
  - 当缓冲区中要发送的数据包大于报文长度是，将发生拆包
- 粘包
  - 当应用层发送过来的数据小于缓冲区剩余空间时，TCP将缓冲区的数据一次发出，将发生粘包
  - 接收端的应用层没有及时读取缓冲区的数据，将发生粘包
- 解决办法
  - 消息定长：消息长度不够补0
  - 设置消息边界：如在尾部加回车换行
  - 消息头消息体：消息头存放一下表示消息长度的字段

### 重传机制

- 超时重传
  - 发送数据时，设定定时器，超时则重发数据

- 快速重传

  - ![Snipaste_2022-02-07_18-32-17](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-02-07_18-32-17.png)
    - 1、seq2因某种原因发送失败
    - 2、发送方连续3次收到同一个ack2
    - 3、发送方快速重新发送seq2
  - SACK
    - SACK在tcp的头部增加了接受方已经缓冲的数据段
    - ![Snipaste_2022-02-07_18-41-59](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-02-07_18-41-59.png)
  - D-SACK
    - 使用了 SACK 来告诉「发送方」有哪些数据被重复接收

### 滑动窗口

- 为了提高通讯效率
- 窗口的大小由接收方的窗口大小来决定

- 发送方滑动窗口
  - ![Snipaste_2022-02-07_18-55-35](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-02-07_18-55-35.png)
- 接受方滑动窗口
  - ![Snipaste_2022-02-07_18-56-02](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-02-07_18-56-02.png)

### 流量控制

- 根据接收发滑动窗口大小，调节发送方滑动窗口的大小，达到流量控制的目的

### 拥塞控制

- 发送发维护一个拥塞窗口变量，根据拥塞窗口变量的大小进行不同的操作
  ​	![拥塞控制](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130134553.png)
  - 慢启动
  - 拥塞避免
  - 超时快重传

### CLOSE_WAIT太多怎么办？

- 被关闭一方收到FIN包，回复ACK后进入close_wait
  - https://blog.huoding.com/2016/01/19/488

- 四次挥手状态转移图
  - ![Snipaste_2022-03-14_10-54-34](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-03-14_10-54-34.png)
  - 原因：
    - 1、服务端代码没有主动调用close
    - 2、服务处理请求时间慢，导致多余的请求在队列中就背客户端关闭

### 为什么需要TIME_WAIT？

- 关闭一方收到FIN包后，回复ACK进入time_wait
  - https://blog.huoding.com/2013/12/31/316
- 原因：
  - 1、若没有time_wait，被关闭一方早些发送的包到达后会发现旧链接已经关闭，只能回复RST包
  - 2、若没有time_wait，被关闭一方早些发送的包到达后，新的连接接已经被重用，干扰新连接
  
### TCP如何保证可靠传输
- 连接管理
- 校验和
- 序列号
- ACK应答
- 超时重传
- 流量控制
- 拥塞控制

### 五层网络模型
- 应用层：fpt、http
- 传输层：tcp、udp
- 网络层：ip
- 数据链路层：以太网协议
- 物理成：

### 为什么说TCP是面向流的协议？而UDP是面向数据报的协议
- TCP
  - 应用层向TCP发送大小不一的数据块
  - TCP把这些数据块看成字节流，TCP协议头无length字段
  - TCP进行粘包、拆包
- UDP
  - 应用层向UDP发送几个数据包，UDP则会发送几个数据包
  - UDP不会发生粘包、拆包
  - UDP协议头有length字段

### 参考

- https://zhuanlan.zhihu.com/p/108822858
- 30张图解： TCP 重传、滑动窗口、流量控制、拥塞控制：https://www.cnblogs.com/xiaolincoding/p/12732052.html