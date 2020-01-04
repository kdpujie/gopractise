package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/entry/thrift/syncer"
	"log"
	"net"
)

func main() {

	// 120.92.44.245 / 192.168.115.109
	transport, err := thrift.NewTSocket(net.JoinHostPort("120.92.44.245", "6666"))
	//transport, err := thrift.NewTSocket("192.168.153.27:6677")
	if err != nil {
		log.Fatalln(err)
	}
	pro := thrift.NewTMultiplexedProtocol(thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(transport), "AdOperationServer")
	if err := transport.Open(); err != nil {
		log.Fatal(err)
	}
	defer transport.Close()
	client := syncer.NewAdOperationServerClientProtocol(transport, pro, pro)
	/*msg , err :=client.OnlineCust(genCust(),syncer.OnlineType_CREATIVE)
	if err != nil {
		fmt.Printf("client.onlineCust failed.%v \n",err)
	}else {
		fmt.Printf("client.onlineCust msg=%v \n",msg)
	}
	//下线创意
	msg , err = client.OfflineCreative("crid05")
	if err != nil {
		fmt.Printf("client.onlineCust failed.%v \n",err)
	}else {
		fmt.Printf("client.onlineCust msg=%v \n",msg)
	}*/
	//广告上线
	/*	msg , err :=client.OnlineCust(genCust(),syncer.OnlineType_CREATIVE)
		if err != nil {
			fmt.Printf("client.onlineCust failed.%v \n",err)
		}else {
			fmt.Printf("client.onlineCust msg=%s,b=%v \n",*msg.Msg, *msg.Success)
		}*/
	gids := []string{"01010", "an0zhwdx", "rnlw0ytr", "u3ed1ier", "vlp7jnss"}
	msg, err := client.GetAdStatus(gids)
	for gid, status := range msg.Adstatus {
		fmt.Printf("广告(%s)的投放状态为：%d \n", gid, status)
	}
}
func genCust() *syncer.Cust {
	cid := proto.String("test001")
	cust := &syncer.Cust{
		Cid:       cid,
		DayBudget: 10000,
		Campaigns: []*syncer.Campaign{genCamp(cid)},
	}
	return cust
}

func genCamp(cid *string) *syncer.Campaign {
	campid := proto.String("camp00101")
	camp := &syncer.Campaign{
		Cid:       cid,
		CampId:    campid,
		DayBudget: 2000,
		Groups:    []*syncer.AdGroup{genGroup(campid)},
	}
	return camp
}

func genGroup(campid *string) *syncer.AdGroup {
	gid := proto.String("group0010101")
	group := &syncer.AdGroup{
		CampId:            campid,
		Gid:               gid,
		Status:            proto.Int64(1),
		AdType:            proto.Int64(21),
		BidMode:           proto.Int64(1),
		Price:             1000,
		InteractionType:   proto.Int64(2),
		AppPackage:        proto.String("com.ksyun.test"),
		AppSize:           proto.Int64(400),
		LandingPage:       proto.String("http://test.apk"),
		ThirdPlatformType: proto.Int64(1),
		ThirdPlatformURL:  proto.String("http://track.com"),
		UserTotalImps:     proto.Int64(0),
		UserTotalClicks:   proto.Int64(0),
		UserDayImps:       proto.Int64(0),
		UserDayClicks:     proto.Int64(0),
		AdDayImps:         proto.Int64(0),
		AdDayClicks:       proto.Int64(0),
		Creatives:         []*syncer.Creative{genCreative(gid)},
	}
	return group
}

func genCreative(gid *string) *syncer.Creative {
	c := &syncer.Creative{
		Gid:           gid,
		Crid:          proto.String("crid06"),
		Title:         proto.String("6666"),
		VideoURL:      proto.String("http://video.test.ksyun.com"),
		VideoDuration: proto.Int64(15),
		CreativeType:  proto.Int64(4),
		AdxCris: []*syncer.AdxCriInfo{
			&syncer.AdxCriInfo{
				AdxCode: proto.String("305"),
				AdxCrid: proto.String("adx-crid001"),
			},
		},
	}
	return c
}
