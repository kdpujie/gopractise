package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	w := NewSlotWhitelist()
	for k, s := range w.slots {
		fmt.Printf("slotid：%s,v=%d \n", k, s)
	}
	for k, g := range w.gids {
		fmt.Printf("gid：%s,v=%d \n", k, g)
	}
	fmt.Printf("白名单中是否包含指定广告：%v \n", w.IsContainGid("jcfjwjo"))
	fmt.Printf("白名单中是否包含指定广告位：%v \n", w.IsContainSlot("zapcc6c85e7f35403ab57c12936e84b3415b089740e"))
}

type SlotWhitelist struct {
	rw    sync.RWMutex
	path  string //白名单所在路径
	gids  map[string]int
	slots map[string]int
}

func NewSlotWhitelist() *SlotWhitelist {
	path := "whitelist.config"
	gids := make(map[string]int, 16)
	slots := make(map[string]int, 32)
	whitelist := &SlotWhitelist{path: path, gids: gids, slots: slots}
	whitelist.readFile()
	return whitelist
}

//读取并解析白名单文件
func (this *SlotWhitelist) readFile() error {
	file, err := os.Open(this.path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//广告配置行，以英文都好分隔
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			line = line[1:strings.Index(line, "]")]
			gids := strings.Split(line, ",")
			for _, gid := range gids {
				this.addGid(gid)
			}
		} else { //否则为把名单-slotID配置行
			slots := strings.Split(line, ",")
			for _, slotId := range slots {
				this.addSlot(slotId)
			}
		}
	}
	return nil
}

//增加需要白名单的广告id
func (this *SlotWhitelist) addGid(gid string) {
	if !this.IsContainGid(gid) {
		this.rw.RLock()
		this.gids[gid] = 1
		this.rw.RUnlock()
	}
}

//像白名单中增加slotID
func (this *SlotWhitelist) addSlot(slotId string) {
	if !this.IsContainSlot(slotId) {
		this.rw.RLock()
		this.slots[slotId] = 1
		this.rw.RUnlock()
	}
}

//白名单是否存在该slotid
func (this *SlotWhitelist) IsContainSlot(slotId string) bool {
	_, ok := this.slots[slotId]
	return ok
}

//白名单中是否已存在该gid
func (this *SlotWhitelist) IsContainGid(gid string) bool {
	_, ok := this.gids[gid]
	return ok
}
