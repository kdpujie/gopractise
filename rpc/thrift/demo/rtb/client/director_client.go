package main

import (
	"fmt"
	thrift "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/entry/thrift/rtb"
	"log"
)

func main() {
	//transport, err := thrift.NewTSocket("120.92.44.245:6666")
	transport, err := thrift.NewTSocket("192.168.115.96:6666")
	if err != nil {
		log.Fatalln(err)
	}
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	defer transport.Close()
	pro := thrift.NewTMultiplexedProtocol(thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(transport), "RtbService")
	rc := rtb.NewRtbServiceClientProtocol(transport, pro, pro)

	mRes, err := rc.GetMobileAd(generateReq())
	if err != nil {
		//p.ReturnTransport(transport,err)
		fmt.Printf("getMobileAd() failed ,error=%s \n", err.Error())
	} else {
		readRes(mRes)
	}
}

func generateReq() *rtb.MReq {
	request := &rtb.MReq{}
	request.RequestID = proto.String("00000001")
	industryBidflorr := &rtb.IndustryBidfloor{
		IndustryID: proto.Int64(4000),
		Bidfloor:   proto.Int32(900),
	}
	app := &rtb.App{
		AppID:     proto.String("test_app_dd"),
		ChannelID: proto.String("304"),
	}
	slot := &rtb.AdSlot{
		AdslotID:      proto.String("8770119859747914936"),
		AdslotType:    []int32{21},
		CreativeTypes: []int32{2, 4},
		AdslotSize: &rtb.Size{
			Width:  640,
			Height: 360,
		},
		Bidfloor:         1000,
		IndustryBidfloor: []*rtb.IndustryBidfloor{industryBidflorr},
		Video: &rtb.Video{
			Title:       proto.String("花千骨10集"),
			Tags:        []string{"仙侠", "连续剧"},
			Minduration: proto.Int64(0),
			Maxduration: proto.Int64(30),
		},
	}
	var connectionType int32 = 2
	ipv4 := "36.6.192.2" //"114.255.44.132"
	network := &rtb.Network{
		ConnectionType: &connectionType,
		Ipv4:           &ipv4,
	}
	var dType int64 = 1
	var osType int64 = 1
	imei := "866192029439468"
	imeiMd5 := "ae9d93d5539ca7d9e1bd3687736ea8f1"
	device := &rtb.Device{
		DeviceType: &dType,
		OsType:     &osType,
		Udid: &rtb.UdId{
			CurrentID: &imei,
			ImeiMd5:   &imeiMd5,
		},
	}
	request.Slots = []*rtb.AdSlot{slot}
	request.App = app
	request.Network = network
	request.Device = device
	return request
}

func readRes(mRes *rtb.MRes) {
	fmt.Printf("返回广告数 %d \n", len(mRes.Ads))
	if len(mRes.Ads) > 0 {
		material := mRes.Ads[0].Materials[0]
		fmt.Printf("ad.cid=%s \n", *mRes.Ads[0].AdKey)
		fmt.Printf("ad.price=%d \n", mRes.Ads[0].MaxCpm)
		fmt.Printf("ad.m.imageSrc=%v \n", material.ImageSrc)
		fmt.Printf("ad.m.creativeType=%d \n", *material.CreativeType)
		fmt.Printf("ad.m.InteractionType=%d \n", *material.InteractionType)
		fmt.Printf("ad.m.LandingPage=%s \n", *material.LandingPage)
		fmt.Printf("ad.m.title=%s \n", *material.Title)
		fmt.Printf("ad.m.desc=%v \n", material.Description)
		fmt.Printf("ad.m.iconSrc=%v \n", material.IconSrc)
		fmt.Printf("ad.m.CustID=%v \n", *material.CustID)
		fmt.Printf("ad.m.CampID=%v \n", *material.Cid)
		fmt.Printf("ad.m.GID=%v \n", *material.Gid)
		fmt.Printf("ad.m.CRID=%v \n", *material.Crid)
		fmt.Printf("ad.m.adx_crid=%v \n", *material.AdxCrid)
		fmt.Printf("ad.m.video_duration=%d \n", *material.VideoDuration)
	}
}
