# var、new、make异同

### var 

- 用于变量声明
- 声明变量，变量默认是该类型的0值
- var a *int，给a赋值会panic，因为a的默认值是null，没有内存空间

### new

- 为变量开辟内存空间
- 返回指针类型
- 可用于所有类型

### make

- 为变量开辟内存空间，也能同时初始化
- 返回的是引用类型本身
- 只能用于slice，map，channel