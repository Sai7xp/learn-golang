package main

import (
	"fmt"
)

func main() {
	fmt.Println("Concurrency Patterns in Go")

	// Fan Out Pattern
	FanOutImplementation()

	// Fan In Pattern
	FanInImplementation()

	// Worker Pool
	WorkerPoolExample()
}
