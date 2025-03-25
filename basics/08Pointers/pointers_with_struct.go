package main

import "fmt"

type OtherThing struct {
	value int
}

type Thing struct {
	otherThing *OtherThing
	mapData    map[int]string
}

func pointersWithStruct() {
	fmt.Println(" ----> BLOCK:START Pointers and Struct Combinations <-----")

	// memory will be allocated and all the values will be initialized to it's zero values,
	//
	// 0 for numeric types, false for booleans, "" for strings,
	// and nil for pointers, slices, maps
	thingg := new(Thing)
	fmt.Println(thingg)
	thingg.mapData = make(map[int]string)
	thingg.mapData[8] = "hi"
	fmt.Println(thingg)

	thingg.otherThing = &OtherThing{}

	if thingg.otherThing != nil { // this if condition is very important
		thingg.otherThing.value++
		fmt.Println(thingg)
		fmt.Println(thingg.otherThing)
		fmt.Println("Incremented pointer value inside a struct")
	}

	fmt.Println(" ----> BLOCK:END Pointers and Struct Combinations <-----")
}
