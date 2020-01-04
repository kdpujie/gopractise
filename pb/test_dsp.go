package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"ksyun.com/dsp-gateway/adaptor/adx_zplay/api/zadx"
	"ksyun.com/mssp-tester/util"
)

func main() {
	defer func() {
		if e, ok := recover().(error); ok {
			log.Println("WARN: panic in %v - %v", e)
			log.Println(string(debug.Stack()))
		}
	}()
	for {
		go test()
		time.Sleep(100 * time.Millisecond)
	}

}
func test() {
	req := `{"id":"0cthUk1DdMAi0UR0ya1VtvNr40oCL2","imp":[{"id":"1","banner":{"w":640,"h":100,"pos":0},"instl":false,"tagid":"slotid","bidfloor":1,"bidfloorcur":"CNY"}],"DistributionchannelOneof":{"App":{"id":"sandbox_app_id_1167","name":"沙盒APP","ver":"0.0.1","bundle":"sandbox_app"}},"device":{"dnt":true,"ua":"Go-http-client/1.1","ip":"10.252.35.192","didsha1":"","dpidsha1":"a7e4a9f759426b2ee5baf434a0b824833a7d37b8","make":"","model":"iPhone6","os":"ios","osv":"9.2.1","w":320,"h":568,"ppi":2,"pxratio":2,"connectiontype":2,"devicetype":4,"macsha1":""}}`

	req_zadx := zadx.BidRequest{}
	req = strings.Replace(req, " ", "", -1)
	req = strings.Replace(req, "\n", "", -1)
	//log.Println(req)
	if false {
		err := json.Unmarshal([]byte(req), &req_zadx)
		if err != nil {
			log.Println("unmarshal err:", err)
			//return
		}
	}
	if true {
		req_zadx.Id = proto.String("0cthUk1DdMAi0UR0ya1VtvNr40oCL2")
		req_zadx.Imp = make([]*zadx.BidRequest_Imp, 1)
		req_zadx.Imp[0] = &zadx.BidRequest_Imp{}
		req_zadx.Imp[0].Id = proto.String("1")
		req_zadx.Imp[0].Bidfloor = proto.Float64(1)
		req_zadx.Imp[0].Bidfloorcur = proto.String("CNY")
		req_zadx.Device = &zadx.BidRequest_Device{}
		req_zadx.Device.Ip = proto.String("0.0.0.0")
	}
	req_http, err := proto.Marshal(&req_zadx)
	if err != nil {
		log.Println("Marshal err:", err)
		return
	}
	body := bytes.NewBuffer(req_http)

	//	res, err := http.Post("http://localhost:8090/api/zplay", "application/json;charset=utf-8", body)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}

	channelR, err := http.NewRequest("POST", "http://localhost:8090/api/zplay", body)
	if err != nil {
		log.Println("NewRequest err:", err)
		return
	}
	client := &http.Client{}
	res, err := client.Do(channelR)
	if nil != err {
		log.Println("client err:", err)
		return
	}

	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ReadAll err:", err)
		return
	}
	res_http := zadx.BidResponse{}
	proto.Unmarshal(result, &res_http)
	log.Println("发送请求:", util.ToJson(req_zadx))
	log.Println("收到反馈:", util.ToJson(res_http))
}
