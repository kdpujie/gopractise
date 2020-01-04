package main

import (
	"flag"
	"log"

	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/robfig/config"
	"ksyun.com/commons/entry/thrift/rtb"
	"learn.com/tplugins/entry"
	"plugin"
	"strconv"
	"strings"
	"crypto/md5"
)

var (
	configFile     = flag.String("configFile", "conf/plugins.conf", "General configuration file")
	EXP_SYMBOL     = "plug.symbol"
	EXP_PATH       = "plug.path"
	EXP_FLOW_TYPE  = "flow.split.type"
	EXP_FLOW_RATIO = "flow.split.ratio"
)

func main() {
	plugs := NewPlugins()
	req := generateReq1()
	plugs.loadConfig()
	fmt.Printf("插件初始化完成 \n")
	for name, plug := range plugs.qp {
		fmt.Printf("执行Q类插件，%s \n", name)
		adlists := []string{"01014", "01012", "01013"}
		mfsAd := plug.p.Filter("exp_whitelist", 0, adlists, req)
		fmt.Printf("mfs选中的广告有：%v \n", mfsAd)
	}
	for name, plug := range plugs.bp {
		fmt.Printf("执行B类插件，%s \n", name)
		gid := "01014" //[]string{"01014","01012","01013"}
		mfsAd := plug.p.GetBidPrice("exp_bidprice", 0, gid, 500, req)
		fmt.Printf("bidprice获取广告%s的价格为%d \n", gid, mfsAd)
	}
}

//插件信息
type pluginInfo struct {
	expName      string
	symbol       string
	path         string
	splitType    int    //1=随机数 2=uid
	splitRatio   []int
	plugType     int    //1 query, 2 bidprice
	shuffle		 string //离散因子
	definePlugin func(string, string)
}

func (this *pluginInfo) loadPlugin() (plugin.Symbol, error) {
	plug, err := plugin.Open(this.path)
	if err != nil {
		return nil, err
	}
	return plug.Lookup(this.symbol)
}

//分桶规则，默认100桶
func (this *pluginInfo) splitBucket(value string) int {
	h := md5.New()
	h.Write([]byte(value + this.shuffle))
	hValue, _ := strconv.ParseUint(hex.EncodeToString(h.Sum(nil))[16:32], 16, 64)
	bucketId := hValue % 100
	fmt.Printf("str=%s, shuffle=%s, bucketId=%d ", value, this.shuffle, bucketId)
	return int(bucketId)
}

//按照流量配比，分配流量，返回分配id
//比如：10:10:80, 分成流量0，1，2;
func (this *pluginInfo) flowId(value string) int {
	var boundary int = 0
	bucketId := this.splitBucket(value)
	for i, ratio := range this.splitRatio {
		boundary = boundary + ratio
		if bucketId < boundary {
			return i
		}
	}
	return len(this.splitRatio) - 1
}

//Q类插件
type QPlugin struct {
	*pluginInfo
	p entry.Query
}

//B类插件
type BPlugin struct {
	*pluginInfo
	p entry.BidPrice
}

//插件集合
type Plugins struct {
	qp map[string]*QPlugin
	bp map[string]*BPlugin
}

func NewPlugins() *Plugins {
	qp := make(map[string]*QPlugin, 4)
	bp := make(map[string]*BPlugin, 4)
	return &Plugins{qp: qp, bp: bp}
}

//
func (this *Plugins) loadPlugin(name string, plugInfo *pluginInfo) {
	s, err := plugInfo.loadPlugin()
	if err != nil {
		fmt.Printf("loadPlugin : %s \n", err.Error())
	}
	switch t := s.(type) {
	case entry.Query:
		fmt.Printf("加载的插件类型为query：%v \n", t)
		plugInfo.plugType = 1
		t.Init()
		this.qp[name] = &QPlugin{p: t, pluginInfo: plugInfo}
	case entry.BidPrice:
		fmt.Printf("加载的插件类型为price：%v \n", t)
		plugInfo.plugType = 2
		t.Init()
		this.bp[name] = &BPlugin{p: t, pluginInfo: plugInfo}
	default:
		fmt.Printf("不支持的插件类型：%v \n", t)
	}
}

func (this *Plugins) loadConfig() *Plugins {
	flag.Parse()
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Failed to find", *configFile, err)
	}
	plugs := NewPlugins()
	for i, section := range cfg.Sections() {
		if section == config.DEFAULT_SECTION {
			continue
		}
		plugInfo := &pluginInfo{expName: section}
		var err error = nil
		if plugInfo.symbol, err = cfg.String(section, EXP_SYMBOL); err != nil {
			continue
		}
		if plugInfo.path, err = cfg.String(section, EXP_PATH); err != nil {
			continue
		}
		if plugInfo.splitType, err = cfg.Int(section, EXP_FLOW_TYPE); err != nil {
			continue
		}
		if splitRatio, err := cfg.String(section, EXP_FLOW_RATIO); err == nil {
			ratios := strings.Split(splitRatio, ":")
			var splitRatio []int = make([]int, 0, len(ratios))
			flag := true
			for i := 0; i < len(ratios); i++ {
				b, err := strconv.ParseInt(ratios[i], 10, 64)
				if err != nil {
					flag = false
					break
				}
				splitRatio = append(splitRatio, int(b))
			}
			if flag {
				plugInfo.splitRatio = splitRatio
			}
		}
		plugInfo.shuffle = strconv.FormatInt(int64(i), 10)//初始化离散因子
		this.loadPlugin(section, plugInfo)
	}
	return plugs
}

func generateReq1() *rtb.MReq {
	slotSize := &rtb.Size{
		Width:  600,
		Height: 500,
	}
	slot := &rtb.AdSlot{
		AdslotID:   proto.String("9030847217791216368"),
		AdslotSize: slotSize,
	}
	req := &rtb.MReq{
		RequestID: proto.String("test000000001"),
		Slots:     []*rtb.AdSlot{slot},
	}
	return req
}
