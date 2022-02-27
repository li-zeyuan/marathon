### defer触发时机
- 包裹着defer语句的函数返回时(先入栈的defer后执行)

  - ```go
     	// 输出结果：
    	// return前执行defer2
    	// return前执行defer1
       func f1() {
           defer fmt.Println("return前执行defer1")
           defer fmt.Println("return前执行defer2")
           return 
       }
    ```

- 当前goroutine发生Panic时

  - ```go
    //输出结果：panic前  第一个defer在Panic发生时执行，第二个defer在Panic之后声明，不能执行到
       func f3() {
           defer fmt.Println("panic前")
           panic("panic中")
           defer fmt.Println("panic后")
       }
    ```
### defer，return，返回值的执行顺序

- **1. 先给返回值赋值**

- **2. 执行defer语句**

- **3. 包裹函数return返回**

  ```go
  // 例1：匿名返回值
  // 可以看作：1、ret = r; 2、r = r *7; 3、return ret
  // 结果：6
  func f1() int { 
          var r int = 6
          defer func() {
                  r *= 7
          }()
          return r
  }
  
  // 例2：有名返回值
  // 可以看作：1、ret = 6; 2、ret = ret *7; 3、return ret
  // 结果：42
  func f2() (ret int) { //有名返回值
          defer func() {
                  ret *= 7
          }()
          return 6
  }
  
  // 例3：有名返回值
  // 可以看作：1、ret = 6; 2、无（因为defer定义时值传递）; 3、return ret
  // 结果：6
  func f3() (ret int) { //有名返回值
      defer func(r int) {
          ret *= 7
      }(ret)
      return 6
  }
  ```

  

### defer源码解析

- ```go
      type _defer struct {
              siz     int32 
              started bool
              sp      uintptr // 函数栈指针
              pc      uintptr // 程序计数器
              fn      *funcval // 函数地址
              _panic  *_panic // Panic是导致运行defer的Panic
          	link    *_defer // (链表)指向自身结构的指针，用于链接多个 defer
      }
  ```

- 新建的延迟函数挂在当前goroutine的_defer的链表上