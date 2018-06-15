package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"time"
)

// Hmac encrypt a string using sha256
func Hmac(key, data string) []byte {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(data))
	return hmac.Sum(nil)
}

// GetHTTPDateTime get a date string formatted for http header
func GetHTTPDateTime() string {
	t := time.Now().UTC()
	return t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
}
