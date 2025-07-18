package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)


// encrypts the given data in bytes and returns base64 encoded string which represents encrypted data in strings
func EncryptAES_GCM(secret string, plainDataInBytes []byte) string {
	rawSecret, _ := base64.StdEncoding.DecodeString(secret)

	block, err := aes.NewCipher(rawSecret)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Nonce - Prevents repeated encryption of the same plaintext from producing the same ciphertext.
	// In AES-GCM, it's typically 12 bytes and acts like a unique IV (Initialization Vector).
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return ""
	}

	encryptedBytes := gcm.Seal(nonce, nonce, plainDataInBytes, nil)

	return base64.StdEncoding.EncodeToString(encryptedBytes)
}

func DecryptAES_GCM(secret, encryptedString string) (plainDataInBytes []byte) {
	rawSecret, _ := base64.StdEncoding.DecodeString(secret)

	cipherData, _ := base64.StdEncoding.DecodeString(encryptedString)

	block, err := aes.NewCipher(rawSecret)
	if err != nil {
		fmt.Println("Error while creating NewCipher", err)
		return nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error while creating NewGCM", err)
		return nil
	}

	nonceSize := gcm.NonceSize()
	if len(cipherData) < nonceSize {
		fmt.Println("cipherText too short")
		return nil
	}

	nonce, cipherData := cipherData[:nonceSize], cipherData[nonceSize:]

	// open the seal to get the original plain data in bytes
	plainDataInBytes, err = gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		fmt.Println("Error while gcm Open", err)
		return nil
	}

	return plainDataInBytes
}
