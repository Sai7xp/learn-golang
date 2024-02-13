/*
* Created on 11 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"fmt"
	"log"
	"os"
)

// Reads the entire file contents in one operation
func readWholeFileInOneGo(filePath string) {
	fileContents, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Whole File Content as String: ")
	fmt.Println(string(fileContents))
}

func main() {
	filePath := "./assets/imp-data-short.txt"
	//
	// 1. Read whole file contents in one go
	//
	readWholeFileInOneGo(filePath)

}
