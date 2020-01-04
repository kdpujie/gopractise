package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	PCTR_1_IMP   = "pctr1_imp"
	PCTR_1_CLICK = "pctr1_click"
	PCTR_2_IMP   = "pctr2_imp"
	PCTR_2_CLICK = "pctr2_click"
)

func main() {
	w := NewBidPriceStrategy()
	for k, s := range w.pctr {
		fmt.Printf("pctr：%s,v=%f \n", k, s)
	}
	for k, g := range w.gids {
		fmt.Printf("gid：%s,v=%d \n", k, g)
	}
	fmt.Printf("bidbase=%f,avg:pctr1=%f,pctr2=%f \n", w.bidbase, w.avgPctr1, w.avgPctr2)
	fmt.Printf("300x250 bidprice=%d \n", w.getBidprice(300, 250))
	fmt.Printf("白名单中是否包含指定广告：%v \n", w.IsContainGid("ad4"))
}

func NewBidPriceStrategy() *BidPriceStrategy {
	path := "bidprce_strategy.conf"
	gids := make(map[string]int)
	pctr := make(map[string]float64)
	whitelist := &BidPriceStrategy{path: path, gids: gids, pctr: pctr}
	whitelist.readFile()
	return whitelist
}

type BidPriceStrategy struct {
	rw       sync.RWMutex
	path     string //白名单所在路径
	bidbase  float64
	gids     map[string]int
	pctr     map[string]float64
	avgPctr1 float64
	avgPctr2 float64
}

//获取竞价价格
//Bid=bidbase*（ctrpre1/ctravg1）*（ctrpre2/ctravg2）
func (this *BidPriceStrategy) getBidprice(width, height int32) int64 {
	size := fmt.Sprintf("%dx%d", width, height)
	if sProb, ok := this.pctr[size]; ok {
		if hProb, ok := this.pctr[strconv.Itoa(time.Now().Hour())]; ok {
			bidPrice := this.bidbase * (sProb / this.avgPctr1) * (hProb / this.avgPctr2)
			return int64(math.Ceil(bidPrice))
		}
	}
	return int64(math.Ceil(this.bidbase))
}

//白名单中是否已存在该gid
func (this *BidPriceStrategy) IsContainGid(gid string) bool {
	_, ok := this.gids[gid]
	return ok
}

//读取并解析白名单文件
func (this *BidPriceStrategy) readFile() error {
	file, err := os.Open(this.path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
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
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			line = line[1:strings.Index(line, "]")]
			gids := strings.Split(line, ",")
			for _, gid := range gids {
				this.gids[gid] = 1
			}
		} else if strings.HasPrefix(line, "bidbase=") { //读取bidbase
			colums := strings.Split(line, "=")
			baseBidF, err := strconv.ParseFloat(colums[1], 10)
			if err == nil {
				this.bidbase = math.Ceil(baseBidF * 100)
			} else {
				fmt.Printf("bidbase读取失败！,读取数据为：%s\n", colums[1])
			}
		} else if pctrFlag != "" { //否则为把名单-slotID配置行
			var imp, click float64
			colums := strings.Split(line, "\t")
			if len(colums) != 4 {
				continue
			}
			prob, err := strconv.ParseFloat(colums[3], 64)
			imp, err = strconv.ParseFloat(colums[1], 64)
			click, err = strconv.ParseFloat(colums[2], 64)
			if err != nil {
				fmt.Printf("parse prob失败！,读取数据行为：%s\n", line)
				continue
			}
			//fmt.Printf("pctr =%s, colums[0]=%s,prob=%f \n",pctrFlag,colums[0],prob)
			key := colums[0]
			if strings.HasPrefix(key, "0") {
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
	return nil
}
