package main

import (
	"github.com/chasex/redis-go-cluster"
	"time"
	"log"
)

func main()  {
	c, err := initCluster()
	if err != nil {
		log.Fatalf("redis cluster init failed:%s \n",err.Error())
	}
	time.Sleep(time.Second * 3)
	c.Do("set","foo","bar")
	reply ,err := redis.Int64(c.Do("get","dsp_ad_day_imp_tehh3gwn_20170524"))
	log.Printf("redis cluster: key=foo, value=%d",reply)
}


func initCluster() (*redis.Cluster, error ) {
	cluster, err := redis.NewCluster(&redis.Options{
		StartNodes: []string{"192.168.153.44:6379", "192.168.153.50:6379", "192.168.153.51:6379"},
		ConnTimeout: 2 * time.Second,
		ReadTimeout: 50 * time.Millisecond,
		WriteTimeout: 50 * time.Millisecond,
		KeepAlive: 16,
		AliveTime: 60 * time.Second,
	})
	if err!= nil {
		return nil ,err
	}
	return cluster,nil
}
