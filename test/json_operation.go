package main

import (
	"encoding/json"
	"fmt"
	"time"
	"utils/common"
)

type IdCard struct {
	Id     string
	Coutry string
	Addr   string
}

type Person struct {
	Name        string `json:"username"` //tag中的第一个参数-用来指定别名,如果不想指定别名,用逗号分隔
	Age         int
	Gender      bool `json:",omitempty"` //如果为空(数字0,字符串"",空数组[]等)值则忽略字段
	Profile     string
	OmitContent string  `json:"-"`       //直接忽略
	Count       int     `json:",string"` //指定序列化后为string类型
	Card        *IdCard //身份证信息
}

func NewPerson() *Person {
	return &Person{Name: "sss"}
}

func JsonStart() {
	jsonMarshal()
	jsonUnmarshal()
}

//序列化
func jsonMarshal() {
	var c *IdCard = &IdCard{
		Id:     "51130219861234",
		Coutry: "中国",
		Addr:   "北京市朝阳区",
	}
	var p *Person = &Person{
		Name:        "brainwu",
		Age:         21,
		Gender:      true,
		Profile:     "I am Wujunbin",
		OmitContent: "OmitContent",
		Card:        c,
	}
	start := time.Now()
	for i := 0; i < 5000; i++ {
		json.Marshal(p)
	}
	common.TimeSpend("序列化", start)
	if bs, err := json.Marshal(p); err != nil {
		panic(err)
	} else {
		fmt.Println("1.person序列化:", string(bs))
	}

	//3.slice序列化为json
	var aStr []string = []string{"Go", "Java", "Python", "Android"}
	if bs, err := json.Marshal((aStr)); err != nil {
		panic(err)
	} else {
		fmt.Println("3.slice序列化为json:", string(bs))
	}
	//4.map序列化为json
	var m map[string]string = make(map[string]string)
	m["Go"] = "No.1"
	m["Java"] = "No.2"
	m["C++"] = "No.3"
	if bs, err := json.Marshal(m); err != nil {
		panic(err)
	} else {
		fmt.Println("4.map序列化:", string(bs))
	}
}

//反序列化
func jsonUnmarshal() {
	var p Person
	//5.反序列化到struct
	var str = `{"Name":"junbin","Age":21,"Gender":true}`
	json.Unmarshal([]byte(str), &p)
	fmt.Println("5.反序列化到struct结果:\n\t Name=", p.Name, ",Age=", p.Age, ",Gender=", p.Gender)
	//6.反序列化到slice-struct
	var sp []Person
	var aJson = `[{"Name":"junbin","Age":21,"Gender":true},
				 {"Name":"蒲杰","Age":29,"Gender":false}]`
	json.Unmarshal([]byte(aJson), &sp)
	fmt.Println("6.反序列化到slice-struct结果:", sp)
	//7.反序列化成map[string] interface{}
	//var obj interface{}
	var m map[string]interface{}
	json.Unmarshal([]byte(str), &m)
	//var m map[string]interface{} = obj.(map[string]interface{})
	fmt.Println("7.反序列化为map - ", m["Name"], ":", m["Age"], ":", m["Gender"])
	//8.反序列化为slice
	var strs string = `["Go","Java","C","Php"]`
	var aStr []string
	json.Unmarshal([]byte(strs), &aStr)
	fmt.Println("8.反序列化为slice:", aStr, "len is", len(aStr))
}
