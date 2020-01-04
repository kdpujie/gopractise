package main

import "fmt"
import (
	"google.golang.org/grpc"
	"time"
)

//import "strings"

func main() {

	//NewJsonRpcSocketServer()
	//Start_Marshal()
	//Start_Unmarshal()
	var ss []string = []string{"a", "b", "c", "d"}
	ss = ss[len(ss):]
	fmt.Println(ss)
	//	strings.Index()
	keys := make([]string, len(ss))
	keys = append(keys, "e")
	for i, v := range keys {
		fmt.Println(i, "-", v)
	}
	/*******************************/
	now := time.Now()
	weekday := int(now.Weekday())
	hourOfDay := now.Hour()
	index := (weekday-1)*24 + hourOfDay - 1
	fmt.Println(index)
	var s string = "010101010101010101010101101010101010101101010101010101010101011010100101"
	fmt.Println(s[index] == 49)
	testAddr()
}

func testAddr() {
	a := grpc.Address{
		Addr:     "127.0.0.1:8080",
		Metadata: "localhost",
	}
	b := grpc.Address{
		Addr:     "127.0.0.1:8080",
		Metadata: "localhost",
	}

	fmt.Printf("a == b: %v \n", a == b)

}
