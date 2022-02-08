# go简介

- 静态类型
- 运行是runtime

### 规范

- 相似的变量放在一起声明
- import包顺序，标准库、第三方库
- 包名全部小写
- map、slice初始化
- 枚举从1开始
- 可以指定slice的容量

### 变量类型

- 值类型：array、int、struct
- 引用类型：map、slice、channel、指针、interface、函数

### 类型比较

- 可比较：int、float、string、bool、**pointe、channel、interface、array**
- 不可比较（编译报错）：slice、map、func
- 复合类型含有不可比较的类型，则该类型也是不可比较；如struct
  - struct含有不可比较类型时，可用reflect.DeepEqual比较
- 浅析go中的类型比较：https://segmentfault.com/a/1190000039005467

### 优点

- 编译快
- 执行效率高
- 内存管理GC
- 海量并发

### 缺点

- 第三方库不够多、不够稳定
- 错误处理代码冗余

### 参考

- https://segmentfault.com/a/1190000022285902