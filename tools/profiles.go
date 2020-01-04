package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var (
	cpuProfile       = flag.String("cpu_profile", "logs/pprof/cpu.prof", "write cpu profile to file")
	memProfile       = flag.String("mem_profile", "logs/pprof/mem.prof", "write mem profile to file")
	blockProfile     = flag.String("block_profile", "logs/pprof/block.prof", "write block profile to file")
	heapProfile      = flag.String("heap_profile", "logs/pprof/heap.prof", "write heap profile to file")
	goroutineProfile = flag.String("goroutine_profile", "logs/pprof/goroutine.prof", "write goroutine profile to file")
	threadProfile    = flag.String("thread_profile", "logs/pprof/thread.prof", "write thread profile to file")
)

//每隔指定时间，保存cpuprof至文件
func LoopSaveCpuprof() {
	go func() {
		for {
			runtime.GC()
			dayFormat := time.Now().Format("20060102")
			cpuProfilePath := "logs/pprof/cpu" + dayFormat + ".prof"
			memProfilePath := "logs/pprof/mem" + dayFormat + ".prof"
			createCPUProfile(cpuProfilePath)
			time.Sleep(24 * time.Hour)
			StopCPUProfile()
			time.Sleep(1 * time.Hour)
			createMemProfile(memProfilePath)
		}
	}()
}

//取cpu概要文件
func StartCPUProfile() {
	createCPUProfile(*cpuProfile)
}

func StopCPUProfile() {
	if *cpuProfile != "" {
		pprof.StopCPUProfile() // 把记录的概要信息写到已指定的文件
	}
}

//保存内存概要文件，每隔512KB保存一次。
func SaveMemProfile() {
	createMemProfile(*memProfile)
}

//调用pprof.Lookup方法取出内存中的阻塞事件记录，保存在文件中
//默认每发生一次阻塞事件时取样一次
func SaveBlockProfile() {
	saveProfile(blockProfile, "block", 1)
}

//调用pprof.Lookup方法取出内存中的堆内存分配情况记录，保存在文件中
//默认每分配512K字节时取样一次
func SaveHeapProfile() {
	saveProfile(heapProfile, "heap", 1)
}

//调用pprof.Lookup方法取出内存中的活跃Goroutine的信息的记录（仅在获取时取样一次），保存在文件中
func SaveGroutineProfile() {
	saveProfile(goroutineProfile, "goroutine", 1)
}

//调用pprof.Lookup方法取出内存中的系统线程创建情况的记录（仅在获取时取样一次），保存在文件中
func SaveThreadProfile() {
	saveProfile(threadProfile, "threadcreate", 1)
}

func saveProfile(profileName *string, ptype string, debug int) {
	if profileName != nil && *profileName != "" {
		f, err := os.Create(*profileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create profile output file: %s", err)
			return
		}
		if err = pprof.Lookup(ptype).WriteTo(f, debug); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", *profileName, err)
		}
		f.Close()
	}
}

//创建cpu profile文件
func createCPUProfile(filePath string) error {

	if filePath != "" {
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",
				err)
			return err
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
			f.Close()
			return err
		} else {
			fmt.Println("start cpu profile. ")
		}
	}
	return nil
}

//创建memory profile文件
func createMemProfile(filePath string) {
	if filePath != "" {
		f, err := os.Create(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
			return
		}
		if err = pprof.WriteHeapProfile(f); err != nil {
			fmt.Fprintf(os.Stderr, "Can not write %s: %s", filePath, err)
		}
		f.Close()
	}
}
