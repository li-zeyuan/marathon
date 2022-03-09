# go简介

- 静态类型
- 运行是runtime

### 基础题
- https://learnku.com/articles/35063

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

### struct值接收者和指针接收者

- 值类型调用者、指针类型调用均可调用值接受者和指针接收者定义的方法

- 指针接收者可以修改调用方的值，而值类型接收者不可以

- 当有定义接口，接收者为指针时，值类型的调用者不能赋给该接口

  ```go
  type People interface {
  	Say() 
  }
  
  type Student struct {
  	Name string
  }
  
  func (stu *Student) Say()  {
  	fmt.Println("say name: ", stu.Name)
  	return 
  }
  
  func Test36(t *testing.T) {
  	var p People
  	stu := Student{"d"}
  	p = stu // 编译报错
  	p.Say()
  }
  ```

  

### iota

##### 不同 const 定义块互不干扰(遇到const会重置成0)

```go
func Test40(t *testing.T) {
	const a = iota

	const (
		b = iota
	)
	fmt.Println(a) // 0
	fmt.Println(b) // 0
}
```

##### 非空行则加一

- 从第一行开始，iota 从 0 逐行加一
- 没有表达式的常量定义复用上一行的表达式
- 同一行，iota值相等

```go
func Test41(t *testing.T) {
	const (
		c    = 10                        // iota:0; 输出:10; 从第一行开始，iota 从 0 逐行加一
		d    = iota                      // iota:1; 输出:1
		e                                // iota:2; 输出:2; 没有表达式的常量定义复用上一行的表达式
		f    = "hello"                   // iota:3; 输出:hello
		g                                // iota:4; 输出:hello
		h    = iota                      // iota:5; 输出:5
		i                                // iota:6; 输出:6
		j    = 0                         // iota:7; 输出:0
		k                                // iota:8; 输出:0
		l, m = iota, iota                // iota:9; 输出:9, 9; 同一行，iota值相等
		n, o                             // iota:10; 输出:10, 10
		p    = iota + 1                  // iota:11; 输出:11 + 1
		q                                // iota:12; 输出:12 + 1
		_                                // iota:13; 输出:
		r    = iota * iota               // iota:14; 输出:14 * 14
		s                                // iota:15; 输出:15 * 15
		tt   = r                         // iota:16; 输出:196
		u                                // iota:17; 输出:196
		v    = 1 << iota                 // iota:18; 输出:1 << 18
		w                                // iota:19; 输出:1 << 19
		x                  = iota * 0.01 // iota:20; 输出:20 * 0.01
		y    float32       = iota * 0.01 // iota:21; 输出:21 * 0.01
		z                                // iota:22; 输出:22 * 0.01
	)

	fmt.Println(c)  
	fmt.Println(d)  
	fmt.Println(e)  
	fmt.Println(f)  
	fmt.Println(g)  
	fmt.Println(h)  
	fmt.Println(i)  
	fmt.Println(j)  
	fmt.Println(k)  
	fmt.Println(l)  
	fmt.Println(m)  
	fmt.Println(n)  
	fmt.Println(o)  
	fmt.Println(p)  
	fmt.Println(q)  
	fmt.Println(r)  
	fmt.Println(s)  
	fmt.Println(tt) 
	fmt.Println(u)  
	fmt.Println(v)  
	fmt.Println(w)  
	fmt.Println(x)  
	fmt.Println(y)  
	fmt.Println(z)  
}
```

### 参考

- https://segmentfault.com/a/1190000022285902