package main

import (
	"fmt"
	//	"strconv"
	"time"

	"github.com/twinj/uuid"
	"github.com/twinj/uuid/savers"
	//	"ksyun.com/commons/util"
)

func startUUID() {
	fmt.Println("-----------------UUID测试-------------------------")
	saver := new(savers.FileSystemSaver)
	saver.Report = true
	saver.Duration = 3 * time.Second

	// Run before any v1 or v2 UUIDs to ensure the savers takes
	uuid.RegisterSaver(saver)
	//	uuid.Init()
	start := time.Now()
	for i := 0; i < 1000; i++ {
		//uuid.NewV3(uuid.NameSpaceURL, uuid.Name("ww.ksyun.com"))
		//uuid.NewV3(uuid.NameSpaceURL, uuid.Name("ww.ksyun.com"))
		uuid.NewV1()
		//fmt.Println(uuid.NewV2())
	}
	//id := uuid.NewV3(uuid.NameSpaceURL, uuid.Name("ww.ksyun.com"))
	id := uuid.NewV1()
	end := time.Now()
	var dur_time time.Duration = end.Sub(start)
	var elapsed_nano int64 = dur_time.Nanoseconds() / 1000
	fmt.Println(id)
	fmt.Printf("生成序列耗时:%d \n", elapsed_nano)
	//uuid_satori()
	uuid_Md5()
}

func uuid_Md5() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		time.Now().UnixNano()
		//util.Md5(strconv.FormatInt(t, 10) + "1472701045108533800")
	}
	end := time.Now()
	var dur_time time.Duration = end.Sub(start)
	var elapsed_nano int64 = dur_time.Nanoseconds() / 1000
	//fmt.Println(id)
	fmt.Printf("md5生成序列耗时:%d \n", elapsed_nano)
	//fmt.Println(ss)
}
