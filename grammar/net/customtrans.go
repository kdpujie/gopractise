package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

/**
在默认的http.Transport之上包一层Transport并实现RoundTrip()方法.
**/

type OutCustomTransport struct {
	Transport http.RoundTripper //组合http.RoundTripper
}

func (t *OutCustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

//传输层(我的理解, 功能类似拦截器)
func (t *OutCustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "Our Custom User-Agent")
	fmt.Println("增加User-Agent")
	return t.transport().RoundTrip(req)
}

//构造client
func (t *OutCustomTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

//自定义Transport
func Start_our_Transport() {
	t := &OutCustomTransport{}
	c := t.Client()
	resp, err := c.Get("http://www.baidu.com")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)
	io.Copy(os.Stdout, resp.Body) //将网页内容打印到标准输出流中
}

//自定义客户端
func Start_Our_Client() {
	var DefaultTransport http.RoundTripper = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	c := &http.Client{Transport: DefaultTransport}
	resp, err := c.Get("http://www.baidu.com")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)
	io.Copy(os.Stdout, resp.Body) //将网页内容打印到标准输出流中
}
