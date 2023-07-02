package config

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"todo_backend/models"
)

type AdminConf struct {
	Uri   string                  `json:"uri,omitempty"`
	Users []models.AdminUserModel `json:"users,omitempty"`
}

type SystemConf struct {
	Admin AdminConf `json:"admin"`
}

type Config struct {
	JwtSecret       string     `json:"jwtSecret"`
	Dsn             string     `json:"dsn"`
	EncryptPassword string     `json:"encrypt_password"`
	ServerPort      uint       `json:"server_port"`
	DefaultResetPwd string     `json:"default_reset_pwd"`
	UploadDir       string     `json:"upload_dir"`    //文件上传路径
	ExtBlacklist    []string   `json:"ext_blacklist"` // 文件类型黑名单
	System          SystemConf `json:"system"`
}

func (c Config) JwtSecretBytes() []byte {
	return []byte(c.JwtSecret)
}

type Aes256cbcKIv struct {
	Key string
	IV  []byte
}

var (
	Cfg Config
	Aes Aes256cbcKIv
)

// InitConfig 初始化配置
func InitConfig(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Println("读取配置文件失败，正在创建配置文件。。。", err)
		// 创建默认配置文件
		configFile = "config.json"
		fp, cfErr := os.Create(configFile)
		if cfErr != nil {
			log.Fatal("创建配置文件失败", cfErr)
		}
		defer fp.Close()
		if err = json.NewEncoder(fp).Encode(DefaultConf()); err != nil {
			log.Fatal("初始化创建配置文件失败", err)
		}
		fmt.Println("创建初始配置文件完成，请配置各字段。。。")
		os.Exit(0)
	}
	// 解析配置文件
	if err = json.Unmarshal(data, &Cfg); err != nil {
		log.Fatal("解析配置文件失败", err)
	}
	// Aes
	m := md5.Sum([]byte(Cfg.EncryptPassword))
	Aes = Aes256cbcKIv{
		Key: hex.EncodeToString(m[:]),
		IV:  []byte("0000000000000000"),
	}
	// 检查上传文件路径
	if Cfg.UploadDir == "" {
		fmt.Println("还没有配置上传路径，将默认为当前路径的upload文件夹")
		if err = os.Mkdir("upload", 0666); err != nil {
			panic(err)
		}
		fmt.Println("创建成功")
	} else {
		info, err := os.Stat(Cfg.UploadDir)
		if err != nil || !info.IsDir() {
			fmt.Printf("没有找到配置上传路径%s，将自动创建...", Cfg.UploadDir)
			if err = os.Mkdir("upload", 0666); err != nil {
				panic(err)
			}
			fmt.Println("创建成功")
		}
	}
}

// DefaultConf 默认配置
func DefaultConf() *Config {
	return &Config{
		JwtSecret:       "此处填写jwt密钥",
		Dsn:             "此处填写mysql连接字符串",
		EncryptPassword: "此处填写加密密钥",
		ServerPort:      8080,
		DefaultResetPwd: "woshizhu",
		UploadDir:       "upload",
		System: SystemConf{
			Admin: AdminConf{
				Uri:   "此处填写管理后台地址，eg. /admin",
				Users: []models.AdminUserModel{{Username: "admin", Password: "admin"}},
			},
		},
	}
}
