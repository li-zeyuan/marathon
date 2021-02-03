# for循环

- 三种形式：for、for+赋值、for+range

- 支持continue、break等操作

- ```go
  func TestFor(t *testing.T) {
  	...
  	for i, v := range l {
  		...
  	}
  	...
  }
  ```

  - i、v是同一块内存

