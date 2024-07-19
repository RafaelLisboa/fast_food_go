package services

import (
	"crypto/sha256"
	"encoding/hex"
)

func encryptPassword(password string) string {
	return hashPassword(password)
}


func hashPassword(password string) string {
    hasher := sha256.New()
    hasher.Write([]byte(password))
    return hex.EncodeToString(hasher.Sum(nil))
}