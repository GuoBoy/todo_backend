package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"todo_backend/config"
)

func AesEncrypt(data []byte) string {
	kiv := config.Aes
	block, err := aes.NewCipher([]byte(kiv.Key))
	if err != nil {
		panic("生成aes cipher失败")
	}
	// 原始数据补码
	data = PKCS7Padding(data, block.BlockSize())
	// 使用cbc模式
	mode := cipher.NewCBCEncrypter(block, kiv.IV)
	// 结果数据, 初始化加密数据接收切片
	ciphertext := make([]byte, len(data))
	mode.CryptBlocks(ciphertext, data)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func AesDecrypt(data []byte) []byte {
	kiv := config.Aes
	block, err := aes.NewCipher([]byte(kiv.Key))
	if err != nil {
		panic("生成aes cipher失败")
	}
	// 加密模式
	mode := cipher.NewCBCDecrypter(block, kiv.IV)
	// 结果数据
	ciphertext := make([]byte, len(data))
	// 解密
	mode.CryptBlocks(ciphertext, data)
	// 去补码
	return PKCS7UnPadding(ciphertext)
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
