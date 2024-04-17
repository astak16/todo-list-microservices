package initialize

import (
	"api-gateway/global"
	"api-gateway/middleware"
	"api-gateway/routes"
	"api-gateway/service"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

func NewHttpServer(conn *grpc.ClientConn) {
	userService := service.NewUserServiceClient(conn)

	clients := []middleware.ProtoClient{{Key: "user", Value: userService}}
	m := middleware.InitMiddleware(clients)
	ginRouter := routes.NewRoute(m)

	server := http.Server{
		Addr:           fmt.Sprintf("%s:%d", global.Server.Host, global.Server.Port),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
