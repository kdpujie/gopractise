package main

import (
	"reflect"
	"fmt"
	"ksyun.com/dsp-engine/index"
)


func main()  {
	var fs = []index.AdFilter{
		index.AdTypeFilter,
	}
	ty := reflect.TypeOf(fs[0])
	fmt.Printf("reflect.typeof.name=%v \n",ty.Name())
	fmt.Printf("reflect.typeof.kind=%v \n",ty.Kind())
	fmt.Println(reflect.ValueOf(fs[0]).Kind())
}