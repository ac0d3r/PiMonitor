package util

import (
	"math/rand"
	"time"
)

var (
	digitSeeds       = []byte("0123456789")
	upperLetterSeeds = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerLetterSeeds = []byte("abcdefghijklmnopqrstuvwxyz")
)

// RandomDigitString 随机生成长度为length的数字字符串
func RandomDigitString(length int) string {
	return randomString(digitSeeds, length)
}

// RandomVerificationCode 随机生成长度为length的验证码（数字+大写字母）字符串
func RandomVerificationCode(length int) string {
	return randomString(append(digitSeeds, upperLetterSeeds...), length)
}

// RandomString 随机生成字符串包含数字及大小写字母
func RandomString(length int) string {
	return randomString(append(append(digitSeeds, upperLetterSeeds...), lowerLetterSeeds...), length)
}

func randomString(seeds []byte, length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]byte, 0)
	for i := 0; i < length; i++ {
		slice = append(slice, seeds[r.Intn(len(seeds))])
	}
	return string(slice)
}
