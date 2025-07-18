package main

import (
	"fmt"
	"learngo/advanced/11-encryption-decryption/symmetric"
)

func main() {
	secretKey, _ := generateKey()
	// secretKey := "ygNO8TglV+IZSqFD81oul7rrULy/IP0UFaA6e9n9918="

	fmt.Println("Secret: ", secretKey)

	message := "sumanth"
	fmt.Println("Data to be encrypted: ", message)

	encryptedData := symmetric.EncryptAES_GCM(secretKey, []byte(message))
	fmt.Println("Encrypted Data: ", encryptedData)

	messageInBytes := symmetric.DecryptAES_GCM(secretKey, encryptedData)
	fmt.Println("DecryptedData ", string(messageInBytes))

}
