package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/entry/uprofile"
	"ksyun.com/commons/util"
	"learn.com/redis/tp"
)

func main() {
	rp := tp.RedisPool()
	con := rp.Get()
	//insertUserProfile(con)
	data, _ := redis.ByteSlices(con.Do("mget", "a2a52583c2c816e01a79d70d4b0f39cb", "8491ce2ca14ea84dd817ea681a82cbdc"))
	userProfile := &uprofile.UserProfile{}
	proto.Unmarshal(data[1], userProfile)
	fmt.Printf("user-profile:%s \n", userProfile.String())
	con.Close()
}

func insertUserProfile(con redis.Conn) {
	imei := "866361021509590"
	imeiMd5 := util.Md5(imei)
	fmt.Printf("uid=%s \n", imeiMd5)
	up := generateUserInfo(imeiMd5)
	data, _ := proto.Marshal(up)
	con.Do("set", imeiMd5, data)

}

func generateUserInfo(uid string) *uprofile.UserProfile {

	up := &uprofile.UserProfile{
		Uid:           uid,
		PayMoney:      "4401003",
		GameFrequency: "2,4",
		GameTimes:     "20",
	}
	return up
}
