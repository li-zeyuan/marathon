# 视频剪辑

### 需求

- 学管需要转发一些上课中的精彩片段给家长观看（监课），并且对学生的课堂表现进行评价

### 流程

```flow
st=>start: Start
op=>operation: 通过回放video_id下载回放视频到本地
op1=>operation: ffmpeg对视频进行剪切，保持片段到本地
op2=>operation: 将片段上传到oss
op3=>operation: 清除回放视频、视频片段
op4=>operation: 更新视频剪辑状态为已完成

e=>end

st->op->op1->op2->op3->op4->e
```

- 执行bash脚本，脚本实现了FFmpeg命令对视频进行剪辑，剪切完成临时保存到本地

  - ```go
    func Execute(execStr string) (err error, lines []string) {
    	cmd := exec.Command("/bin/sh", "-c", execStr) // 执行脚本
    	stdout, err := cmd.StdoutPipe()
    	err = cmd.Start()
    	reader := bufio.NewReader(stdout)
    
    	//实时循环读取输出流中的一行内容
    	for {
    		line, err2 := reader.ReadString('\n')
    		if err2 != nil || io.EOF == err2 {
    			break
    		}
    		fmt.Println(line)
    		lines = append(lines, strings.Replace(line, "\n", "", -1))
    	}
    
    	err = cmd.Wait()
    	return
    }
    ```

    

### 相关接口

- 剪辑接口
- 轮询剪辑状态接口
- 剪辑视频列表接口

### 遇到问题

- 1、学管反应有时剪辑速度慢（30分钟以上）
  - 原因：回放视频合成与视频剪辑使用的是同一个队列同一个消费者，回放视频合成任务使用了redis做了缓存，合成任务串行执行，导致视频剪辑任务阻塞。
  - 解决：回放视频合成，视频剪辑分开两个队列，用两个消费者进行处理

### 参考

- FFmpeg视频处理：https://www.ruanyifeng.com/blog/2020/01/ffmpeg.html