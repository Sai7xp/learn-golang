package main

import (
	"fmt"
	"strings"
)

// fmt.Stringer interface

type IPAddr [4]byte

func (ip IPAddr) String2() string {
	var sb strings.Builder
	for i, x := range ip {
		if i > 0 {
			sb.WriteString(".")
		}
		sb.WriteString(fmt.Sprint(int(x)))
	}
	return sb.String()
}

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

/*
We can define this String() method any struct as well,
when we do fmt.Println(personStruct) whatever we have defined in that String() receiver fn will be printed
*/

func runStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
