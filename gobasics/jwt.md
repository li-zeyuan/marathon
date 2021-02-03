# JWT认证

### 组成

- ```
  header.payload.signature
  ```

- 头部（header）：指明令牌的类型、加密算法等。Base64 编码

- 载荷（payload）：保存用户信息。Base64 编码

- 签名（signature）：base64后的头信息+"."+base64后的载荷信息+Secret，加密生成

  - ```
    HMACSHA256(
      base64UrlEncode(header) + "." +
      base64UrlEncode(payload),
      Secret
    )
    ```

  - HMACSHA256加密算法不可逆

  - token传到服务端后，用同样的算法对base64后的头信息+"."+base64后的载荷信息加密，和传过来的Secret比对。起到防篡改作用

### 单设备登录

- 服务端生成token的同时，在rides中保存token的白名单为最后一个token
- 中间件校验token时，需要和rides中的token白名单对比

### 参考

- https://www.cnblogs.com/fengzheng/p/13527425.html

