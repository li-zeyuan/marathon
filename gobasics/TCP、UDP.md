# TCP与UDP

- TCP

  - 面向连接
  - 可靠交付、流量控制、拥塞控制、全双工
  - 字节流传输
  - 只能一对一

- UDP

  - 面向报文，不拆包、组包
  - 可一对一、一对多、多对多

- TCP三次握手
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

- TCP粘包、拆包

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

- TCP流量控制

  - 根据接收发滑动窗口大小，调节发送方滑动窗口的大小，达到流量控制的目的

- TCP拥塞控制

  - 发送发维护一个拥塞窗口变量，根据拥塞窗口变量的大小进行不同的操作
    ​	![拥塞控制](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210130134553.png)
- 慢启动
    - 拥塞避免
    - 快重传
  
- 参考

  - https://zhuanlan.zhihu.com/p/108822858