# panic
### 数据结构

```go
type _panic struct {
    argp      unsafe.Pointer
    arg       interface{}    // panic 的参数
    link      *_panic        // 链接下一个 panic 结构体
    recovered bool           // 是否恢复，到此为止？
    aborted   bool           // the panic was aborted
}
```

```go
// panic函数
// runtime/panic.go
func gopanic(e interface{}) {
    // 在栈上分配一个 _panic 结构体
    var p _panic
    // 把当前最新的 _panic 挂到链表最前面
    p.link = gp._panic
    gp._panic = (*_panic)(noescape(unsafe.Pointer(&p)))
    
    for {
        // 取出当前最近的 defer 函数；
        d := gp._defer
        if d == nil {
            // 如果没有 defer ，那就没有 recover 的时机，只能跳到循环外，退出进程了；
            break
        }

        // 进到这个逻辑，那说明了之前是有 panic 了，现在又有 panic 发生，这里一定处于递归之中；
        if d.started {
            if d._panic != nil {
                d._panic.aborted = true
            }
            // 把这个 defer 从链表中摘掉；
            gp._defer = d.link
            freedefer(d)
            continue
        }

        // 标记 _defer 为 started = true （panic 递归的时候有用）
        d.started = true
        // 记录当前 _defer 对应的 panic
        d._panic = (*_panic)(noescape(unsafe.Pointer(&p)))

        // 执行 defer 函数
        reflectcall(nil, unsafe.Pointer(d.fn), deferArgs(d), uint32(d.siz), uint32(d.siz))

        // defer 执行完成，把这个 defer 从链表里摘掉；
        gp._defer = d.link
        
        // 取出 pc，sp 寄存器的值；
        pc := d.pc
        sp := unsafe.Pointer(d.sp)
        // 如果 _panic 被设置成恢复，那么到此为止；
        if p.recovered {
            // 摘掉当前的 _panic
            gp._panic = p.link
            // 如果前面还有 panic，并且是标记了 aborted 的，那么也摘掉；
            for gp._panic != nil && gp._panic.aborted {
                gp._panic = gp._panic.link
            }
            // panic 的流程到此为止，恢复到业务函数堆栈上执行代码；
            gp.sigcode0 = uintptr(sp)
            gp.sigcode1 = pc
            // 注意：恢复的时候 panic 函数将从此处跳出，本 gopanic 调用结束，后面的代码永远都不会执行。
            mcall(recovery)
            throw("recovery failed") // mcall should not return
        }
    }

    // 打印错误信息和堆栈，并且退出进程；
    preprintpanics(gp._panic)
    fatalpanic(gp._panic) // should not return
    *(*int)(nil) = 0      // not reached
}
```



### panic 究竟是啥？是一个结构体？还是一个函数？
- 背后执行gopanic函数

### 为什么 panic 会让 Go 进程退出的 ？
- 因为gopanic函数调用了**exit(2)** 
### 为什么 recover 一定要放在 defer 里面才生效？
- 因为gopanic函数执行会从当前的挂载的_defer链表取出defer延迟函数执行
### 为什么 recover 已经放在 defer 里面，但是进程还是没有恢复？

- 嵌套panic导致
- defer只对当前的goroutine有效

### 为什么 panic 之后，还能再 panic ？有啥影响？

- defer嵌套panic

### 总结

- panic会新建一个_panic对象，放在链表表头。通过_panic.link下一个panic对象

- panic嵌套，从链表尾部向上递归打印，如

  - ```go
    func main() {
    	defer func() { // 延迟函数
    		panic("panic again")
    	}()
    	panic("first")
    }
    
    // panic: first
    //       panic: panic again
    
    ```

### 参考
- https://jishuin.proginn.com/p/763bfbd651e3
# recover

### 源码

```go
// runtime/panic.go
func recovery(gp *g) {
    // 取出栈寄存器和程序计数器的值
    sp := gp.sigcode0
    pc := gp.sigcode1
    // 重置 goroutine 的 pc，sp 寄存器；
    gp.sched.sp = sp
    gp.sched.pc = pc
    // 重新投入调度队列
    gogo(&gp.sched)
}
```

- 1、`_panic.recovered` 字段被设置成 true 
- 2、修改pc、sp寄存器

### 总结

- panic 的恢复，就是重置 pc 寄存器，直接跳转程序执行的指令，跳转到原本 defer 函数执行完该跳转的位置（`deferreturn` 执行），从 `gopanic` 函数中跳出，不再回来，自然就不会再 `fatalpanic` ；