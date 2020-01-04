package main

import (
	"github.com/samuel/go-zookeeper/zk"
	"pujie.org/practise/bigdata/zookeep/util"
	"time"
)



func main()  {
	//test_zkconn()
	//test_zk_Ephemeral()
	zkList := []string{"10.69.56.55:2181"}
	conn := util.GetConnect(zkList)
	defer conn.Close()
	/*conn.Create("/thrift-rpc/AdOperationService",nil,0,zk.WorldACL(zk.PermAll))
	conn.Create("/thrift-rpc/AdOperationService/providers",nil,0,zk.WorldACL(zk.PermAll))
	conn.Create("/thrift-rpc/AdOperationService/configurators",nil,0,zk.WorldACL(zk.PermAll))
	conn.Create("/thrift-rpc/AdOperationService/configurators/dispatch",[]byte("2"),0,zk.WorldACL(zk.PermAll))*/
	//conn.Create("/test0/test1/test2",nil,0,zk.WorldACL(zk.PermAll))
	test_zk_Ephemeral(conn)
	time.Sleep(60 * time.Second)
}



/**
测试连接
**/
func test_zkconn(){
	zkList := []string{util.Zk_host}
	conn := util.GetConnect(zkList)
	defer conn.Close()
	conn.Create(util.Zk_work_path,nil,0,zk.WorldACL(zk.PermAll))
	time.Sleep(20 * time.Second)
}

/**
 * 测试临时节点
 */
func test_zk_Ephemeral(conn *zk.Conn)  {
	conn.Create(util.Zk_work_path+"/127.0.0.1:8080", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	time.Sleep(5 * time.Second)
	conn.Create(util.Zk_work_path+"/192.168.153.12:8080", nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))

}



