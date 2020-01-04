package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"learn.com/redis/tp"
	"os"
	"log"
	"bufio"
	"strings"
	"ksyun.com/commons/util"
)

var (
	server string = "172.31.252.134:6379"
	//server string = "120.92.59.212:6379" //120.92.59.212
	pool   *redis.Pool
)

func main()  {
	pool = tp.RedisPool()
	conn := pool.Get()
	//key,value := "pujie3","3"
	//inserKeyValue(pool,key,value)
/*	for {
		searchKey(pool,key)
		time.Sleep(100 * time.Millisecond)
	} */
	//mget(conn,"pujie1","pujie2","pujie3","pujie4")
	initUserBlacklist(conn)
	//readUserblacklist(conn)
	conn.Close()
}

//从redis server直接获取连接.
func inserKeyValue(pool *redis.Pool,key,value string) {
	conn := pool.Get()
	_, err := conn.Do("set",key,value)
	if err != nil {
		fmt.Printf("inserKeyValue err=%s \n",err.Error())
	}
	conn.Close()
}

func searchKey (pool *redis.Pool,key string) {
	conn := pool.Get()
	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		fmt.Printf("searchKey err=%s \n",err.Error())
	}else{
		fmt.Printf("searchKey 返回值%s \n", reply)
	}
	conn.Close()
}

func mget(conn redis.Conn, key... interface{})  {
	reply, err := redis.Ints(conn.Do("mget", key...))

	if err != nil {
		fmt.Printf("searchKey err=%s \n",err.Error())
	}else{
		for key, value := range reply {
			fmt.Printf("key=%d, value=%d \n", key, value)
		}
	}

}
//读取用户黑名单配置文件 -> redis
func initUserBlacklist(conn redis.Conn)  {
	configPath := "conf/user_blacklist.conf"
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("readUserBlacklist：%s \n", err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		colums := strings.Split(line, "\t")
		var key string
		if len(colums) <2 {
			key = util.Md5(line)
		}else {
			if strings.ToLower(colums[1]) != "md5" {
				break
			}
			key = util.Md5(colums[0])
		}
		reply, err := conn.Do("sadd", "kir_blacklist_user", key)
		if err != nil {
			fmt.Printf("readUserBlacklist key=%s, err=%s \n", key, err.Error())
		}else{
			fmt.Printf("readUserBlacklist key=%s, 返回值%d \n", key, reply)
		}
	}
}

func readUserblacklist(conn redis.Conn)  {
	members , err := redis.Strings(conn.Do("smembers", "kir_blacklist_user"))
	if err != nil {
		log.Printf("readUserblacklist err message=%s \n", err.Error())
		return
	}
	for i, member := range members {
		log.Printf("member %d : %s \n",i ,member)
	}
}