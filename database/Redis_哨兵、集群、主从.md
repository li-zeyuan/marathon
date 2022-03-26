# 主从
- 一主多从；主负责写，从负责读
### 主从复制原理

![Snipaste_2022-03-25_10-40-40](![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20220325104800.png))

- 1、全量同步
- 2、增量同步

### 优点

- 读写分离，提高并发

### 缺点

- 不具备容灾能力


# 哨兵

### 总体架构

![Snipaste_2022-03-11_18-48-53](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-03-11_18-48-53.png)

- 哨兵模式是Redis的高可用方式
- 哨兵节点是特殊的redis服务，不提供读写功能

### 作用
- 监控：监控redis node是否正常工作
- 告警：redis出现故障，发出告警
- 故障转移：自动选举新的master，并向客户端发布新的master配置

### 优点

- 读写分离
- 自动故障转移

### 缺点

- 每台数据一致，内存使用率低

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
  
# 集群

![Snipaste_2022-03-25_11-17-11](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20220325111742.png)

- 每个master节点存储的数据都不一样
- client请求集群查找目标节点采用**slots 插槽**，不是一致性哈稀

### 工作原理

- 1、每个master节点，都有**slots 插槽**
- 2、存储key时，采用CRC16算法得出结果，对16384取模得到哈稀槽，根据哈稀槽找到master节点
- 3、当masterA宕机，salveA会充当master；若masterA和salveA都宕机，集群不可用

### 优点

- 去中心话
- 可线性扩展到1000多个节点，节点可动态添加或删除

### 缺点

- slave充当“冷备”，不能缓解读压力
- 3.0推出，成功案例不多