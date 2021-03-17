# 包舒适化顺序

![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210317175120.png)

- main.go先执行import的包
- import的顺序为深度优先
- 同一个包中先执行const -> var ->init(可以有多个)
- 同一个包只能初始化一次

### 参考

- https://blog.csdn.net/claram/article/details/77745665