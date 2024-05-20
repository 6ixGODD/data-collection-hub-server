package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func MD5WithSalt(text, salt string) string {
	hash := md5.New()
	hash.Write([]byte(text + salt))
	return hex.EncodeToString(hash.Sum(nil))
}
