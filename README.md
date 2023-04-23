# 微服务项目搭建

## 一、创建用户项目

### 1、添加 proto 文件
定义rpc接口
```
qsmall项目根木目录下执行如下命令：

kratos proto add api/user/user.proto

在api/user目录下创建一个 user.proto 文件

├── README.md
├── api
│   └── user
│       └── user.proto
├── app
│   └── user


```
user.proto 文件内容如下：
```
syntax = "proto3";

package api.user;

option go_package = "qsmall/api/user;user";
option java_multiple_files = true;
option java_package = "api.user";

service User {
	rpc CreateUser (CreateUserRequest) returns (CreateUserReply);
	rpc UpdateUser (UpdateUserRequest) returns (UpdateUserReply);
	rpc DeleteUser (DeleteUserRequest) returns (DeleteUserReply);
	rpc GetUser (GetUserRequest) returns (GetUserReply);
	rpc ListUser (ListUserRequest) returns (ListUserReply);
}

message CreateUserRequest {}
message CreateUserReply {}

message UpdateUserRequest {}
message UpdateUserReply {}

message DeleteUserRequest {}
message DeleteUserReply {}

message GetUserRequest {}
message GetUserReply {}

message ListUserRequest {}
message ListUserReply {}
```
当然也可以自己定义DML


### 2、生成*.pb 和 *.grpc.pb 代码
生成rpc服务代码
```
qsmall项目根木目录下执行如下命令：

kratos proto client api/user/user.proto


├── README.md
├── api
│   └── user
│       ├── user.pb.go
│       ├── user.proto
│       └── user_grpc.pb.go

```

### 3、生成实现 grpc service 的代码
生成实现好的rpc服务代码
```
kratos proto server api/user/user.proto -t app/user/internal/service

```

### 4、grpc 和 http 服务实例的创建
grpc 和 http 服务实例的创建

app/user/internal/server/grpc.go
```
func NewGRPCServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			tracing.Server(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	user.RegisterUserServer(srv, greeter)
	return srv
}
```

app/user/internal/server/http.go
```
func NewHTTPServer(c *conf.Server, greeter *service.UserService, logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		)),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	user.RegisterUserHTTPServer(srv, greeter)
	return srv
}

```
使用 wire 来管理依赖的 服务
app/user/internal/server/server.go
```
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

```

### 5、服务入口 

app/user/cmd/user/main.go
```

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "user"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}


func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {

	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}



func main() {
	flag.Parse()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
	
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
```

使用wire 来管理依赖
app/user/cmd/user/wire.go
```
package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"qsmall/app/user/internal/biz"
	"qsmall/app/user/internal/conf"
	"qsmall/app/user/internal/data"
	"qsmall/app/user/internal/server"
	"qsmall/app/user/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

```


### 4、服务之间的依赖关系
service->biz->repo->data->{mysql,redis,mq,etcd}
```
service 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
biz     业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
repo    业务接口定义    
data    业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来。
```
^_^  开发的时候应该从依赖少的对象开发，在这里我们应该从mysql,redis,mq 开始干活

完整的目录结构：
```
   .
├── Dockerfile  
├── LICENSE
├── Makefile  
├── README.md
├── api // 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
│   └── helloworld
│       └── v1
│           ├── error_reason.pb.go
│           ├── error_reason.proto
│           ├── error_reason.swagger.json
│           ├── greeter.pb.go
│           ├── greeter.proto
│           ├── greeter.swagger.json
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd  // 整个项目启动的入口文件
│   └── server
│       ├── main.go
│       ├── wire.go  // 我们使用wire来维护依赖注入
│       └── wire_gen.go
├── configs  // 这里通常维护一些本地调试用的样例配置文件
│   └── config.yaml
├── generate.go
├── go.mod
├── go.sum
├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
│   │   ├── README.md
│   │   ├── biz.go
│   │   └── greeter.go
│   ├── conf  // 内部使用的config的结构定义，使用proto格式生成
│   │   ├── conf.pb.go
│   │   └── conf.proto
│   ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
│   │   ├── README.md
│   │   ├── data.go
│   │   └── greeter.go
│   ├── server  // http和grpc实例的创建和配置
│   │   ├── grpc.go
│   │   ├── http.go
│   │   └── server.go
│   └── service  // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
│       ├── README.md
│       ├── greeter.go
│       └── service.go
└── third_party  // api 依赖的第三方proto
    ├── README.md
    ├── google
    │   └── api
    │       ├── annotations.proto
    │       ├── http.proto
    │       └── httpbody.proto
    └── validate
        ├── README.md
        └── validate.proto

```
好了现在我们分别定义这些服务
(1)创建mysql客户端实例


```

```



### 4、错误码文件生成

(1)安装错误插件

```
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest

```

(2)错误定义
api/user/user.proto

```
enum ErrorReason {
	// 设置缺省错误码
	option (errors.default_code) = 500;

	// 为某个枚举单独设置错误码
	USER_NOT_FOUND = 0 [(errors.code) = 404];

	CONTENT_MISSING = 1 [(errors.code) = 400];
}

```

(3)根目录下执行 make api 或 make errors 生成错误代码文件 user_errors.pb
如果使用make errors,需要在根目录下的Makefile文件下定义该生成命令：
protoc --proto_path=. \
--proto_path=./third_party \
--go_out=paths=source_relative:. \
--go-errors_out=paths=source_relative:. \
$(API_PROTO_FILES)

然后执行 make erros 生成 user_errors.pb 文件，部分代码如下：

```
// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package user

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

// 为某个枚举单独设置错误码
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}


```

### 5、参数校验

```


```


