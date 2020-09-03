/**
@description json实例
**/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// str := mapMarshal()
	// mapUnmarshal(str)
	marshalStruct()
}

type ServiceInfo struct {
	EndPoint string `json:"endPoint"`
	Env      string `json:"-"`
}

// map marshal
func mapMarshal() string {
	var m = make(map[string][]string, 0)
	iosVersion := []string{"7.1", "8.0", "99.0"}
	m["1"] = iosVersion
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintf(os.Stdout, "map json.marshal err=%s \n", err.Error())
		return ""
	}
	fmt.Fprintf(os.Stdout, "map marshal result: %s\n", string(data))
	return string(data)
}

func mapUnmarshal(str string) {
	var m = make(map[string][]string, 0)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Fprintf(os.Stdout, "map json.marshal err=%s \n", err.Error())
		return
	}
	for key, value := range m {
		fmt.Fprintf(os.Stdout, "key=%s, valaue=%s\n", key, value)
	}
}

func marshalStruct() {
	s := &ServiceInfo{EndPoint: "127.0.1.1", Env: "dev"}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("marshal err: %v \n", err)
	}
	fmt.Printf("%s \n", string(data))
}
