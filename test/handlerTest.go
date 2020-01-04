package main

import "fmt"

type TypeFunc func(int, int) //定义一个函数类型

func (f TypeFunc) Serve(a, b int) { //此函数类型有一个Serve方法.
	f(a, b)
}

func f(a, b int) { //满足函数类型TypeFunc的参数要求的普通函数.
	fmt.Println("a + b =", (a + b))
}

func Start_Type_func() {
	test := TypeFunc(f) //把普通函数f转换为TypeFunc
	test(2, 2)          // 直接传递相应参数, 访问函数本身
	test.Serve(1, 2)    //也可以访问函数类型的Serve方法.
}
