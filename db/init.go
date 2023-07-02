package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"todo_backend/config"
	"todo_backend/env"
	"todo_backend/models"
	"todo_backend/version"
)

var (
	DB                 *gorm.DB
	initialFileContent = []byte(fmt.Sprintf("%s@%s\n%s",
		config.AppName, version.GetV(version.Server), time.Now().Format(config.TimeLayout)))
)

// 迁移表结构
func migrateTable() {
	fmt.Println("正在迁移表格。。。")
	err := DB.AutoMigrate(&models.User{},
		&models.GroupModel{}, &models.ItemModel{},
		&models.BookNote{}, &models.BookAttachment{},
		&models.AccessLogModel{},
		&models.AppVersion{}, &models.AppHistoryVersion{},
		&models.FileStoreModel{},
		&models.FeedBackModel{},
		&models.QTodoModel{})
	if err != nil {
		log.Fatal("创建表结构失败", err)
	}
}

// 创建initial文件
func createInitialFile() {
	fmt.Println("正在创建.initial...")
	if err := os.WriteFile(".initial", initialFileContent, 0666); err != nil {
		log.Fatal("创建.initial失败", err)
	}
}

// InitDB 初始化数据库
func InitDB() {
	gormCfg := &gorm.Config{}
	// 当开发环境时
	if env.Env.DevelopmentEnv {
		// 开启日志
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
		//删除initial
		if err := os.Remove(".initial"); err != nil {
			fmt.Println("删除.initial文件失败！")
		}
	}
	db, err := gorm.Open(mysql.Open(config.Cfg.Dsn), gormCfg)
	if err != nil {
		log.Fatal("数据库初始化失败", err)
	}
	DB = db
	// 初始化连接数据库和模型
	data, err := os.ReadFile(".initial")
	if err != nil {
		migrateTable()
	}
	// 版本修改，重新初始化
	dt := string(data)
	ok, err := regexp.MatchString("ToDo Server@[.\\d]+", dt)
	if !ok || err != nil {
		fmt.Println(".initial文件错误，正在重新创建。。。")
		createInitialFile()
	}
	dt = string(initialFileContent)
	if strings.Split(strings.Split(dt, "\n")[0], "@")[1] != version.GetV(version.Server) {
		migrateTable()
		createInitialFile()
	}
}
