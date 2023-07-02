package controller

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"todo_backend/utils"
)

type CryptForm struct {
	HC   string `json:"hc,omitempty"`   // 数据加密前hash
	Data string `json:"data,omitempty"` // 加密后的数据
}

// Decrypt 解密数据
func (c *CryptForm) Decrypt() ([]byte, error) {
	dt, err := base64.StdEncoding.DecodeString(c.Data)
	if err != nil {
		return nil, err
	}
	res := utils.AesDecrypt(dt)
	if utils.MakeMd5(res) != c.HC {
		return nil, errors.New("数据已被篡改")
	}
	return res, nil
}

// EncryptToForm 加密响应数据
func EncryptToForm(data any) (c CryptForm) {
	dt, _ := json.Marshal(data)
	c.Data = utils.AesEncrypt(dt)
	c.HC = utils.MakeMd5(dt)
	return c
}

// DecryptRequest 解析请求
func DecryptRequest[T any](ctx *gin.Context, res *T) error {
	var form CryptForm
	if err := ctx.ShouldBind(&form); err != nil {
		return err
	}
	dt, err := form.Decrypt()
	if err != nil {
		return err
	}
	if err = json.Unmarshal(dt, &res); err != nil {
		return err
	}
	return nil
}

// DecryptWsRequest DecryptRequest 解析请求
func DecryptWsRequest[T any](data CryptForm) (res T, err error) {
	dt, err := data.Decrypt()
	if err != nil {
		return
	}
	if err = json.Unmarshal(dt, &res); err != nil {
		return
	}
	return
}
