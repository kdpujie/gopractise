package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var logger *log.Logger

const (
	FileName = "test.txt"
)

func main() {
	//Log()
	//	readConfig()
	ReadFile()
}

func Log() {
	logger = OpenLogFile(FileName)
	var index int = 1
	for {
		logger.Printf(fmt.Sprintf("\tworld\tamerica %d", index))
		index++
		time.Sleep(1000000)
	}
}

func Print(args string, logger *log.Logger) {
	if !isExist(FileName) {
		logger = OpenLogFile(FileName)
	}

	logger.Println(args)

}

func OpenLogFile(fileName string) *log.Logger {
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("%s\r\n", err.Error())
		os.Exit(-1)
	}
	w := bufio.NewWriter(logFile)
	w.WriteString("")
	//runtime.SetFinalizer(logFile, logFile.Close())
	return log.New(logFile, "", log.LstdFlags)

}

//判断文件是否存在
func isExist(fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
