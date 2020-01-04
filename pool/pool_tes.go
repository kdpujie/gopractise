package main

import (
	"time"
	"git.apache.org/thrift.git/lib/go/thrift"
	tt "ksyun.com/commons/tool/thrift"
	"log"
	"learn.com/rpc/thrift/protocol/thriftrpc"
	"fmt"
	"runtime/debug"
	"sync"
	"os"
)

var pool * tt.Pool

func main()  {
	defer func() {
		if e, ok := recover().(error); ok {
			log.Printf("error panic: %s \n", e.Error())
			log.Println(string(debug.Stack()))
		}
	}()

	zkList := []string{"123.59.14.199:2181"}
	watchPath := "/thrift-rpc/readbook/providers"
	registry := tt.NewZkRegistry(zkList, 2*time.Second, watchPath)
	pool = tt.NewPool(10,10, 60*time.Second,registry)

	go func() {
		for {
			time.Sleep(100*time.Millisecond)
			pool.Status()
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(100)
	var flag bool = true
	for i:=0 ;i<100;i++ {
		f := func() {
			err := requestClient()
			wg.Done()
			if err != nil {
				os.Exit(0)
			}
		}
		if flag {
			go f()
		}
		//time.Sleep( 100 * time.Millisecond)
	}
	wg.Wait()
}

/**
无重试机制，和server的网络出问题后，正在读写的链接会失败。
**/
func requestClient() error  {
	transport := pool.Get()
	defer transport.Close()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := thriftrpc.NewBookServiceClientFactory(transport, protocolFactory)

	aaa := "aaa"
	name, err := client.ReadBook("Go web 编程", &thriftrpc.Work{1, 2, &aaa})
	if err != nil {
		return err
	}
	fmt.Println("ReadBook", " ", name)
	return nil
}