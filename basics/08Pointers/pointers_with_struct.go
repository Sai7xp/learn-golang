package main

import "fmt"

type OtherThing struct {
	value int
}

type Thing struct {
	otherThing *OtherThing
}

func pointersWithStruct() {
	fmt.Println(" ----> BLOCK:START Pointers and Struct Combinations <-----")

	thingg := Thing{}
	if thingg.otherThing != nil { // this if condition is very important
		thingg.otherThing.value++
		fmt.Println(thingg.otherThing)
	}

	fmt.Println(" ----> BLOCK:END Pointers and Struct Combinations <-----")
}
