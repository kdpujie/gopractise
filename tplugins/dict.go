package main

import (
	"plugin"
	"fmt"
	"learn.com/tplugins/entry"
)

func main()  {

	paths := map[string]string{
		"./eng/eng.so":"Greeter",
		"./chi/chi.so":"Greeter",
	}
	plugins := loadPlugins(paths)
	user := &entry.User{}
	user.Name = "jack"
	for path,greeter := range plugins {
		fmt.Printf("%s \n\t", path)
		greeter.Greet(user)
	}

}

func loadPlugins(paths map[string]string) map[string]entry.Greeter {
	var greeterPlugins map[string]entry.Greeter = make(map[string]entry.Greeter,8)
	for path,  symbol:= range paths {
		fmt.Printf("开始open动态库文件%s ....\n",path)
		plug, err :=  plugin.Open(path)
		if err != nil {
			fmt.Printf("Open插件(%s)失败: %v \n",path ,err)
			continue
		}
		symPlug, err := plug.Lookup(symbol)
		if err != nil {
			fmt.Printf("插件(%s) Lookup symbol(%s)失败: %v \n",path, symbol ,err)
			continue
		}else {
			fmt.Printf("插件(%s) Lookup symbol(%s) \n",path, symbol)
		}
		var greeter entry.Greeter
		greeter, ok := symPlug.(entry.Greeter)
		if !ok {
			fmt.Println("unexpected type from module symbol")
			continue
		}
		greeterPlugins[path] = greeter
	}
	return greeterPlugins
}
