package main

import "fmt"

// https://gobyexample.com/closures
func main() {

	// Go supports anonymous functions, which can form closures.
	// Anonymous functions are useful when you want to define a
	// function inline without having to name it.
	closures()
}

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

// This function [calculateOddNumbers] returns another function,
// which we define anonymously in the body.
// The returned function closes over the variable `num` to form a `closure`
func calculateOddNumbers() func() int {
	num := 1
	return func() int {
		num += 2
		return num
	}
}

func closures() {

	// call the outer function
	message := greet()

	// call the inner function
	fmt.Println(message()) // Hi John
	fmt.Println(message()) // Hi Hi John
	fmt.Println(message()) // Hi Hi Hi John

	nextOddNumber := calculateOddNumbers()
	// fmt.Println(nextOddNumber) // 0x1022b86b0 ADDRESS WILL BE PRINTED
	fmt.Println("List of Odd Numbers ")
	fmt.Println(nextOddNumber()) // 3
	fmt.Println(nextOddNumber()) // 5
	fmt.Println(nextOddNumber()) // 7
	fmt.Println(nextOddNumber()) // 9

	nextOddNumber2 := calculateOddNumbers()
	fmt.Println(nextOddNumber2()) // 3 again starts from scratch

}
