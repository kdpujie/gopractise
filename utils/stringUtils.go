package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"ksyun.com/commons/util"
)

func main()  {
	testSplit("4~9")
	fmt.Printf("本机IP是：%s \n",util.GetLocalIpByTcp())
	s := "0.0"
	section := "0~0"
	re := isInSectionWithLimit(s,section,"~",5)
	fmt.Printf("匹配结果：%v \n",re)
	testSubstr()
	testContains()
}

func testSplit(str string)  {
	values := strings.Split("/thrift-rpc/BlacklistServer/providers", "/")
	for i,v := range values {
		fmt.Printf("splist: value[%d]=%s \n",i,v)
	}
}

func startStrings() {
	str := `a bc d src="{CLICK_URL}"`
	writeString()
	fmt.Println(strings.TrimSpace(str))
	str = strings.Replace(str, "{CLICK_URL}", "http://baidu.com/", 1)
	fmt.Println(str)
	fmt.Println("格式化字符串:" + fmt.Sprintf("1\t %d", 333))
	number_util()
	var ss []string = make([]string, 0, 10)
	ss = append(ss, "a", "b", "c")
	fmt.Printf("addr:%p \t len:%v content:%v \n", ss, len(ss), ss)
	testSlice(&ss)
	fmt.Printf("addr:%p \t len:%v content:%v \n", ss, len(ss), ss)
	s1 := strings.Split("3", ",")
	fmt.Println(s1[len(s1)-1:][0])
	fmt.Println("slice:")
	fmt.Println(ss)
	string_join()
	data, _ := json.Marshal(ss)
	fmt.Printf("字符切片序列化后:%s \n", string(data))
	str1 := replaceKscdn2https("http://kscdn.ksyun.com/a/badfaf.jpg")
	fmt.Printf("修改https后,连接地址为:%s \n", str1)
	str2 := "	1 "
	fmt.Printf("str是否为空串:%d", IsBlank(str2))
	testContains()
}
func string_join() {
	str := []string{"Hello", "World", "Good"}
	fmt.Println(strings.Join(str, ","))

	for i := len(str) - 1; i > -1; i-- {
		if i == 1 {
			str = append(str[:i], str[i+1:]...)
		} else {
			fmt.Printf("str-%d = %s \n", i, str[i])
		}
	}
}

func number_util() {
	var f float64 = 1.3
	var d64 int64 = 32423423
	ds := strconv.FormatInt(d64, 10)
	ff := strconv.FormatFloat(f, 'f', 2, 32)
	fmt.Println("1.3转换成字符串:" + ff)
	fmt.Println("32423423转换成字符串:" + ds)
}

func writeString() {
	b := bytes.Buffer{}
	b.WriteString("a")
	b.WriteString("b")
	b.WriteString("c")
	b.WriteString("d")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("1")
	b.WriteString("2")
	b.WriteString("3")
	b.WriteString("4")
	b.WriteByte(byte(6))
	fmt.Println("字符串相加: " + b.String())
	v := tabSeparateParams("3", "17b95790c506623af4ba09036ca58753", "c2cf9021-73db-11e6-8e50-507b9dade6b1", "oobk2q80", "epw9kos9", "275",
		"asdfasdf", "1", "866926020248380")
	fmt.Println(v)

}

func tabSeparateParams(args ...string) string {
	buffer := &bytes.Buffer{}
	for index, arg := range args {
		buffer.WriteString(arg)
		if index != len(args)-1 {
			buffer.WriteString("|")
		}
	}
	return buffer.String()
}
func testSlice(ss *[]string) {

	*ss = append(*ss, "z")
	fmt.Printf("addr:%p \t len:%v content:%v \n", ss, len(*ss), ss)
	//ss[0] = "f"
}

//把kscdn.ksyun.com域名改为https方式访问.
func replaceKscdn2https(src string) (dir string) {
	if strings.Contains(src, "kscdn.ksyun.com") {
		dir = strings.Replace(src, "http:", "https:", 1)
	}
	return dir
}

//判断字符串是否为空串
func IsBlank(str string) bool {
	str = strings.TrimSpace(str)
	if len(str) < 1 {
		return true
	}
	return false
}

//判断给定的数字，是否包含在指定section中。
//如果section的上边界大于指定值upBoundary, 则表示不设上限
func isInSectionWithLimit(s, section, sep string,upBoundary float64)  bool {
	vs := strings.Split(section, sep)
	if len(vs) != 2 {
		return false
	}
	if start, err := strconv.ParseFloat(vs[0],64);  err == nil {

		if end, err := strconv.ParseFloat(vs[1],64);  err == nil {
			fmt.Printf("s=%s,section=%s,sep=%s,upBoundary=%f,start=%f ,end=%f \n",s,section,sep,upBoundary,start,end)
			if mid ,err := strconv.ParseFloat(s,64); err == nil {
				return mid >=start && (mid <= end || end > upBoundary)
			}
		}
	}
	return false
}

func testSubstr()  {
	str := "p=dsp&s=%7B%22down_x%22%3A%22-999%22%2C%22down_y%22%3A%22-999%22%2C%22up_x%22%3A%22-999%22%2C%22up_y%22%3A%22-999%22%7D&v=1&z=aHR0cDovL2JhaWR1LmNvbQ=="
	i := strings.Index(str,"&z")
	fmt.Printf("index=%d \n",i)
	fmt.Printf("%s \n",str[:i])
}

func testContains() {
	now := time.Now()
	now = now.AddDate(0,0,2)
	weekday := int(now.Weekday())

	var timeInterval int32
	timeInterval = int32(10 + weekday * 24)
	base :=  weekday * 24
	var ss string = "101011111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	bytes := []byte(ss)
	bytes[0] = 48
	fmt.Println(strings.Contains(ss, "1"))
	fmt.Printf("now.hour=%d,weekday=%d,day=%d,timeInterval=%d, ss[%d]=%d \n", now.Hour(), weekday, now.Day(), timeInterval, timeInterval,ss[timeInterval])
	fmt.Printf("今天的投放时段：len(ss)=%d, %s \n",len(ss) ,ss[base:base+24])
	elapsedTimeRate(ss[base:base+24], now.Hour())
}

func elapsedTimeRate(src string, index int)  {
	traversal(src)
	total,pre := calculateOnes(src, index)
	t := float64(time.Duration(total) * time.Hour)
	p := float64(time.Duration(pre) * time.Hour)
	p = p + currentHourMili()
	rate := p/t
	fmt.Printf("总时间%f,已过时间%f, 投放时间已过比率：%f \n", t, p, rate)
}

func traversal (src string) {
	for i, v := range src {
		fmt.Printf("index=%d, value=%s \n",i, v)
	}
}

func calculateOnes(src string, index int) (total, indexPre int) {
	for i,s := range src {
		if s == 49 {
			total ++
			if i < index {
				indexPre ++
			}
		}
	}
	return total, indexPre
}

func currentHourMili() float64 {
	now := time.Now()
	hourStart := time.Date(now.Year(),now.Month(),now.Day(),now.Hour(),0,0,0, now.Location())
	return float64(now.Sub(hourStart))
}