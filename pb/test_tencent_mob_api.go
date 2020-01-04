package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

/***
移动ssp_api测试
**/
func StartTencentMobApi() {
	var url string = "http://mi.gdt.qq.com/gdt_mview.fcg?"
	var param string = "adposcount=1&count=1&posid=8070816290074563&posw=300&posh=250&charset=utf8&datafmt=html&ext={%22req%22:{%22apiver%22:%221.4%22,%22appid%22:%221105395473%22,%22c_os%22:%22android%22,%22muidtype%22:1,%22muid%22:%22a4ffbb171718e578a1a42969da59d010%22,%22c_device%22:%22unknown%22,%22c_pkgname%22:%22android.testA.test%22,%22posttype%22:1,%22conn%22:1}}"

	response, err := http.Get(url + param)
	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	log.Println("1.status code = ", response.StatusCode)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	log.Println(string(data))

}
