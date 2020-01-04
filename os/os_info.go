/**
@description	获取主机名称
@author pujie
@data	2018-01-18
**/

package main

import (
	"fmt"
	"os"
)

func main() {
	h, err := os.Hostname()
	if err != nil {
		fmt.Printf("os.hostname() err: %s \n", err.Error())
	} else {
		fmt.Printf("hostname=%s \n", h)
	}
}
