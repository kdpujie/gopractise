package main

import (
	"time"
	ct "ksyun.com/commons/tool"
	"ksyun.com/commons/util"
	"io"
	"log"
	"net/http"
	"fmt"
	"runtime/pprof"
	_ "net/http/pprof"
)

//简单的HTTP服务端

func helloHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
	defer r.Body.Close()
	ip := r.FormValue("ip")
	ipNumber ,err:= util.Ip2Long(ip)
	if err == nil {
		io.WriteString(w,fmt.Sprintf("ip=%s,city_code=%d \n",ip, ct.DefaultIpLib.SearchIp(ipNumber).CityCode))
	}else {
		io.WriteString(w,fmt.Sprintf("ip=%s,查询失败, msg=%s \n",ip, err.Error()))
	}
}

func goroutineHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func heapHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ip := "192.168.3.2"
	ipNumber ,err:= util.Ip2Long(ip)
	c := ct.DefaultIpLib.SearchIp(ipNumber).CityCode
	if err == nil {
		//io.WriteString(w,fmt.Sprintf("ip=%s,city_code=%d \n",ip, c))
		fmt.Printf("ip=%s,city_code=%d \n",ip, c)
	}else {
		//io.WriteString(w,fmt.Sprintf("ip=%s,查询失败, msg=%s \n",ip, err.Error()))
		fmt.Printf("ip=%s,查询失败, msg=%s \n",ip, err.Error())
	}
	p := pprof.Lookup("heap")
	p.WriteTo(w, 1)
}

func main() {
	server := &http.Server{ //自定义server
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/sip", helloHandler)
	http.HandleFunc("/pprof/goroutine", goroutineHandler)
	http.HandleFunc("/pprof/heap", heapHandler)
	//http.Serve(":8080",nil)  默认server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}

}
