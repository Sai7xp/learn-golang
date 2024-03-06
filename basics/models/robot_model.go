package models

import (
	"fmt"
	"strings"
	"time"
)

type Robot struct {
	Name             string
	IsActive         bool
	ManufacturedDate time.Time
	Power            float64 // 1 - 10

}

// / First lette should be capital to make it publicly available to other packages
func CreateRobot(name string, power float64) (r Robot) {
	r = Robot{Name: name, IsActive: true, ManufacturedDate: time.Now(), Power: power}
	return
}

/// Receiver Functions
// In Go, receiver functions, also known as methods, can be used on user-defined types.
// These are types created using the type keyword followed by a type declaration.

// here r is the receiver
func (r Robot) PrintFormattedRobotDetails() {
	res := ""
	res = strings.Join([]string{"🤖 Robot Name : ", r.Name, " and this robot is 🔩 manufactured in ", r.ManufacturedDate.Format("01/02/2006")}, "")
	fmt.Println(res)
	// return res
}
