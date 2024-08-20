/*
* Created on 20 Aug 2024
* @author Sai Sumanth
**/

package main

// Let's see if Go supports Pass By Reference or Not
// https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions

import "fmt"

func goHasOnlyPassByValueNoPassByRef() {
	fmt.Println("---------- goHasOnlyPassByValueNoPassByRef() ------------")

	count := 0
	fmt.Println("count value BEFORE passing to increment fn :", count)
	fmt.Printf("count address : %v\n", &count) // 0x14000110018 (this is just the address format, value may not be same everytime)
	increment(&count)
	fmt.Println("count value AFTER passing to increment fn :", count) // 200
}

func increment(a *int) {
	/*
		💡 Go technically has only pass-by-value

		'a' is complete different value from the passed argument(basically a copy got created)
		so in GoLang Pointers are not pass by ref unlike c++
	*/
	fmt.Printf("a value : %v\n", a)   // 0x14000110018 (a holds the address that is passed in arg)
	fmt.Printf("&a value : %v\n", &a) // 0x14000116020 (address of a itself is different)
	fmt.Printf("*a value : %v\n", *a) // 0
	*a = 199                          // modifies the arg value as well
	*a++
	eg := 76
	a = &eg // 💡 has no impact on the caller because we overwrote the pointer value!

}
