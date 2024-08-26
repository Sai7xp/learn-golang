/*
* Created on 26 Aug 2024
* @author Sai Sumanth
 */

package main

import "fmt"

func newVsMake() {
	fmt.Println("---------------- new vs make ---------------")
	/*
		what does make() do?
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

	fmt.Println("---------------- new vs make ---------------")
}
