package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ReadFile() {
	mapping := read3("mapping.txt")
	//	for k, v := range mapping {
	//		fmt.Println(k, v)
	//	}
	writeIpLib(mapping)
}

func writeIpLib(mapping map[string]string) {
	//读取文件,当不存时创建
	wf, err := os.OpenFile("ip_lib_my.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	defer wf.Close()
	rf, err := os.Open("ip.txt")
	if err != nil {
		fmt.Println("file open failed.", err)
	}
	defer rf.Close()
	scanner := bufio.NewScanner(rf)
	for scanner.Scan() {
		line := scanner.Text()
		colums := strings.Split(line, "\t")
		mapValue := mapping[colums[8]]
		//		fmt.Printf("mapValue=%s,colums=%s", mapValue, colums[8])
		values := strings.Split(mapValue, "\t")
		temp := &bytes.Buffer{}
		temp.WriteString(colums[0])
		temp.WriteString("\t")
		temp.WriteString(colums[1])
		temp.WriteString("\t")
		temp.WriteString(colums[2])
		temp.WriteString("\t")
		temp.WriteString(colums[3])
		temp.WriteString("\t")
		temp.WriteString(colums[4])
		temp.WriteString("\t")
		temp.WriteString(colums[5])
		temp.WriteString("\t")
		temp.WriteString(values[0])
		temp.WriteString("\t")
		temp.WriteString(values[1])
		temp.WriteString("\t")
		temp.WriteString(values[2])
		temp.WriteString("\t")
		temp.WriteString(colums[9])
		temp.WriteString("\n")
		wf.WriteString(temp.String())
	}
	wf.Sync()
}

//逐行读取
func read3(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open failed.", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var mapping map[string]string = make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		worlds := strings.Split(line, "\t")
		mapping[worlds[3]] = line
		//		for _, v := range worlds {
		//			fmt.Print(v, ":")
		//		}
		//		fmt.Println()
	}
	//	return scanner.Err()
	return mapping
}

//逐行读取
func read2(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("file open failed.", err)
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
	}
}

func read1(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	fmt.Println(string(fd))
}
