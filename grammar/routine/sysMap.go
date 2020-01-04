/**
@description 测试golang 1.9版本提供的sys.Map
@author 蒲杰
@date	2017-09-25
 **/
package main

import (
	"sync"
	"fmt"
)

func main()  {
	normalMap := map[string]interface{}{
		"name":          "田馥甄",
		"birthday":      "1983年3月30日",
		"age":           34,
		"hobby":         []string{"听音乐", "看电影", "电视", "和姐妹一起讨论私人话题"},
		"constellation": "白羊座",
	}
	var cMap sync.Map
	fmt.Printf("cMap是否有内容：%v \n", hasContainkey(cMap))
	for key, value := range normalMap { //初始化
		cMap.Store(key,value)
	}
	fmt.Printf("cMap是否有内容：%v \n", hasContainkey(cMap))
	//遍历
	cMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//查找
	_, ok := cMap.Load("name")
	fmt.Printf("是否存在key=name,result=%v \n", ok)
	_, ok  = cMap.Load("sex")
	fmt.Printf("是否存在key=sex,result=%v \n",ok)
	num, ok := cMap.Load("intKey")
	//n := num.(int)
	fmt.Printf("整形的默认值: %d, %v \n",num, ok)
}

func hasContainkey (sMap sync.Map) bool {
	flag := false
	sMap.Range(func(key, value interface{}) bool {
		flag = true
		fmt.Println("遍历。。。。")
		return false
	})
	return flag
}