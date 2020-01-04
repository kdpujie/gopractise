package util

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

const (
	Zk_host      string = "10.69.56.55:2181"
	Zk_work_path        = "/thrift-rpc/RtbService/providers"
)

/**
和获取zk连接
@return zk.Conn
**/
func GetConnect(zkList []string) (conn *zk.Conn) {
	conn, _, err := zk.Connect(zkList, 10*time.Second)
	if err != nil {
		fmt.Printf("zk.Connect() err ,errMes=%s", err.Error())
	}
	return
}
