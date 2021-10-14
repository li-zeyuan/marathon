# atomic包与原子操作

## 操作系统基础

##### COW（写时复制）

##### CAS（对比和交换）

##### LOCK 指令

- 总线锁：该CPU独享共享内存，其他CPU对内存的读写请求会阻塞，开销大
- 缓存锁：若访问的内存区域已经存在CPU缓存行，则会对缓存行进行锁定，其他CPU将不能缓存此数据，开销小

参考：https://albk.tech/%E8%81%8A%E8%81%8ACPU%E7%9A%84LOCK%E6%8C%87%E4%BB%A4.html

##### 缓存一致性协议


## func AddT(addr *T, delta T)

- 对addr地址数据原子增加/减少delta
- 实现原理：http://www.cyub.vip/2021/04/05/Golang%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90%E7%B3%BB%E5%88%97%E4%B9%8Batomic%E5%BA%95%E5%B1%82%E5%AE%9E%E7%8E%B0/#Add%E6%93%8D%E4%BD%9C
  - 1、使用`LOCK`指令保证原子性操作
  
    ```
    TEXT runtime∕internal∕atomic·Xadd64(SB), NOSPLIT, $0-24
    	MOVQ	ptr+0(FP), BX
    	MOVQ	delta+8(FP), AX
    	MOVQ	AX, CX
    	LOCK                      // LOCK指令进行锁住操作，实现对共享内存独占访问
    	XADDQ	AX, 0(BX)
    	ADDQ	CX, AX
    	MOVQ	AX, ret+16(FP)
    	RET
    ```

## func CompareAndSwapT(addr *T, old, new T) (swapped bool)

- 对比addr地址的值是否为old，是则交换成new，否则不交换

- 实现原理：http://www.cyub.vip/2021/04/05/Golang%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90%E7%B3%BB%E5%88%97%E4%B9%8Batomic%E5%BA%95%E5%B1%82%E5%AE%9E%E7%8E%B0/#CompareAndSwap%E6%93%8D%E4%BD%9C

  - **CMPXCHGQ**比较并交换指令是原子操作

    ```
    TEXT runtime∕internal∕atomic·Cas64(SB), NOSPLIT, $0-25
    	MOVQ	ptr+0(FP), BX
    	MOVQ	old+8(FP), AX
    	MOVQ	new+16(FP), CX
    	LOCK
    	CMPXCHGQ	CX, 0(BX)            // 比较并交换指令
    	SETEQ	ret+24(FP)
    	RET
    ```

## func LoadT(addr *T) (val T)

- GO语言实现
- 直接取addr地址的值，并返回

## func StoreT(addr *T, val T)

- 向addr地址原子存储值

- 实现原理：http://www.cyub.vip/2021/04/05/Golang%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90%E7%B3%BB%E5%88%97%E4%B9%8Batomic%E5%BA%95%E5%B1%82%E5%AE%9E%E7%8E%B0/#Store%E6%93%8D%E4%BD%9C

  - **XCHGQ**交换指令，用于交换源操作数和目的操作数

    ```
    TEXT runtime∕internal∕atomic·Store64(SB), NOSPLIT, $0-16
    	MOVQ	ptr+0(FP), BX
    	MOVQ	val+8(FP), AX
    	XCHGQ	AX, 0(BX)             // 交换指令
    	RET
    ```

## func SwapT(addr *T, new T) (old T)

- 交换addr地址的值成new，并返回old值

- 实现原理：http://www.cyub.vip/2021/04/05/Golang%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90%E7%B3%BB%E5%88%97%E4%B9%8Batomic%E5%BA%95%E5%B1%82%E5%AE%9E%E7%8E%B0/#Swap%E6%93%8D%E4%BD%9C

  - **XCHGQ**交换指令，用于交换源操作数和目的操作数

    ```
    TEXT runtime∕internal∕atomic·Xchg64(SB), NOSPLIT, $0-24
    	MOVQ	ptr+0(FP), BX
    	MOVQ	new+8(FP), AX
    	XCHGQ	AX, 0(BX)           // 交换指令
    	MOVQ	AX, ret+16(FP)
    	RET
    ```

## Value

- 数据结构

  ```go
  type Value struct {
  	v interface{}
  }
  ```

  - 空接口结构
  - 由type、data两部分组成

- func (v *Value) Store(x interface{})

  ```go
  func (v *Value) Store(x interface{}) {
  	if x == nil { // atomic.Value类型变量不能是nil
  		panic("sync/atomic: store of nil value into Value")
  	}
  	vp := (*ifaceWords)(unsafe.Pointer(v)) // 将指向atomic.Value类型指针转换成*ifaceWords类型
  	xp := (*ifaceWords)(unsafe.Pointer(&x)) // xp是*faceWords类型指针，指向传入参数x
  	for { // for循环自旋
  		typ := LoadPointer(&vp.typ) // 原子性返回vp.typ
  		if typ == nil { // 第一次调用Store时候，atomic.Value底层结构体_type部分是nil
  			runtime_procPin() // pin process处理，防止M被抢占
  			if !CompareAndSwapPointer(&vp.typ, nil, unsafe.Pointer(^uintptr(0))) { // 通过cas操作，将atomic.Value的_type部分存储为unsafe.Pointer(^uintptr(0))，若没操作成功，继续操作
  				runtime_procUnpin() // unpin process处理，释放对当前M的锁定
  				continue
  			}
  
  			// vp.data == xp.data
  			// vp.typ == xp.typ
  			StorePointer(&vp.data, xp.data)
  			StorePointer(&vp.typ, xp.typ)
  			runtime_procUnpin()
  			return
  		}
  		if uintptr(typ) == ^uintptr(0) { // 此时说明第一次的Store操作未完成，正在处理中，此时其他的Store等待第一次操作完成
  			continue
  		}
  
  		if typ != xp.typ { // 再次Store操作时进行typ类型校验，确保每次Store数据对象都必须是同一类型
  			panic("sync/atomic: store of inconsistently typed value into Value")
  		}
  		StorePointer(&vp.data, xp.data) // vp.data == xp.data
  		return
  	}
  }
  ```

  - ^uintptr(0)：64位系统下，为max uint64，用来标记Value正在进行Store操作
  - for循环实现自旋，对比与交换&vp.typ不成功 或 第一次的Store操作未完成 的情况下自旋
  - runtime_procUnpin( )：禁止 M被强占；runtime_procUnpin( )：解禁止；保证了Value的type部分和data部分的原子存储
  - 多次调Store，v的类型必须一样，否则panic

- func (v *Value) Load() (x interface{})

  ```go
  func (v *Value) Load() (x interface{}) {
  	vp := (*ifaceWords)(unsafe.Pointer(v)) // 将指向v指针转换成*ifaceWords类型
  	typ := LoadPointer(&vp.typ)
  	if typ == nil || uintptr(typ) == ^uintptr(0) { // typ == nil 说明Store方法未调用过
  	// uintptr(typ) == ^uintptr(0) 说明第一Store方法调用正在进行中
  		return nil
  	}
  	data := LoadPointer(&vp.data)
  	xp := (*ifaceWords)(unsafe.Pointer(&x))
  	xp.typ = typ
  	xp.data = data
  	return
  }
  ```

  - 若uintptr(typ) == ^uintptr(0)；说明第一次Store方法正在调用

## 总结

- Store与Swap都可以实现原子性存储值，区别在与两个函数的返回值不同
- Store/Load是否必要？和普通赋值/取值的区别？
  - 解决不用系统架构问题：https://stackoverflow.com/questions/46556857/is-atomic-loaduint32-necessary

## 参考
- atomic.Value 设计与实现：https://juejin.cn/post/6934217861389008904#heading-6

- Golang 并发赋值的安全性探讨：https://cloud.tencent.com/developer/article/1810536
