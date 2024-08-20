/*
* Created on 20 Aug 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

// Loops -> classic for loop, for-in, using for loop as while loop, Range over integers
func flowControlInGO() {
	fmt.Println("\n------------ START: FLOW CONTROL IN GO LANG ------------")
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	println("sum calculated using for loop", sum)

	fruits := []string{"🍒", "🍑", "🥑"}
	duplicateFruits := make([]string, 3, 300) /// Creating slice using make
	copy(duplicateFruits, fruits)
	fmt.Println("duplicateFruits ", duplicateFruits)
	fmt.Println("duplicateFruits length ", len(duplicateFruits))
	fmt.Println("duplicateFruits capacity ", cap(duplicateFruits))

	// similar to while loop
	for len(fruits) > 0 {
		fmt.Println("Last of fruits ", fruits[len(fruits)-1])
		fruits = fruits[:len(fruits)-1] // pop last element
	}

	// for in loop
	fmt.Println("For in range Loop")
	for i, val := range duplicateFruits {
		fmt.Printf("Fruit at position %d is %v || ", i, val)
		if i == len(duplicateFruits)-1 {
			fmt.Println()
		}
	}
	/// new feature in go lang 1.22
	fmt.Println("Go 1.22 New Feature -> Ranging over integers✨")
	for i := range 12 {
		fmt.Println(i)
	}

	fmt.Println("------------ END: FLOW CONTROL IN GO LANG ------------")
}
