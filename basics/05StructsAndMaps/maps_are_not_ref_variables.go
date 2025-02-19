/*
* Created on  Aug 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
)

// https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
func mapsAreNotReferenceVariables() {
	fmt.Println("------------ mapsAreNotReferenceVariables() ---------------")

	// 💡 Maps are NOT reference variables - let's see that with an example
	// A map value is a pointer to a runtime.hmap structure
	modifyMap := map[int]string{0: "Default Value", 1: "Default Value"}
	fmt.Println("Map inside the main: ", modifyMap)
	mapModificationReflectionTest(modifyMap)

	// 🤯 map[0:Default Value 1:Modified]
	fmt.Println("Map inside the main after modification: ", modifyMap)

	// -----------------
	var n map[int]int
	fn(n)
	fmt.Println("n == nil ? ", n == nil) // true
	fmt.Printf("&n in main - %p\n", &n)
	fmt.Printf(" n in main - %p\n", n)

	fmt.Println()
	var m map[int]int
	fnp(&m)
	fmt.Printf("&m in main - %p\n", &m)
	fmt.Printf(" m in main - %p\n", m)
	fmt.Println("m == nil ? ", m == nil) // false
}

func mapModificationReflectionTest(m map[int]string) {
	fmt.Printf("Map received as argument: %+v\n", m)
	m[1] = "Modified" // this will modify the original argument as well 🥷
	fmt.Printf("Map after doing modification: %+v\n", m)

	// this will NOT modify the original arg. so maps are not ref variables
	m = make(map[int]string)
	fmt.Printf("Map after reassigning: %+v\n", m) // map[] 🪹 it's empty here
}

func fn(n map[int]int) {
	fmt.Printf("&n in fn - %p\n", &n)
	fmt.Printf(" n in fn - %p\n", n)
	// 🌻 if map is reference variable then the below assignment will have effect on passed argument
	n = make(map[int]int)
}

func fnp(m *map[int]int) {
	fmt.Printf("m, &m, *m -> %p %p %p\n", m, &m, *m)

	fmt.Printf("&m in fnp - %p\n", &m)
	fmt.Printf(" m in fnp - %p\n", m)
	*m = make(map[int]int)
	fmt.Printf("m, &m, *m -> %p %p %p\n", m, &m, *m)
}
