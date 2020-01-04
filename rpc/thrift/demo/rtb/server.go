package main

import (
	"encoding/json"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/entry/thrift/rtb"
	"ksyun.com/commons/entry/thrift/syncer"
	"log"
)

func main() {
	var hostPost string = "172.31.16.14:6666"
	//var hostPost string = "192.168.153.27:6666"
	serverTransport, err := thrift.NewTServerSocket(hostPost)
	if err != nil {
		log.Fatalln(err)
	}
	processors := thrift.NewTMultiplexedProcessor()
	rtb := rtb.NewRtbServiceProcessor(&RtbService{})
	sync := syncer.NewAdOperationServerProcessor(&AdOperationService{})
	processors.RegisterProcessor("RtbService", rtb)
	processors.RegisterProcessor("AdOperationService", sync)
	server := thrift.NewTSimpleServer2(processors, serverTransport)
	fmt.Println("thrift server in", hostPost)
	err = server.Serve()
	if err != nil {
		log.Fatalf("服务启动失败！%s", err.Error())
	}
}

type RtbService struct {
}

//移动端广告请求接口
func (this *RtbService) GetMobileAd(request *rtb.MReq) (*rtb.MRes, error) {
	data, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("请求格式异常！%s \n", err.Error())
		return nil, nil
	}
	fmt.Printf("收到请求：%s \n", string(data))
	r := &rtb.MRes{}
	r.ErrorCode = proto.Int64(0)
	r.RequestID = request.RequestID
	var creativeType int64 = 2
	var interactionType int64 = 2
	var htmlSnippet string = "<html></html>"
	var title string = "王者荣耀"
	var desc string = "王者荣耀》是腾讯第一5V5英雄公平对战手游,拥有超过2亿注册用户,每天有5000万玩家在王者峡谷中开团战斗!"
	var imgSrc string = "http://shp.qpic.cn/ishow/2735032715/1490598138_1644740874_15701_sProdImgNo_2.jpg/0"
	var appPackage string = "com.tencent.w.test"
	var custId string = "cust_001"
	var gid string = "12345678"
	var crid string = "987654321"
	var landingPage string = "http://pvp.qq.com"
	var clickTrack string = "http://track.ksyun.com/click"
	ad := &rtb.Ad{
		SlotID:          request.Slots[0].SlotID,
		MaxCpm:          100,
		CreativeType:    &creativeType,
		InteractionType: &interactionType,
		HTMLSnippet:     &htmlSnippet,
		Title:           &title,
		Description:     &desc,
		ImageSrc:        &imgSrc,
		AppPackage:      &appPackage,
		CustID:          &custId,
		Gid:             &gid,
		Crid:            &crid,
		LandingPage:     &landingPage,
		ThirdClickTrack: &clickTrack,
	}
	r.Ads = []*rtb.Ad{ad}
	data, err = json.Marshal(r)
	if err != nil {
		fmt.Printf("相应格式异常！%s \n", err.Error())
		return nil, nil
	}
	fmt.Printf("响应：%s \n", string(data))
	return r, nil
}

type AdOperationService struct {
}

//客户级广告上线,包括下属所有可投放的广告
func (this *AdOperationService) OnlineCust(c *syncer.Cust, tp syncer.OnlineType) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//下线客户及下属所有广告
func (this *AdOperationService) OfflineCust(custId string) (msg *syncer.Message, err error) {
	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//下线计划及所有广告
func (this *AdOperationService) OfflineCampaign(camId string) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//下线单个广告
func (this *AdOperationService) OfflineAdGroup(gid string) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//下线单个创意
func (this *AdOperationService) OfflineCreative(crid string) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//广告主信息跟新
func (this *AdOperationService) UpdateCust(cust *syncer.Cust) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//投放计划信息跟新.
func (this *AdOperationService) UpdateCampaign(campaign *syncer.Campaign) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//广告信息跟新.
func (this *AdOperationService) UpdateAdGroup(group *syncer.AdGroup) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}

//创意信息跟新
func (this *AdOperationService) UpdateCreative(creative *syncer.Creative) (msg *syncer.Message, err error) {

	return &syncer.Message{Success: proto.Bool(true)}, nil
}
