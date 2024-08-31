/*
* Created on 31 Aug 2024
* @author Sai Sumanth
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
 */
func differentWaysToReadInputFromUser() {

	usingScanAndScanf()

	usingBufioAndScanf()

	// string formatting using different format specifiers
	stringFormatting()
}

// Reading input using fmt.Scan & fmt.Scanf
func usingScanAndScanf() {
	fmt.Println("-------- Reading Input from using Scan & Scanf ---------")
	/*
		-------> fmt.Scan <-------
		1. Reads space separated values from input. And it won't stop reading input even
		   if Enter key is pressed while reading multiple inputs.
		2. Doesn't require a format string, attempts to read input according
		   to type of variables provided.

	*/
	var a, b int
	fmt.Print("Enter two integer values a and b : ")
	fmt.Scan(&a, &b)
	fmt.Println("a:", a, "b:", b)

	/*
		--------> fmt.Scanf <-------
		1. It can read space-separated values.
		2. Requires a format string, like "%d" for integer, "%s" for string to specify what
		   type of input it expects and how to parse it
		3. More control over how the input is read and formatted.

	*/
	var c, d int
	fmt.Print("Enter two integer values c and d : ")
	fmt.Scanf("%d %d", &c, &d)
	fmt.Println("c:", c, "d:", d)

}

// bufio package
/*
   -------> Scanf vs bufio <---------
   1. Scanf requires exact Format matching. For example if you use `fmt.Scanf("%d")` to read integer
	  but user types a non-integer value. it will result in error.

   2. `bufio.Scanner` reads input line by line. ideal for reading large inputs.
      It reads any king of text input as string or Bytes. Then we need to convert
	  input into desired format.


*/
func usingBufioAndScanf() {
	fmt.Println("-------- Reading Input from using bufio ---------")
	sc := bufio.NewScanner(os.Stdin) // scans input from Standard Input (keyboard)
	fmt.Print("Please enter your name: ")
	sc.Scan() // wait for user to input some text and press enter
	input := sc.Text()
	fmt.Println("Hi", input)

	fmt.Print("Enter Your Year of Birth: ")
	sc.Scan()
	yob, err := strconv.ParseInt(sc.Text(), 10, 64) // convert text to an integer
	if err == nil {
		fmt.Println("Your age is", time.Now().Year()-int(yob))
	} else {
		fmt.Println(err)
	}

	favNum := 0
	fmt.Print("Enter Your Fav Number: ")
	fmt.Scanf("%d", &favNum) // read int from standard input
	fmt.Println("Your Fav Number is", favNum, "and it is", func() string {
		if favNum%2 == 0 {
			return "even"
		}
		return "odd"
	}())
}

type point struct {
	x, y int
}

// different format specifiers
// https://gobyexample.com/string-formatting
func stringFormatting() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p) // {1, 2}

	fmt.Printf("struct2: %+v\n", p) // {x:1 y:2}

	fmt.Printf("struct3: %#v\n", p) // main.point{x:1, y:2}

	// Type of a variable
	fmt.Printf("type: %T\n", p)  // main.point
	fmt.Printf("type: %T\n", 78) // int

	fmt.Printf("bool: %t\n", true) // true

	fmt.Printf("int: %d\n", 123) // 123

	fmt.Printf("binary: %b\n", 14) // 1110

	binary := fmt.Sprintf("%b", 15)
	fmt.Println("15 in binary:", binary)

	fmt.Printf("char: %c\n", 33) // !

	fmt.Printf("float1: %f\n", 78.9) // 78.900000

	fmt.Printf("pointer: %p\n", &p) // 0x1400000e130
}
