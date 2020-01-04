package main

import (
	"fmt"
	//"github.com/samuel/go-zookeeper/zk"
	//"learn.com/zookeep/util"
	"github.com/samuel/go-zookeeper/zk"
	"learn.com/bigdata/zookeep/util"
	"learn.com/rpc/thrift/demo/thriftdemo/server/imp"
)

func main() {
	var server01 = "127.0.0.1:6688"
	zkList := []string{util.Zk_host}
	conn := util.GetConnect(zkList)
	defer conn.Close()
	conn.Create(util.Zk_work_path+"/"+server01, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	err := imp.StartServer(server01)
	if err != nil {
		fmt.Printf("startServer() started failed. message=%s \n", err.Error())
	}
}
