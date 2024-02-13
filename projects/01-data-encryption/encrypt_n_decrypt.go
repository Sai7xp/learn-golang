/*
	Created on 11 Feb 2024
	@author Sai Sumanth

Problem Description:

	Read huge text file (1 GB) and encrypt the contents using an
	encryption key and store the encrypted contents back to a file. Use concepts like concurrency & Parallelism
	to read the file efficiently.

Steps:
 1. Read txt file in Chunks
 2. Encrypt/Decrypt each Chunk separately on a routine (Maintaining order is important here)
 3. Write en/de data to file(1 chunk at a time)
*/
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"
	"os"
)

const secretKey = "12345678912345678912345678912345"

var iv = []byte("1234567891234567")

// Reads the entire file contents in one operation
func readWholeFileInOneGo(filePath string) {
	fileContents, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Whole File Content as String: ")
	fmt.Println(string(fileContents))
}

func encrypt(eachChunk []byte, secretKey []byte) []byte {
	/// create new cipher block
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := make([]byte, len(eachChunk)+aes.BlockSize)
	/// encrypter
	cipher := cipher.NewCFBEncrypter(block, iv)
	// performs xor operation
	cipher.XORKeyStream(ciphertext[aes.BlockSize:], eachChunk)

	// str := base64.StdEncoding.EncodeToString(ciphertext)
	return ciphertext

}

func decrypt(encryptedChunk []byte, secretKey []byte) []byte {
	// / create new cipher block
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	decryptedText := make([]byte, len(encryptedChunk)+aes.BlockSize)
	/// decrypter
	cipher := cipher.NewCFBDecrypter(block, iv)
	// performs xor operation
	cipher.XORKeyStream(decryptedText[aes.BlockSize:], encryptedChunk)

	return decryptedText
}

func main() {
	filePath := "./assets/imp-data-short.txt"
	encryptTo := "./assets/encrypted-imp-data.txt"
	decryptTo := "./assets/decrypted-imp-data.txt"
	chunkSize := 1024 * 64 // 64 KB
	//
	//  Read whole file contents in one go
	//
	// readWholeFileInOneGo(filePath)

	encryptFileChunkByChunk(filePath, encryptTo, chunkSize)

	decryptFileChunkByChunk(encryptTo, decryptTo, chunkSize)

}

// / performs encryption for whole file
func encryptFileChunkByChunk(filePath string, encryptTo string, chunkSize int) {
	// Open the input file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	/// close the file at the end
	defer file.Close()

	/// created the output file
	outFile, err := os.Create(encryptTo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()

	for {
		eachChunk := make([]byte, chunkSize)
		/// read [eachChunk] of file
		n, error := file.Read(eachChunk)
		if error == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		/// get the cipher text for each chunk
		encryptedChunk := encrypt(eachChunk[:n], []byte(secretKey))
		/// write to a file
		_, err = outFile.Write(encryptedChunk)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

// / performs decryption for whole file
func decryptFileChunkByChunk(encryptedFilePath string, decryptTo string, chunkSize int) {
	// open the encrypted file first
	file, err := os.Open(encryptedFilePath)
	if err != nil {
		fmt.Println("Error while opening file:", err)
		return
	}
	/// close the file at the end
	defer file.Close()

	// created the output file
	outputFile, err := os.Create(decryptTo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	for {
		eachChunk := make([]byte, chunkSize)
		/// read each encrypted chunk
		n, error := file.Read(eachChunk)
		if error == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		/// get the plain text for each encrypted chunk
		decryptedChunk := decrypt(eachChunk[:n], []byte(secretKey))
		/// write to a file
		_, err = outputFile.Write(decryptedChunk)
		if err != nil {
			fmt.Println(err)
			return
		}

	}
}
