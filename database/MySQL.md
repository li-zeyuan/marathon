# MySQL

### 范式

- 1NF：原子性，列不可以再拆分。

- 2NF：1、表必须有主键。2、非主键列必须完全依赖主键，而不能只依赖主键的一部分。

  - ```
    例：订单明细表：【OrderDetail】（OrderID，ProductID，UnitPrice，Discount，Quantity，ProductName）。 
    因为我们知道在一个订单中可以订购多种产品，所以单单一个 OrderID 是不足以成为主键的，主键应该是（OrderID，ProductID）。显而易见 Discount（折扣），Quantity（数量）完全依赖（取决）于主键（OderID，ProductID），而 UnitPrice，ProductName 只依赖于 ProductID。所以 OrderDetail 表不符合 2NF。不符合 2NF 的设计容易产生冗余数据。 
    可以把【OrderDetail】表拆分为【OrderDetail】（OrderID，ProductID，Discount，Quantity）和【Product】（ProductID，UnitPrice，ProductName）来消除原订单表中UnitPrice，ProductName多次重复的情况。
    ```

- 3NF：在满足2NF的情况下，非主键列必须直接依赖主键，而不能依赖非主键类。也就是不能传递依赖。

- 反三范式：为了优化查询效率，可以增加冗余字段。

- 参考
  - https://thinkwon.blog.csdn.net/article/details/104778621
  - 三范式：https://blog.csdn.net/Dream_angel_Z/article/details/45175621

### 引擎

|        | MyISAM                                                       | Innodb                                                       |
| :----: | :----------------------------------------------------------- | :----------------------------------------------------------- |
|  外键  | 不支持                                                       | 支持                                                         |
|  事务  | 不支持                                                       | 支持                                                         |
| 锁粒度 | 表级锁                                                       | 行级锁、表级锁                                               |
|  CURD  | select更优                                                   | insert、update、delete更优                                   |
|  索引  | 1、不支持哈希索引<br />2、支持全文索引<br />3、非聚簇索引<br />4、索引叶子节点存储的是行数据地址，需要再次寻址才能得到数据 | 1、支持哈希索引<br />2、不支持全文索引<br />3、聚簇索引<br />4、主键索引的叶子节点存储行数据，不需要再次寻址 |
|  场景  | 以读、插入为主的应用程序，如博客、新闻                       | 更新、删除频繁的应用，如op系统                               |

### 索引类型

- 主键索引
- 普通索引
- 唯一索引
- 组合索引
- 全文索引

### 聚簇索引、非聚簇索引

- 聚簇索引：数据和索引放在一起，找到了索引也就找到数据，不需要回表查询。如Innodb的B+tree索引
- 非聚簇索引，数据和索引是分开的，索引结构只是保存了指向数据对应的行，需要回表查询。如MyISAM的B+tree索引结构

### 索引数据结构

- b+tree
  - ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210204144458.png)
  - 大多数情况下走的是b+tree索引
  - 非叶子节点不保存数据
  - 所有数据都保存在叶子节点
  - 叶子节点形成有序的列表结构，方便范围查询
- hash
  - 单条数据查询，如where id，使用的是哈希索引，其他走b+tree索引
  - 底层数据结构是哈希表，通过哈希算法实现
- b-tree
  - ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210204145443.png)
  - 所有叶子节点都有保存key和数据
- 为什么使用b+tree，而不是b-tree
  - b+tree非叶子节点存储更多的元素，IO查询次数更少
  - 所有的查询都到叶子节点，性能更稳定
  - 叶子节点形成有序链表，便于范围查询

### 索引使用场景

- where
- order by
- join...on...

### 索引使用方式

- 全值匹配最佳
- 复合索引要遵从 最佳左前缀法则
- 在索引列上操作（计算、函数、类型转换），引起索引失效
- 范围查询（bettween/</>/in）右边的列索引失效
- select尽量覆盖索引列
- is null、is not null、!=、<>、or 导致索引失效
- like 以通配符开头（'%字符串'）导致索引失效

### 事务



### 参考

- MySQL索引背后的数据结构及算法原理：https://blog.codinglabs.org/articles/theory-of-mysql-index.html
- MySQL面试汇总：https://thinkwon.blog.csdn.net/article/details/104778621
- 索引：https://blog.csdn.net/wuseyukui/article/details/72312574

