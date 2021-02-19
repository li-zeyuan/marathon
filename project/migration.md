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

### 难点

- 脚本迁移的顺序问题
  - 问题
    - 1、若按照创建顺序执行，同事没有拉最新代码，导致创建序号相同的脚本
    - 2、若顺序顺序记录到数据库，想要修改执行顺序需要修改数据库，正式服比较麻烦
  - 解决
    - 创建迁移脚本记录到index.json文件，执行按照index.json文件中的顺序进行执行。

### 参考

- 命令行包：https://github.com/urfave/cli