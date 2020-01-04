package main

import "C"

import (
	"fmt"
	"learn.com/tplugins/entry"
)

type greetingChi string

func (g greetingChi) Greet(user *entry.User) {
	fmt.Printf("%s: 你好宇宙, 你被金山云录用了 \n", user.Name)
	fmt.Printf("age=%d, company=%s \n", user.Age, user.Company)
	user.Company = "金山云"
}
// exported
var Greeter greetingChi