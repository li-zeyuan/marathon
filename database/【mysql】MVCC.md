# MVCC
### 是什么？
- Multiversion concurrency control (MCC or MVCC)
- 多版本并发控制（MCC 或 MVCC）是一种并发控制方法，通常被数据库管理系统用来提供对数据库的并发访问，并以编程语言来实现事务存储。
### 解决了什么？
- 不加锁的情况下解决了脏读、不可重复读和快照读下的幻读问题（幻读问题最终就是使用间隙锁解决）
### 如何实现？
- 隐式字段
    - DB_TRX_ID：记录改条数据修改它的事务 ID
    - DB_ROLL_PTR：回滚指针，指向这条记录的上一个版本
- undo_log：回滚日志
- read-view：执行select语句时生成的视图
### 参考 
- *MVCC：听说有人好奇我的底层实现*：https://xie.infoq.cn/article/eff93ec47b54a5069e0bd1726
- MySQL · 引擎特性 · InnoDB undo log 漫游 http://mysql.taobao.org/monthly/2015/04/01/
- redo log和bin log https://blog.csdn.net/qq_40194399/article/details/120862971
- MVCC多版本并发控制机制—包你学会 https://www.bianchengquan.com/article/231348.html 