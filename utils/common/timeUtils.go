package common

import (
	"log"
	"time"
	"fmt"
)

const c_1  int= 1
var v_1 int = c_1 * 2

func init()  {
	fmt.Printf("common.timeUtils init():c_1=%d, v_1=%d \n",c_1, v_1)
}

func main()  {
	fmt.Printf("common.timeUtils main():c_1=%d, v_1=%d \n",c_1,v_1)
}

// 写超时警告日志 通用方法
func TimeoutWarning(tag, detailed string, start time.Time, timeLimit float64) {
	dis := time.Now().Sub(start).Seconds()
	if dis > timeLimit {
		log.Println(tag, " detailed:", detailed, "TimeoutWarning using", dis, "s")
		//pubstr := fmt.Sprintf("%s count %v, using %f seconds", tag, count, dis)
		//stats.Publish(tag, pubstr)
	}
}

//执行耗时
func TimeSpend(tag string, start time.Time) {
	dis := time.Now().Sub(start).Seconds()
	log.Println(tag, " using", dis, " s")
}

func Excute()  {
	fmt.Println("common.timeUtils excute!")
}