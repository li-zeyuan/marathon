# 哨兵

### 总体架构

![Snipaste_2022-03-11_18-48-53](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-03-11_18-48-53.png)

- 哨兵模式是Redis的高可用方式
- 哨兵节点是特殊的redis服务，不提供读写功能

### 作用
- 监控：监控redis node是否正常工作
- 告警：redis出现故障，发出告警
- 故障转移：自动选举新的master，并向客户端发布新的master配置

### 工作原理

- 1、心跳机制
  - Sentinel 向 Redis Node发送心跳包
  - Sentinel与Sentinel：基于发布订阅
- 2、判断master节点是否下线 
  - 主观下线：master回包异常
  - 客观下线：多数Sentinel判定成主观下线
- 3、基于Raft算法选举领头sentinel
- 4、选举一个slave成为master
  - 网络质量最好
  - 与master数据相似度最高
- 5、修改配置
  - 领头sentinel向redis node广播新的master配置
  - 领头sentinel向客户端广播新的master配置