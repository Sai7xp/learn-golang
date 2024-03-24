/*
* Created on 24 Feb 2024
* @author Sai Sumanth
 */

package main

import "fmt"

func main() {
	interfacesInGo()
}

// In Go, an interface is a collection of method signatures that a type can implement.
// Interfaces provide a way to specify the behavior of an object:
// what methods it should have, without defining how those methods should be implemented.
//  This promotes polymorphism and allows different types to be treated
// in a uniform way based on their behavior rather than their concrete implementation.

type Shape interface {
	/// interface can contain only method signatures
	calculateArea() float64
	toString()
}

func displayAreaInfo(shape Shape) {
	fmt.Printf("Area of %T is : %v\n", shape, shape.calculateArea())
}

// Square 🟧
type Square struct {
	side float64
}

// 🟧 Method to calculate square area
// receiver function
func (s Square) calculateArea() float64 {
	return s.side * s.side
}

// 🟧 to string method for square
func (s Square) toString() {
	fmt.Printf("Details of Square : %+v\n", s)
}

// Rectangle 🟪
type Rectangle struct {
	length float64
	width  float64
}

// 🟪 method to calculate rectangle area
func (r Rectangle) calculateArea() float64 {
	return r.length * r.width
}

// 🟪 to string method for rectangle
func (s Rectangle) toString() {
	fmt.Printf("Details Rectangle : %+v\n", s)
}

func interfacesInGo() {
	fmt.Printf("\n -------> BLOCK:START INTERFACE in Go Lang <------- \n")

	square1 := Square{20}
	displayAreaInfo(square1)

	rect := new(Rectangle) // var rect *Rectangle
	rect.length = 90
	fmt.Println(rect)
	rect2 := Rectangle{}
	var rect3 Rectangle
	fmt.Println("is rect equals to rect2? ", *rect == rect2)
	fmt.Println("is rect2 equals to rect3? ", rect2 == rect3)

	displayAreaInfo(rect)
	displayAreaInfo(rect2)
	displayAreaInfo(rect3)

	compareSquare := Square{40}
	compareSquare2 := Square{4}
	fmt.Println("is compareSquare with side 40 and compareSquare2 with side 4 are equal? ", compareSquare == compareSquare2) // false

	/// group of shapes
	multipleShapes := []Shape{
		Square{4},
		new(Rectangle),
		Square{side: 8},
		Rectangle{20, 10},
	}

	fmt.Println("Slice of type Shape")
	for _, val := range multipleShapes {
		val.toString()
	}

	fmt.Printf("\n -------> BLOCK:END INTERFACE in Go Lang <------- \n")
}
