package initialize

import (
	"context"
	"net"
	"strconv"
	"task/etcd"
	"task/global"
	"task/handler"
	"task/service"

	"google.golang.org/grpc"
)

var (
	ctx = context.Background()
	// prefix     = global.EtcdConfig.Prefix
)

func InitGrpc() *grpc.Server {
	port := strconv.Itoa(global.Server.Port)
	etcdServer := global.EtcdConfig.Host + ":" + strconv.Itoa(global.EtcdConfig.Port)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	client, err := etcd.NewClient(ctx, []string{etcdServer}, etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	r := etcd.NewRegistrar(client, etcd.Service{
		Key:   global.EtcdConfig.Prefix,
		Value: "http://" + global.Server.Host + ":" + port,
	})

	service.RegisterTaskServiceServer(s, handler.NewTaskService())

	r.Register()
	defer r.Deregister()
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

	return s
}
