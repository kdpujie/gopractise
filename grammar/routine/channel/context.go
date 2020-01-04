package main

import (
	"fmt"
	"context"
	"time"
	cu "ksyun.com/commons/util"
	"sync"
)

func main()  {
	outChan:=make(chan int, 1)
	ctx ,cancel:= context.WithTimeout(context.Background(),5 * time.Second)

	defer cancel()
	var wg sync.WaitGroup
	wg.Add(10)
	testFilter := func(index int) {
		defer wg.Done()
			select {
			case outChan <- filter(ctx, index):
				fmt.Printf("====filter成功返回,index = %d \n",index)
			case <-ctx.Done():
				fmt.Printf("----filter过期返回,index = %d \n",index)
			}
	}
	for i :=0; i < 11 ; i++ {
		go testFilter(i)
	}
	go func() {
		wg.Wait()
		close(outChan)
	}()
	for i := range outChan {
		fmt.Printf("消费者：接收filter%d的数据。。。。\n",i)
	}

}



func filter(ctx context.Context, index int) int  {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	outChan:=make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(10)
	inner := func(i int) {
		defer wg.Done()
		select {
		case outChan <- delay(i):
			fmt.Printf("inner成功返回: successed,parent = %d,i=%d \n",index,i)
		case <-ctx.Done():
			fmt.Printf("inner过期返回,parent = %d,i=%d \n",index,i)
		}
	}

	for i :=10; i < 21 ; i++ {
		go inner(i)
	}
	go func() {
		wg.Wait()
		close(outChan)
	}()
	for i := range outChan {
		fmt.Printf("消费者：接收inner%d的数据。。。。\n",i)
		if i >15 {
			cancel()
			return 1000
		}
	}
	return index
}

func delay(i int) int{
	num := time.Duration(cu.RandomInt(10000))
	time.Sleep(num * time.Millisecond)
	return i
}