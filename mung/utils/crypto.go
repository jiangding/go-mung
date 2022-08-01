package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 微信签名算法方式
const (
	SignTypeMD5        = `MD5`
	SignTypeHMACSHA256 = `HMAC-SHA256`
)
// EncryptMsg 加密消息
func EncryptMsg(random, rawXMLMsg []byte, appID, aesKey string) (encrtptMsg []byte, err error) {
	defer func() {
		if e := recover(); e != any(nil) {
			err = fmt.Errorf("panic error: err=%v", e)
			return
		}
	}()

	return
}

// Md5 加密
func Md5(s string) string {
	m := md5.New()
	m.Write([]byte(s))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// Sha256 加密
func Sha256(s string) string {
	m := sha256.New()
	m.Write([]byte(s))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

