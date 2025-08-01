package ciphers

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 returns the MD5 hash of the input string as a hexadecimal string.
func MD5(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
