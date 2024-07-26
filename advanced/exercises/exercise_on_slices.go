package main

import "golang.org/x/tour/pic"

/// Exercise from A Tour from Go Website


func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)

	for y := range res {
		res[y] = make([]uint8, dx)
		for x := range res[y] {
			res[y][x] = uint8((x + y) /2)
		}
	}
	return res
}

func main2() {
	pic.Show(Pic)
}
