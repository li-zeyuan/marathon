### 背景

- 当前业务代码中，通过埋点发送消息的方式去处理后续的业务逻辑，埋点的代码分散，容易出现漏埋点的情况
- 用CDC的方式直接监控pg表数据的变化，然后触发指定的逻辑执行

### 实现原理

```flow
st=>start: Start
op=>operation: pg开启发布
op1=>operation: pg将WAL log写入复制槽
op2=>operation: 基于amazonriver库实现订阅
op3=>operation: 解析WAL log，根据update、del、insert生成事件发送到mns
op4=>operation: mns发送事件到消费者执行业务逻辑
e=>end

st->op->op1->op2->op3->op4->e
```

- 基于amazonriver库的解析，解析的结果为
  - `{"type":update", "data":{"$column1":$value1, "column2":$value2}, "old_data":"{$Column1:$value3}}`
- 订阅者接收到WAL log后，回复一个消息类型个pg，使得log确认消费，WAL log就不会停留在复制槽中

### 存在问题

- 停掉订阅，WAL log一直堆积，导致pg内存不足

- 基于amazonriver库解析出来的struct准确性有待认证
  - 现在是使用的是从struct中提取出id，然后回表查询的方式
  - 回表查询的方式也是存在并发问题，如：回表查询是，其他goroutine已经对数据进行修改，导致回表查询的结果不正确

### 参考

- amazonriver：https://github.com/hellobike/amazonriver
- PostgreSQL发布订阅：http://postgres.cn/docs/10/logical-replication.html