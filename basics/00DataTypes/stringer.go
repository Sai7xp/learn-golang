package main

import (
	"fmt"
	"strings"
)

// IPAddr represents an IP address as an array of 4 bytes
type IPAddr [4]byte

// String implements the fmt.Stringer interface for IPAddr
// This method is called automatically when IPAddr is used in fmt.Println, fmt.Printf, etc.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func (ip IPAddr) String2() string {
	// usually strings.Builder is preferred over string concatenation for performance
	var sb strings.Builder
	for i, x := range ip {
		if i > 0 {
			sb.WriteString(".")
		}
		sb.WriteString(fmt.Sprint(int(x)))
	}
	return sb.String()
}

// runStringer demonstrates basic IPAddr Stringer usage
func runStringer() {
	fmt.Println("=== IPAddr Stringer Example ===")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
		"localhost": {0, 0, 0, 0},
	}

	// fmt.Printf automatically calls String() method
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	fmt.Println()
}

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, City: %s}", p.Name, p.Age, p.City)
}

type ValidationError struct {
	Field   string
	Message string
	Code    int
}

func (e ValidationError) String() string {
	return fmt.Sprintf("ValidationError[%d]: Field '%s' - %s", e.Code, e.Field, e.Message)
}
