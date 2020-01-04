package main

import (
	"fmt"
	//"sync"
	"time"
)

/**
生产者-消费者模式
go语言目前版本,一下代码强制是异步IO, 所以无法固定receiv和send的输出顺序

**/
//生产者
func Producer(queue chan<- int) {
	//lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		//lock.Lock()
		queue <- i
		fmt.Println("send:", i)
		//lock.Unlock()
	}
}

//消费者
func Consumer(queue <-chan int) {
	//lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		//lock.Lock()
		v := <-queue
		fmt.Println("receiv:", v)
		//lock.Unlock()
	}
}
func startProducer() {

	queue := make(chan int, 1)
	go Producer(queue)
	go Consumer(queue)
	time.Sleep(3 * time.Second)
}
