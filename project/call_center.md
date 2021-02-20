# 外呼功能

### 需求

- 在op系统中可以直接呼叫学员手机进行通话
- 网络电话模式，需要借助桌面虚拟话机软件

### 实现逻辑

```sequence
participant 前端
participant 后端
participant 智齿
前端->后端: 获取智齿pass token
后端->智齿: 获取pass token
Note over 后端: redis缓存
后端->前端: 返回pass token
Note over 前端: 建立ws，发起外呼
前端->后端: 发起外呼成功后，向后端传call_id
Note over 后端: pg存下call_id
后端->前端: 响应成功
Note over 前端: 通话中记时
前端->后端: 通话结束后，添加跟进记录-保存
Note over 后端: pg存下前端记时的通话时长
后端->>前端: 响应
智齿->>后端: 通话结束后，推通话记录详情
Note over 后端: 根据call_id存下智齿的通话时长
后端->>智齿: 响应
```

- redis缓存智齿的token，避免频繁获取
- 通话结束后，通过回调通话详细，数据库记录，方便排查问题
- 外显号码动态配置

### 问题点

- 出现问题，排查难，涉及前端、第三方、后端
  - 解决：增加回调接口推送通话记录，数据库保存，方便提供call_id等
- 外显号码经常更换，更换是需要找运维执行SQL
  - 解决：增加动态配置接口，自己调用接口即可更改外显号码

### 参考

- 接入文档：https://www.sobot.com/developerdocs/service/call_center.html