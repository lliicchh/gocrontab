package main

import (
	"flag"
	"fmt"
	"godemo/crontab/master"
	"runtime"
)

var(
	confFile string
)
//解析命令参数
func initArgs()  {
	// usage : master -config ./master.json
	// master -h
	flag.StringVar(&confFile, "config", "./master.json" ,"指定master.json")
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
	if err =master.InitConfig(confFile);err!=nil{
		goto ERR
	}
	// 启动api server
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}
	fmt.Println(master.G_config.ApiReadTimeout)

	return

ERR:
	fmt.Println(err)
}
