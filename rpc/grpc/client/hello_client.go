/**
@description grpc helloworld client示例程序
**/
package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"pujie.org/rpc/grpc/entry"
	"time"
)

func sayHello(client entry.GreeterClient) {
	r := &entry.Request{
		Id:   "0001",
		Name: "pujie",
		Uid:  "123456789987654321",
		Ip:   "10.69.58.201",
	}
	aggregation := &entry.Aggregation{
		Request: r,
	}
	resAggregation, err := client.GetStatus(context.Background(), aggregation)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("==========request info=========%p \n", aggregation)
	printAggregation(aggregation)
	log.Printf("==========response info===========%p \n", resAggregation)
	printAggregation(resAggregation)
}

func main() {
	conn, err := grpc.Dial(entry.HostPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := entry.NewGreeterClient(conn)
	for i := 0; i < 10; i++ {
		start := time.Now()
		sayHello(c)
		sub := time.Now().Sub(start)
		log.Printf("第%d次调用，耗时%d us \n", i, sub/1000)
	}

}

func printAggregation(a *entry.Aggregation) {
	log.Printf("err_code:%s \n", a.ErrorCode)
	if r := a.Request; r != nil {
		log.Printf("request.id:%s \n", r.Id)
		log.Printf("request.Uid:%s \n", r.Uid)
		log.Printf("request.name:%s \n", r.Name)
		log.Printf("request.ip:%s \n", r.Ip)
	}
	if res := a.Response; res != nil {
		log.Printf("response.id:%s \n", res.Id)
		log.Printf("response.message:%s \n", res.Message)
		log.Printf("response.status:%s \n", res.Status)
	}
}
