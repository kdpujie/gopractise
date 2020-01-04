package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	str := "http://track.sigmob.cn/track?&c=Cgx0ZXN0XzAwMDAwMDESGHRlc3RfMDAwMDAwMV9zbG90X2lkLTAwMSIBMSoLc2xvdF9pZC0wMDE6EgoBMxICMzgaATEgASjoBzDoB0gBWi0KGmlkZmEtYWRzZmFzZGYtYXNkZmRTQWRmYWRmEg91ZGlkLTA5ODc2NTQzMjFiCWFwcF9pZF8wMQ&e=active&p=_PRICE_"
	fmt.Printf("移位:%d\n", 64>>6)
	l := url.QueryEscape(str)
	fmt.Println(l)
	l = strings.Replace(l, "_PRICE_", "5", -1)
	r, _ := url.QueryUnescape(l)
	fmt.Println(r)

}
