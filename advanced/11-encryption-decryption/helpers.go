package main

import (
	"crypto/rand"
	"encoding/base64"
)

// Generate a 32-byte AES key for AES-256
func generateKey() (string, error) {
	key := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(key), nil
}
