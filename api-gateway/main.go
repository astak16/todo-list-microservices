package main

import (
	"api-gateway/initialize"
	"fmt"
)

func main() {
	initialize.InitConfig()
	conn, err := initialize.InitGrpc()
	if err != nil {
		fmt.Println("和rpc建立连接失败：", err)
	}

	initialize.NewHttpServer(conn)
}
