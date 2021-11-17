# Kafka

### 总体架构
- Producer：生成消息，push到Topic
- Broker：每个节点就是一个Broker，负责创建Topic，并将Topic中消息持久化到磁盘
- Topic：同一个Topic可以分布在一个或多个Broker，一个Topic包含一个或多个Partition
- Partition：存储消息的单元，由Topic创建
- Consumer：从订阅的Topic主动拉取消息并消费
- ZooKeeper：维护集群节点状态信息

### Topic 与 Partition
- 分区策略：顺序分发、Hash分区
- 每个Partition即是一个文件夹，包含.index、.log文件，读取消息时：
    - 1、从.index文件获取消息在.log文件中的offset值
    - 2、从.log文件的offset位置开始读取消息
    - 3、消息定长，即到offset+len(消息定长)处结束读取
    
### Consumer 与 Consumer Group

### Consumer Group 与 topic 订阅

### 故障转移

### 数据一致性

### Q&A
- 为什么Producer不在Zookeeper中注册？
- 如何保障Kafka吞吐率？
    - 顺序写磁盘
    
- 消费者获取消息是pull，而不使用push？
    - 消费者根据自身的处理能力去拉取消息并处理，若采用push方式，可能会push消息速率过高而压垮消费者
    

### 参考
- Kafka 详解：https://www.modb.pro/db/105106