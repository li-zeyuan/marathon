# error

### 数据结构

- ```go
  type error interface {
  	Error() string
  }
  ```

- 本质是一个接口，有一个Error方法

### 自定error

- 1、通过errors包，它有一个New()方法
- 2、通过fmt.Errorf()
- 3、自定义errorenum模块，实现Error()方法