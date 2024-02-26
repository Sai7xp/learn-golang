package utils

import (
	"fmt"
	"os"
)

func readFile() {

	data, err := os.ReadFile("data.txt")
	if err != nil {

		fmt.Println("Error occurred while reading file")
		// return
	}
	if err == nil {
		fmt.Println(data)
	}
}
