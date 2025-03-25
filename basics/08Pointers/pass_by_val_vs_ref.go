/*
* Created on 20 Aug 2024
* @author Sai Sumanth
**/

package main

// Let's see if Go supports Pass By Reference or Not
// https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions

import "fmt"

/*
	-------- WHAT IS PASS BY REFERENCE IN C++ -------

	// 💡 'b' is a reference to the original variable
	void modify(int &b) {
    	b = 90; // This changes the original variable directly
	}

	int main() {
    int a = 20;
    modify(a); // 'a' is passed by reference
    std::cout << a << std::endl; // 'a' is now changed to 90
    return 0;
	}


	------ PASS BY REF & PASS BY POINTER -----

	void modify(int &b) { b = 90; }

	void modify(int *b) { *b = 90; }

	What is the use of pass by reference in c++ if we can modify the value using pass by pointer
	1. Pointers can be null but References cannot be null (eliminates null checks)
	2. References can be more efficient because they avoid the overhead of pointer dereferencing

*/

func goHasOnlyPassByValueNoPassByRef() {
	fmt.Println("---------- goHasOnlyPassByValueNoPassByRef() ------------")

	count := 0
	var intUsingPointer *int
	fmt.Println("count value BEFORE passing to increment fn :", count)
	fmt.Printf("count address : %v\n", &count) // 0x14000110018 (this is just the address format, value may not be same everytime)
	increment(&count)
	fmt.Println("count value AFTER passing to increment fn :", count) // 200

	fmt.Println("intUsingPointer: ", intUsingPointer)
	// 💀 this will cause nill pointer dereference error since intUsingPointer is not initialized
	// increment(intUsingPointer)
	// to fix the above nil pointer issue assign value intUsingPointer = &count
}

func increment(a *int) {

	// if a == nil {
	// 	// handling nil pointer dereference issue
	// 	return
	// }

	/*
		💡 Go technically has only pass-by-value

		'a' is complete different value from the passed argument(basically a copy got created)
		so in GoLang Pointers are not pass by ref unlike c++
	*/
	fmt.Printf("a value : %v\n", a)   // 0x14000110018 ('a' holds the address that is passed in arg)
	fmt.Printf("&a value : %v\n", &a) // 0x14000116020 (address of 'a' itself is different)
	fmt.Printf("*a value : %v\n", *a) // 0
	*a = 199                          // modifies the arg value as well
	*a++
	eg := 76
	a = &eg // 💡 has no impact on the caller because we overwrote the pointer value!

}

// C++ EXAMPLE

/*

#include <iostream>

// https://www.programiz.com/online-compiler/7NwxJhJwiBe0F

// PASS BY POINTER
void modify(int *a)
{
    std::cout << a << std::endl;
    int z = 999;
    *a = 78;
    a = &z;
}

// PASS BY REF
void modifyPro(int &a)
{
    a = 56;
}

int main()
{
    // Write C++ code here
    std::cout << "Try programiz.pro" << std::endl;
    int a = 90;
    std::cout << a << std::endl;
    modifyPro(a);
    std::cout << a << std::endl;

    return 0;
}

*/
