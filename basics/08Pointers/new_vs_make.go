/*
* Created on 26 Aug 2024
* @author Sai Sumanth
 */

package main

import "fmt"

func newVsMake() {
	fmt.Println("---------------- new vs make ---------------")
	/*
		🌻 what does make() do?
		It allocates & initialize memory for slices, maps, and channels only.
		Unlike new, make does not return a pointer.

	*/

	slice := make([]int, 0)      // For slices
	mapp := make(map[string]int) // For maps
	channel := make(chan int)    // For channels (unbuffered channel)

	fmt.Println(slice)
	fmt.Println(mapp)
	fmt.Println(channel)

	go func() {
		channel <- 9
	}()
	fmt.Println(<-channel)

	/*
		🌻 what does new() do?
		Allocates the memory to a variable and returns pointer to that memory.
		It does not initialize the memory beyond zeroing it, so the allocated memory is zero-valued.
		new is used to allocate memory for any type, including custom types

		p := new(T)  // T can be any type

		This is commonly used when working with user-defined types (e.g., structs),
		when we want to create an instance but don’t want to initialize it with specific values.
	*/

	p := new(int)
	fmt.Println("p created using new(): ", p)

	type Person struct {
		Name string
		Age  int
	}

	per := new(Person) // p is of type *Person, with Name = "" and Age = 0
	fmt.Printf("default person : %#v\n", per)

	fmt.Println("---------------- new vs make ---------------")
}
