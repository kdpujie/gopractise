package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"learn.com/gopractise/pb/baidu"
)

/***
移动ssp_api测试
**/
func StartBaiduMobApi() {
	//	request := generateAndroidRequest()
	request := generateIosRequest()
	//序列化
	data, err := proto.Marshal(request)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}
	//log.Println("1. 序列化后:", request.String())
	//log.Println("1. 序列化后:", data)
	//发送
	//	response, err := http.Post("http://debug.mobads.baidu.com/api_5", "application/octet-stream", bytes.NewBuffer(data))
	//response, err := http.Post("http://123.59.14.199:8889/api/def", "application/octet-stream", bytes.NewBuffer(data))
	response, err := http.Post("http://111.202.114.80/api_5", "application/octet-stream", bytes.NewBuffer(data))
	//response, err := http.Post("http://localhost:8080/api/def", "application/octet-stream", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("request err:", err)
		return
	}
	log.Println("1.status code = ", response.StatusCode)
	data, err = ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal("response err:", err)
		return
	}
	res := &baidu.MobadsResponse{}
	err = proto.Unmarshal(data, res)
	if err != nil {
		log.Fatal("Unmarshal error:", err)
	}

	readResponse(res)

}

//读取广告信息
func readResponse(response *baidu.MobadsResponse) {
	log.Println("1.requestid=", response.GetRequestId())
	log.Println("2.error_code=", response.GetErrorCode())
	log.Println("3.expiration_time=", response.GetExpirationTime())
	for _, ad := range response.GetAds() {
		log.Println("4.slot_id = ", ad.GetAdslotId())
		log.Println("5.ad_key = ", ad.GetAdKey())
		log.Println("6.html_snippet = ", ad.GetHtmlSnippet())
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
		log.Println("22.mob_adtext=", ad.GetMobAdtext())
		log.Println("23.mob_adlogo=", ad.GetMobAdlogo())
	}
}

/***
ks request生成
**/
func generateAndroidRequest() *baidu.MobadsRequest {
	//slot info
	slotSize := &baidu.Size{
		Width:  proto.Uint32(360),
		Height: proto.Uint32(300),
	}
	slot := &baidu.AdSlot{
		AdslotId:   proto.String("3137381"),
		AdslotSize: slotSize,
	}
	//app info
	app := &baidu.App{
		AppId:      proto.String("bc1d2ec3"),
		AppVersion: &baidu.Version{Major: proto.Uint32(3), Minor: proto.Uint32(4)},
	}
	//device info
	uId := &baidu.UdId{
		//Idfa:      proto.String("FC0F3445-0FCE-40EE-8646-3CA8BB2663EA"),
		Mac:       proto.String("12:34:56:78:90:ab"),
		Imei:      proto.String("3Bd52A121ba64c8"),
		AndroidId: proto.String("3fd5h2fg2sgr64h4"),
	}
	deviceSize := &baidu.Size{
		Width:  proto.Uint32(1080),
		Height: proto.Uint32(1920),
	}
	deviceType := baidu.Device_PHONE
	osType := baidu.Device_ANDROID
	device := &baidu.Device{
		DeviceType: &deviceType,
		OsType:     &osType,
		OsVersion:  &baidu.Version{Major: proto.Uint32(6), Minor: proto.Uint32(0)},
		Vendor:     []byte("MEIZU"),
		Model:      []byte("MX5"),
		Udid:       uId,
		ScreenSize: deviceSize,
	}
	//network info
	connType := baidu.Network_WIFI
	operatorType := baidu.Network_CHINA_MOBILE
	network := &baidu.Network{
		Ipv4:           proto.String("114.255.44.132"),
		ConnectionType: &connType,
		OperatorType:   &operatorType,
	}
	//gps info
	coordinateType := baidu.Gps_WGS84
	gps := &baidu.Gps{
		CoordinateType: &coordinateType,
		Longitude:      proto.Float64(40.7127),
		Latitude:       proto.Float64(74.0059),
		Timestamp:      proto.Uint32(123456),
	}

	request := &baidu.MobadsRequest{
		RequestId:  proto.String("PlCe8PBnrK3sUZ1mmC0gkbyUP4j8XfZp"),
		ApiVersion: &baidu.Version{Major: proto.Uint32(5), Minor: proto.Uint32(0)},
		Adslot:     slot,
		App:        app,
		Device:     device,
		Network:    network,
		Gps:        gps,
		IsDebug:    proto.Bool(true),
	}
	return request
}

/***
ks request生成
**/
func generateIosRequest() *baidu.MobadsRequest {
	//slot info
	slotSize := &baidu.Size{
		Width:  proto.Uint32(360),
		Height: proto.Uint32(300),
	}
	slot := &baidu.AdSlot{
		AdslotId:   proto.String("2563829"),
		AdslotSize: slotSize,
	}
	//app info
	app := &baidu.App{
		AppId:      proto.String("e0884746"),
		AppVersion: &baidu.Version{Major: proto.Uint32(3), Minor: proto.Uint32(4)},
	}
	//device info
	uId := &baidu.UdId{
		Idfa: proto.String("FC0F3445-0FCE-40EE-8646-3CA8BB2663EA"),
		Mac:  proto.String("12:34:56:78:90:ab"),
	}
	deviceSize := &baidu.Size{
		Width:  proto.Uint32(1080),
		Height: proto.Uint32(1920),
	}
	deviceType := baidu.Device_PHONE
	osType := baidu.Device_IOS
	device := &baidu.Device{
		DeviceType: &deviceType,
		OsType:     &osType,
		OsVersion:  &baidu.Version{Major: proto.Uint32(4), Minor: proto.Uint32(2)},
		Vendor:     []byte("IPHONE"),
		Model:      []byte("IPHONE4S"),
		Udid:       uId,
		ScreenSize: deviceSize,
	}
	//network info
	connType := baidu.Network_WIFI
	operatorType := baidu.Network_CHINA_MOBILE
	network := &baidu.Network{
		Ipv4:           proto.String("114.255.44.132"),
		ConnectionType: &connType,
		OperatorType:   &operatorType,
	}
	//gps info
	coordinateType := baidu.Gps_WGS84
	gps := &baidu.Gps{
		CoordinateType: &coordinateType,
		Longitude:      proto.Float64(40.7127),
		Latitude:       proto.Float64(74.0059),
		Timestamp:      proto.Uint32(123456),
	}

	protocalType := baidu.MobadsRequest_HTTPS_PROTOCOL_TYPE
	request := &baidu.MobadsRequest{
		RequestId:           proto.String("PlCe8PBnrK3sUZ1mmC0gkbyUP4j8XfZp"),
		ApiVersion:          &baidu.Version{Major: proto.Uint32(5), Minor: proto.Uint32(0)},
		Adslot:              slot,
		App:                 app,
		Device:              device,
		Network:             network,
		Gps:                 gps,
		IsDebug:             proto.Bool(true),
		RequestProtocolType: &protocalType,
	}
	return request
}
