package main

import "C"
import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
	"ksyun.com/commons/util"
	"ksyun.com/commons/entry/thrift/rtb"
	"sync"
)

var nilAdList []string = make([]string,0,0)

//广告位级 白名单
type slotWhitelist struct {
	rw 			sync.RWMutex
	online		bool   //实验开关
	configPath	string     //白名单所在路径
	gids 		map[string]int
	slots   	map[string]int
}
// exported(symbol)
var Whitelist slotWhitelist

//系统启动或者重新加载配置时，由框架调用
func (this *slotWhitelist) Init()  {
	this.configPath = "conf/plugs/whitelist.config"
	this.slots = make(map[string]int, 32)
	this.gids = make(map[string]int, 32)
	this.readFile(this.configPath) //初始化配置
/*	fmt.Println("白名单加载的广告有：")
	for gid, _ := range this.gids {
		fmt.Printf("\tgid=%s \n",gid)
	}*/
/*	fmt.Println("白名单加载的slotId有：")
	for slotId, _ := range this.slots {
		fmt.Printf("\tslotId=%s \n",slotId)
	}*/
}

func (this *slotWhitelist) Filter(expId string, flowId int, inputAdlist []string, req *rtb.MReq) []string{
	slotId := req.GetSlots()[0].GetAdslotID()
	fmt.Printf("mfs插件：收到inputAdlist=%v, slotId=%s, flowId=%d \n",inputAdlist, slotId, flowId)
	if this.online && flowId==0 {
		var outputAdList []string = make([]string,0,len(inputAdlist))
		if  prob,ok := this.slots[slotId];ok { //命中广告位
			r := util.RandomInt(101)
			if prob >= r { //命中概率
				for _, gid := range inputAdlist {
					if _, ok := this.gids[gid]; ok { //命中广告
						outputAdList = append(outputAdList, gid)
					}
				}
			}
		}
		return outputAdList
	}
	return nilAdList
}



//读取并解析白名单文件
func (this *slotWhitelist) readFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	this.off() //关闭实验，更新配置数据
	this.clearCache()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//广告配置行，以英文都好分隔
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line,"]") {
			line = line[1:strings.Index(line,"]")]
			gs := strings.Split(line, ",")
			for _, gid := range gs {
				this.gids[gid] = 1
			}
		}else { //否则为把名单-slotID配置行
			colums := strings.Split(line, "\t")
			if len(colums) == 2 {
				prob , err := strconv.ParseFloat(colums[1],10)
				if err ==nil {
					this.slots[colums[0]] = int(math.Ceil(prob * 100))
				}
			}
		}
	}
	this.on() //配置数据更新完毕，开启实验
	return nil
}

//广告位清除
func (this *slotWhitelist) clearCache()  {
	for k, _ :=range this.slots {
		delete(this.slots,k)
	}
	for k, _ :=range this.gids {
		delete(this.gids,k)
	}
}

//打开实验
func (this *slotWhitelist) on()  {
	this.rw.RLock()
	this.online = true
	this.rw.RUnlock()
}

//关闭实验
func (this *slotWhitelist) off()  {
	this.rw.RLock()
	this.online = false
	this.rw.RUnlock()
}
