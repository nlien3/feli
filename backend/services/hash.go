package services

import (
	"crypto/md5"
	"encoding/hex"
)

// Hashear la contraseña usando MD5
func hashPassword(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}