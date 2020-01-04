package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func Start_Marshal() {
	var param *Param = &Param{
		Action:  "A",
		Id:      "g0001",
		Content: `{"Type":1,"Weight":8,"SalesMap":{"a":2,"b":3}}`,
	}
	var s *Strategy = &Strategy{
		Id:       "sid0001",
		SlotType: 2,
		Width:    600,
		height:   800,
	}

	var mode = &SaleMode{
		StrategyId: "sid0001",
		Type:       1,
		Weight:     8,
		SalesMap: map[string]int{
			"a": 2,
			"b": 3,
		},
	}
	var baidu *Api = &Api{
		Code:   "1",
		AppId:  "12345678",
		SlotId: "Bs8s0sa7",
	}
	var tencent *Api = &Api{
		Code:   "2",
		AppId:  "12345678",
		SlotId: "Ts8s0sa7",
	}
	var apis []*Api = []*Api{baidu, tencent}
	var slot *AdSlot = &AdSlot{
		SlotId:   "s8s0sa7",
		SlotType: 9,
		Ads:      5,
		Apis:     apis,
	}
	var adGroup *AdGroup = &AdGroup{
		Gid:                "g0001",
		Status:             2,
		StartTime:          1424657784,
		EndTime:            1434657784,
		AdType:             9,
		ChargeMode:         1,
		ChargePrice:        4.4,
		ActiveTrackingType: 1,
		ActiveTrackingUrl:  "http://talkingdata.com/t?idfs=xxxx",
		LimitTotalImp:      40000,
		LimitDayImp:        4000,
		TargetDevice:       0,
	}
	img := []string{"baidu.com", "jd.com"}
	var adCreative *AdCreative = &AdCreative{
		Cid:         "c000001",
		Title:       []byte("ab"),
		IconSrc:     "http://adstatic.ksyun.com/landingpage/20160825/resources/content/tem/logo.png",
		ImageSrc:    img,
		HtmlSnippet: "<html lang=\"zh-cn\"></html>",
	}
	var m map[string]int = map[string]int{
		"01": 2,
		"02": 9,
		"03": 8,
	}
	var ss []*Strategy = []*Strategy{s}
	var slots []*AdSlot = []*AdSlot{slot}
	var groups []*AdGroup = []*AdGroup{adGroup}
	var modes []*SaleMode = []*SaleMode{mode}
	var creatives []*AdCreative = []*AdCreative{adCreative}
	binaryParam, err := json.Marshal(param) //rpc数据格式
	binaryStrategy, err := json.Marshal(ss) //流量策略
	binarySlot, err := json.Marshal(slots)  //广告位
	binaryAdGroup, err := json.Marshal(groups)
	binaryAdCreative, err := json.Marshal(creatives)
	binarymap, err := json.Marshal(m)
	binaryModes, err := json.Marshal(modes)
	if err != nil {
		log.Println("序列化异常:", err)
		return
	}
	log.Println("1.rpc数据格式:", string(binaryParam))
	log.Println("1.策略:", string(binaryStrategy))
	log.Println("2.结算方式-api:", string(binaryModes))
	log.Println("3.广告位:", string(binarySlot))
	log.Println("4.广告:", string(binaryAdGroup))
	log.Println("5.创意:", string(binaryAdCreative))
	log.Println("6.map:", string(binarymap))
}

func Start_Unmarshal() {
	var apiMode SaleMode
	var param Param
	var groups []*AdGroup
	var creative AdCreative
	var apiModeStr string = `{"Type":1,"Weight":8,"SalesMap":{"a":2,"b":3}}`
	var paramStr string = `{"Action":"A","Id":"g0001","Content":"{\"Type\":1,\"Weight\":8,\"SalesMap\":{\"a\":2,\"b\":3}}"}`
	var groupStr string = `[{"Gid":"g0001","Status":2,"StartTime":1424657784,"EndTime":1434657784,"AdType":9,"ChargeMode":1,"ChargePrice":4.4,"ActiveTrackingType":1,"ActiveTrackingUrl":"http://talkingdata.com/t?idfs=xxxx","LimitTotalImp":40000,"LimitDayImp":4000,"TargetDevice":0,"TargetOS":0,"TargetTimeInterval":"","TargetCity":null,"AppPackage":"","AppSize":0,"ClickUrl":""}]`
	var creativeStr string = `{"Cid":"c000001","Title":"YWI=","Description":null,"IconSrc":"http://adstatic.ksyun.com/landingpage/20160825/resources/content/tem/logo.png","ImageSrc":["http://adstatic.ksyun.com/res//201610/a23fe4a0-e805-4fad-bccb-9a3b211c66c2.gif"],"VideoUrl":"","VideoDuration":0,"Width":0,"Height":0,"HtmlSnippet":"\u003chtml lang=\"zh-cn\"\u003e\u003c/html\u003e"}`
	json.Unmarshal([]byte(apiModeStr), &apiMode)
	json.Unmarshal([]byte(paramStr), &param)
	json.Unmarshal([]byte(groupStr), &groups)
	json.Unmarshal([]byte(creativeStr), &creative)
	log.Printf("1.售卖方式:type=%d,weight=%d,map[a]=%d", apiMode.Type, apiMode.Weight, apiMode.SalesMap["a"])
	log.Printf("2.rpc数据格式:action=%s,id=%s,content=%s", param.Action, param.Id, param.Content)
	log.Print("3.广告:")
	for _, g := range groups {
		log.Printf("\t gid=%s,status=%d", g.Gid, g.Status)
	}
	log.Printf("4.创意:cid=%s,icon=%s,ImageSrc=%s,VidelUrl=%s", creative.Cid, creative.IconSrc, creative.ImageSrc, creative.VideoUrl)
}

//rpc数据格式
type Param struct {
	Action  string
	Id      string
	Content string
}

//售卖方式
type SaleMode struct {
	StrategyId string         //所属的流量策略ID
	Type       uint32         //类型. 1:api; 2:直投;
	Weight     int            //权重大的, 能获得更大的分发机会.
	SalesMap   map[string]int //售卖目标列表: 如果类型为api则(渠道编码:权重);如果类型为直投,则(CID:权重)
}

//为*Person添加String()方法，便于输出
func (this *SaleMode) String() string {
	return fmt.Sprintf("(%s)", this.SalesMap)
}

//可排序的"售卖方式"切片
type SaleModeList []*SaleMode

//实现 sort.Interface接口
func (list SaleModeList) Len() int {
	return len(list)
}

//实现 sort.Interface接口. 排序规则：按Priority排序（由小到大），
func (list SaleModeList) Less(i, j int) bool {
	if list[i].Weight <= list[j].Weight {
		return true
	} else {
		return false
	}
}

//实现 sort.Interface接口
func (list SaleModeList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

//策略信息
type Strategy struct {
	Id        string       //策略ID
	Status    int          // 状态: 0-生效, 1-关闭
	SlotType  uint32       //广告位类型
	Width     uint32       //广告尺寸:宽
	height    uint32       //广告尺寸:宽
	SaleModes SaleModeList //可排序的售卖方式集合
}

//配置的渠道信息
type Api struct {
	Code   string //渠道编码
	AppId  string //渠道分配的appId
	SlotId string //渠道分配的广告位Id
}

//广告位信息
type AdSlot struct {
	SlotId      string    //广告位业务ID
	SlotType    uint32    //广告位类型
	ChargeMode  int       //结算方式
	ChargePrice float64   //结算单价
	Ads         uint32    //返回广告数量,目前支持信息流. 其他广告默认返回一条.
	Apis        []*Api    //广告位绑定的渠道信息
	StrategyId  string    //绑定的策略ID
	Strategy    *Strategy //绑定的流量策略
}

//推广组
type AdGroup struct {
	Gid                string            //推广组ID
	Status             int               //广告投放状态. 1:未投放 2:投放中 3:投放完成 4:投放过期
	StartTime          int64             //投放开始时间(秒)
	EndTime            int64             //投放结束时间(秒)
	AdType             uint32            //广告类型
	ChargeMode         uint32            //结算方式
	ChargePrice        float64           //结算单价
	ActiveTrackingType uint32            //激活检测平台.
	ActiveTrackingUrl  string            //激活检测连接地址
	LimitTotalImp      int               //总曝光
	LimitDayImp        int               //日曝光
	TargetDevice       int               //定向: 设备类型
	TargetOS           int               //定向: 操作系统
	TargetTimeInterval string            //定向: 时段
	TargetCity         map[string]string //定向: 地域
	AppPackage         string            //推广应用包名
	AppSize            uint32            //推广应用大小
	ClickUrl           string            //点击地址
}

//广告创意
type AdCreative struct {
	Cid           string   //创意ID
	Title         []byte   //推广标题，中文需要UTF-8编码
	Description   []byte   //广告描述，中文需要UTF-8编码
	IconSrc       string   //广告图标地址
	ImageSrc      []string //广告图片地址
	VideoUrl      string   //广告视频物料地址
	VideoDuration uint32   //广告视频物料时长
	Width         uint32   //物料的宽度
	Height        uint32   //物料高度
	HtmlSnippet   string   //激励视频广告时,存放落地页面模板H5代码.
}
