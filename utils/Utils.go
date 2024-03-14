package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

func GenerateRandomMd5() string {
	str := make([]byte, 32)
	rand.Read(str)

	// 计算 MD5 值
	md5 := md5.Sum(str)

	// 输出 MD5 值
	return hex.EncodeToString(md5[:])
}
