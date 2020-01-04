package main

import (
	"fmt"
	thrift "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/entry/thrift/rtb"
	"ksyun.com/commons/entry/thrift/syncer"
	pt "ksyun.com/commons/tool/thrift"
	"ksyun.com/commons/util"
	"time"
)

func main() {
	var zkHost []string = []string{"10.69.56.55:2181"}
	var rtbProviderPath = "/thrift-rpc/RtbService/providers"
	//var hostPort string = "120.92.44.245:6666"
	/*	var hostPort string = "192.168.115.169:6666"
		transport, err := thrift.NewTSocket(hostPort)
		if err != nil {
			log.Fatalln(err)
		}*/
	var p *pt.Pool
	registry := pt.NewZkRegistry(zkHost, 1*time.Second, rtbProviderPath)
	p = pt.NewPool(1, 2, 65*time.Second, registry)
	fmt.Println("thrift pool初始化完成。")

	go func() {
		for {
			p.Status()
			time.Sleep(2000 * time.Millisecond)
		}

	}()
	time.Sleep(200 * time.Second)
	for i := 1; i <= 1; i++ {
		//fmt.Printf("=====testRtbService 开始第%d次调用..... \n",i)
		testRtbService(p)
		//fmt.Printf("=====testRtbService 第%d次调用完成 \n",i)
		//time.Sleep(500 * time.Millisecond)
	}
}

func testOperationService(transport thrift.TTransport) {
	syncPro := thrift.NewTMultiplexedProtocol(thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(transport), "AdOperationService")
	sc := syncer.NewAdOperationServerClientProtocol(transport, syncPro, syncPro)
	//client := rtb.NewRtbServiceClientFactory(,)

	//////
	m, err := sc.OfflineCampaign("001")
	if err != nil {
		fmt.Println("offlineCampaign() failed!")
		return
	}
	fmt.Printf("offlineCampaign() success result=%v \n", m.Success)
}

func testRtbService(p *pt.Pool) {
	transport := p.Get()
	//transport, err := thrift.NewTSocket(net.JoinHostPort("192.168.115.169", "6666"))
	//transport, err := thrift.NewTSocket("192.168.153.27:6677")
	/*	if err != nil {
			log.Fatalln(err)
		}
		if err := transport.Open(); err != nil {
			log.Fatal(err)
		}*/
	//defer transport.Close()

	pro := thrift.NewTMultiplexedProtocol(thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(transport), "RtbService")
	rc := rtb.NewRtbServiceClientProtocol(transport, pro, pro)

	request := &rtb.MReq{}
	request.RequestID = proto.String("00000001")
	app := &rtb.App{
		AppID:     proto.String("test_app_dd"),
		ChannelID: proto.String("302"),
	}

	slot := &rtb.AdSlot{
		AdslotID:      proto.String("4007443935195074342"),
		AdslotType:    []int32{2, 21, 22, 23},
		CreativeTypes: []int32{2},
		AdslotSize: &rtb.Size{
			Width:  600,
			Height: 500,
		},
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
	imei := "866361021509590"
	imeiMd5 := util.Md5(imei)
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
	mRes, err := rc.GetMobileAd(request)
	defer p.ReturnTransport(transport, err)
	if err != nil {
		//p.ReturnTransport(transport,err)
		fmt.Printf("getMobileAd() failed ,error=%s \n", err.Error())
	} else {
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
}
