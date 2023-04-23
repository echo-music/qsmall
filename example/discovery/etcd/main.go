package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"qsmall/api/user"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	dis := etcd.New(client)

	endpoint := "discovery:///qsmall.user.service"
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := user.NewUserClient(conn)

	res, err := c.CreateUser(context.Background(), &user.CreateUserRequest{Name: "dewd"})

	fmt.Println(res, err)

}

