### 基于gin的web后端api服务

 1. 生成公私钥
```
生成私钥
openssl genrsa -out rsa/github.com/zhangshanwen/the_one.rsa
根据私钥生成公钥
openssl rsa -in rsa/github.com/zhangshanwen/the_one.rsa -pubout > rsa/github.com/zhangshanwen/the_one.rsa.pub
```

2.目录树
```markdown
├── api               # api接口
├── bin               # 生成的二进制文件目录
├── cmd               # 程序入口
├── code              # 返回的状态码
├── common            # 常量
├── env               # 配置文件
├── initialize
│   ├── conf          # 配置
│   ├── db            # 初始化db(mysql,redis)
│   ├── init.go      
│   ├── logger        # 初始化日志
│   ├── app           # 初始化api
│   └── service       # api响应与请求
├── internal          
│   ├── header        # 请求头参数
│   ├── param         # 请求参数
│   └── response      # 响应参数
├── middleware        # 中间件
│   ├── jwt.go        # jwt校验
├── model             # 模型
├── router            # 路由
│   ├── api
├── rsa               # jwt 公私钥
└── tools             # 工具
    ├── jwt.go
    ├── jwt_test.go
    └── rsa

```