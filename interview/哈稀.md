# first

### 读写锁的实现

- 用一个全局的int变量，32位
- 高16位表示读锁，加锁则加1
- 低16位表示写锁，加锁则加1

# second

### 1、逻辑题

- 1、答案：8次
  - 一、30/5=6次，找出前六名
  - 二、6/5 = 1次，找出前两名
  - 三、前两名和另一匹马比，找到前二名

- 2、答案：最少三小时
  - ·刚好第一只小白鼠就喝到有毒的水，用时3小时

### 2、算法题

- 1、
  - 一、对801个号码排序
  - 二、中间这个号码就是落单的号码
  - 时间复杂度：O(logN)
- 2、
  - 从头到尾遍历卖票记录
  - 一个快指针，一个慢指针
  - 若快指针所指向的数据等于慢指针指向的数据，则重复票就是这张
  - 时间复杂度：O(N)

### 3、Go

- 1

  - 类型：有缓冲区channel，无缓冲区channel
  - 区别：有缓冲区channel的缓冲区没有满时，发送者不阻塞；无缓冲区channel发送者阻塞，直到有接收者从channel中取出数据

- 2

  - 存在问题：并发情况下数据不准确

  - 解决：加互斥锁

  - ```go
    import (
    	"fmt"
    	"net/http"
    	"net/url"
    	"sync"
    )
    
    var (
    	userCount = make(map[string]int64, 0)
    	lock      sync.Locker
    )
    
    func main() {
    	http.HandleFunc("/", HttpHandle)
    	// 启动服务器
    	err := http.ListenAndServe(":20520", nil)
    	if err != nil {
    		panic(err)
    	}
    }
    func HttpHandle(w http.ResponseWriter, r *http.Request) {
    	if URL, err := url.ParseRequestURI(r.RequestURI); err == nil {
    		query, err1 := url.ParseQuery(URL.RawQuery)
    		if err1 != nil {
    			return
    		}
    
    		username := query.Get("username")
    		lock.Lock()
    		defer lock.Unlock()
    		if count, ok := userCount[username]; ok {
    			count++
    			userCount[username] = count
    		} else {
    			userCount[username] = 1
    		}
    	}
    }
    ```

- 3

  - ```go
    func CheckSrt(str string) bool {
    	tempList := make([]string, 0)
    	for _, s := range str {
    		l := len(tempList)
    
    		switch string(s) {
    		case ")":
    			if tempList[l-1] == "(" {
    				tempList = tempList[:l-1]
    				continue
    			}
    
    			return false
    		case ">":
    			if tempList[l-1] == "<" {
    				tempList = tempList[:l-1]
    				continue
    			}
    
    			return false
    		case "}":
    			if tempList[l-1] == "{" {
    				tempList = tempList[:l-1]
    				continue
    			}
    
    			return false
    		case "]":
    			if tempList[l-1] == "[" {
    				tempList = tempList[:l-1]
    				continue
    			}
    
    			return false
    		default:
    			tempList = append(tempList, string(s))
    		}
    	}
    
    	return true
    }
    ```

### 4、MySQL

- 1、
  - 类型：主键索引、唯一索引、组合索引、全文索引、普通索引
  - innodb的索引底层的数据结构一般是B+tree
  - 非叶子节点不保存数据，仅保存指向下一个节点的指针
  - 叶子节点保存数据，并且形成链表，方便范围查询

- 2、
  - 原因：
    - 1、数据量大，count时，这张表有800W-1000W的数据
    - 2、每5分钟跑的定时任务删除数据占用了数据库cpu
    - 3、删除的数据后，数据库需要重新建立索引
  - 解决：
    - 不删除表的数据
    - 水平分表，没天的数据存到一张表中
- 3、
  - 原因：
    - 两个事物锁住一批资源，相互等待对方释放资源后才能继续进行
  - 解决：
    - 一般是在事物中嵌套事物

### 5、Redis

- 1、
  - string：缓冲token
  - hash：对象属性
  - zset：排行榜
  - list：缓存个人名下数据

- 2、
  - 1、在Redis中设置一个key，并设置过期时间
  - 2、监听key的过期时间
  - 3、过期则执行对应的逻辑

