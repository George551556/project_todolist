package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 将传入的字符串生成为一个md5码
func My_md5(str string) string {
	str_1 := []byte(str)
	md5New := md5.New()
	md5New.Write(str_1)
	md5string := hex.EncodeToString(md5New.Sum(nil))
	return md5string
}
