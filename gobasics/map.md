# map

- hash冲突：map的底层数据结构是数组，当向map中存储一个kv时，通过hash计算得出这个kv应该存储在底层数组的哪个下标，如果在始之前该数组下标已经存在kv（前后两个kv的hash值一样），这时就产生了冲突。

- hash冲突解决：

  - 开放定址法：当存储kv产生hash冲突时，就从数组冲突下标往后查找，找到一个空值下标就存储该kv。
  - 拉链法：当产生hash冲突时，就在冲突的下标形成一个链表，通过指针相连接。

- go map的实现原理：

  - 底层是一个bucket数组，每个bucket可以存储8个kv，当超过8个后，会产生一个新的bucket，并通过overflow指针指向新bucket。tophash通常包含该bucket中每个键的hash值的高八位。

    ```go
    // bucket的结构
    type bmap{
    	//tophash通常包含该bucket中每个键的hash值的高八位
    	tophash [bucketCnt]uint8
        overflow *[]*bmap
    }
    ```

    

- kv存储的过程：当往map中存储kv时，对k进行hash，定位到底层数组的下标（bucket），k的hash值高8位和bucket的tophash对比，判断k是否已经存在。将kv存储到该bucket中，若bucket满了，新建一个新的bucket，并用overflow指向新的bucket。

- go 中map不安全，并发读写map会：fatal error: concurrent map read and map write

  - 解决：用sync.Map，基于sync.RWMutex
  - 参考：https://www.jianshu.com/p/10a998089486

- 参考

  - https://learnku.com/articles/35019
  - https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/