# map

### hash冲突

- hash冲突：map的底层数据结构是数组，当向map中存储一个kv时，通过hash计算得出这个kv应该存储在底层数组的哪个下标，如果在始之前该数组下标已经存在kv（前后两个kv的hash值一样），这时就产生了冲突。

- hash冲突解决：

  - 开放定址法：当存储kv产生hash冲突时，就从数组冲突下标往后查找，找到一个空值下标就存储该kv。
  - 拉链法：当产生hash冲突时，就在冲突的下标形成一个链表，通过指针相连接。

### go map的实现原理：

- 底层是一个bucket数组，每个bucket可以存储8个kv。
- 拉链法 解决hash冲突

bucket：

  ![Snipaste_2022-01-23_15-56-02](https://raw.githubusercontent.com/li-zeyuan/access/master/img/Snipaste_2022-01-23_15-56-02.png)![]()

  ```go
   // bucket的结构
   type bmap{
   	//tophash通常包含该bucket中每个键的hash值的高八位
   	tophash [bucketCnt]uint8
       overflow *[]*bmap
   }
  ```
  - tophash：用于快速查找key是否在该bucket中
  - overflow：指向扩容后的bucket
  - 按k1k2k3、v1v2v3存储，内存对齐

### make map

### set key

- 1、对k进行hash， 
- 2、hash值的低8位和bucket数组长度取余，定位到bucket数组的下标
- 3、k的hash值高8位和bucket的tophash对比，判断k是否已经存在
- 4、将kv存储到该bucket中，若bucket满了，新建一个新的bucket，并用overflow指向新的bucket。

### get key

- 类似set key

### del key

- 类似set key

### 扩容

- 条件1、装载因子超过阈值，源码里定义的阈值是 6.5。
- 条件2、overflow 的 bucket 数量过多：
- 扩容会导致新旧bucket key搬迁

### 为什么range不是有序的
- 1、map 在扩容后，会发生 key 的搬迁
- 2、每次从一个随机值序号的 bucket 开始遍历，并且是从这个 bucket 的一个随机序号的 cell 开始遍历

### 注意点

- go 中map不安全，并发读写map会：fatal error: concurrent map read and map write

  - 解决：用sync.Map，基于sync.RWMutex
  - 参考：https://www.jianshu.com/p/10a998089486
  
- go中map的value若是值类型，不能修改，如：

  ```go
  type Student struct {
      Age int
  }
  func main() {
      kv := map[string]Student{"menglu": {Age: 21}}
      kv["menglu"].Age = 22 // 报错，需要换成map[string]&Student{"menglu": {Age: 21}}
      s := []Student{{Age: 21}}
      s[0].Age = 22
      fmt.Println(kv, s)
  }
  ```

  

- 参考

  - https://learnku.com/articles/35019
  - https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/
  - go-map源码简单分析（map遍历为什么时随机的）：https://www.helloworld.net/p/3714029944