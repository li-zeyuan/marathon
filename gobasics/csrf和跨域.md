# CSRF和跨域

### CSRF

- ![](https://raw.githubusercontent.com/li-zeyuan/access/master/img/20210203152634.png)

- 跨站请求伪造。攻击者伪造正常用户的身份，发送恶意请求。
- from表单POST提交，不受浏览器同源策略的限制。
- cookie保存认证信息不安全，攻击者容易伪造。
- 参考
  - https://www.cnblogs.com/hyddd/archive/2009/04/09/1432744.html

### 跨域

- 因为有浏览器同源策略的限制，ajax向其他源发送请求，响应会被浏览器拦截

- 解决

  - CORS

    - 当前端需要发送复杂请求时，需要发"预检"请求，方法为option
    - 后端响应头返回Access-Control-Allow*

  - nginx反向代理

- 参考

  - https://juejin.cn/post/6844903767226351623#heading-13