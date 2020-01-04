package main

import (
	"fmt"
	"log"
	//"net"

	"git.apache.org/thrift.git/lib/go/thrift"
	"learn.com/rpc/thrift/protocol/thriftrpc"
	"runtime/debug"
)

const HostPort = "localhost:50051"

func main() {
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(HostPort)
	//transport, err := thrift.NewTSocket("192.168.153.27:6677")
	if err != nil {
		log.Fatalln(err)
	}

	//useTransport := transportFactory.GetTransport(transport)
	client := thriftrpc.NewBookServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	defer transport.Close()
	getBook(client)
}

func getBook(client *thriftrpc.BookServiceClient) {
	aaa := "aaa"
	_, err := client.ReadBook("Go web 编程", &thriftrpc.Work{1, 2, &aaa})
	if err != nil {
		fmt.Printf("初始化广告失败! errMes=%s \n", err.Error())
		log.Println(string(debug.Stack()))
	}
}
