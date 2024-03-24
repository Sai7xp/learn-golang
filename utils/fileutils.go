package utils

import (
	"fmt"
	"os"
)

func ReadFileContents(path string) (string, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error occurred while reading file")
		// panic(err)
		return "", err
	} else {
		// fmt.Println(string(data)) // convert []byte data to string format
		return string(data), nil
	}
}
