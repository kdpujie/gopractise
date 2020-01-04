package main

import (
	"fmt"
	"time"
	"learn.com/tools"
)

/**
用筛法求素数的基本思想是：
	把从1开始的、某一范围内的正整数从小到大顺序排列， 1不是素数，首先把它筛掉。
	剩下的数中选择最小的数是素数，然后去掉它的倍数。依次类推，直到筛子为空时结束。
**/
var index int
func main()  {
	tools.StartCPUProfile()
	defer tools.StopCPUProfile()

	//testChan()
	testSelect()
}
func prime()  {
	origin , wait := make(chan int), make(chan struct{})
	processor(origin, wait)
	for num :=2; num<1000; num ++ {
		origin <- num
	}
	time.Sleep(1*time.Second)
	close(origin)
	<-wait
}
/**
并发协程，过滤能被当前prime整除的数(非素数)，不能过滤掉的数，发送给下一个协程处理。
**/
func processor(seq chan int, wait chan struct{})  {
	index++
	go func(i int) {
		prime , ok := <-seq
		if !ok {
			close(wait)
			return
		}
		//fmt.Printf("%d-%d : \n",i,prime)
		out := make(chan int)
		processor(out, wait) //开启下一层过滤。
		for num := range seq {
			//fmt.Printf("\t%d seq:num=%d,prime=%d \n",i,num,prime)
			if num % prime != 0 {
				out <- num
			}
		}
		close(out)
	}(index)
}

func testChan()  {
	origin := make(chan int,10)
	go print(origin)
	for i:=1 ; i< 10; i++ {
		origin <- i
	}
	close(origin)
	fmt.Println("origin生产数据完毕，关闭写入。")
	time.Sleep(5 * time.Second)
}

func print(seq chan int)  {
	for  {
		num , ok :=<- seq
		if !ok { //chan 被关闭
			fmt.Printf("退出：%d，ok=%v \n",num, ok)
			//break
		}
		fmt.Printf("读取值：%d，ok=%v \n",num, ok)
	}

}
func testSelect()  {
	a :=make(chan bool)
	LOOP:for i:=0;i<5; i++ {
		select {
		case f:= <- a:
				fmt.Printf("c输出:%v\n",f)
		case <-time.After(1 * time.Second):
			fmt.Println("超时退出")
			break LOOP
		}
		fmt.Printf("退出select，所在循环 i =%d",i)
	}
	fmt.Println("退出循环")
}