package main

import (
	"fmt"
	"runtime"
)

//gosched让出时间片
func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, " 让出时间片")
		runtime.Gosched() //让出时间片,下一个使用完cpu时间片会返回通知该线程.
		fmt.Println(s, "  被唤醒")
	}
}
func GoschedStart() {
	go say("world")
	say("hello")
}
