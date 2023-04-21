package server

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"qsmall/app/user/internal/conf"
	"time"
)

var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {

	c, err := etcdv3.New(etcdv3.Config{
		Endpoints:   conf.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return etcd.New(c)
}
