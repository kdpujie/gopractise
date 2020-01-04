package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
)

//计数
type Counter struct {
	Sum int
}

func (this *Counter) Add(i int, r *int) error {
	this.Sum += i
	*r = this.Sum
	return nil
}

func (this *Counter) Echo(args string, reply *string) error {
	*reply = "echo:" + args
	return nil
}

func (this *Counter) TestMap(args []string, reply *string) error {
	fmt.Println("ddddddddd")
	for _, k := range args {
		*reply = *reply + k
	}

	return nil
}

type User struct {
	Name string
	Age  int
}

func (this *User) Print(user User, reply *string) error {
	*reply = "姓名:" + user.Name + ",年龄:" + strconv.Itoa(user.Age)
	fmt.Println(*reply)
	return nil
}

func NewJsonRpcSocketServer() {
	rpc.Register(Counter{})
	l, err := net.Listen("tcp", ":3333")
	if err != nil {
		fmt.Printf("Listener tcp err : %s", err)
		return
	}
	var index int = 0
	for {
		fmt.Printf("wating.....\n")
		conn, err := l.Accept()
		if err != nil {
			fmt.Sprintf("accept connect err:%s \n", conn)
		}
		index = index + 1
		fmt.Printf("服务请求数%d \n", index, conn.RemoteAddr().String())

		go jsonrpc.ServeConn(conn)
	}

}
