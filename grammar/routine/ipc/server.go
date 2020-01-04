package ipc

import (
	"encoding/json"
	"fmt"
)

/**
定义server的统一接口
**/
type Request struct {
	Method string "method"
	Params string "params"
}
type Response struct {
	Code string "code"
	Body string "body"
}
type Server interface {
	Name() string
	Handle(method, params string) *Response
}
type IpcServer struct {
	Server
}

//当做构造函数
func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0) // unbuffered channel of strings
	go func(c chan string) {
		for {
			request := <-c
			if request == "close" { //关闭连接
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b) //返回结果
		}
		fmt.Println("Session closed.")
	}(session)
	fmt.Println("A new session has been created successfully.")
	return session
}