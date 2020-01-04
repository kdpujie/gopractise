package main

import (
	"os"
	//"fmt"
	"io"
	"net/http"
)

func start_http_get() {
	resp, err := http.Get("http://www.baidu.com") //当你得到一个重定向的错误时，两个变量都将是non-nil。这意味着你最后依然会内存泄露。
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)
	io.Copy(os.Stdout, resp.Body) //将网页内容打印到标准输出流中
}
