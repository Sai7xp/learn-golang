/*
* Created on 20 Aug 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

func slicesDeepCopyShallowCopy() {
	fmt.Println("Slice Shallow Copy and Deep Copy")
	/*
		Assigning one slice to another (e.g., b := a) creates a shallow copy, not a deep copy
		This means that both a and b will reference the same underlying array in memory.
		Modifying elements in b will affect a, and vice versa.
	*/
	sliceA := make([]int, 5)
	sliceB := sliceA
	sliceB[3] = 99
	fmt.Println("sliceA :", sliceA) // [0 0 0 99 0] 🤯
	fmt.Println("sliceB :", sliceB) // [0 0 0 99 0]

	/*
		Assigning one slice variable to another does not make a copy of the slice contents.
		This is because a slice does not directly hold its contents.
		Instead a slice holds a pointer to its underlying array
		which holds the contents of the slice.
	*/

	orgSlice := make([]int, 5)
	orgSlice[2] = 876

	copySlice := make([]int, len(orgSlice))
	copy(copySlice, orgSlice)
	copySlice[1] = 999
	fmt.Println("orgSlice :", orgSlice)   // [0 0 876 0 0] 😌 original slice is safe
	fmt.Println("copySlice :", copySlice) // [0 999 876 0 0]

	/*
		Copying/Cloning slice using append method
	*/
	nums := []int{1, 2, 3, 4, 5}

	// copy slice using
	duplicateSlice := append(make([]int, 0), nums...)

	// modify the original slice
	nums[0] = 100
	fmt.Println("Original Slice: ", nums)            // [100, 2, 3, 4, 5]
	fmt.Println("Duplicate Slice: ", duplicateSlice) // [1, 2, 3, 4, 5] (remains unchanged)

}
