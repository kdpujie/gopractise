package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	testSwitch()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2)
	fmt.Println("进程号:", os.Getpid())
	//阻塞直至有信号传入
	s := <-c
	fmt.Println("get signal:", s)
}

func testSwitch() {
	var i int = 2
	switch {
	case i > 1:
		fmt.Println("i > 1")
		fallthrough
	case i > 2:
		if i <= 2 {
			break
		}
		fmt.Println("i > 2")
		fallthrough
	case i > 3:
		fmt.Println("i > 3")
	}
}
