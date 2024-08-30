/*
* Created on 30 Aug 2024
* @author Sai Sumanth
 */

package main

import "fmt"

func scopeAndShadowing() {
	fmt.Println("Scope and Shadowing in Golang")
	scopes()

	/*
		Shadowing occurs in Go when a variable declared in a nested scope has the same name
		as a variable declared in an outer scope.
		In such cases, the variable in the inner scope shadows or hides
		the variable in the outer scope.
	*/
	shadowing()
}

var x int // global scope

func scopes() {
	fmt.Println("x value :", x) // 0
	x := 90
	fmt.Println("x value :", x) // 90
	fn(x)
}

func fn(x int) {
	fmt.Println("X value before loop starts :", x) // 90

	for x = 0; x < 10; x++ {
		fmt.Println(x)
	}
	fmt.Println("X value outside of loop :", x) // x becomes 10 here

	for x := 0; x < 5; x++ {

		fmt.Println("shadowed x value : ", x) // prints 0....4
	}

	// x will be 10 here, not 5. x declared in above for loop is scoped only for that for loop only
	fmt.Println("X value after two for loops :", x)

	/// 🌻 Tricky Part
	for x := 0; x < 5; x++ {
		x := 999
		fmt.Println("Tricky - x value : ", x) /// prints 999 all the time
	}
}

func shadowing() {
	x := 10
	fmt.Println("Outer x:", x) // Outer x: 10

	{
		// remove the short hand notation and simply assign the 20 to 'x'(unshadowing) and observe the output
		x := 20                    // shadowing
		fmt.Println("Inner x:", x) // Inner x: 20
	}

	fmt.Println("Outer x:", x) // Outer x: 10
}
