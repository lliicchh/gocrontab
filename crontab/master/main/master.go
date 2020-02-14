package main

import (
	"flag"
	"fmt"
	"gocrontab/master"
	"runtime"
	"time"
)

var (
	confFile string
)

//解析命令参数
func initArgs() {
	// usage : master -config ./master.json
	// master -h
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
}

func main() {
	var (
		err error
	)
	//初始化线程
	initEnv()

	// 加载配置
	initArgs()
	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 启动api server
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 启动JobMgr
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	time.Sleep(30 * time.Second)

	return

ERR:
	fmt.Println(err)
}
