package gox

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(value string) string {
	data := md5.Sum([]byte(value))
	return hex.EncodeToString(data[:])
}
