# 1 架构图
缩短从需求到上线的距离
go-zero 是一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

go-zero 包含极简的 API 定义和生成工具 goctl，可以根据定义的 api 文件一键生成 Go, iOS, Android, Kotlin, Dart, TypeScript, JavaScript 代码，并可直接运行。

使用 go-zero 的好处：

轻松获得支撑千万日活服务的稳定性
内建级联超时控制、限流、自适应熔断、自适应降载等微服务治理能力，无需配置和额外代码
微服务治理中间件可无缝集成到其它现有框架使用
极简的 API 描述，一键生成各端代码
自动校验客户端请求参数合法性
大量微服务治理和并发工具包
!["架构图"](./architecture.png)


# 2 单体服务
最快生成一个服务
```shell
$ mkdir go-zero-demo
$ cd go-zero-demo
$ go mod init go-zero-demo
$ goctl api new greet
$ go mod tidy
Done.
```
查看一下greet服务的目录结构

```shell
$ tree greet
greet
├── etc
│   └── greet-api.yaml
├── greet.api
├── greet.go
└── internal
    ├── config
    │   └── config.go
    ├── handler
    │   ├── greethandler.go
    │   └── routes.go
    ├── logic
    │   └── greetlogic.go
    ├── svc
    │   └── servicecontext.go
    └── types
        └── types.go
```

复杂一点的单体服务,主要区别还是在api文件的定义上,可以通过划分多个api文件

```
@server(
    jwt: Auth
    middleware: CheckUrl
    group: order/cart
    prefix: /api/order/cart
)
service admin-api {
    @handler CartItemAdd
    post /add (addCartItemReq) returns (addCartItemResp)
}
```

# 3 goctl
# 4 api
# 5 rpc
# 6 组件深入

