# channel

- CSP的设计理念：channel

- 在go语音中实现goroutine间通信，分有缓存区和无缓存区

- 分单向和双向模式

- 结构

  ```go
      type hchan struct {
       qcount   uint      	// 队列中剩余元素数量
       dataqsiz uint     		// 循环队列的长度（channel的大小）
       buf      unsafe.Pointer // 长度为dataqsiz的底层数组指针，缓存型channel特有
       elemsize uint16		// 元素大小	
       closed   uint32		// 是否关闭
       elemtype *_type 		// 接收、发送的的元素类型
       sendx    uint  		// 已发送元素在循环队列中的索引位置
       recvx    uint  		// 已接收元素在循环队列中的索引位置
       recvq    waitq  		// 接收者sudog等待队列（阻塞等待接收的goroutine）
       sendq    waitq  		// 发送者sudog等待对列（阻塞等待接收的goroutine）
      
       lock mutex				// 互斥锁
      }
  ```

- 发送数据

  ```flow
    st=>start: Start
    op=>operation: make(chan T, size)
    op1=>operation: 发送
    op2=>operation: 数据写入buf
    op3=>operation: 取出一个G
    op4=>operation: 数据写入G
    op5=>operation: 唤醒G
    op6=>operation: 将G加入sendq队列
    op7=>operation: 待换醒
    
    cond=>condition: recvq非空？
    cond1=>condition: buf有空位？
    e=>end
    
    st->op->op1->cond
    cond(yes)->op3->op4->op5->e
    cond(no)->cond1
    cond1(yes)->op2->e
    cond1(no)->op6->op7->e
  ```

  - 情况1：recvq有G；直接取出G，写入数据并唤醒。
  - 情况2：recvq没有G，buf有空位；将数据写入buf。
  - 情况3：recvq没有G，buf没有空位；当前goroutine阻塞，并加入sendq队列

- 接收数据

  ```flow
  st=>start: Start
  op1=>operation: 接收
  op3=>operation: 取出一个G
  op4=>operation: 从G中读取数据
  op5=>operation: 唤醒G
  op6=>operation: 将G加入sendq队列
  op7=>operation: 待换醒
  op8=>operation: 从buf队首取出数据
  op9=>operation: 取出一个sendG
  op10=>operation: 将sendG数据写入buf队尾
  op11=>operation: 唤醒sendG
  op12=>operation: 从buf中取出数据
  op13=>operation: 将当前的recveG加入recveq
  op14=>operation: recveG待唤醒
  
  cond=>condition: sendq非空？
  cond1=>condition: 有buf？
  cond2=>condition: buf非空？
  e=>end
  
  st->op1->cond
  cond(yes)->cond1
  cond(no)->cond2
  
  cond1(yes)->op8->op9->op10->op11->e
  cond1(no)->op3->op4->op5->e
  
  cond2(yes)->op12->e
  cond2(no)->op13->op14->e
  
  ```

  - 情况1（无缓冲channel，发送G阻塞）：取出sendG，获取sendG的数据并唤醒。
  - 情况2（有缓冲channel，发送G阻塞）：从buf队首取数据，从sendq取出sendG，将sendG的数据写入channel，并唤醒。
  - 情况3（有缓冲channel，缓冲区有数据，无发送G阻塞）:从缓冲区中取出数据。
  - 情况4（无阻塞sendG，缓冲区无数据）：将当前的recveG加入recveq，阻塞当前recveG。

- 思考

  - 这里思考一个问题，那 goroutine1 和 goroutine2 又怎么互相知道自己的数据 ”到“ 了呢？
    - channel结构中的recvq、sendq保存着阻塞等待的goroutine，但goroutine1向环型队列中发送数据时，就会从recvp取出goroutine并唤醒。

- 参考

  - https://mp.weixin.qq.com/s/ZXYpfLNGyej0df2zXqfnHQ
  - https://www.cnblogs.com/-wenli/p/12710361.html
  - https://segmentfault.com/a/1190000019172554