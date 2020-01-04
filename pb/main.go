package main

//import "fmt"
//import "time"

//"fmt"
//	"strings"
//配置的渠道信息
type Adaptor interface {
	RequestAd()
}

func main() {
	//	StartSDKServer() //聚合SDK测试server
	//	StartTestSize() /*百度广告尺寸测试方法（线上运行时请注释！）by hanse*/
	//setMarshal() /**批量设置广告隐射关系（线上运行时请注释！）by hanse**/
	//	start_Channel()
	//	StartKsMobApi()
	//startRedis2Map() //
	StartBaiduMobApi()
	//startMap()
	//	Start_Type_func()
	//start_marshal()
	//start_unmarshal()

	//	var str string = "ssp_adid_config_tftl_native01"
	//	s := strings.Split(str, "ssp_adid_config_")
	//	fmt.Printf("长度=%d; 内容=%s; slotId=%s \n", len(s), s, s[1])
	//	s = strings.SplitN(str, "_", -1)
	//	fmt.Printf("长度=%d; 内容=%s \n", len(s), s)
	//	s = strings.SplitAfter(str, "_")
	//	fmt.Printf("长度=%d; 内容=%s \n", len(s), s)

	//	start_random()
	/***
	var m map[string]string = make(map[string]string)
	m["a"] = "97"
	m["b"] = "98"
	m["c"] = "99"
	for k := range m {
		fmt.Println(k)
	}
	fmt.Println("  切片操作:")
	var ss []string
	//切片尾部追加元素append elemnt
	for i := 0; i < 1; i++ {
		ss = append(ss, fmt.Sprintf("s%d", i))
	}
	index := len(ss) - 1
	ss = append(ss[:index], ss[index+1:]...)
	fmt.Println("删除最后一个元素后: ", ss)

	for i := 0; i < 10; i++ {
		fmt.Println("i=", i)
	LOOP:
		for j := 0; j < 5; j++ {
			if j == 4 {
				break LOOP
			}
			fmt.Println("----j=", j)
		}

	}
	fmt.Println("时间操作")
	t := time.Now()
	fmt.Printf("时间:%d \n", t.Unix())
	fmt.Printf("今天星期%d \n", int(t.Weekday()))
	fmt.Printf("时间%d \n", t.Hour())
	**/
}
