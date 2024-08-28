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
		new is used to allocate memory for any type, including custom types.

		Zero-Valued
		------------
		Int   	-> 0
		Float   -> 0.0
		bool    -> false
		String  -> ""

		Pointer Type  -> *T for any type T: nil
		Function Type -> nil
		Slice         -> []T nil
		Map 		  -> map[K][V] nil
		Channel Type  -> nil
		Interface     -> nll



		p := new(T)  // T can be any type

		This is commonly used when working with user-defined types (e.g., structs),
		when we want to create an instance but don’t want to initialize it with specific values.
	*/

	p := new(int)
	fmt.Println("p created using new(): ", p, "and the value after dereferencing: ", *p)

	mapP := new(map[string]int) // map is initialized to zero-value of it. nil
	fmt.Println(mapP == nil)    // false
	fmt.Println(*mapP == nil)   // true
	fmt.Println("mapP map created using new: ", mapP)
	(*mapP)["hello"] = 90 // 🚨 this will throw an error
	fmt.Println(mapP)

	type Person struct {
		Name string
		Age  int
	}

	per := new(Person) // p is of type *Person, with Name = "" and Age = 0
	fmt.Printf("default person : %#v\n", per)

	fmt.Println("---------------- new vs make ---------------")
	// Key differences:
	// 1. `new` returns a pointer, `make` returns a value
	// 2. `new` can be used with any type, `make` is only for slices, maps, and channels
	// 3. `new` only allocates memory, `make` initializes the data structure

}
