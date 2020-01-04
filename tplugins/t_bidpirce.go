package main

import (
	"plugin"
	"fmt"
	"learn.com/tplugins/entry"
	"ksyun.com/commons/entry/thrift/rtb"
	"github.com/golang/protobuf/proto"
)

func main()  {
	path := "./bidprice/bidprice.so"
	symbol := "BidPrice"
	plug, err :=  plugin.Open(path)
	if err != nil {
		fmt.Printf("Open插件(%s)失败: %v \n",path ,err)
	}
	symPlug, err := plug.Lookup(symbol)
	if err != nil {
	}else {
		fmt.Printf("插件(%s) Lookup symbol(%s) \n",path, symbol)
	}
	bid, ok := symPlug.(entry.BidPrice)
	if !ok {
		fmt.Println("unexpected type from module symbol")
	}
	bid.Init()
	req := generateReq()
	gid := "01014" //[]string{"01014","01012","01013"}
	mfsAd := bid.GetBidPrice("exp_bidprice",0,gid, 500, req)
	fmt.Printf("bidprice获取广告%s的价格为%d \n", gid, mfsAd)
}

func generateReq() *rtb.MReq  {
	slotSize := &rtb.Size{
		Width:600,
		Height:500,
	}
	slot := &rtb.AdSlot{
		AdslotID:proto.String("9030847217791216368"),
		AdslotSize:slotSize,
	}
	req := & rtb.MReq{
		RequestID: proto.String("test000000001"),
		Slots:[]*rtb.AdSlot{slot},
	}
	return req
}