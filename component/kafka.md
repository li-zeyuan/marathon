# Kafka

### 总体架构

![649269d80d7b278699b49d3279511721](https://raw.githubusercontent.com/li-zeyuan/access/master/img/649269d80d7b278699b49d3279511721.png)

- Producer：生成消息，push到Topic
- Broker：每个节点就是一个Broker，负责创建Topic，并将Topic中消息持久化到磁盘
- Topic：同一个Topic可以分布在一个或多个Broker，一个Topic包含一个或多个Partition
- Partition：存储消息的单元，由Topic创建，分leader partition和follower partition
- Consumer：从订阅的Topic主动拉取消息并消费
- ZooKeeper：维护集群节点状态信息

### Topic 与 Partition
- 分区策略：顺序分发、Hash分区
- 每个Partition即是一个文件夹，包含.index、.log文件，读取消息时：
    - 1、从.index文件获取消息在.log文件中的offset值
    - 2、从.log文件的offset位置开始读取消息
    - 3、消息定长，即到offset+len(消息定长)处结束读取
  
### 数据一致性

- ISR（in-sync Replica）：与leader Broker数据保持**一定程度同步**的follower
- OSR：与leader Broker数据**滞后过多** 的follower
- LEO：每个broker消息偏移量
- HW：所有broker的LEO最小值，Consumer只能读取HW之前的消息

producer生产消息至broker后，HW和LEO变化过程：

- 1、Producer向broker发送消息

  ![Snipaste_2021-11-19_11-18-33](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-11-19_11-18-33.png)

- 2、Leader更新LEO，Follower开始同步Leader消息

  ![Snipaste_2021-11-19_11-23-08](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-11-19_11-23-08.png)

- 3、其中一个Follower完全同步了Leader的消息，另一个只同步了部分消息

  ![Snipaste_2021-11-19_11-28-56](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-11-19_11-28-56.png)

- 4、所有的Follower完成消息同步

  ![Snipaste_2021-11-19_11-31-19](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2021-11-19_11-31-19.png)

### 故障转移
- 1、producer提交消息时，同步ISR中的所有follower，才会回复ACK
- 2、ZooKeeper维护节点的alive状态
- 3、leader节点宕机后，从ISR列表中选举一个follower节点成为leader

### 可用性和持久性保证
- 1、禁用unclean leader选举机制：ISR副本全部宕机情况下，不允许非ISR副本选举leader
- 2、指定最小的ISR集合大小，只有当ISR的大小大于最小值，分区才能接受写入操作

### Q&A
- 如何保证消息传输？
    - broker commit成功，有副本机制(replication)的存在，保证消息不丢
    - broker commit不成功，producer会重试，可能导致重复消息

- 如何保证消息顺序？
    - 同一个partition消息是有序的
    - 不同partition消息无序

- 为什么Producer不在Zookeeper中注册？
    - Producer直接由Broker中的Coordinator协调、管理，并进行rebalance
    - 减少Zookeeper的rebalance负担

- 如何保障Kafka吞吐率？
    - partition顺序读写磁盘
    - broker持久化数据采用mmap页缓存
    - customer从broker读取数据采用0拷贝（sendfile）
    - broker数据批处理，压缩，减少io，非强制刷新缓存写操作
    - customer并行读取partition消息
    - https://xie.infoq.cn/article/49bc80d683c373db93d017a99
- 消费者获取消息是pull，而不使用push？
    - 消费者根据自身的处理能力去拉取消息并处理，若采用push方式，可能会push消息速率过高而压垮消费者

- kafka怎么保证数据 一致性？
    - 引入ISR、OSR、LEO、HW
    - 既不是完全的同步复制，也不是单纯的异步复制，平衡吞吐量和确保消息不丢

- 为什么不采用Quorum读写机制？
    - Quorum：如果选择写入时候需要保证一定数量的副本写入成功，读取时需要保证读取一定数量的副本，读取和写入之间有重叠。
    - 优点：延迟取决与最快的节点
    - 优点：保证了读取和写入之间有重叠部分节点包含所有的数据
    - 缺点：多数的节点挂掉不能选择 leader
    - 缺点：单点故障需要3份数据，要冗余2个故障需要5份，降低吞吐量，大数据量下成本高
  
- ZooKeeper高可用？

- v2.8版本后移除ZooKeeper，采用KRaft方式
    - raft协议：https://xie.infoq.cn/article/57da6912139339e5098afb9cb

### 参考
- Kafka 详解：https://www.modb.pro/db/105106
- 官文-设计思路：https://kafka.apachecn.org/documentation.html#design
- kafka为什么这么快：https://xie.infoq.cn/article/49bc80d683c373db93d017a99
- Kafka 的基础架构：https://xie.infoq.cn/article/eabce320fb1d710db0e4fc9f9