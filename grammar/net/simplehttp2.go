package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

/**
建立TCP链接来实现初步的HTTP协议(net.DialTCP()函数), 通过向网络主机发送HTTP Head请求, 读取网络主机返回的信息.
**/
func start_simplehttp2() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service) //地址对象
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0/r/n/r/n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
