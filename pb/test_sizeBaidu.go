package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"learn.com/pb/ks"
	//	"time"

	"github.com/golang/protobuf/proto"
)

const (
	IS_DEBUG = true
)

var iconmap, imagemap map[string]bool
var isOK = false                                                           //搜索DOWNLOAD类型
var reqFun func(os string, adtype int, w uint32, h uint32) = startKsMobApi //访问方法设置//startBaiduMobApi//startKsMobApi

/**
压力测试
**/
func Push(c chan string) {

	for i := 0; i < 100; i++ {
		go finsh(c)
	}
	for i := 0; i < 100; i++ {
		s := <-c
		fmt.Println(s)
	}
	os.Exit(0)
}
func finsh(c chan string) {
	defer func() {
		c <- "finsh()"
	}()
	StartTestSize()
}

/***
移动ssp_api测试
**/
func StartTestSize() {
	//采用map记录所有尺寸
	iconmap = make(map[string]bool)
	imagemap = make(map[string]bool)

	if false { //搜索DOWNLOAD类型
		for isOK == false {
			reqFun("android", 8, 0, 0)
		}
		return
	}

	for i := 0; i < 100; i++ {
		reqFun("android", 1, 580, 90)
		reqFun("android", 1, 360, 300)
		reqFun("android", 1, 480, 60)
		reqFun("android", 1, 640, 270)
		time.Sleep(1 * 100)

		reqFun("android", 4, 640, 960)
		reqFun("android", 4, 720, 1280)
		reqFun("android", 4, 1024, 768)
		time.Sleep(1 * 100)

		reqFun("android", 2, 360, 300)
		time.Sleep(1 * 100)

		reqFun("android", 8, 0, 0)
		time.Sleep(1 * 100)
	}
	for i := 0; i < 100; i++ {
		reqFun("ios", 1, 580, 90)
		reqFun("ios", 1, 360, 300)
		reqFun("ios", 1, 480, 60)
		reqFun("ios", 1, 640, 270)
		time.Sleep(1 * 100)

		reqFun("ios", 4, 640, 960)
		reqFun("ios", 4, 720, 1280)
		reqFun("ios", 4, 1024, 768)
		time.Sleep(1 * 100)

		reqFun("ios", 2, 360, 300)
		time.Sleep(1 * 100)

		reqFun("ios", 8, 0, 0)
		time.Sleep(1 * 100)
	}

	fmt.Println("iconmap")
	for e, _ := range iconmap {
		fmt.Println(e)
	}
	fmt.Println("imagemap")
	for e, _ := range imagemap {
		fmt.Println(e)
	}
}
func startKsMobApi(os string, adtype int, w uint32, h uint32) {
	var logstr, adsid, appid string
	var request *ks.MobadsRequest
	if os == "ios" {
		appid = "ay884746"
		switch adtype {
		case 1:
			adsid = "87034385"
		case 2:
			adsid = "77333366"
		case 4:
			adsid = "17333387"
		case 8:
			adsid = "47132318"
		}
		request = _generateKsIosRequest(adsid, appid, w, h)
	} else {
		appid = "ay45d93d"
		switch adtype {
		case 1:
			adsid = "88550654"
		case 2:
			adsid = "28759613"
		case 4:
			adsid = "28557681"
		case 8:
			adsid = "88955698"
		}
		request = _generateKsAndroidRequest(adsid, appid, w, h)
	}

	logstr += os + "\t"
	logstr += "adtype=" + strconv.Itoa(adtype) + "\treqSize=(" + strconv.Itoa(int(w)) + "," + strconv.Itoa(int(h)) + ")\t"

	//序列化
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	baiduR, err := http.NewRequest("POST", "http://120.92.16.234/api/def", bytes.NewBuffer(data)) //KS
	response, err := client.Do(baiduR)
	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	res := &ks.MobadsResponse{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	_readKsResponse(res, logstr)
}
func startBaiduMobApi(os string, adtype int, w uint32, h uint32) {
	var logstr, adsid, appid string
	var request *ks.MobadsRequest
	if os == "ios" {
		appid = "e0884746"
		switch adtype {
		case 1:
			adsid = "2563818"
		case 2:
			adsid = "2563825"
		case 4:
			adsid = "2563824"
		case 8:
			adsid = "2563829"
		}
		request = _generateKsIosRequest(adsid, appid, w, h)
	} else {
		appid = "ebce05f9"
		switch adtype {
		case 1:
			adsid = "2542725"
		case 2:
			adsid = "2542728"
		case 4:
			adsid = "2542727"
		case 8:
			adsid = "2542729"
		}
		request = _generateKsAndroidRequest(adsid, appid, w, h)
	}

	logstr += os + "\t"
	logstr += "adtype=" + strconv.Itoa(adtype) + "\treqSize=(" + strconv.Itoa(int(w)) + "," + strconv.Itoa(int(h)) + ")\t"

	//序列化
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	baiduR, err := http.NewRequest("POST", "http://debug.mobads.baidu.com/api_5", bytes.NewBuffer(data)) //百度
	response, err := client.Do(baiduR)
	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	res := &ks.MobadsResponse{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	_readKsResponse(res, logstr)
}

//单个测试
func testOne(os string, adsid string, appid string, w uint32, h uint32) {
	var logstr string
	var request *ks.MobadsRequest
	if os == "ios" {
		request = _generateKsIosRequest(adsid, appid, w, h)
	} else {
		request = _generateKsAndroidRequest(adsid, appid, w, h)
	}

	logstr += os + "\t"
	logstr += "\treqSize=(" + strconv.Itoa(int(w)) + "," + strconv.Itoa(int(h)) + ")\t"

	//序列化
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	baiduR, err := http.NewRequest("POST", "http://debug.mobads.baidu.com/api_5", bytes.NewBuffer(data)) //百度
	response, err := client.Do(baiduR)
	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	res := &ks.MobadsResponse{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	_readKsResponse(res, logstr)
}

//读取广告信息
func _readKsResponse(response *ks.MobadsResponse, logstr string) {
	logger := getLogger("./Log/baidu.log")
	var logstr1, logstr2 string
	logstr1 += "\tgeticonSize="
	logstr2 += "\tgetimageSize="

	for _, ad := range response.GetAds() {
		if IS_DEBUG {
			log.Println("slot_id = ", ad.GetAdslotId())
		}
		groups := ad.GetMetaGroup()
		for i, group := range groups {
			if IS_DEBUG {
				log.Printf("\t%d creative_type=%s", i, group.GetCreativeType())
				log.Printf("\t%d interaction_type=%s", i, group.GetInteractionType())
				log.Printf("\t%d win_notice_url=%s", i, group.GetWinNoticeUrl())
				log.Printf("\t%d click_url=%s", i, group.GetClickUrl())
				log.Printf("\t%d title=%s", i, group.GetTitle())
				log.Printf("\t%d brand_name=%s", i, group.GetBrandName())
				log.Printf("\t%d description=%s", i, group.GetDescription())
				log.Printf("\t%d icon_src=%s", i, group.GetIconSrc())
				log.Printf("\t%d image_src=%s", i, group.GetImageSrc())
				log.Printf("\t%d app_size=%d", i, group.GetAppSize())
			}
			icons := group.GetIconSrc()
			for _, k := range icons {
				x1, y1 := GetImgSize(k)
				tmpstr := "(" + strconv.Itoa(x1) + "," + strconv.Itoa(y1) + ")"
				iconmap[tmpstr] = true
				logstr1 += tmpstr + ","
			}
			images := group.GetImageSrc()
			for _, k := range images {
				x2, y2 := GetImgSize(k)
				tmpstr := "(" + strconv.Itoa(x2) + "," + strconv.Itoa(y2) + ")"
				imagemap[tmpstr] = true
				logstr2 += tmpstr + ","
			}

			if string(group.GetInteractionType()) == "DOWNLOAD" { //搜索DOWNLOAD类型
				isOK = true
				//				return
			}
		}
	}

	logstr += logstr1 + logstr2
	logger.Println(logstr)
}
func _generateKsAndroidRequest(adsid string, appid string, w uint32, h uint32) *ks.MobadsRequest {
	//slot info
	slotSize := &ks.Size{
		Width:  proto.Uint32(w),
		Height: proto.Uint32(h),
	}
	slot := &ks.AdSlot{
		AdslotId:   proto.String(adsid), //47435394(原生)	27433301(横幅)
		AdslotSize: slotSize,
	}
	//app info
	app := &ks.App{
		AppId:      proto.String(appid),
		AppVersion: &ks.Version{Major: proto.Uint32(3), Minor: proto.Uint32(4)},
	}
	//device info
	uId := &ks.UdId{
		//Idfa:      proto.String("FC0F3445-0FCE-40EE-8646-3CA8BB2663EA"),
		Mac:       proto.String("12:34:56:78:90:ab"),
		Imei:      proto.String("3Bd52A121ba64c8"),
		AndroidId: proto.String("3fd5h2fg2sgr64h4"),
	}
	deviceSize := &ks.Size{
		Width:  proto.Uint32(1080),
		Height: proto.Uint32(1920),
	}
	deviceType := ks.Device_PHONE
	osType := ks.Device_ANDROID
	device := &ks.Device{
		DeviceType: &deviceType,
		OsType:     &osType,
		OsVersion:  &ks.Version{Major: proto.Uint32(6), Minor: proto.Uint32(0)},
		Vendor:     []byte("MEIZU"),
		Model:      []byte("MX5"),
		Udid:       uId,
		ScreenSize: deviceSize,
	}
	//network info
	connType := ks.Network_WIFI
	operatorType := ks.Network_CHINA_MOBILE
	network := &ks.Network{
		Ipv4:           proto.String("114.255.44.132"),
		ConnectionType: &connType,
		OperatorType:   &operatorType,
	}
	//gps info
	coordinateType := ks.Gps_WGS84
	gps := &ks.Gps{
		CoordinateType: &coordinateType,
		Longitude:      proto.Float64(40.7127),
		Latitude:       proto.Float64(74.0059),
		Timestamp:      proto.Uint32(123456),
	}
	request := &ks.MobadsRequest{
		RequestId:  proto.String("PlCe8PBnrK3sUZ1mmC0gkbyUP4j8XfZp"),
		ApiVersion: &ks.Version{Major: proto.Uint32(5), Minor: proto.Uint32(0)},
		Adslot:     slot,
		App:        app,
		Device:     device,
		Network:    network,
		Gps:        gps,
		IsDebug:    proto.Bool(true),
	}
	return request
}
func _generateKsIosRequest(adsid string, appid string, w uint32, h uint32) *ks.MobadsRequest {
	//slot info
	slotSize := &ks.Size{
		Width:  proto.Uint32(w),
		Height: proto.Uint32(h),
	}
	slot := &ks.AdSlot{
		AdslotId:   proto.String(adsid),
		AdslotSize: slotSize,
	}
	//app info
	app := &ks.App{
		AppId:      proto.String(appid),
		AppVersion: &ks.Version{Major: proto.Uint32(3), Minor: proto.Uint32(4)},
	}
	//device info
	uId := &ks.UdId{
		Idfa:      proto.String("FC0F3445-0FCE-40EE-8646-3CA8BB2663EA"),
		Mac:       proto.String("12:34:56:78:90:ab"),
		Imei:      proto.String("3Bd52A121ba64c8"),
		AndroidId: proto.String("3fd5h2fg2sgr64h4"),
	}
	deviceSize := &ks.Size{
		Width:  proto.Uint32(1080),
		Height: proto.Uint32(1920),
	}
	deviceType := ks.Device_PHONE
	osType := ks.Device_IOS //之前代码为安卓，需采用IOS
	device := &ks.Device{
		DeviceType: &deviceType,
		OsType:     &osType,
		OsVersion:  &ks.Version{Major: proto.Uint32(4), Minor: proto.Uint32(2)},
		Vendor:     []byte("IPHONE"),
		Model:      []byte("IPHONE4S"),
		Udid:       uId,
		ScreenSize: deviceSize,
	}
	//network info
	connType := ks.Network_WIFI
	operatorType := ks.Network_CHINA_MOBILE
	network := &ks.Network{
		Ipv4:           proto.String("114.255.44.132"),
		ConnectionType: &connType,
		OperatorType:   &operatorType,
	}
	//gps info
	coordinateType := ks.Gps_WGS84
	gps := &ks.Gps{
		CoordinateType: &coordinateType,
		Longitude:      proto.Float64(40.7127),
		Latitude:       proto.Float64(74.0059),
		Timestamp:      proto.Uint32(123456),
	}
	request := &ks.MobadsRequest{
		RequestId:  proto.String("PlCe8PBnrK3sUZ1mmC0gkbyUP4j8XfZp"),
		ApiVersion: &ks.Version{Major: proto.Uint32(5), Minor: proto.Uint32(0)},
		Adslot:     slot,
		App:        app,
		Device:     device,
		Network:    network,
		Gps:        gps,
		IsDebug:    proto.Bool(true),
	}
	return request
}

func GetImgSize(path string) (int, int) {
	if len(path) > 0 {
		ppath := strings.Split(path, "?")
		//	fmt.Println("GetImgSize路径为:", path)
		src, err := getURLImg(ppath[0])
		if err == nil {
			bound := src.Bounds()
			dx := bound.Dx()
			dy := bound.Dy()

			//		fmt.Println("正确获取图片，获取尺寸！")

			return dx, dy
		}
	}

	return 0, 0
}
func getURLImg(url string) (img image.Image, err error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	pix, err := ioutil.ReadAll(resp.Body)
	rd := bytes.NewReader(pix)
	img, err = jpeg.Decode(rd)
	if err != nil {
		img, err = png.Decode(rd)
		if err != nil {
			//			fmt.Println("解析图片失败！")
		}
	}

	return img, err
}

//根据名称获取logger
func getLogger(name string) (tp *log.Logger) {
	defer func() {
		if e, ok := recover().(error); ok {
			log.Println("WARN: panic in %v", e)
			log.Println(string(debug.Stack()))
		}
	}()
	tp = safeLogger(name)

	return tp
}
func safeLogger(name string) *log.Logger {
	// 不存在目录时创建目录
	p, _ := path.Split(name)
	d, err := os.Stat(p)
	if err != nil || !d.IsDir() {
		if err := os.MkdirAll(p, 0777); err != nil {
			log.Println("Creat dir faile!")
		}
	}
	// 不存在文件时创建文件
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("%s\r\n", err.Error())
	}
	//runtime.SetFinalizer(file, file.Close) //GC时,释放资源

	return log.New(file, "", log.LstdFlags)
}

//**************************************************************************************//
//序列化测试
func _start_marshal(selfid string, name string, appid string, slotid string, typ uint32, wt int) {
	pool = poolInit()
	var strategy *StrategyEntry = &StrategyEntry{
		Name:            name,
		AppId:           appid,
		ChannelAdSlotId: slotid,
		SlotType:        typ,
		Weight:          wt,
	}

	data, err := json.Marshal(strategy)
	if err != nil {
		log.Println("序列化异常:", err)
		return
	}
	if IS_DEBUG {
		log.Println("json串:", "selfid=", selfid, string(data))

	} else {
		log.Println("json串:", "selfid=", selfid, string(data))
		save2RedisAsMap(selfid, strategy.Name, string(data))
	}
}
func setMarshal() {
	var selfid string
	var name string
	var appid string
	var slotid string
	var typ uint32
	var wt int

	{
		//******百度*******/
		name = "baidu"
		wt = 9

		arryAppid := [...]string{
			"df45d93d", "df45d93d", "df45d93d", "df45d93d", "df45d93d", "df45d93d", "df45d93d", "df45d93d", "df45d93d", "e16f710b",
			"f2bcc457", "f2bcc457", "f2bcc457", "f2bcc457", "f80d4259", "f80d4259", "f80d4259", "f80d4259", //"", "",/*AYang广告平台聚合id列表160707*/

			"e0884746", "e0884746", "e0884746", "e0884746", "ebce05f9", "ebce05f9", "ebce05f9", "ebce05f9", //"","",/*广告位映射-测试专用聚合广告位id-0629*/

			"a3372c8a", "fc29d2a3", "b3fdcf77", "f418a072", "a3372c8a", "fc29d2a3", "b3fdcf77", "f418a072", "a3372c8a", //"", /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			"b8b160e8", //"", "", "", "", "", "", "", "", "",/*广告位映射-TAD广告平台（ADX）广告位申请表-银橙-免费电子书*/

			"a3372c8a", "a3372c8a", "a3372c8a", //"","","","","","","",/*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/

			//"","","","","","","","","","",/*空的10个用于批量添加时复制*/
		}
		arrySlotid := [...]string{
			"2683587", "2683594", "2683608", "2683602", "2683612", "2683613", "2683616", "2683590", "2685081", "2682697",
			"2675892", "2675894", "2675900", "2675902", "2675940", "2675942", "2675944", "2684898", //"","",/*AYang广告平台聚合id列表160707*/

			"2563829", "2563824", "2563825", "2563818", "2542729", "2542727", "2542728", "2542725", //"", "",/*广告位映射-测试专用聚合广告位id-0629*/

			"2579883", "2579868", "2579860", "2579848", "2579899", "2579877", "2579863", "2579855", "2624359", //"", /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			"2706637", //"", "", "", "", "", "", "", "", "",/*广告位映射-TAD广告平台（ADX）广告位申请表-银橙-免费电子书*/

			"2713225", "2715650", "2639917", //"","","","","","","",/*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/

			//"","","","","","","","","","",/*空的10个用于批量添加时复制*/
		}
		arrySelfid := [...]string{
			"28557681", "28759613", "88550654", "88955698", "58558615", "98159646", "48459677", "68352612", "08956689", "91001412",
			"91106012", "90096722", "40499703", "80198754", "80691705", "60897716", "80699758", "31601150", //"", "",/*AYang广告平台聚合id列表160707*/

			"47132318", "17333387", "77333366", "87034385", "47435394", "17030393", "37835382", "27433301", //"","",/*广告位映射-测试专用聚合广告位id-0629*/

			"80274558", "80275682", "40972686", "20272700", "50276526", "70573680", "70279684", "70072658", "77739350", //"", /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			"82706637", //"", "", "", "", "", "", "", "", "",/*广告位映射-TAD广告平台（ADX）广告位申请表-银橙-免费电子书*/

			"80274558", "50276526", "77739350", //"","","","","","","",/*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/

			//"","","","","","","","","","",/*空的10个用于批量添加时复制*/
		}
		arryTyp := [...]uint32{
			4, 2, 1, 8, 1, 1, 1, 4, 8, 8,
			8, 1, 2, 4, 4, 1, 2, 8, /*AYang广告平台聚合id列表160707*/

			8, 4, 2, 1, 8, 4, 2, 1, /*广告位映射-测试专用聚合广告位id-0629*/

			8, 8, 8, 8, 4, 4, 4, 4, 8, /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			8, /*广告位映射-TAD广告平台（ADX）广告位申请表-银橙-免费电子书*/

			8, 4, 8, /*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/
		}
		for i, s := range arryAppid {
			appid = s
			slotid = arrySlotid[i]
			selfid = arrySelfid[i]
			typ = arryTyp[i]
			_start_marshal(selfid, name, appid, slotid, typ, wt)
			if IS_DEBUG {
				fmt.Printf("%s\t%s\t%s\t%s\t\t\t%d\n", selfid, name, appid, slotid, typ)
			}
		}
	}

	{
		//******腾讯*******/
		name = "tencent"
		wt = 9

		arryAppid := [...]string{
			"1105509910", "1105509910", "1105509910", "1105509910", "1105509910", "1105509910", "1105509910", "1105509910", "1105509910", "1105517494",
			"1105517534", "1105451817", "1105520928", "1105520928", "1105520928", "1105520928", "1105447583", "1105447583", "1105447583", "1105447583", /*AYang广告平台聚合id列表160707*/

			"1105444822", "1105444822", "1105444822", "1105444822", "1105444780", "1105444780", "1105444780", "1105444780", //"","",/*广告位映射-测试专用聚合广告位id-0629*/

			"1105395539", "1105395553", "1105468712", "1105468722", "1105395539", "1105395553", "1105468712", "1105468722", "1105395539", //"", /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			"1105395539", "1105395539", "1105395539", //"","","","","","","",/*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/

			//"","","","","","","","","","",/*空的10个用于批量添加时复制*/
		}
		arrySlotid := [...]string{
			"1070811228557681", "7060512228759613", "7010415288550654", "2040114288955698", "3050210258558615", "5060811298159646", "3020611248459677", "1070116268352612", "3050710208956689", "1050619311902026",
			"4030411351602097", "2060416391001412", "8000819391106012", "2010715390096722", "8080812340499703", "9080010380198754", "3070115380691705", "4020713360897716", "4020611380699758", "4010116331601150", /*AYang广告平台聚合id列表160707*/

			"8070214247132318", "5080311217333387", "2010917277333366", "3040219287034385", "6030519247435394", "4010415217030393", "2000017237835382", "6070610227433301", //"","",/*广告位映射-测试专用聚合广告位id-0629*/

			"3070015280274558", "9030918280275682", "4040614240972686", "4020911220272700", "6030310250276526", "8050118200573680", "8070017270279684", "3050011270072658", "1080617277739350", //"", /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			"3070015280274558", "6030310250276526", "1080617277739350", //"","","","","","","",/*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/

			//"","","","","","","","","","",/*空的10个用于批量添加时复制*/
		}
		arrySelfid := [...]string{
			"28557681", "28759613", "88550654", "88955698", "58558615", "98159646", "48459677", "68352612", "08956689", "11902026",
			"51602097", "91001412", "91106012", "90096722", "40499703", "80198754", "80691705", "60897716", "80699758", "31601150", /*AYang广告平台聚合id列表160707*/

			"47132318", "17333387", "77333366", "87034385", "47435394", "17030393", "37835382", "27433301", //"","",/*广告位映射-测试专用聚合广告位id-0629*/

			"80274558", "80275682", "40972686", "20272700", "50276526", "70573680", "70279684", "70072658", "77739350", //"", /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			"80274558", "50276526", "77739350", //"","","","","","","",/*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/

			//"","","","","","","","","","",/*空的10个用于批量添加时复制*/
		}
		arryTyp := [...]uint32{
			4, 2, 1, 8, 1, 1, 1, 4, 8, 8,
			8, 8, 8, 1, 2, 4, 4, 1, 2, 8, /*广告位映射-AYang广告平台聚合id列表160707*/

			8, 4, 2, 1, 8, 4, 2, 1, /*广告位映射-测试专用聚合广告位id-0629*/

			8, 8, 8, 8, 4, 4, 4, 4, 8, /*广告位映射-TAD广告平台（ADX）广告位申请表（铃声多多）-附广告位id-0628*/

			8, 4, 8, /*广告位映射-TAD广告平台广告位（多多）-修改绑定id-0715.xlsx*/
		}
		for i, s := range arryAppid {
			appid = s
			slotid = arrySlotid[i]
			selfid = arrySelfid[i]
			typ = arryTyp[i]
			_start_marshal(selfid, name, appid, slotid, typ, wt)
			if IS_DEBUG {
				fmt.Printf("%s\t%s\t%s\t%s\t%d\n", selfid, name, appid, slotid, typ)
			}
		}
	}
	os.Exit(0)
}
