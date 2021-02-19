# migration迁移工具

### 背景

- 版本上线时，数据处理处理脚本需要逐个依次手动执行，麻烦，容易出错
- 没有脚本执行记录，容易出现重复执行

### 实现功能

- 模仿Django的migration，命令行实现创建脚本模本，执行脚本，重复执行脚本，跳过脚本
- 数据库记录脚本执行，执行记录可查询

### 流程

```flow
st=>start: Start
op=>operation: 创建命令行对象，app := &cli.App{}
op1=>operation: 自定义flag
op2=>operation: 绑定flag的行为（输入命令行后的执行的函数）
op3=>operation: app.Run(os.Args)
op4=>operation: 根据模板文件生成脚本模板
op5=>operation: 记录顺序到index.json文件
op6=>operation: 读取index.json文件，按顺序执行脚本
op7=>operation: pg记录执行结果
op8=>operation: 记录log文件

cond=>condition: --make？
e=>end

st->op->op1->op2->op3->cond
cond(yes)->op4->op5->e
cond(no)->op6->op7->op8->e
```

### 参考

- 命令行包：https://github.com/urfave/cli