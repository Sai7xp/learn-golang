/*
* Created on 23 Aug 2024
* @author Sai Sumanth
 */

package main

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
}

func (a Address) printAddress() {
	fmt.Printf("Street: %s, City: %s, State: %s\n", a.Street, a.City, a.State)
}

type Person struct {
	Name    string
	Age     int
	Address Address // Composition: Address is a field of Person
}

func main() {
	// RunEmbeddingUsage()

	p := Person{
		Name: "Jane",
		Age:  28,
		Address: Address{
			Street: "123 Main St",
			City:   "Gotham",
			State:  "NY",
		},
	}

	fmt.Println(p)
	fmt.Println(p.Address.City) // Accessing Address's field through Person

	// accessing the Address method
	p.Address.printAddress()

}
