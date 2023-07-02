package env

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ModelEnv 环境变量解析
type ModelEnv struct {
	DevelopmentEnv bool `env:"development"` // 开发环境
}

// Env 环境变量
var Env ModelEnv

func init() {
	// 读取环境变量文件
	data, err := os.ReadFile(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("read env err, ", err)
	}
	// 不存在设置默认值
	if os.IsNotExist(err) {
		Env.DevelopmentEnv = false
		return
	}
	// 解析
	parseOut := make(map[string]string) // 文件解析结果键值对
	dts := strings.Split(string(data), "\n")
	for _, item := range dts {
		temp := strings.Split(item, "=")
		if len(temp) == 1 {
			temp = append(temp, "")
		}
		parseOut[temp[0]] = temp[1]
	}
	// 赋值
	typeEnv := reflect.TypeOf(Env)
	valueEnv := reflect.ValueOf(&Env).Elem()
	for i := 0; i < typeEnv.NumField(); i++ {
		field := typeEnv.Field(i)
		// 当配置文件中存在时赋值
		if value, ok := parseOut[field.Tag.Get("env")]; ok {
			switch field.Type.Kind() {
			case reflect.Bool:
				t, err := strconv.ParseBool(value)
				if err != nil {
					fmt.Println("解析bool失败")
					continue
				}
				valueEnv.FieldByName(field.Name).SetBool(t)
				break
			case reflect.String:
				valueEnv.FieldByName(field.Name).SetString(value)
				break
			default:
				fmt.Println("未知类型，无法解析")
				break
			}
		}
	}
}
