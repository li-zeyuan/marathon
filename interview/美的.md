# first

### goroutine内存泄漏

- channel读写阻塞

- select阻塞

- 死循环

  ```
  解决方式：
  1、借助阿里云内存监控，查看内存是否飙升
  2、用go pprof
  ```

  

### goroutine终止

- main退出
- channel通知退出
- content通知退出
- panic退出
- 执行完成后退出

### linux查找大文件

- du 
- 参考：https://blog.csdn.net/CL_YD/article/details/79458092

### 查看系统配置

- uname -a

### 安全关闭channel

- 关闭已经关闭的channel会panic
- 向已经关闭的channel发送数据会panic

### 常使用的一些包

- fmt
- io
- buff
- strconv
- os
- sync
- josn
- http

### protobuf 和json 的区别

- protobuf的编码解码比json快
- protobuf的内存占用更少

### 系统线程和用户线程的区别

- 参考：https://blog.csdn.net/dan15188387481/article/details/49450491

### Nginx如何配置

- 参考：https://www.runoob.com/w3cnote/nginx-setup-intro.html

