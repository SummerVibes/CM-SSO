# 创明单点登录服务

### 第三方库一览
名称|说明
---|---
[viper]("https://github.com/spf13/viper")|配置信息处理框架
[gorm]("https://jasperxu.github.io/gorm-zh")|数据持久化框架
[gin]("https://gin-gonic.com/zh-cn/")|高性能web框架和路由库
[MySQL]("https://www.mysql.com/cn/")|数据库
[go-cache]("https://github.com/patrickmn/go-cache")|内存键值存储/缓存库，适用于单机程序
[jwt-go]("https://github.com/dgrijalva/jwt-Go")|jwt的go实现

### 部署
1. 编写对应环境的配置文件,生产环境修改config-pro.yaml文件
2. 编译项目 `go build -i -o 路径 .`
3. 将配置文件和可执行程序放在同一目录
4. 设置环境变量`CNF_ENV=dev`或者pro(对应开发和生产环境),设置环境变量`CNF_PATH=`
5. 启动可执行程序

### 功能描述
1. 为其他服务提供登录注册功能,接入第三方平台,提供多种方式
2. 为其他服务提供[单点登录](https://baike.baidu.com/item/%E5%8D%95%E7%82%B9%E7%99%BB%E5%BD%95)功能
3. 以RestAPI提供用户信息的CRUD功能

### 运行流程


### 数据层
数据库的相关操作写在各个model中,供service层调用
定义数据模型,数据序列化器

一共有两张数据表
1. *用户信息表*

存放用户个人信息

id|nick_name|avatar|
---|---|---
1|XueBling|https://sadasd.cn/asdasdas.jpg

2. *授权信息表*

存放用户授权信息

id|user_id|auth_type|identifier|credential|last_login_time|role
---|---|---|---|---|---|---
1|6|phone|12121211212|12312321|2019-1-1|Admin


### Service层
利用dao层实现业务逻辑,为controller层提供服务
1. 用户服务
* 单点登录(核心)
* 退出登录态
* 展示所有用户信息
* 更新用户信息
* 删除用户
* 找回密码
2. 邮箱服务
* 邮箱登录
* 邮箱注册
3. QQ服务
* QQ登录



### Controller层
利用service层处理http请求

### 中间件
1. 验证token,并刷新token过期时间


### 界面设计
1. 功能:登录,注册,找回密码
2. 登录,注册方式有邮箱,QQ
3. 需要考虑页面的自适应布局