package main

import "fmt"

func main() {
	fmt.Println("Concurrency Patterns in Go")

	// Fan Out
	FanOutImplementation()

	// Fan In
	FanInImplementation()
}
