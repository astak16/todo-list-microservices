package main

import (
	"user/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitDB()
	initialize.NewEtcdServer().Register(10)
	server := initialize.InitGrpc()
	defer server.Stop()
}
