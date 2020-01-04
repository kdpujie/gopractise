package main

import "C"
import (
	"fmt"
	"learn.com/tplugins/entry"
)

type greetingEng string

func (g greetingEng) Greet(user *entry.User) {
	fmt.Printf("%s: Hello Universe \n", user.Name)
	fmt.Printf("age=%d, company=%s \n", user.Age, user.Company)
	user.Company = "ksyun"
}

// exported
//var Greeter greeting
var Greeter greetingEng
