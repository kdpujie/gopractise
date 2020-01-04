package main

import (
	"fmt"
)

func main()  {
	var sl []string = nil //= []string{"a","b"}
	fmt.Println("slice开始遍历......")
	for i, s :=range sl {
		fmt.Printf("\tindex =%d , value=%s \n", i, s)
	}
}

func startMap() {
	var test map[string]string = map[string]string{"001": "刘备", "002": "诸葛亮", "003": "关羽", "004": "张飞", "005": "赵云", "006": "马超", "007": "黄忠"}
	fmt.Println("初始值:", test)
	delete(test, "001")
	fmt.Println("删除001后:", test)
	clear(test)
	fmt.Println("clear后:", test)
}

func clear(m map[string]string) {
	for key, _ := range m {
		delete(m, key)
	}
}
