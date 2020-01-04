/**
@参考信息
[深入理解Go Channel](http://blog.csdn.net/kongdefei5000/article/details/75209005)
[3种优雅的Go channel用法](http://blog.csdn.net/andylau00j/article/details/53934142)
**/

package main

import "fmt"
import "strconv"

/**
利用channel的阻塞机制,等待所有协程计算完毕
**/

func count1(value string, ch chan string) {
	fmt.Println("进入channel-", value)
	ch <- value
	fmt.Println("channel-", value, " save over.")
}

func ChannelStart() {
	chs := make([]chan string, 10) //chan数组
	for i := 0; i < 10; i++ {
		chs[i] = make(chan string)
		go count1(strconv.Itoa(i), chs[i])
	}
	for _, ch := range chs {
		j := <-ch
		fmt.Println("channel - ", j, " get over.")
	}
}
