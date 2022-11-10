package md5Hash

import (
	"crypto/md5"
	"encoding/hex"
)

func Get(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
