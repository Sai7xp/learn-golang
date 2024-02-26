/*
* Created on 21 Feb 2024
* @author Sai Sumanth
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

// read input from user
func readInputFromUserUsingBufio() {
	scaner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Your Name ")
	scaner.Scan()
	input := scaner.Text()
	fmt.Println("Hi", input)
	fmt.Print("Enter Your Year you were born ")
	scaner.Scan()
	yob, e := strconv.ParseInt(scaner.Text(), 10, 64)
	if e == nil {
		fmt.Println("Your age is", time.Now().Year()-int(yob))
	} else {
		fmt.Println(e)
	}

	favNum := 0
	fmt.Print("Enter Your Fav Number ")
	// fmt.Scan(&favNum) // reading input using Scan method
	fmt.Scanf("%d", &favNum) // using scanf method
	fmt.Println("Your Fav Number is ", favNum)

}
func dataStructuresInGo() {
	fmt.Println("→→→→→→→→→→→ START:BLOCK DATA STRUCTURES IN GO LANG ←←←←←←←←←")
	// readInputFromUserUsingBufio()

	// Arrays in Go
	fmt.Println("ARRAYS in Go Lang")
	var justInitArr [3]int
	var initAndDeclare = [3]int{1, 3, 4}
	fmt.Println("Just Initialized Array ", justInitArr)
	fmt.Println("Initialized & Declared Array ", initAndDeclare)

	arr := [5]int64{1, 2, 3} // remaining two values will be 0, 0 -> {1, 2, 3, 0, 0}
	fmt.Println(arr)
	arr[4] = 90
	fmt.Println(arr[4])
	fmt.Println(arr)

	var noSizeArr = [...]string{"name1", "name2", "name3"} /// compiler count the array elements
	noSizeArr[2] = "nameThree"
	fmt.Println(noSizeArr)

	evenNumberRemZeros := [...]int{0: 2, 2: 4, 4: 6, 6: 8, 8: 10}
	fmt.Println("Even Numbers at Even Indexes ", evenNumberRemZeros)
	var sumOfEvenNumbers int

	for i := 0; i < len(evenNumberRemZeros); i++ {
		sumOfEvenNumbers += evenNumberRemZeros[i]
	}
	fmt.Println("Sum of Even Numbers: ", sumOfEvenNumbers)
	fmt.Printf("Type %T and value %v and size is : %d %U\n", evenNumberRemZeros, evenNumberRemZeros, len(evenNumberRemZeros), 'A')

	/// 2D arrays
	arr2D := [4][3]int{{1, 2}, {3, 4}, {5, 6, 7}}
	fmt.Println("2Dimentional Array ", arr2D)

	/// SLICES
	slicesInGo()
	/// Loops, if, else, switch and defer
	flowControlInGO()

	ticTacToeUsingSlices()

	fmt.Println("→→→→→→→→→→→ END:BLOCK DATA STRUCTURES IN GO LANG ←←←←←←←←←")

}

func slicesInGo() {
	// Slices in Go
	fmt.Println("------------ START: SLICES IN GO LANG ------------")

	// this is an array
	numbers := [5]int{1, 2, 3, 4, 5}

	// this is a slice
	// numbers := []int{1, 2, 3, 4, 5}

	var nums = []int{1, 2, 3, 4}
	negNums := []int{-1, -2, -3}
	negNums[0] = -0
	fmt.Println("Nums Slice values ", nums)
	nums = append(nums, 23, 45) // add two values
	fmt.Println("Updated Nums : ", nums)
	nums = append(nums, negNums...)
	fmt.Println("Again some values appended to Nums : ", nums)
	sort.Ints(nums)
	fmt.Println("Sorted Nums : ", nums)

	/// Creating slice from an Array
	sliceFromArray := numbers[1:5]
	fmt.Println("slice created from array ", sliceFromArray)
	sliceFromArray = numbers[2:]
	fmt.Println("slice created from array ", sliceFromArray)
	fmt.Println()

	duplicateSlice := negNums
	fmt.Println("Duplicate Slice ", duplicateSlice)
	fmt.Println("Original Slice ", negNums)
	negNums = append(negNums, -10, -11)
	negNums[0] = -999999
	fmt.Println("Updated negNums Slice ", negNums)
	fmt.Println("Duplicate Slice ", duplicateSlice)

	var unInitSlice []int // un initialized
	fmt.Println("unInitSlice ", unInitSlice, unInitSlice == nil, len(unInitSlice))

	var initSliceWithNoElements = []int{}
	inSliceNoEle := []int{}
	fmt.Println("initSliceWithNoElements ", initSliceWithNoElements, initSliceWithNoElements == nil, len(initSliceWithNoElements))
	fmt.Println("inSliceNoEle ", inSliceNoEle, inSliceNoEle == nil, len(inSliceNoEle))

	/// LENGTH and CAPACITY of SLICE
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // len=6 cap=6 [2 3 5 7 11 13]

	// Slice the slice to give it zero length
	s = s[:0]
	printSlice(s) // len=0 cap=6 []

	// Extend its length
	s = s[:4]
	printSlice(s) // len=4 cap=6 [2 3 5 7]

	// Drop its first two values
	s = s[2:]
	printSlice(s) // len=2 cap=4 [5 7]
	fmt.Println("------------ END: SLICES IN GO LANG ------------")
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

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

	fmt.Println("------------ END: FLOW CONTROL IN GO LANG ------------")
}

func ticTacToeUsingSlices() {
	fmt.Println("\n------> BLOCK:START Tic Tac Toe Board created using 2d arrays <-------")
	board := [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	/// few moves
	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"

	for _, v := range board {
		fmt.Println(v)
	}
	fmt.Printf("------> BLOCK:END Tic Tac Toe Board created using 2d arrays <-------\n\n")

}
