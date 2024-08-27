package main

import "fmt"

type OtherThing struct {
	value int
}

type Thing struct {
	otherThing *OtherThing
	slice      map[int]string
}

func pointersWithStruct() {
	fmt.Println(" ----> BLOCK:START Pointers and Struct Combinations <-----")

	// memory will be allocated and all the values will be initialized to it's zero values,
	//
	// 0 for numeric types, false for booleans, "" for strings,
	// and nil for pointers, slices, maps
	thingg := new(Thing)
	fmt.Println(thingg.slice)
	thingg.slice = make(map[int]string)
	thingg.slice[8] = "hi"
	fmt.Println(thingg.slice)

	if thingg.otherThing != nil { // this if condition is very important
		thingg.otherThing.value++
		fmt.Println(thingg.otherThing)
	}

	fmt.Println(" ----> BLOCK:END Pointers and Struct Combinations <-----")
}
