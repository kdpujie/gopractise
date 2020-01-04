package main

import (
	"google.golang.org/grpc"
	"log"
	"pujie.org/rpc/grpc/entry"
	"testing"
)

func BenchmarkLoops(b *testing.B) {
	conn, err := grpc.Dial(entry.HostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := entry.NewGreeterClient(conn)
	for i := 0; i < b.N; i++ {
		sayHello(c)
	}
}
