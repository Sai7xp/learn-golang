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
}

// Reading input using fmt.Scan & fmt.Scanf
func usingScanAndScanf() {
	fmt.Println("-------- Reading Input from using Scan & Scanf ---------")
	var a, b int
	fmt.Print("Enter a and b values: ")
	fmt.Scan(&a, &b)
	fmt.Println("a:", a, "b:", b)

	var c, d int
	fmt.Print("Enter c and d values: ")
	fmt.Scanf("%d %s", c, d)
	fmt.Println("c:", c, "d:", d)
}

// reads user input from keyboard
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
