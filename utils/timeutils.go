package main

import (
	"fmt"
	"strconv"
	"time"
	"learn.com/utils/common"
)

const (
	c_1 int= 1
)
var(
	v_1 int = c_1 * 2
)

func init()  {
	fmt.Printf("init(): c_1=%d, v_1=%d \n",c_1,v_1)
}

func main() {
	fmt.Println("main()")
	common.Excute()
	fmt.Printf("self.exucute: c_1=%d, v_1=%d \n",c_1,v_1)
	//时间戳
	t := time.Now().Unix()
	ts := strconv.FormatInt(t, 10)
	fmt.Println(t)
	fmt.Printf("时间戳:%d \n", t)
	fmt.Printf("时间戳转换成uint32:%d \n", uint32(t))
	fmt.Printf("时间戳转换成字符串:%s \n", ts)
	//带纳秒的时间戳
	t = time.Now().UnixNano()
	fmt.Println(t)
	fmt.Println("------------------")

	//基本格式化的时间表示
	fmt.Println(time.Now().String())
	fmt.Printf("当前的hour值:%d, %d\n \n",time.Now().Weekday(), time.Now().Hour())
	fmt.Println(time.Now().Format("2006year 01month 02day"))
	time_format()
	parse_time()
}

func time_format() {
	fmt.Printf("格式化当前时间:%s \n", time.Now().Format("20060102"))

}

func add_time() {
	t := time.Duration(2) * time.Hour
	now := time.Now().Add(t)
	fmt.Println("时间测试:", now)
}

func parse_time() {
	tm := time.Unix(1483007246, 0)
	fmt.Println("parse_time:")
	fmt.Printf("before time:%s \n",tm.Format("2006-01-02 03:04:05 PM"))
	tm1 := tm.Add(6 * time.Second)
	fmt.Printf("after time:%s \n",tm1.Format("2006-01-02 03:04:05 PM"))

}

func ini() string  {
	return ""
}