package main

import (
	"fmt"
	"runtime"
	"sync"
)

var a string
var once sync.Once

/**
单例
**/
func setup() {
	a = "abcd"
	b := []byte(a)
	fmt.Println(b)
}
func start_once() {
	go once.Do(setup)
	go once.Do(setup)
	//go setup()
	//go setup()
	runtime.Gosched()

}
