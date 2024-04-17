package main

import (
	"context"
	"fmt"
	"user/etcd"
	"user/etcd/proto"

	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

const ServerName = "greeter"

var (
	etcdServer = "etcd-server:2379"  // in the change from v2 to v3, the schema is no longer necessary if connecting directly to an etcd v3 instance
	prefix     = "/services/foosvc/" // known at compile time
	// instance   = "127.0.0.1:10001"    // taken from runtime or platform, somehow
	port     = "8080"
	host     = "127.0.0.1"
	instance = host + ":" + port
	key      = prefix // should be globally unique
	// value    = "http://" + instance // based on our transport
	ctx = context.Background()
)

func main() {

	for {
		sayHello()
		time.Sleep(10 * time.Second)
	}

}

func sayHello() {

	client, err := etcd.NewClient(ctx, []string{etcdServer}, etcd.ClientOptions{})
	if err != nil {
		panic(err)
	}

	d := etcd.NewDiscovery(client, etcd.Service{Key: prefix})
	resolver.Register(d)
	conn, err := grpc.Dial(d.Scheme()+":///"+prefix,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println("和rpc建立连接失败：", err)
		return
	}

	cli := proto.NewGreeterClient(conn)
	in := &proto.HelloRequest{Msg: "Hello Service"}

	r, err := cli.SayHello(ctx, in)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Msg)
}
