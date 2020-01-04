package main

import (
	"fmt"
	"runtime"
	"sync"
)

//用锁的机制实现,主线程等待所有协程计算完毕
//用锁的机制实现同步
var counter int = 0

//同步访问共享数据counter
func count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func ThreadStart() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go count(lock) //提交10个任务=10个并发
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched() //让出时间片
		if c >= 10 {
			break
		}
	}
}
