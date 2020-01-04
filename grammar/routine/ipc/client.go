package ipc

import (
	"encoding/json"
)

//客户端
type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

//调用服务器的服务
func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	client.conn <- string(b) //发消息
	str := <-client.conn     //读取服务端返回的消息
	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

//向服务器发送关闭请求
func (client *IpcClient) close() {
	client.conn <- "close"
}
