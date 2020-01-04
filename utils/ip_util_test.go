package main

import (
	"sort"
	"testing"
)

//测试ip检索方法时延
func Benchmark_IpSearch(b *testing.B) {
	b.StopTimer()
	ipLib, _ := loadIpLib()
	sort.Sort(ipLib)
	ip := "1.180.208.35"
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ipLong, _ := Ip2Long(ip)
		binarySearch(ipLib, ipLong)
	}
	b.StopTimer()
}

//并发测试ip检索方法
func Benchmark_ParallelIpSearch(b *testing.B) {
	b.StopTimer()
	ipLib, _ := loadIpLib()
	sort.Sort(ipLib)
	ip := "1.180.208.35"
	b.SetParallelism(1000)
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ipLong, _ := Ip2Long(ip)
			binarySearch(ipLib, ipLong)
		}
	})
	b.StopTimer()
}
