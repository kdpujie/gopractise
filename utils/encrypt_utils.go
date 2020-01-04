package main

import (
	"ksyun.com/commons/util"
	"fmt"
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"hash"
	"strings"
)
//
func main(){
	h := md5.New()
	splitRatio := "10:10:80"
	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"FF:FF:FF:FF:FF:FF", "0"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"FC:FF:FF:FF:FF:FF", "1"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"02:00:00:00:00:00", "2"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"00:00:00:00:00:00","3"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"00000000-0000-0000-0000-000000000000","4"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"012345678912345", "4"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"00000000000000","5"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"000000000000000","6"))

	fmt.Printf("\t flowid=%d \n", flowId(h,splitRatio,"fafasdfadsdf","7"))
	testMd5("898600710115")
}
func testMd5(src string) {
	fmt.Printf("MD5: %s md5后的值为：%s \n",src, util.Md5(src))
}
//18446744073709551615
func hashValue(str string) {
	hash := md5.New()
	hash.Write([]byte(str))
	hexMd5 := hex.EncodeToString(hash.Sum(nil))
	bb, _ := strconv.ParseUint(hexMd5[16:32],16, 64)
	fmt.Printf("str=%s, hash=%d \n ",hexMd5 , bb)
}

func splitBucket(h hash.Hash, value, shuffle string) int {
	h.Reset()
	h.Write([]byte(value + shuffle))
	hValue, _ := strconv.ParseUint(hex.EncodeToString(h.Sum(nil))[16:32],16, 64)
	bucketId := hValue % 100
	fmt.Printf("str=%s, shuffle=%s, bucketId=%d ", value, shuffle, bucketId)
	return int(bucketId)
}
//按照流量配比，分配流量，返回分配id
//比如：10:10:80, 分成流量0，1，2;
func flowId(h hash.Hash, splitRatio, value, shuffle string) int  {
	var boundary int = 0
	bucketId := splitBucket(h, value, shuffle)
	ratios := strings.Split(splitRatio, ":")
	for i, ratio := range ratios {
		b , _ := strconv.ParseInt(ratio, 10 , 64)
		boundary = boundary + int(b)
		if bucketId < boundary {
			return i
		}
	}
	return len(ratios) - 1
}