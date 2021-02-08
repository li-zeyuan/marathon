# PostgreSQL

### 索引

- Btree、Hash、GIN、GiST、SP-GiST、BRIN
- `CREATE INDEX`默认是创建Btree索引
- 主键默认建Btree索引

### Btree

- Btree索引的实现类似MySQL的B+tree实现
- 适用`< <= = >= >`、between、in、is null、is not null

### Hash

- 类似MySQL的Hash索引
- 底层实现是哈希表
- 只适用等值比较

### Gin

- 是一种倒排索引
- 存储的是键值对结构（key，posting list）

### 正排索引、倒排索引

- 正排索引
   - ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210208111826.png)
  - 文档id为关键字，文档内容为记录

- 倒排索引

  - ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210208164256.png)
  - 分词为关键字，文档id为记录

### json和jsonb

- json
  - 是对输入完整拷贝，保留了空格、键的顺序等
  - 使用的时候再去解析
  - 存储快，因为不需要解析，使用慢

- jsonb

  - 对输入进行解析成二进制后保存，不保留空格、键顺序等
  - 使用时不要再次解析
  - 存储慢，因为需要解析，使用快

- ```
  两者的区别：
  1、存储，使用效率
  2、空格、键顺序是否保留
  ```

- 可以在jsonb列上建gin索引

### 引擎

- heap
- 行级锁

### 参考

- gin索引：https://blog.csdn.net/ctypyb2002/article/details/108865908?ops_request_misc=%25257B%252522request%25255Fid%252522%25253A%252522161275168616780271587140%252522%25252C%252522scm%252522%25253A%25252220140713.130102334..%252522%25257D&request_id=161275168616780271587140&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~sobaiduend~default-2-108865908.pc_search_result_before_js&utm_term=postgresql+gin

- 官方文档：http://postgres.cn/docs/10/datatype-json.html#JSON-INDEXING