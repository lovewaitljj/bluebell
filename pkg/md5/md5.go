package md5

import (
	"crypto/md5"
	"encoding/hex"
)

var secret = ("夏天夏天悄悄过去留下小秘密")

// Encrypt encrypts a string using MD5.
func Encrypt(text string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	hash.Write([]byte(text))
	// Convert the hash to a hexadecimal string
	return hex.EncodeToString(hash.Sum(nil))
}
