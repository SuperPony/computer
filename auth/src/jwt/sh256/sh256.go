package sh256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

const (
	secretKey = "abcdefg"
)

func Encode(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}

func Decode(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func Encryption(data []byte) string {
	m := hmac.New(sha256.New, []byte(secretKey))
	m.Write(data)
	return hex.EncodeToString(m.Sum(nil))
}
