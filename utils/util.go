package utils

import (
	"crypto/sha256"
	"time"
	"crypto/hmac"
)

func Hmac(key, data string) []byte {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(data))
	return hmac.Sum(nil)
}

func GetHttpDateTime() string {
	t := time.Now().UTC()
	return t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
}