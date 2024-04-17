package main

import (
	"user/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitDB()
	server := initialize.InitGrpc()
	defer server.Stop()
}
