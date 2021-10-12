# atomic包与原子操作

## 操作系统基础

##### COW（写时复制）

##### CAS（对比和交换）

##### LOCK 指令

- 总线锁：该CPU独享共享内存，其他CPU对内存的读写请求会阻塞，开销大
- 缓存锁：若访问的内存区域已经存在CPU缓存行，则会对缓存行进行锁定，其他CPU将不能缓存此数据，开销小

参考：https://albk.tech/%E8%81%8A%E8%81%8ACPU%E7%9A%84LOCK%E6%8C%87%E4%BB%A4.html

##### 缓存一致性协议


## func AddXXX(addr *T, delta T)

- 对addr地址数据原子增加/减少delta
- 实现原理：http://www.cyub.vip/2021/04/05/Golang%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90%E7%B3%BB%E5%88%97%E4%B9%8Batomic%E5%BA%95%E5%B1%82%E5%AE%9E%E7%8E%B0/#Add%E6%93%8D%E4%BD%9C
  - 1、使用`LOCK`指令保证原子性操作

## CompareAndSwapXXX

## LoadXXX

## StoreXXX

## SwapXXX

## Value

## 参考
- atomic.Value 设计与实现：https://juejin.cn/post/6934217861389008904#heading-6

