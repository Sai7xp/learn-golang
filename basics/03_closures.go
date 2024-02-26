package main

import "fmt"

// outer function
func greet() func() string {

	// variable defined outside the inner function
	name := "John"

	// return a nested anonymous function
	return func() string {
		name = "Hi " + name
		return name
	}
}

func closures() {

	// call the outer function
	message := greet()

	// call the inner function
	fmt.Println(message())

	nextOddNumber := calculateOddNumbers()
	nextOddNumber2 := calculateOddNumbers()
	// fmt.Println(nextOddNumber) // 0x1022b86b0 ADDRESS WILL BE PRINTED
	fmt.Println("List of Odd Numbers ")
	fmt.Println(nextOddNumber()) // 3
	fmt.Println(nextOddNumber()) // 5
	fmt.Println(nextOddNumber()) // 7
	fmt.Println(nextOddNumber()) // 9

	fmt.Println(nextOddNumber2()) // 3 again starts from scratch

}

func calculateOddNumbers() func() int {
	num := 1
	return func() int {
		num += 2
		return num
	}
}
