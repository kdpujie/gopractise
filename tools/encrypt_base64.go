/**
@description	加解密联系
@author pujie
@data	2018-02-01
**/

package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func TestBase64() {
	targetStr := "fgnsfsfgsergcxcfb"
	encodeStr := base64.StdEncoding.EncodeToString([]byte(targetStr))
	fmt.Fprintf(os.Stdout, "encode:%s\n", encodeStr)
	decodeBytes, _ := base64.StdEncoding.DecodeString("+")
	fmt.Fprintf(os.Stdout, "decode: %s\n", string(decodeBytes))

	uEnc := base64.URLEncoding.EncodeToString([]byte(targetStr))
	fmt.Fprintf(os.Stdout, "urlEncode:%s\n", uEnc)

	uDec, _ := base64.URLEncoding.DecodeString("-")
	fmt.Fprintf(os.Stdout, "urlDecode:%s\n", string(uDec))
}

//
// func TestByteSlice() {
// 	t := &track.TrackLog{}
// 	t.RequestId = "rid_0000111"
// 	t.Vid = "vid_alsdgioj23r134234asdf"
// 	t.ChannelId = "1"
// 	t.AdslotId = "slotId_345kafds"
// 	t.TrackingEventType = "start"
// 	t.OsType = 1
// 	t.AdType = 1
// 	t.AppId = "appId_234dsf"
// 	data, err := proto.Marshal(t)
// 	if err != nil {
// 		fmt.Printf("序列化出错:%v\n", err)
// 		return
// 	}
// 	encodeStr := base64.URLEncoding.EncodeToString(data)
// 	fmt.Fprintf(os.Stdout, "encode:%s len=%d\n", encodeStr, len(encodeStr))
//
// 	d, err := base64.URLEncoding.DecodeString(encodeStr)
// 	if err != nil {
// 		fmt.Printf("DecodeString:%v\n", err)
// 		return
// 	}
// 	t1 := &track.TrackLog{}
// 	err = proto.Unmarshal(d, t1)
// 	if err != nil {
// 		fmt.Printf("Unmarshal:%v\n", err)
// 	}
// 	fmt.Printf("解开的值：requestId=%s\n", t1.RequestId)
// }
