package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	wait := sync.WaitGroup{}
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wait.Done()
			wait.Add(1)
			cond.L.Lock()
			fmt.Println("Waiting start...")
			cond.Wait()
			fmt.Println("Waiting end...")
			cond.L.Unlock()

			fmt.Println("Goroutine run. Number:", i)
		}(i)
	}

	time.Sleep(time.Second * 5)
	cond.Signal()

	time.Sleep(time.Second * 5)
	cond.Signal()

	time.Sleep(time.Second * 5)
	cond.Signal()

	wait.Wait()
}