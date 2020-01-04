package main

import (
	"fmt"
	"learn.com/rpc/thrift/demo/thriftdemo/server/imp"
)

const (
	hostPort = "localhost:50051"
)

func main() {

	err := imp.StartServer(hostPort)
	if err != nil {
		fmt.Printf("startServer() started failed. message=%s \n", err.Error())
	}
}
