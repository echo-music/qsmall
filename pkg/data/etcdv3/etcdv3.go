package etcdv3

import (
	etcdclient "go.etcd.io/etcd/client/v3"
	"sync"
)

var cli *etcdclient.Client
var once sync.Once

func Cli() *etcdclient.Client {

	once.Do(func() {
		var err error
		cli, err = etcdclient.New(etcdclient.Config{
			Endpoints: []string{"127.0.0.1:2379"},
		})
		if err != nil {
			panic(err)
		}
	})

	if cli == nil {
		panic("etcd 未初始化客户端连接")
	}
	return cli
}
