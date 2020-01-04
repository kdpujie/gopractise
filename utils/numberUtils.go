package main

import (
	"fmt"
	"strconv"
	"math"
)

func main() {
	var index int
	var ads []int = make([]int,0,10)
	fmt.Printf("before:len= %d,cap=%d \n",len(ads),cap(ads))
	for k,v :=range ads {
		fmt.Printf("ads[%d]=%d  ",k,v)
	}
	for i:=0;i<10;i++{
		//ads[i] = index
		ads = append(ads,index)
		index ++

	}
	fmt.Printf("\nafter:len= %d,cap=%d \n",len(ads),cap(ads))
	for k,v :=range ads {
		fmt.Printf("ads[%d]=%d  \n",k,v)
	}
	var index01 uint16 = 9
	fmt.Printf("index value=%d, addr=%v \n",index,&index)
	fmt.Printf("index01 value=%d, addr=%v \n",index01,&index01)
	var temp = 40 / 1.5

	fmt.Printf("0 / 0 = %f,向下取整 = %f \n", temp,math.Floor(temp))
	logic()
	cRatio := math.Ceil(float64(160 / 160))
	sRatio := math.Ceil(float64(1280) / float64(720))
	fmt.Printf("cRatio=%f,sRation=%f,1280/720=%f \n",cRatio,sRatio,1280/720)
	fmt.Println("取整数:",fmt.Sprintf("%.0f",40.6))

}

func startNumber() {

	var p3 uint64 = 0x0000000010000002
	var p4 uint64 = 8589934600
	p5, _ := strconv.ParseUint("0000000010000002", 16, 64)
	fmt.Printf("p3=%d in of city %v \n", p3, containsCity(p3))
	fmt.Printf("p4=%d in of city %v \n", p4, containsCity(p4))
	fmt.Printf("p5=%d in of city %v \n", p5, containsCity(p5))

}

func containsCity(index uint64) bool {
	var p1 uint64 = 0x0000000010000002
	var p2 uint64 = 0x0000000200000008
	sum := p1 | p2
	h36 := sum & 0xfffffffff0000000
	if h36&index > 0 {
		l7 := sum & 0x000000000ffffff
		if l7&index > 0 {
			return true
		}
	}
	return false
}

func logic()  {
	num ,_ :=strconv.ParseInt("01",10,64)
	fmt.Println("逻辑与：",3 & 1)
	fmt.Printf("字符转换成数字：%d \n",num)
	fmt.Printf("逻辑或运算：1 | 2 = %d \n",1 | 2)
	fmt.Printf("平方数：2=%d \n", 1<<1)
}