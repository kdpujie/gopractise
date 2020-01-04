package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	RedisServer string = "123.59.166.189:6379"
)

var (
	server string = "123.59.166.189:6379"
	pool   *redis.Pool
)

type StrategyEntry struct {
	Name            string //渠道名称(唯一)
	AppId           string //渠道分配的appID
	ChannelAdSlotId string //渠道分配的广告ID
	SlotType        uint32 `json:",string"` //广告位类型
	Weight          int    `json:",string"` //权重大的, 能获得更大的分发机会.}
	Ads             uint32 `json:",string"` //返回广告数
}

func startRedis2Map() {
	pool = poolInit()
	conn := pool.Get()
	defer conn.Close()
	reply, err := redis.StringMap(conn.Do("hgetall", "ssp_adid_config_91106012"))
	if err != nil {
		log.Println("redis 查找err", err)
		return
	}
	for k, v := range reply {
		log.Printf("key=%s, value=%s", k, v)
	}
}

func startMap() {
	var m map[string][]StrategyEntry = make(map[string][]StrategyEntry)
	var baidu *StrategyEntry = &StrategyEntry{
		Name:            "baidu",
		AppId:           "ebce05f9",
		ChannelAdSlotId: "2542725",

		Weight: 9,
	}
	var tencent *StrategyEntry = &StrategyEntry{
		Name:            "tencent",
		AppId:           "ebce05f9",
		ChannelAdSlotId: "2542725",
		Weight:          4,
	}
	var ss []StrategyEntry = []StrategyEntry{}
	ss = append(ss, *baidu)
	ss = append(ss, *tencent)

	m["channel"] = ss
	for k, v := range m {
		log.Println("key = ", k)
		for i, e := range v {
			log.Print("\t", i, e.Name)
		}
	}
	var s StrategyEntry
	log.Println("1.", s)
}

func start_random() {
	rand.Seed(time.Now().Unix())
	var baidu *StrategyEntry = &StrategyEntry{
		Name:            "baidu",
		AppId:           "ebce05f9",
		ChannelAdSlotId: "2542725",
		Weight:          9,
	}
	var tencent *StrategyEntry = &StrategyEntry{
		Name:            "tencent",
		AppId:           "ebce05f9",
		ChannelAdSlotId: "2542725",
		Weight:          4,
	}
	var strategys map[string]StrategyEntry = make(map[string]StrategyEntry)
	strategys["baidu"] = *baidu
	strategys["tencent"] = *tencent
	for i := 0; i < 20; i++ {
		channeName := unEqualRandom(strategys)
		log.Println("被选中的渠道:", channeName)
	}

}

func unEqualRandom(strategys map[string]StrategyEntry) string {
	var sum, index, name = 0, 0, ""
	for _, v := range strategys {
		sum += v.Weight
	}
	random := rand.Intn(sum)
	for k, v := range strategys {
		index += v.Weight
		if random < index {
			name = k
			break
		}
	}
	return name
}

//序列化测试
func start_marshal() {
	pool = poolInit()
	var strategy *StrategyEntry = &StrategyEntry{
		Name:            "baidu",
		AppId:           "ebce05f9",
		ChannelAdSlotId: "2542728",
		SlotType:        2,
		Weight:          9,
	}

	data, err := json.Marshal(strategy)
	if err != nil {
		log.Println("序列化异常:", err)
		return
	}
	log.Println("json串:", string(data))
	//save2RedisAsMap("37835382", strategy.Name, string(data))

}

//反序列化测试
func start_unmarshal() {
	var s StrategyEntry
	var str = `{"Ads":"5","AppId":"esqugqdy","ChannelAdSlotId":"pujie-tencent-slot-8","Weight":"2","Name":"tencent","SlotType":"8"}`
	json.Unmarshal([]byte(str), &s)
	log.Println(s)
	log.Println(s.AppId)
	//	pool = poolInit()
	//	conn := pool.Get()
	//	defer conn.Close()

	//	reply, err := redis.Strings(conn.Do("keys", "ssp_adid_config_*"))
	//	if err != nil {
	//		log.Print(err)
	//	}
	//	var s StrategyEntry
	//	for _, v := range reply {
	//		s, _ = unmarsha2Strategy(v, "baidu")

	//		log.Printf("name=%s, AppId=%s,ChannelAdSlotId=%s,Weight=%d,Ads=%d", s.Name, s.AppId, s.ChannelAdSlotId, s.Weight, s.Ads)
	//	}

}

//查找并反序列化StrategyEntry
func unmarsha2Strategy(key, field string) (s StrategyEntry, err error) {
	conn := pool.Get()
	defer conn.Close()
	reply, err := redis.String(conn.Do("hget", key, field))
	if err != nil {
		return s, err
	}
	var strategy StrategyEntry
	json.Unmarshal([]byte(reply), &strategy)
	return strategy, nil
}

//保存map值到redis
func save2RedisAsMap(slotId, channelName, value string) {
	conn := pool.Get()
	defer conn.Close()
	conn.Do("hset", "ssp_adid_config_"+slotId, channelName, value)
}

//初始化连接池
func poolInit() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", RedisServer)
			if err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", "Ksc-redis-6379"); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}
