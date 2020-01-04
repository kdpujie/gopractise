package main

import (
	"fmt"
	"runtime"
)

type Vector []float64

/**
分解任务,并发执行
**/
//计算任务
func (v *Vector) caculate(start, end int, ch chan int) {
	for ; start < end-1; start++ {
		(*v)[end-1] += (*v)[start]
	}
	fmt.Println(v, " start=", start, " end=", end)
	ch <- 1 //发送信号告诉leader,计算完毕
}

//按cpu核数分解任务
func (v *Vector) splitTask() float64 {
	var num int = runtime.NumCPU()
	ch := make(chan int, num)
	lens := len(*v)
	size := lens / num
	for i := 0; i < num; i++ {
		start, end := i*size, (i+1)*size
		if i == (num - 1) {
			end = lens
		}
		go v.caculate(start, end, ch) //并发(并行?)
	}
	for j := 0; j < num; j++ {
		<-ch //获取到一个数据, 表示一个cpu计算完成.
	}
	for i := 1; i < num; i++ { //最后一个数为 总和
		index := i * size
		(*v)[lens-1] += (*v)[index-1]
	}
	return (*v)[lens-1]
}

func vtest() {

	var u Vector = Vector{1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0}
	value := u.splitTask()
	fmt.Println("总和value=", value)
}
