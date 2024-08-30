/*
* Created on  Aug 2024
* @author Sai Sumanth
 */

package main

import "fmt"

/*
https://gobyexample.com/variadic-functions

Variadic Functions -> this fn can be called as sumOfN(2, 3, 4, 2, 1) 😎

values ... int should be at the end of the parameters list.
sumOfN(values ...int, s string) ❌ not possible
sumOfN(s string, values ...int) ✅ possible
*/
func sumOfN(values ...int) (total, count int) {
	count = len(values)
	for i := 0; i < count; i++ {
		total += values[i]
	}
	return

}

func RunVariadicFunctions() {
	// Calling Variadic Functions
	sum, n := sumOfN(2, 6, 8, 9, 24)
	fmt.Println("Sum of ", n, "numbers (calculated inside Variadic Fn) : ", sum)
	// even this Println is also a Variadic Function `fmt.Println(a ...any)`
	fmt.Println("Hello", 123, true)
}
