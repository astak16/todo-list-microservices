package initialize

import (
	"fmt"
	"net"
	"user/global"
	"user/handler"
	"user/service"

	"google.golang.org/grpc"
)

func InitGrpc() *grpc.Server {
	server := grpc.NewServer()
	service.RegisterUserServiceServer(server, handler.NewAuthService())
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", global.Server.Port))
	if err != nil {
		panic(err)
	}
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
	return server
}
