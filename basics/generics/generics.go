/*
* Created on 29 Feb 2024
* @author Sai Sumanth
 */
package main

// https://go.dev/doc/tutorial/generics
import (
	"fmt"
)

// function for adding integers
func addIntegers(a int, b int) int {
	return a + b
}

// another function for adding float values
func addFloat(a float64, b float64) float64 {
	return a + b
}

// 😎 Generic function for adding Numbers
func genericAdd[T int | float64](a T, b T) T {
	return a + b
}

// generic func to calculate the sum of arrays of different types
func calculateSum[T int | int8 | int32 | uint64 | float32 | float64 | string](arr []T) T {
	var res T
	for _, val := range arr {
		res += val
	}
	return res
}

type Nums interface {
	int | int8 | int32 | uint64 | float32 | float64
}

// Generic function for finding product of [Nums]
func findProduct[T Nums](arr []T) T {
	var res T = 1
	for _, v := range arr {
		res *= v
	}
	return res
}

// 💎 Best use case of Generics
//
// # Normal map function which can be used only for int type
//
// Usage:
// r := mapValues(nums, func(i int) int { return i * 3 }) // [3 6 9 12 15 18 21 24 27 30]
func mapValues(values []int, operation func(i int) int) []int {
	newValues := []int{}
	valueAfterPerformingOperation := 0
	for _, v := range values {
		valueAfterPerformingOperation = operation(v)
		newValues = append(newValues, valueAfterPerformingOperation)
	}
	return newValues
}

// 😻 A Generic func which can be used on any [Nums] type
//
// Usage:
//
// floatValues := []float64{1.1, 2.2, 3.3}
// genericMapValues(floatValues, func(f float64) float64 { return f * 2 }) // [2.2 4.4 6.6]
func genericMapValues[T Nums](values []T, operation func(T) T) []T {
	newValues := []T{}
	var valueAfterPerformingOperation T
	for _, v := range values {
		valueAfterPerformingOperation = operation(v)
		newValues = append(newValues, valueAfterPerformingOperation)
	}
	return newValues
}

// entry point
func main() {
	fmt.Println("GO")
	fmt.Println(60.0 / 29)
	fmt.Println("Addition two integers ", addIntegers(1, 39))
	fmt.Println("Addition two float values ", addFloat(1.00, 39.09))
	fmt.Println("genericAdd :", genericAdd(1, 39.78))
	fmt.Println("genericAdd :", genericAdd(1, 39))

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	floatValues := []float64{1.1, 2.2, 3.3}
	/// generate 3 table
	multiplesOfThree := mapValues(nums, func(i int) int { return i * 3 })
	fmt.Println("Multiples of 3 are : ", multiplesOfThree)
	multiplesOfThree = genericMapValues(nums, func(i int) int { return i * 3 })
	fmt.Println("Multiples of 3 generated using Generic Function are : ", multiplesOfThree)
	floatMultiples := genericMapValues(floatValues, func(f float64) float64 { return f * 2 })
	fmt.Println("2 * float values generated using Generic Function are : ", floatMultiples)

	intArr := []int{1, 2, 3, 4, 5, 29}   // sum is 44
	int8Arr := []int8{1, 2, 3, 4, 5, 29} // max value can be 127
	var int8Number int8 = 127
	// value becomes 0 (once the value overflows 127 it starts again from -128)
	fmt.Println("sum of (int8)127+127+2 is: ", int8Number+127+2)

	int32Arr := []int32{1, 2, 3, 4, 5, 29}
	uint64Arr := []uint64{1, 2, 3, 4, 5, 18446744073709551412}
	float32Arr := []float32{1.2, 2.66, 3 / 2, 4.78, 5.54, 29.66, 0.54345}
	float64Arr := []float64{1, 2, 3, 4, 5, 999.0 / 29}
	strinArr := []string{"Generics", " In", " Go", " Lang!😻"}

	fmt.Println("\nA GENERIC FUNCTION TO ADD NUMBERS OF AN ARRAY")
	/// SUM of all nums in array
	fmt.Printf("Sum of %T is %v\n", intArr, calculateSum[int](intArr))
	fmt.Printf("Sum of %T is %v\n", int8Arr, calculateSum[int8](int8Arr))
	fmt.Printf("Sum of %T is %v\n", int32Arr, calculateSum[int32](int32Arr))
	fmt.Printf("Sum of %T is %v\n", uint64Arr, calculateSum[uint64](uint64Arr))
	fmt.Printf("Sum of %T is %v\n", float32Arr, calculateSum[float32](float32Arr))
	fmt.Printf("Sum of %T is %v\n", float64Arr, calculateSum[float64](float64Arr))
	/// sum operation can be used on strings as well
	fmt.Printf("Sum of %T is %v\n", strinArr, calculateSum[string](strinArr))

	fmt.Println("A GENERIC FUNCTION TO MULTIPLY NUMBERS OF AN ARRAY")
	/// Product of all nums in array
	fmt.Printf("Product of %T is %v\n", intArr, findProduct[int](intArr))
	fmt.Printf("Product of %T is %v\n", int8Arr, findProduct[int8](int8Arr))
	fmt.Printf("Product of %T is %v\n", int32Arr, findProduct[int32](int32Arr))
	fmt.Printf("Product of %T is %v\n", uint64Arr, findProduct[uint64](uint64Arr))
	fmt.Printf("Product of %T is %v\n", float32Arr, findProduct[float32](float32Arr))
	fmt.Printf("Product of %T is %v\n", float64Arr, findProduct[float64](float64Arr))
}
