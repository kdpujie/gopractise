package main

import (
	"fmt"
)

func main() {
	deferCall()
}

func deferCall() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}

/**
	以下代码有什么问题吗?
**/
type student struct {
	Name string
	Age  int
}

func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for key, v := range m {
		fmt.Printf("key=%s, v=%v \n", key, v)
	}
}
