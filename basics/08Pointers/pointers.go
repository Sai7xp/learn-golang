/*
* Created on 23 Feb 2024
* @author Sai Sumanth
 */
package main

/*
# Memory Management, Pointers
							--->> Resources <<---
1. https://www.udacity.com/blog/2021/05/c-pass-by-reference-explained.html
2. https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions

*/

import "fmt"

func main() {
	fmt.Println("Inside pointers.go")
	var i int
	j := 0
	fmt.Printf("i: %v and is i == 0 ? %v\n", i, (i == 0)) // true
	fmt.Printf("j: %v and is j == 0 ? %v\n", j, (j == 0)) // true
	j = 90

	p := &j // here p is a pointer (contains address)
	fmt.Printf("p value: %v\n", p)
	fmt.Printf("p pointing to value : %v\n", *p)
	*p = 99
	fmt.Printf("j value: %v\n", j) // 99 j value got changed

	counter := 0
	// 6 secs timer
	for i := 1; i <= 6; i++ {
		fmt.Printf("counter: %v\n", counter)
		// prints 0 everytime ever after calling incCounter method
		// because when we are passing `counter` to [incCounter] copy is being created
		// and copy is being modified
		incCounter(counter)
	}
	fmt.Println("Timer Started again. This time it should work properly")
	for i := 0; i < 6; i++ {
		fmt.Printf("counter: %v\n", counter)
		incCounterReal(&counter) // passing the address of counter variable
	}

	/////////////////
	intUsingNew := new(int)
	fmt.Println("intUsingNew ", intUsingNew)
	// intUsingNew

	var intUsingPointer *int
	fmt.Printf("intUsingPointer: %v\n", intUsingPointer) // nil
	// *intUsingPointer = 20 // 💀 dereferencing the nil pointer. horribly goes wrong
	num := 20
	intUsingPointer = &num
	// Now it's safe to dereference intUsingPointer since it points to a valid memory address
	fmt.Printf("intUsingPointer: %v\n", *intUsingPointer) // 20
	*intUsingPointer = 20

	/// Make can be used only for creating maps, slices and channels
	sliceUsingMake := make([]int, 10, 15)
	fmt.Println(sliceUsingMake) // [0 0 0 0 0 0 0 0 0 0]

	pointersWithStruct()

	/*
		💎 Does go supports Pass By Reference ?
	*/
	goHasOnlyPassByValueNoPassByRef()
}

func incCounter(val int) {
	val++
}

func incCounterReal(c *int) {
	// c contains the address of passed argument, copy gets created

	*c++ // modifying the value using dereferencing pointer.

	// btw this is not Pass by Reference. Go doesn't support Pass by ref unlike C++.
	// Golang supports only Pass by Value
}
