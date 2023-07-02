package main

import (
	"flag"
	"fmt"
	"todo_backend/config"
	"todo_backend/db"
	"todo_backend/routers"
	"todo_backend/version"
)

var (
	configFile string
)

func init() {
	version.CheckVersionFile()
	flag.StringVar(&configFile, "c", "config.json", "json类型配置文件路径")
	flag.Parse()
	fmt.Printf("Welcome to use %s@%s\n", config.AppName, version.GetV(version.Server))
}

func main() {
	config.InitConfig(configFile)
	db.InitDB()
	// 启动web服务
	routers.Run()
}
