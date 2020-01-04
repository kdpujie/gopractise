package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	//"os"
	"time"
)

func startRandom() {
	//rand.Seed(int64(time.Now().Nanosecond()))
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("随机数生成\n")
	for i := 0; i < 18; i++ {

		randNum := random(20)
		fmt.Printf("  %d,", randNum)
	}
	fmt.Println()
	fmt.Println("获取本机IP")
	getIp()
}

func random(num int) int {
	return rand.Intn(65536) % num
}

func getIp() {
	//	addrs, err := net.InterfaceAddrs()

	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}

	//	for _, address := range addrs {

	//		// 检查ip地址判断是否回环地址
	//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//			if ipnet.IP.To4() != nil {
	//				fmt.Println(ipnet.IP.String())
	//			}

	//		}
	//	}
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr().String())
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}
