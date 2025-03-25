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

	// Fan out simple example
	FanOutEaxmple2()
}
