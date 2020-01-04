package main

import "log"
import "github.com/golang/protobuf/proto"
import "learn.com/pb/ks"

func start_Channel() {
	ch := make(chan int)
	request := generateKsAndroidRequest()
	go func() {
		//request.Adslot.AdslotId = proto.String("after")
		changeValue(*request)
		ch <- 1
	}()
	<-ch
	log.Println(*request.Adslot.AdslotId)
}

func changeValue(request ks.MobadsRequest) {
	request.Adslot.AdslotId = proto.String("after")
}
