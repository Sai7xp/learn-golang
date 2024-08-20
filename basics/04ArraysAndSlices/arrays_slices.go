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

func main() {

	arraysInGo()

	slicesInGo()

	slicesDeepCopyShallowCopy()

	readInputFromUserUsingBufio()

}

// read input from user
func readInputFromUserUsingBufio() {
	scaner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Your Name ")
	scaner.Scan()
	input := scaner.Text()
	fmt.Println("Hi", input)
	fmt.Print("Enter Year you were born ")
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
func arraysInGo() {
	fmt.Println("→→→→→→→→→→→ START:BLOCK DATA STRUCTURES IN GO LANG ←←←←←←←←←")

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

	ticTacToeUsingSlices()

	fmt.Println("→→→→→→→→→→→ END:BLOCK DATA STRUCTURES IN GO LANG ←←←←←←←←←")

}

/*
ARRAY
var s [2]int
fmt.Println(s) // prints [0,0]

SLICE
var sl []int
fmt.Println(sl) // prints nil
*/
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
	sliceUsingMake := make([]int, 0)
	fmt.Println("initSliceWithNoElements ", initSliceWithNoElements, initSliceWithNoElements == nil, len(initSliceWithNoElements))
	fmt.Println("inSliceNoEle ", inSliceNoEle, inSliceNoEle == nil, len(inSliceNoEle))
	fmt.Println("sliceUsingMake ", sliceUsingMake, sliceUsingMake == nil, len(sliceUsingMake))

	/// LENGTH and CAPACITY of SLICE
	/*
		When a slice grows via append, it takes time for the Go runtime to allocate new mem‐ory
		and copy the existing data from the old memory to the new.
		The old memory also needs to be garbage collected.

		For this reason, the Go runtime usually increases a slice by more than one each time
		it runs out of capacity. The rules as of Go 1.14 are to double the size of the slice
		when the capacity is less than 1,024 and then grow by at least 25% afterward
	*/
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

	s = append(s, []int{91, 92, 93, 94}...)
	printSlice(s) // len=6 cap=8 [5 7 91 92 93 94] 🌻 Capacity increased from 4 to 8(doubled)
	fmt.Println("------------ END: SLICES IN GO LANG ------------")
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
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