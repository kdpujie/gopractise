/**
	以下代码打印出来什么内容，说出为什么。。。
**/
package main

import (
	"fmt"
)

type People2 interface {
	//Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People2 {
	var stu *Student

	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}
