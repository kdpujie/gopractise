package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/golang/protobuf/proto"
	"learn.com/pb/ks"
)

func StartSDKServer() {
	http.HandleFunc("/sdk/test/def", safeHandler(DefaultWebHandler))
	http.HandleFunc("/test/get/ip", safeHandler(SelectIP))
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}

	log.Println("服务启动成功.")
}

//闭包避免程序运行时出错崩溃: 所有handler都经过此方法,统一处理handler抛出的异常panic
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				log.Println("WARN: panic in %v - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Headers", "x-requested-with,content-type")
		fn(w, r)
	}
}

func SelectIP(w http.ResponseWriter, r *http.Request) {
	var body string = "<html><body>"
	for k, v := range r.Header {
		body = body + k + "=" + strings.Join(v, "-") + "<br/>"
	}
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("x-forwarded-for")
	}
	if ip == "" {
		ip = r.Header.Get("Remote_addr")
	}
	if ip == "" {
		ip = r.Header.Get("Proxy-Client-IP")
	}
	if ip == "" {
		ip = r.Header.Get("WL-Proxy-Client-IP")
	}
	if ip == "" {
		ip = r.RemoteAddr
	}
	body = body + "ip=" + ip + "<br/>"
	io.WriteString(w, body+"</body></html>")
	return

}

//ssp web js默认流量接收器
func DefaultWebHandler(w http.ResponseWriter, r *http.Request) {
	response := &ks.MobadsResponse{}

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Unmarshal error:", err)
		response.ErrorCode = proto.Uint64(1001) //读取body失败
	} else {
		//log.Println(data)
		request := &ks.MobadsRequest{}
		err = proto.Unmarshal(data, request)
		if request.GetAdslot().GetAdslotId() == "haj95fzy" && request.GetApp().GetChannelId() == "1" {
			generateResponse(response)
		} else {
			response.ErrorCode = proto.Uint64(200000)
		}
	}

	data, err = proto.Marshal(response)
	if err != nil {
		log.Println("分发失败,response errorCode", response.GetErrorCode())
	} else {
		w.Write(data)
	}
}

func generateResponse(response *ks.MobadsResponse) {
	adStrategys := []*ks.AdStrategy{}
	vungleWeight := ks.AdStrategy_CHANNEL_VUNG
	unityWeight := ks.AdStrategy_CHANNEL_UNIT
	ksWeight := ks.AdStrategy_CHANNEL_KS
	vungle := &ks.AdStrategy{
		AppId:       proto.String(""),
		AdslotId:    proto.String("586f184e2944d5950b000400"),
		Weight:      proto.Uint32(10),
		ChannelType: &vungleWeight,
	}
	unity := &ks.AdStrategy{
		AppId:       proto.String(""),
		AdslotId:    proto.String("1257148"),
		Weight:      proto.Uint32(10),
		ChannelType: &unityWeight,
	}
	ks := &ks.AdStrategy{
		AppId:       proto.String("qxidp2e6"),
		AdslotId:    proto.String("k4lkcwyv"),
		Weight:      proto.Uint32(10),
		ChannelType: &ksWeight,
	}

	adStrategys = append(adStrategys, vungle)
	adStrategys = append(adStrategys, unity)
	adStrategys = append(adStrategys, ks)

	response.ErrorCode = proto.Uint64(0)
	response.AdStrategy = adStrategys
	response.AutomaticPro = proto.Bool(false)
}
