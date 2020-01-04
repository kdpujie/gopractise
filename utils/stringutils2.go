package main

import (
	"fmt"
	"ksyun.com/commons/util"
	"math/rand"
	"strings"
)

func main() {
	testSlice2()
}

func testSlice2() {
	str := []string{"a", "b", "c", "d"}
	fmt.Printf("str:%v \n", str)
	/*	str1 := append(str[:0],str[1:]...)
		fmt.Printf("str1:%v \n",str1)
		fmt.Printf("str:%v \n",str)*/
	str2 := append(str[:1], str[2:]...)
	fmt.Printf("str2:%v \n", str2)
	fmt.Printf("str:%v \n", str)
	/*	str3 := append(str[:3],str[4:]...)
		fmt.Printf("str3:%v \n",str3)*/
	testContain()
	expStr := "whitelist|2,bidprice|0,whitelist-prob|1,"
	containExp(expStr, "bidprice|")
	containExp(expStr, "whitelist|")
	containExp(expStr, "whitelist-prob|")
	fmt.Printf("testkey1 md5 =%s \n", util.Md5("testkey1"))
	label := "ul_theme"
	value := "801001,801002,801003,801004,801005,801006,801007,801008,801009,801010,801011,801012,801013,801014,801015,802001,802002,802003,802004,802005,802006,803001,803002,803003,803004,804001,804002,804003,805001,805002,805003,805004,805005,805006,806001,806002,806003,807001,807002,807003,807004,807005,808001,808002,809001,809002,809003,809004,809005,809006,809007,809008,809009,809010,809011,809012,809013,809014,809015,809016,809017,809018,809019,809020,809021,809022,809023,809024,809025,809026,809027,809028,809029,809030,809031,809032,809033,809034,809035,809036,809037,809038,809039,810001,810002,810003,810004,810005,810006,810007,810008,810009,810010,810011,810012,810013,810014,810015,810016,810017,810018,810019,810020,810021,810022,810023,810024,,812001,812002,812003,812004,812005,812006,812007,812008,812009,812010,812011,812012,812013,812014,812015,812016,812017,812018,812019,812020,812021,812022,812023,812024,812025,812026,812027,812028,812029,812030,812031,812032,812033,812034,812035,812036,812037,812038,812039,812040,812041,812042,812043,812044,812045,812046,812047,812048,812049,812050,812051,812052,812053,812054,812055,812056,812057,812058,812059,812060,812061,812062,812063,812064,812065,812066,812067,812068,812069,812070,812071,812072,812073,812074,812075,812076,812077,812078,812079,812080,812081,812082,812083,812084,812085,812086,812087,812088,812089,812090,812091,812092,812093,812094,812095,812096,812097,812098,812099,812100,812101,812102,812103,812104,812105,812106,812107,812108,812109,812110,812111,812112,812113,812114,812115,812116,812117,812118,812119,812120,812121,812122,812123,812124,812125,812126,812127,812128,812129,813001,813002,813003,813004,814001,814002,815001,815002,816001,816002,816003,816004,816005,816006,816007,816008,816009,816010,816011,816012,816013,816014,816015,816016,816017,817001,817002,817003,817004,817005,817006"
	userProfile := "pay_money:\"nul\" game_frequency:\"nul\" game_times:\"nul\" theme:\"nul\" play_type:\"nul\" standalone_type:\"nul\" game_mode:\"nul\" subject_element:\"nul\" gameplay_element:\"nul\" direction:\"nul\" visual_angle:\"nul\" regional_preference:\"nul\" picture_style:\"nul\" head_and_body:\"nul\" art_features:\"nul\" picture_type:\"nul\" screen:\"nul\" operation_mode:\"nul\" game_scene:\"nul\" network_environment:\"nul\" game_charge:\"nul\" user_appeal:\"nul\""
	fmt.Printf("UProfileTarget: lablel=%s,value=%s,ufs=%s, result=%v \n", label, value, userProfile,
		isUProfileContainLable(userProfile, label[3:], value))
	var s1 []string
	fmt.Printf("即将打印的切片大小为：%d \n", len(s1))
	s1 = append(s1, "1", "2", "3")
	fmt.Printf("切片s1[:2]：%v \n", s1[:2])
	for i, v := range s1 {
		fmt.Printf("i=%d, v=%s \n", i, v)
	}
	remove()
	fmt.Printf("trimSuffix：%s \n", strings.Replace("adsf==", "=", "", -1))
}
func testContain() {
	fmt.Printf("rand.fload32=:%d \n", rand.Float64())
}
func containExp(expStr, subStr string) {
	expStr = expStr[strings.Index(expStr, subStr):]
	start := strings.Index(expStr, "|")
	end := strings.Index(expStr, ",")
	flowId := expStr[start+1 : end]
	fmt.Printf("expStr=%s; targetExp=%s; flowId=%s; start=%d, end=%d \n",
		expStr, subStr, flowId, start, end)
}

func remove() {
	adList := "1,"
	newAdList := strings.Replace(adList, "1,", "", 1)
	fmt.Printf("%s\n%s \n", adList, newAdList)
	if newAdList == "" {
		fmt.Printf("adlist == kong \n")
	}
}

//判断ufs标签中是否包含指定的标签
func isUProfileContainLable(uProfile, label, labelValue string) bool {
	fmt.Printf("label = %s \n", label)
	if strings.Contains(uProfile, label) {
		values := strings.Split(labelValue, ",")
		fmt.Printf("\t uProfile conains lablel valus-size=%d \n", len(values))
		for _, lValue := range values {
			fmt.Println(lValue)
			if lValue != "" && strings.Contains(uProfile, lValue) {
				fmt.Printf("\t %s conains %s \n", uProfile, lValue)
				return true
			}
		}
	}
	return false
}
