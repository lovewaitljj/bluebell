package md5

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = []byte("夏天夏天悄悄过去留下小秘密")

// MD5Encrypt encrypts a string using MD5.
func MD5Encrypt(text string) string {
	hash := md5.New()
	hash.Write(secret)
	hash.Write([]byte(text))
	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash.Sum(nil))
}
