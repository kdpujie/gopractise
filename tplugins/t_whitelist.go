package main

import (
	"plugin"
	"fmt"
	"learn.com/tplugins/entry"
)

func main()  {
	path := "./mfs/mfs.so"
	symbol := "Whitelist"
	plug, err :=  plugin.Open(path)
	if err != nil {
		fmt.Printf("Open插件(%s)失败: %v \n",path ,err)
	}
	symPlug, err := plug.Lookup(symbol)
	if err != nil {
	}else {
		fmt.Printf("插件(%s) Lookup symbol(%s) \n",path, symbol)
	}
	var query entry.Query
	query, ok := symPlug.(entry.Query)
	if !ok {
		fmt.Println("unexpected type from module symbol")
	}
	query.Init()
	req := generateReq()
	adlists := []string{"01014","01012","01013"}
	mfsAd := query.Filter("exp_whitelist",0,adlists, req)
	fmt.Printf("mfs选中的广告有：%v \n", mfsAd)
}
