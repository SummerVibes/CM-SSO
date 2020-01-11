package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

func Verify(code, key, str string) bool {
	return Encrypt(code,key)==str
}

func Encrypt(code, key string) string {
	return MD5(MD5(code)+key)
}

func MD5(code string) string {
	w := md5.New()
	_, err := io.WriteString(w, code)//将str写入到w中
	if err != nil {
		log.Printf("生成MD5值失败",err)
	}
	md5Str := fmt.Sprintf("%x", w.Sum(nil))//w.Sum(nil)将w的hash转成[]byte格式
	return md5Str
}


//生成随机字符串
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
var src = rand.NewSource(time.Now().UnixNano())
func GenRandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
