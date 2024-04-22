package initialize

import (
	"api-gateway/etcd"
	"api-gateway/global"
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

var (
	ctx = context.Background()
)

func InitGrpc() (*grpc.ClientConn, error) {
	etcdServer := global.EtcdConfig.Host + ":" + strconv.Itoa(global.EtcdConfig.Port)
	prefix := global.EtcdConfig.Prefix

	client, err := etcd.NewClient(ctx, []string{etcdServer}, etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	d := etcd.NewDiscovery(client, etcd.Service{Key: prefix})
	resolver.Register(d)

	conn, err := grpc.Dial(d.Scheme()+":///"+prefix,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		fmt.Println("和rpc建立连接失败：", err)
		return nil, err
	}
	return conn, nil
}
