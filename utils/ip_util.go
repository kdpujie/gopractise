package main

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
	"learn.com/tools"
	"ksyun.com/commons/util"
)

//ip段
type IpSegment struct {
	startIp  uint64 //ip段开始ip地址
	endIp    uint64 //IP段结束ip地址
	cityCode uint64 //ip段对应城市编码
	cityName string
}

type IpLib []*IpSegment

func (list IpLib) Len() int {
	return len(list)
}

//排序规则：按startIp升序排列
func (list IpLib) Less(i, j int) bool {
	if list[i].startIp <= list[j].startIp {
		return true
	} else {
		return false
	}
}

func (list IpLib) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func main()  {
	tools.StartCPUProfile()
	defer tools.StopCPUProfile()
	startIpTest()
	tools.SaveMemProfile()
	tools.SaveBlockProfile()
	tools.SaveGroutineProfile()
	tools.SaveThreadProfile()
	tools.SaveHeapProfile()
	time.Sleep(4 * time.Second)
}

func startIpTest() {
	ips := []string{"223.104.38.11"}
	bMap := initArea()

	ipLib, _ := loadIpLib()
	sort.Sort(ipLib)
	start := time.Now()
	for _, ip := range ips {
		ipLong, err := Ip2Long(ip)
		if err != nil {
			fmt.Println("ip2long error")
		} else {
			//			fmt.Printf("ip %s 转化为数字:%d \n", ip, ipLong)
		}
		ipSegment := binarySearch(ipLib, ipLong)
		//		bMap.get(ipSegment.cityCode)
		fmt.Printf("%s \t startip=%s,endip=%s,cityCode=%d,cityName=%s,isContains=%t \n", ip, Long2Ip(ipSegment.startIp), Long2Ip(ipSegment.endIp), ipSegment.cityCode, ipSegment.cityName, bMap.Get(ipSegment.cityCode))
	}
	//	time.Sleep(1 * time.Second)
	end := time.Now()
	var dur_time time.Duration = end.Sub(start)
	var elapsed_nano int64 = dur_time.Nanoseconds()
	fmt.Printf("匹配耗时:%d \n", elapsed_nano)
	//	printIpLib()
}

func initArea() *util.BitSet {
	bMap := util.NewBitMap(8)
	bMap.Set(362)
	bMap.Set(1390511)
	for i := 1; i < 151; i++ {
		bMap.Set(uint64(i))
	}
	return bMap
}

func printIpLib() {
	ipLib, _ := loadIpLib()
	sort.Sort(ipLib)
	var i int
	for _, v := range ipLib {
		i++
		fmt.Printf("start=%d,end=%d,cityCode=%X \n", v.startIp, v.endIp, v.cityCode)
	}
	fmt.Println("共初始化数据:", i)
}

//二分查找ip所属ip段.
func binarySearch(ipLib IpLib, ipLong uint64) *IpSegment {
	begin, end := 0, len(ipLib)-1
	if ipLong < ipLib[begin].startIp || ipLong > ipLib[end].endIp {
		return &IpSegment{
			startIp:  ipLong,
			endIp:    ipLong,
			cityCode: 371, //未识别
		}
	}
	var mid int
	for begin <= end {
		mid = (begin + end) / 2
		if ipLib[mid].startIp > ipLong {
			end = mid - 1
		} else if ipLib[mid].endIp < ipLong {
			begin = mid + 1
		} else {
			return ipLib[mid]
		}

	}
	return &IpSegment{
		startIp:  ipLong,
		endIp:    ipLong,
		cityCode: 371, //未识别
	}
}

//load ip_lib.txt文件,初始化ip库
func loadIpLib() (IpLib, error) {
	file, err := os.Open("ip_lib_my.txt")
	if err != nil {
		fmt.Println("file open failed.", err)
	}
	defer file.Close()
	ipLib := IpLib([]*IpSegment{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colums := strings.Split(line, "\t")
		startIP, err := Ip2Long(colums[0])
		endIp, err := Ip2Long(colums[1])
		cityCode, err := strconv.ParseUint(colums[8], 10, 64)
		if err == nil {
			var ipSegement *IpSegment = &IpSegment{}
			ipSegement.startIp = startIP
			ipSegement.endIp = endIp
			ipSegement.cityCode = cityCode
			ipSegement.cityName = colums[4]
			ipLib = append(ipLib, ipSegement)
		}
	}
	return ipLib, scanner.Err()
}

//点分十进制ip转化为数字
func Ip2Long(ip string) (uint64, error) {
	ipSeg := strings.Split(ip, ".")
	if len(ipSeg) != 4 {
		return 0, errors.New("ip format error!")
	}
	b0, err := strconv.ParseUint(ipSeg[0], 10, 64)
	b1, err := strconv.ParseUint(ipSeg[1], 10, 64)
	b2, err := strconv.ParseUint(ipSeg[2], 10, 64)
	b3, err := strconv.ParseUint(ipSeg[3], 10, 64)
	if err != nil {
		return 0, err
	}
	return (b0<<24 | b1<<16 | b2<<8 | b3), nil
}

//ip的数字形式转化为点分十进制
func Long2Ip(ipLong uint64) string {
	temp := &bytes.Buffer{}
	temp.WriteString(strconv.FormatUint(ipLong>>24, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint((ipLong&0x00FFFFFF)>>16, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint((ipLong&0x0000FFFF)>>8, 10))
	temp.WriteString(".")
	temp.WriteString(strconv.FormatUint(ipLong&0x000000FF, 10))
	return temp.String()
}
