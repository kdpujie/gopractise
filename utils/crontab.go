package main

import (
	"log"
	//"time"

	"github.com/robfig/cron"
)

func main() {
	startCrontab()
}

func startCrontab() {
	var c *cron.Cron
	sec := "1/1 * * * * ?"
	c = cron.New()
	c.AddFunc(sec, timerFunc)
	log.Println("1.启动crontab....")
	c.Start()
	log.Println("2.crontab启动完毕....")
	select {}
	//time.Sleep(20 * time.Second)
	log.Println("主线程退出....")
}
func timerFunc() {
	log.Println("callYourFunc come here.")
}
