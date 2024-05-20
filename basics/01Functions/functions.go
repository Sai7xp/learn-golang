package main

import (
	"fmt"
	m "math"
	"math/rand"
)

const (
	Big = 1 << 200
)

func main() {
	functionsInGo()
}

// Function with one argument of type int
func power(x, y int) int {
	if y == 0 {
		return 1
	}
	res := 1
	for times := 1; times <= y; times++ {
		res *= x
	}
	return res
}

// Function with Multiple Return Statements
func swap(x string, y string) (string, string) {
	return y, x
}

// Function with two return values of different types
func divideBy2(dividend int) (float64, int) {
	r := float64(dividend) / 2
	q := dividend % 2
	return r, q
}

// 🤩 Function with Named Return Values
func divideBy5(dividend int) (r float64, q int) {
	// A return statement without arguments returns the named return values.
	// This is known as a "naked" return
	r = float64(dividend) / 5
	q = dividend % 5
	return
}

// https://gobyexample.com/variadic-functions
func sumOfN(values ...int) (total, count int) {
	count = len(values)
	for i := 0; i < count; i++ {
		total += values[i]
	}
	return

}

func sumOfFirstNUsingRecursion(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sumOfFirstNUsingRecursion(n-1)
	}

}

// recursive function for finding the factorial of a number
// Recursion
// 5 -> 1 * 2 * 3 * 4 * 5
func factorialOfN(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorialOfN(n-1)
	}
}

type Input struct {
	A int
	B int
}

// function with naked return
func sumOfTwo(i Input) (s int) {
	s = i.A + i.B
	return
}

func functionsInGo() {
	fmt.Println("→→→→→→→→→→→ START:BLOCK FUNCTIONS IN GO LANG ←←←←←←←←←")
	fmt.Println("🏁 Welcome To Functions in Go Lang. Random Number is", rand.Intn(10))
	fmt.Printf("Type of Big %T and value is %v\n", float64(Big), float64(Big))
	fmt.Println("Square of 8 is : ", power(8, 2))
	fmt.Println("Square of 8 is : ", m.Pow(8, 2))
	fmt.Println("SquareRoot of 8 is : ", m.Sqrt(m.Pow(8, 2)))
	fmt.Println("pi value is: ", m.Pi)

	a, b := swap("Hello", "Meow")
	fmt.Println(a, " & ", b)
	value := 31
	r, q := divideBy5(value)
	fmt.Println("When ", value, " is Divided by 5 Remainder will be : ", r, "& Quotient will be : ", q)
	rem, quo := divideBy2(value)
	fmt.Println("When ", value, " is Divided by 2 Remainder will be : ", rem, "& Quotient will be : ", quo)

	sum, n := sumOfN(2, 6, 8, 9, 24)
	fmt.Println("Sum of ", n, "numbers is ", sum)
	fmt.Println("Sum of 2 Numbers 23 & 89 :  ", sumOfTwo(Input{A: 23, B: 89}))

	// Recursive Functions
	fmt.Println("Recursive Function Examples")
	fmt.Println("Sum of first 10 Natural Number : ", sumOfFirstNUsingRecursion(5))
	factorial := 0
	fmt.Print("Enter Number to find the Factorial : ")
	// Read input from User
	fmt.Scan(&factorial)
	fmt.Printf("Factorial of %v is : %v\n", factorial, factorialOfN(factorial))
	fmt.Println("→→→→→→→→→→→ END:BLOCK FUNCTIONS IN GO LANG ←←←←←←←←←")

	/// Anonymous functions
	func() {
		fmt.Println("Inside Anonymous Function without Name")
	}()

	func(s string) {
		fmt.Println("My fav programming language is", s)
	}("Go")

	var ano = func(n int) {
		fmt.Println("Hello from Anonymous Function", n)
	}
	ano(12)
}
