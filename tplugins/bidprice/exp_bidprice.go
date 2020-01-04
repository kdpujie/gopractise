package main

import "C"
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
	"ksyun.com/commons/entry/thrift/rtb"
	"sync"
	"time"
)

const (
	PCTR_1_IMP = "pctr1_imp"
	PCTR_1_CLICK = "pctr1_click"
	PCTR_2_IMP = "pctr2_imp"
	PCTR_2_CLICK = "pctr2_click"
)

type BidPriceStrategy struct {
	rw sync.RWMutex
	configPath	string     //白名单所在路径
	online		bool   //实验开关
	bidbase 	float64
	gids		map[string]int
	pctr 		map[string]float64
	avgPctr1 	float64
	avgPctr2 	float64
}

// exported(symbol)
var BidPrice BidPriceStrategy

//系统启动或者重新加载配置时，由框架调用
func (this *BidPriceStrategy) Init()  {
	this.configPath = "conf/plugs/bidprice_strategy.conf"
	this.pctr = make(map[string]float64, 32)
	this.gids = make(map[string]int, 32)
	this.readFile(this.configPath) //初始化配置
/*	fmt.Println("竞价策略加载的广告有：")
	for gid, _ := range this.gids {
		fmt.Printf("\tgid=%s \n",gid)
	}
	fmt.Println("竞价策略加载的pos有：")
	for pos, ctr := range this.pctr {
		fmt.Printf("\tkey=%s, ctr=%f \n",pos, ctr)
	}*/
}
/**
获取竞价价格
Bid=bidbase*（ctrpre1/ctravg1）*（ctrpre2/ctravg2）
expId: 实验id
flowId：流量ID
gid: 广告ID
advertiserPrice：广告主设置价格
req：请求
**/
func (this *BidPriceStrategy) GetBidPrice(expId string, flowId int, gid string, advertiserPrice int64, req *rtb.MReq) int64 {
	size := req.GetSlots()[0].AdslotSize
	sizeStr := fmt.Sprintf("%dx%d",size.Width,size.Height)
	fmt.Printf("bidprice插件：expId=%s, flowId=%d gid=%s, advertiserPrice=%d,size=%s \n",expId, flowId, gid, advertiserPrice, sizeStr)
	if this.online && flowId ==0 && this.IsContainGid(gid) {
		if sProb,ok:= this.pctr[sizeStr];ok{
			if hProb, ok := this.pctr[strconv.Itoa(time.Now().Hour())];ok {
				bidPrice := this.bidbase * (sProb / this.avgPctr1) * (hProb / this.avgPctr2)
				return int64(math.Ceil(bidPrice))
			}
		}
		return int64(math.Ceil(this.bidbase))
	}
	return advertiserPrice
}

//白名单中是否已存在该gid
func (this *BidPriceStrategy) IsContainGid(gid string) bool {
	_,ok := this.gids[gid]
	return ok
}
func (this *BidPriceStrategy) clearCache()  {
	for k, _ :=range this.pctr {
		delete(this.pctr,k)
	}
	for k, _ :=range this.gids {
		delete(this.gids,k)
	}
}
//打开实验
func (this *BidPriceStrategy) on()  {
	this.rw.RLock()
	this.online = true
	this.rw.RUnlock()
}

//关闭实验
func (this *BidPriceStrategy) off()  {
	this.rw.RLock()
	this.online = false
	this.rw.RUnlock()
}
//读取并解析白名单文件
func (this *BidPriceStrategy) readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	this.off() //关闭实验，更新配置
	this.clearCache()
	pctrFlag := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "pctr1=") { //标识行
			pctrFlag = "pctr1"
			continue
		}
		if strings.HasPrefix(line, "pctr2=") {
			pctrFlag = "pctr2"
			continue
		}
		if strings.HasPrefix(line, "}") {
			pctrFlag = ""
			continue
		}
		//广告配置行，以英文都好分隔
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line,"]") {
			line = line[1:strings.Index(line,"]")]
			gids := strings.Split(line, ",")
			for _, gid := range gids {
				this.gids[gid] = 1
			}
		}else if strings.HasPrefix(line, "bidbase=") { //读取bidbase
			colums := strings.Split(line, "=")
			baseBidF , err := strconv.ParseFloat(colums[1],10)
			if err == nil {
				this.bidbase = math.Ceil(baseBidF * 100)
			}else {
				fmt.Printf("bidbase读取失败！,读取数据为：%s\n",colums[1])
			}
		}else if pctrFlag != "" { //否则为把名单-slotID配置行
			var imp, click float64
			colums := strings.Split(line, "\t")
			if len(colums) != 4 {
				continue
			}
			prob , err := strconv.ParseFloat(colums[3],64)
			imp , err  = strconv.ParseFloat(colums[1],64)
			click , err  = strconv.ParseFloat(colums[2],64)
			if err !=nil {
				fmt.Printf("parse prob失败！,读取数据行为：%s\n",line)
				continue
			}
			//fmt.Printf("pctr =%s, colums[0]=%s,prob=%f \n",pctrFlag,colums[0],prob)
			key := colums[0]
			if strings.HasPrefix(key,"0") {
				key = key[1:]
			}
			this.pctr[key] = prob
			if pctrFlag == "pctr1" {
				this.pctr[PCTR_1_IMP] = this.pctr[PCTR_1_IMP] + imp
				this.pctr[PCTR_1_CLICK] = this.pctr[PCTR_1_CLICK] + click
			}
			if pctrFlag == "pctr2" {
				this.pctr[PCTR_2_IMP] = this.pctr[PCTR_2_IMP] + imp
				this.pctr[PCTR_2_CLICK] = this.pctr[PCTR_2_CLICK] + click
			}
		}
	}
	this.avgPctr1 = this.pctr[PCTR_1_CLICK] / this.pctr[PCTR_1_IMP]
	this.avgPctr2 = this.pctr[PCTR_2_CLICK] / this.pctr[PCTR_2_IMP]
	this.on() //更新配置完毕，开启配置
	return nil
}

