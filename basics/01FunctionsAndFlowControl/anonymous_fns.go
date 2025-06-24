package main

import "fmt"

func filterEvenNumbers() {
	fmt.Println("------> Anonymous Functions Example <--------")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumbers := filterSlice(nums, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println("Even Numbers are : ", evenNumbers)
	fmt.Println()
}

// function as parameter
func filterSlice(slice []int, filterFn func(int) bool) []int {
	res := []int{}
	for _, v := range slice {
		if filterFn(v) {
			res = append(res, v)
		}
	}

	return res
}
