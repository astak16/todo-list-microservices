package main

import (
	"context"
	"etcd"
	"etcd/proto"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println(req.Msg)
	return &proto.HelloResponse{Msg: "Hello Client"}, nil
}

var (
	etcdServer = "etcd-server:2379"  // in the change from v2 to v3, the schema is no longer necessary if connecting directly to an etcd v3 instance
	prefix     = "/services/foosvc/" // known at compile time
	// instance   = "127.0.0.1:10001"    // taken from runtime or platform, somehow
	port     = "8080"
	host     = "127.0.0.1"
	instance = host + ":" + port
	key      = prefix + instance    // should be globally unique
	value    = "http://" + instance // based on our transport
	ctx      = context.Background()
)

func main() {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	client, err := etcd.NewClient(ctx, []string{etcdServer}, etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}
	registrar := etcd.NewRegistrar(client, etcd.Service{
		Key:   key,
		Value: value,
	})
	go registrar.Register()
	defer registrar.Deregister()

	proto.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
