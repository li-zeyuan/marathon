# http与https

### http

- 无状态，80端口
- 明文传输

### https

- 443端口
- 是http+SSL/TLS
- 需要CA证书
- 采用对称加密和非对称加密的方式

##### 通讯过程

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210321135107.png)

- 客户端请求建立SSL连接
- 服务端返回证书信息（服务端公钥）
- 客户端SSL/TLS校验证书书
- 客户端生成会话秘钥（对称秘钥），用公钥加密发送给服务端
- 服务端用私钥解密得到会话秘钥
- 客户端、服务端通讯用会话秘钥加密