package main

import (
	"fmt"
	"ksyun.com/commons/tool/thrift"
	"time"
)

func main() {
	zkList := []string{"192.168.153.44:2181"}
	watchPath := "/go_servers/bidder"
	registry := thrift.NewZkRegistry(zkList, 2*time.Second, watchPath)
	time.Sleep(2 * time.Second)
	var i int
	for {
		i++
		time.Sleep(1 * time.Second)
		endpoint, _, err := registry.ServiceEndpoint()
		if err != nil {
			fmt.Printf("ERR: %s\n", err.Error())
			continue
		}
		fmt.Printf("times %d: EndPoint=%s \n", i, endpoint)

	}
}