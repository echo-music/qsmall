# 微服务项目搭建

## 一、创建用户项目
 
### 1、添加用户proto文件

```
qsmall项目根木目录下执行如下命令：

kratos proto add api/user/user.proto

├── README.md
├── api
│   └── user
│       └── user.proto
├── app
│   └── user


```
输出的proto文件内容如下
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
### 2、生成proto代码
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

### 3、生成service代码
```
kratos proto server api/user/user.proto -t app/user/internal/service


```
