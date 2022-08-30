### 技术栈
- 前端
   - Vue3.2
   - Ant Design Vue
   - Axios
   - RSA
   - Vuex
   - Vue-router
   - TS
   - less
- 后端
   - gin
   - gorm
   - goini
   - redis
   - RSA
   - mysql

---

### 注册

- 前端
   - 参数规则
      - 用户为全英文 5-12位
      - 密码 5-12位
      - 验证码 4 位
   - 获取验证码
   - 预校验
   - 请求后端
   - 页面响应
- 后端
   - 配置路由
   - 获取参数
   - 生成验证码
   - 验证码存入redis
   - 返回验证码
   - 校验验证码
   - 删除redis验证码
   - 密码加密
   - 用户信息存入数据库
   - 返回结果

---

### 登录
**前端**

- 参数预校验
- 密码RSA加密
- 请求后端
- 存储token
   - Vuex
   - localstorge
- 响应页面
- 跳转路由 
   - router.replace('/home')  

**后端**

- 配置路由
- 获取参数
- RSA解密
- 校验密码
- 生成 token
- 返回token和结果

---
### 参考资料
1. [前端加密](http://sjsj)
