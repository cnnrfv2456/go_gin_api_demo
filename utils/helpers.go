package utils

import (
	"crypto/md5"
	"fmt"
)

//將字串轉md5
func StrToMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
