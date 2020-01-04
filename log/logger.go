package main

import (
	"log"
	"os"
	"time"
)

func main() {
	lg := New("test_logg.txt")
	var index int
	for {
		index++
		lg.Printf("%d 行数 %s 猎术 ", index, "/")
		time.Sleep(time.Second * 1)
		lg.SetFlags(0)
	}

}

type Logg struct {
	*log.Logger
	name string
	file *os.File
}

func New(logName string) *Logg {
	file, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("%s\r\n", err.Error())
	}
	l := log.New(file, "", log.LstdFlags)
	return &Logg{name: logName, Logger: l}
}

func (this *Logg) openFile() error {
	var err error
	this.file, err = os.OpenFile(this.name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	this.Logger.SetOutput(this.file)
	return err
}

//判断文件是否存在
func (this *Logg) isExist() bool {
	var exist = true
	if _, err := os.Stat(this.name); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (this *Logg) Printf(format string, v ...interface{}) {
	if !this.isExist() {
		this.openFile()
	}
	this.Logger.Printf(format, v...)
}

// SetFlags sets the output flags for the logger
func (this *Logg) SetFlags(flag int) {
	this.Logger.SetFlags(flag)
}
