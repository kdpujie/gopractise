package main
import "fmt"

//包装int
type Integer int
func (a Integer) compare(b Integer) int {//两数比大小
	var flag int = 0
	if a < b {
		flag = -1
	}else if a > b {
		flag = 1
	}
	return flag
}
func (a *Integer) add(b Integer) {//两数相加
	 *a += b
}
//定义接口: 可对数据进行大小比较 和 增加操作
type LessAdder interface {
	compare(b Integer) int
	add(b Integer)
}
//定义接口: 可对数据进行大小比较
type Lesser interface {
	compare(b Integer) int
}

func main(){
	var a,b Integer = 1,3
	var c1 LessAdder = &a    //只能这么赋值
	var d1,d2 Lesser = b,&b //两种赋值都可以
	//var c2 LessAdder = a  // 这种赋值会导致编译错误, 为什么?
	c1.add(b) //指针是可以直接调用方法的
	a.add(b)
	fmt.Println("c1=",c1,",a=",a,",d1=",d1,",d2=",d2)
	fmt.Println(d1.compare(b))
	fmt.Println(d2.compare(b))
}