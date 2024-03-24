package main

import (
	"fmt"
	"io"
	"learngo/utils"
	"os"
)

func main() {
	fmt.Printf(" \n------>Inside the basics/Files folder <-------- \n")
	content := "- store this text message in a md file"
	/// CREATE File
	outFile, err := os.Create("message.md")
	if err != nil {
		panic(err)
	}
	/// closing the file after performing all other operations
	defer outFile.Close()
	/// WRITE string to file
	_, err1 := io.WriteString(outFile, content)
	/// error occurred while writing content to file
	utils.CheckForNilError(err1)

	io.WriteString(outFile, "\n") // add next line inside the file

	/// another way of writing
	_, err2 := outFile.Write([]byte(content))
	utils.CheckForNilError(err2)

	/// READ the file contents
	fileContent, _ := utils.ReadFileContents("message.md")
	fmt.Println("file contents are : ")
	fmt.Println(fileContent)
 
	/// DELETE File 
	removeError := os.Remove("message.md")

	if removeError != nil{
		fmt.Println("File Removed successfully from the given Path")
	}
}
