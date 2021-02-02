# cron

- 实现了Linux中crontab这个命令的效果
- 秒级别，Linux中的crontab是分钟级别
- 通过表达指定定时任务的执行周期
- AddFunc增加任务，每个任务开启一个goroutine来执行
- 参考：https://segmentfault.com/a/1190000023029219