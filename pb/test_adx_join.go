package main

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/proto"
	"ksyun.com/commons/entry"
	"ksyun.com/commons/util"
	"os"
	"strings"
)

func main() {
	readFile()
}

func readFile() {
	file, err := os.Open("conf/adx_jion.data")
	if err != nil {
		fmt.Printf("file err : %s \n", err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colums := strings.Split(line, "\t")
		pbStr, err := util.Base64decode(colums[1], "")
		if err != nil {
			fmt.Printf("base64 err : %s \n", err.Error())
			return
		}
		sl := &entry.SearchLog{}
		err = proto.Unmarshal([]byte(pbStr), sl)
		if err != nil {
			fmt.Printf("unmarshal err : %s \n", err.Error())
			return
		}
		fmt.Printf("sl.sid=%s;  \n", sl.Sid)
		if i := sl.Request.RequestIqiyi; i != nil {
			fmt.Printf("iqiyi:os=%s; %v \n", i.GetDevice().GetOs(), i.GetSite().GetContent().GetKeyword())
		}
		if t := sl.Request.RequestToutiao; t != nil {
			fmt.Printf("toutiao:os=%s; str=%s \n", t.GetDevice().GetOs(), t.String())
		}
		//fmt.Printf("%s,%s\n", colums[0], colums[1])
	}
}
