package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"practise.com/learn/pb/ks"

	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/util"
)

var client *http.Client = util.DefaultClient

/***
移动ssp_api测试
**/
func main() {
	request := generateKsAndroidRequest()
	//request := generateKsIosRequest()
	//序列化
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	//log.Println("1. 序列化后:", request.String())
	//log.Println("1. 序列化后:", data)
	//发送  http://kspi.ksyun.com/api/json   http://kspi.ksyun.com/api/def
	//baiduR, err := http.NewRequest("POST", "http://192.168.153.8:8080/api/def", bytes.NewBuffer(data))
	//baiduR, err := http.NewRequest("POST", "http://120.92.44.245/api/def", bytes.NewBuffer(data))
	//	baiduR, err := http.NewRequest("POST", "http://api.ssp.ksyun.com/api/def", bytes.NewBuffer(data))
	baiduR, err := http.NewRequest("POST", "http://kspi.ksyun.com/api/json", bytes.NewBuffer(data))
	//baiduR, err := http.NewRequest("POST", "http://debug.mobads.baidu.com/api_5", bytes.NewBuffer(data)) //百度
	//baiduR, err := http.NewRequest("POST", "http://mobads.baidu.com/api_5", bytes.NewBuffer(data)) //百度
	//baiduR, err := http.NewRequest("POST", "http://localhost:8080/api/def", bytes.NewBuffer(data)) //本地
	//baiduR, err := http.NewRequest("POST", "http://123.59.14.199:8084/api/test/9", bytes.NewBuffer(data))
	baiduR.Header.Add("Host", "mi.gdt.qq.com")
	baiduR.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	baiduR.Header.Add("Upgrade-Insecure-Requests", "1")
	baiduR.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.101 Safari/537.36")
	baiduR.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
	baiduR.Header.Add("Accept-Language", "zh-CN,zh;q=0.8")
	baiduR.Header.Add("Connection", "keep-alive")
	baiduR.AddCookie(&http.Cookie{Name: "uin", Value: "o2014861221"}) //填写后获取成功率大大提高(生产时需要删除！！！！)
	response, err := client.Do(baiduR)

	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	header := response.Header
	for key, value := range header {
		log.Println(key, value)
	}
	log.Println("1.status code = ", response.StatusCode)

	//gzip处理@2016-7-11 by chenbintao
	if response != nil {
		defer response.Body.Close()
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			data = make([]byte, 0, 1024)
			reader, err := gzip.NewReader(response.Body)
			if err != nil {
				log.Fatal("获取gzipReader失败,", err)
			}
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)
				if err != nil && err != io.EOF {
					panic(err)
				}
				if n == 0 {
					break
				}
				buf = buf[:n]
				data = append(data, buf...)
			}
			data = data[:len(data)]
		default:
			data, err = ioutil.ReadAll(response.Body)
		}
		//			log.Println(string(data))
	}

	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	res := &ks.MobadsResponse{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	readKsResponse(res)

}

//读取广告信息
func readKsResponse(response *ks.MobadsResponse) {
	log.Println("1.requestid=", response.GetRequestId())
	log.Println("2.error_code=", response.GetErrorCode())
	log.Println("3.expiration_time=", response.GetExpirationTime())
	for _, ad := range response.GetAds() {
		ad.GetMetaGroup()[0].GetInteractionType()
		log.Println("4.slot_id = ", ad.GetAdslotId())
		log.Println("5.ad_key = ", ad.GetAdKey())
		log.Println("6.html_snippet = ", string(ad.GetHtmlSnippet()))
		log.Println("******************************start tracking********************************")
		trackings := ad.GetAdTracking()
		for i, tracking := range trackings {
			log.Printf("\t%d tracking_event=%s", i, tracking.GetTrackingEvent().String())
			log.Printf("\t%d tracking_url=%s", i, tracking.GetTrackingUrl())
		}
		log.Println("******************************end tracking信息********************************")
		/**
		log.Println("******************************start meta信息********************************")
		meta := ad.GetMaterialMeta()
		log.Println("7.creative_type=", meta.GetCreativeType())
		log.Println("8.interaction_type=", meta.GetInteractionType())
		log.Println("9.win_notice_url=", meta.GetWinNoticeUrl())
		log.Println("10.click_url=", meta.GetClickUrl())
		log.Println("11.title = ", string(meta.GetTitle()))
		log.Println("12.brand_name=", meta.BrandName)
		for i, des := range meta.GetDescription() {
			log.Printf("13.%d description=%s", i, string(des))
		}

		log.Println("14.icon_src=", meta.GetIconSrc())
		log.Println("15.image_src=", meta.GetImageSrc())
		log.Println("16.app_package=", meta.GetAppPackage())
		log.Println("17.app_size=", meta.GetAppSize())
		log.Println("18.video_url=", meta.GetVideoUrl())
		log.Println("19.meta_index=", meta.GetMetaIndex())
		log.Println("20.material_width=", meta.GetMaterialWidth())
		log.Println("21.material_height=", meta.GetMaterialHeight())
		log.Println("******************************end meta信息********************************")
		**/
		log.Println("******************************start meta group信息********************************")
		groups := ad.GetMetaGroup()
		for i, group := range groups {
			log.Printf("\t%d creative_type=%s", i, group.GetCreativeType())
			log.Printf("\t%d interaction_type=%s", i, group.GetInteractionType())
			log.Printf("\t%d win_notice_url=%s", i, group.GetWinNoticeUrl())
			log.Printf("\t%d click_url=%s", i, group.GetClickUrl())
			log.Printf("\t%d title=%s", i, group.GetTitle())
			log.Printf("\t%d brand_name=%s", i, group.GetBrandName())
			log.Printf("\t%d description=%s", i, group.GetDescription())
			log.Printf("\t%d icon_src=%s", i, group.GetIconSrc())
			log.Printf("\t%d image_src=%s", i, group.GetImageSrc())
			log.Printf("\t%d vidio_src=%s", i, group.GetVideoUrl())
			log.Printf("\t%d app_size=%d", i, group.GetAppSize())
			log.Printf("\t%d material_width=%d", i, group.GetMaterialWidth())
			log.Printf("\t%d material_height=%d", i, group.GetMaterialHeight())

		}
		log.Println("******************************end meta group信息********************************")
		log.Println("5.ad_key = ", ad.GetAdKey())
	}
}

/***
 request生成
**/
func generateKsAndroidRequest() *ks.MobadsRequest {
	//slot info
	//	slotSize := &ks.Size{
	//		Width:  proto.Uint32(0),
	//		Height: proto.Uint32(0),
	//	}
	slot := &ks.AdSlot{
		AdslotId: proto.String("jinleuqc"), //47435394(原生)	27433301(横幅) 11902026(wiffi万能钥匙-信息流)
		//AdslotSize: slotSize,
		//Ads: proto.Uint32(2),
	}
	//app info
	app := &ks.App{
		AppId:      proto.String("qxidp2e6"),
		AppVersion: &ks.Version{Major: proto.Uint32(2), Minor: proto.Uint32(1)},
		AppPackage: proto.String("com.nd.android.pandahome2"),
		ChannelId:  proto.String("dd"),
	}
	//device info
	uId := &ks.UdId{
		//Idfa:      proto.String("FC0F3445-0FCE-40EE-8646-3CA8BB2663EA"),
		//Mac:       proto.String("38:bc:1a:ff:f0:9f"),
		Imei: proto.String("866912032434185"),
		//AndroidId: proto.String("eedc1f100ee44ee8"),
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
		OsVersion:  &ks.Version{Major: proto.Uint32(5), Minor: proto.Uint32(0), Micro: proto.Uint32(2)},
		Vendor:     []byte("MI"),
		Model:      []byte("m1 note"),
		Udid:       uId,
		ScreenSize: deviceSize,
	}
	//network info
	connType := ks.Network_WIFI
	operatorType := ks.Network_CHINA_MOBILE
	network := &ks.Network{
		Ipv4:           proto.String("123.116.62.221"),
		ConnectionType: &connType,
		OperatorType:   &operatorType,
	}
	//gps info
	/**
	coordinateType := ks.Gps_WGS84
	gps := &ks.Gps{
		CoordinateType: &coordinateType,
		Longitude:      proto.Float64(40.7127),
		Latitude:       proto.Float64(74.0059),
		Timestamp:      proto.Uint32(123456),
	}**/
	request := &ks.MobadsRequest{
		RequestId:  proto.String("17d95790c506623af4ba09036ca58355"),
		ApiVersion: &ks.Version{Major: proto.Uint32(5), Minor: proto.Uint32(3), Micro: proto.Uint32(0)},
		Adslot:     slot,
		App:        app,
		Device:     device,
		Network:    network,
		//Gps:        gps,
		IsDebug: proto.Bool(false),
	}
	return request
}

/***
 request生成
**/
func generateKsIosRequest() *ks.MobadsRequest {
	//slot info
	slotSize := &ks.Size{
		Width:  proto.Uint32(800),
		Height: proto.Uint32(480),
	}
	slot := &ks.AdSlot{
		AdslotId:   proto.String("mfyrrfeb"),
		AdslotSize: slotSize,
	}
	//app info
	app := &ks.App{
		AppId:      proto.String("j0g2phg7"),
		AppVersion: &ks.Version{Major: proto.Uint32(3), Minor: proto.Uint32(4)},
	}
	//device info
	uId := &ks.UdId{
		Idfa: proto.String("FC0F3445-0FCE-40EE-8646-3CA8BB2663EA"),
		Mac:  proto.String("12:34:56:78:90:ab"),
	}
	deviceSize := &ks.Size{
		Width:  proto.Uint32(414),
		Height: proto.Uint32(736),
	}
	deviceType := ks.Device_PHONE
	osType := ks.Device_IOS //之前代码为安卓，需采用IOS
	device := &ks.Device{
		DeviceType: &deviceType,
		OsType:     &osType,
		OsVersion:  &ks.Version{Major: proto.Uint32(3), Minor: proto.Uint32(1)},
		Vendor:     []byte("IPHONE"),
		Model:      []byte("IPHONE4S"),
		Udid:       uId,
		ScreenSize: deviceSize,
	}
	//network info
	connType := ks.Network_WIFI
	operatorType := ks.Network_CHINA_MOBILE
	network := &ks.Network{
		Ipv4:           proto.String("10.10.20.60"),
		ConnectionType: &connType,
		OperatorType:   &operatorType,
	}
	//gps info
	//	coordinateType := ks.Gps_WGS84
	//	gps := &ks.Gps{
	//		CoordinateType: &coordinateType,
	//		Longitude:      proto.Float64(40.7127),
	//		Latitude:       proto.Float64(74.0059),
	//		Timestamp:      proto.Uint32(123456),
	//	}
	request := &ks.MobadsRequest{
		RequestId:  proto.String("PlCe8PBnrK3sUZ1mmC0gkbyUP4j8XfZr"),
		ApiVersion: &ks.Version{Major: proto.Uint32(5), Minor: proto.Uint32(0), Micro: proto.Uint32(0)},
		Adslot:     slot,
		App:        app,
		Device:     device,
		Network:    network,
		//Gps:        gps,
		IsDebug: proto.Bool(true),
	}
	return request
}
