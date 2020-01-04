package main

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"ksyun.com/dsp-gateway/channel/ksyun"
	"log"
	"net/http"
)

func main() {
	request := generateAndroidRequest()
	//序列化
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	//log.Println("1. 序列化后:", request.String())
	//log.Println("1. 序列化后:", data)
	//发送
	//	response, err := http.Post("http://debug.mobads.baidu.com/api_5", "application/octet-stream", bytes.NewBuffer(data))
	//response, err := http.Post("http://123.59.14.199:8889/api/def", "application/octet-stream", bytes.NewBuffer(data))
	response, err := http.Post("http://192.168.153.12:8080/api/def", "application/octet-stream", bytes.NewBuffer(data))
	//response, err := http.Post("http://localhost:8080/api/def", "application/octet-stream", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	log.Println("1.status code = ", response.StatusCode)
	data, err = ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	res := &ksyun.BidResponse{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	readResponse(res)
}

func readResponse(res *ksyun.BidResponse) {
	log.Printf("1.requestid=%s \n", res.GetRequestId())
	log.Printf("2.error_code=%d \n", res.GetErrorCode())
	log.Printf("3.request_process_time=%f \n", res.GetRequestTimeS())
	log.Printf("4. ads=%d\n", len(res.Ads))
	for i, ad := range res.Ads {
		log.Printf("\t %d slotId=%s \n", i, ad.GetAdslotId())
		log.Printf("\t %d maxCpm=%d \n", i, ad.GetMaxCpm())
		log.Printf("\t %d html=%s \n", i, ad.GetHtmlSnippet())
		log.Printf("\t %d img_src=%s \n", i, ad.GetMetaGroup()[0].ImageSrc)
	}
}

func generateAndroidRequest() *ksyun.BidRequest {
	slot := &ksyun.AdSlot{
		AdslotId:   proto.String("slot001"),
		AdslotType: proto.Uint32(8),
		MinimumCpm: proto.Int32(100),
	}
	app := &ksyun.App{
		AppId:      proto.String("app01"),
		ChannelId:  proto.String("channel01"),
		AppPackage: proto.String("com.ksyun.kuaishou.test"),
	}
	deviceType := ksyun.Device_PHONE
	osType := ksyun.Device_ANDROID
	dev := &ksyun.Device{
		DeviceType: &deviceType,
		OsType:     &osType,
		Vendor:     []byte("mi ui"),
		Model:      []byte("note 2"),
	}
	conType := ksyun.Network_CELL_4G
	netWork := &ksyun.Network{
		Ipv4:           proto.String("114.255.44.132"),
		ConnectionType: &conType,
	}
	rtb := &ksyun.Rtb{
		MediaKeyword: [][]byte{[]byte("旅游爱好者")},
	}
	request := &ksyun.BidRequest{
		RequestId: proto.String("test-000000000000001"),
		Adslot:    slot,
		App:       app,
		Device:    dev,
		Network:   netWork,
		Rtb:       rtb,
	}
	return request
}
