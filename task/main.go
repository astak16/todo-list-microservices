package main

import "task/initialize"

func main() {
	initialize.InitConfig()
	initialize.InitDB()
	server := initialize.InitGrpc()
	defer server.Stop()
}
