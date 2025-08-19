package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateStreamKey() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}
