package main

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
)

func main() {
	hosts := []*as.Host{
		as.NewHost("10.14.41.54", 3000),
		as.NewHost("10.14.41.55", 3000),
		as.NewHost("10.14.41.56", 3000),
	}
	client, _ := as.NewClientWithPolicyAndHost(nil, hosts...)
	defer client.Close()

	spolicy := as.NewScanPolicy()
	spolicy.ConcurrentNodes = true
	spolicy.Priority = as.LOW
	spolicy.IncludeBinData = true

	recs, _ := client.ScanAll(spolicy, "dsp", "indexer")
	// deal with the error here

	for res := range recs.Results() {
		if res.Err != nil {
			// handle error here
			// if you want to exit, cancel the recordset to release the resources
		} else {
			// process record here
			fmt.Printf("%s, %v \n", res.Record.Key.Namespace(), res.Record.Bins)
		}
	}

}
