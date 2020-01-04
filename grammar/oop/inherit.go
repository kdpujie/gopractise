package main
import "fmt"
import "time"
/**
go语言面向对象特性(为类型添加方法): 继承. 和java相比, 子类方法需要明确传入java的this. 当需要修改对象的时候, 必须用指针.因为值传递.
**/
type Integer int
func (a Integer) less(b Integer) bool {
	return a < b
}
func (a *Integer) add(b Integer) {
	*a = *a + b
}

func main(){
	var a,b Integer = 1,2
	fmt.Println("1. 定义Interger结构体 继承于int,并增加add方法:")
	if a.less(b) {
		fmt.Println("\t a小于b:a=",a,",b=",b)
	}else{
		fmt.Println("\t a不小于b:a=",a,",b=",b)
	}
	a.add(b)
	fmt.Println("\t a + b后 a =",a) 
	fmt.Println("2. 引用传递")
	var c = [3]int{1,2,3}
	var d = &c
	fmt.Println("\t",*d)
	//reac1 := new(Reac)
	log := &Logger{"矩形"}
	reac2 := &Reac{width:100,height:50,Logger:log}
	fmt.Println(reac2.GetArea())	
}
/**
结构体,地位相当于java的class.
**/
type Reac struct { //矩形
	x,y float64
	width,height float64
	*Logger
}
//计算矩形的面积
func (r *Reac) GetArea() float64 {
	r.Log("GetArea()被调用,width=",r.width)
	return (*r).width * (*r).height
}
//匿名组合. 组合Logger的类, 可以直接使用Logger的方法
type Logger struct {
	name string
}
func (l *Logger) Log(message ...interface{}){
	fmt.Println((*l).name,"-",time.Now(),"-",": ",message)
}