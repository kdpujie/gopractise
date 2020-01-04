package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"taoche.com/aps/arch/kconsumers/persist/vo"
)

func main() {
	rf, err := os.Open("config/vehicle.txt")
	if err != nil {
		fmt.Println("file open failed.", err)
	}
	defer rf.Close()
	scanner := bufio.NewScanner(rf)
	temp := &bytes.Buffer{}
	for scanner.Scan() {
		line := scanner.Text()
		temp.WriteString(line)
	}
	v := vo.VehicleVO{}
	err = json.Unmarshal([]byte(temp.String()), &v)
	if err != nil {
		fmt.Println("json.Unmarshal.", err)
		os.Exit(1)
	}
	fmt.Println(v.VehicleInfo)
}
