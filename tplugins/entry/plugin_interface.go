package entry

import "ksyun.com/commons/entry/thrift/rtb"

//demo
type Greeter interface {
	Greet(user *User)
}

//AR阶段 筛选广告
type Query interface {
	Init()
	Filter(expId string, flowId int, adlist []string, req *rtb.MReq) []string
}

//ranking阶段，获取竞价价格
type BidPrice interface {
	Init()
	GetBidPrice(expId string, flowId int, gid string, advertiserPrice int64, req *rtb.MReq) int64
}

type User struct {
	Name string
	Age int
	Company string
}