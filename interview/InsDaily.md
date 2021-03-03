# first

- 翻转切片

  - ```go
    func reverse(s []int) []int {
        for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
            s[i], s[j] = s[j], s[i]
        }
        return s
    }
    ```
- goroutine同步
  - ```go
    func main() {
    		threads := []string{"A", "B", "C", "D", "E"}
    		n := len(threads)
    		// 开启 n 个 goroutine, 每个 goroutine 只输出自身内容, 各自重复 n 遍, 要求所有 goroutine 的输出结果是有序的, 以上面 threads 为例, 要求输出 "ABCDEABCDEABCDEABCDE", threadA 只输出 "AAAAA", threadB 只输出 "BBBBB"
    		wg := sync.WaitGroup{}
    		c := make(chan string, 1)
    		for i := 0; i < n; i++ {
    			for _, t := range threads {
    				wg.Add(1)
    				b := t
    				c <- b
    				go func() {
    					defer wg.Done()
    					a := <-c
    					fmt.Print(a)
    				}()
    			}
    		}
    
    		wg.Wait()
    		close(c)
    }
    ```

  - 