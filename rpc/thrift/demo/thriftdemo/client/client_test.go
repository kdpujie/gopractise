package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"learn.com/rpc/thrift/protocol/thriftrpc"
	"log"
	"testing"
)

func Benchmark_ReadBook(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		getBook(client)
	}
}
