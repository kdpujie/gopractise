/**
	Type:  通过调用TypeOf获取其动态类型信息，该函数返回一个Type类型值
	Value: 调用ValueOf函数返回一个Value类型值，该值代表运行时的数据
参考资料：
	[浅谈Go语言中的结构体struct & 接口Interface & 反射](https://mp.weixin.qq.com/s/qXsQC1TmJzcbMNxFjeS_Pg)
**/
package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	//reflect.Kind() //type枚举值
	//baseType()
	testSearch()
}

func baseType() {
	var a, b = 3.14, 520.00
	fmt.Println("type:", reflect.TypeOf(a)) //type: float64
	av := reflect.ValueOf(&a)
	fmt.Println("value = ", av)                      // 0xc420014050
	fmt.Println("av.type = ", av.Type())             // *float64
	fmt.Println("av.elem.type = ", av.Elem().Type()) // float64
	fmt.Println("av.kind = ", av.Kind())             // ptr
	fmt.Println("av.interface = ", av.Interface())   // 0xc420014050

	av.Elem().SetFloat(b)
	fmt.Println("av.new.value = ", av.Elem()) // 520
}

func testCopy() {
	index := 2
	s1 := []string{"a", "b", "c", "d", "e"}
	s2 := s1[index+1:]
	s3 := s1[index:]
	fmt.Printf("line1 :s1=%v; s2 = %v; s3=%v \n", s1, s2, s3)
	copy(s2, s3)
	fmt.Printf("line2: after copy(s2, s3), s2 = %v \n", s2)
	fmt.Printf("line3: s1 = %v \n", s1)
}

func testSearch() {
	a := []int{10, 11, 13, 14, 15, 19, 21}
	b := sort.Search(len(a), func(i int) bool { return a[i] >= 30 })
	fmt.Printf("line1: index=%d\n", b)
	b1 := sort.Search(len(a), func(i int) bool { return a[i] > 18 })
	fmt.Printf("line2: index=%d\n", b1)
	// 升序排列，必须使用> or >=; 降序使用：< or =<
	b2 := sort.Search(len(a), func(i int) bool {
		fmt.Printf("a[%d] = %d \n", i, a[i])
		return a[i] == 13
	})
	fmt.Printf("line3: index=%d\n", b2)
	b3 := sort.Search(len(a), func(i int) bool { return a[i] <= 11 })
	fmt.Printf("line4: index=%d\n", b3)
	b4 := sort.Search(len(a), func(i int) bool { return a[i] < 10 })
	fmt.Printf("line5: index=%d\n", b4)
}
