package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func StrMd5(pwd string) string {
	temp := md5.Sum([]byte(pwd))
	return hex.EncodeToString(temp[:])
}

func MakeMd5(st []byte) string {
	m := md5.Sum(st)
	return hex.EncodeToString(m[:])
}
