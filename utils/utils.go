package utils

var Three = 69

// name should start with CAPITAL Letter otherwise this Add function
// can't be imported in other packages
func Add(x, y int) (z int) {
	z = x + y
	return
}
