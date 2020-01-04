package main

import (
	"container/list"
	"fmt"
)

func main()  {
	var l list.List
	l.Init()
	l.PushFront(1)
	l.PushFront(2)
	//l.PushFront(3)
	//l.PushFront(4)
	fmt.Printf("list.back()=%d \n",l.Back().Value)
	fmt.Printf("list.front()=%d",l.Front().Value)
}
