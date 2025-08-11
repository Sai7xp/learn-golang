package main

import (
	"fmt"
	"learngo/utils"
	"strconv"
)

/*
🌻🌻🌻🌻🌻🌻🌻

1. Basic type: Numbers, Strings, and Boolean
2. Aggregate type: Array and Structs(Composite Type)
3. Reference type: Pointers, slices, maps, functions, and channels
4. Interface type

🌻🌻🌻🌻🌻🌻🌻
*/

// variable declarations without initializers 🪹
var java, kotlin int // 0
var c bool           // false
var python string    // ""

// Variables with Initializers 🫄
var dart, js float32 = 8.543, 9
var typescript = "BOSS" // No Need to Specify the type here

// TYPED vs UNTYPED Constants
const totalPlanets = 8     // untyped constant
var y int64 = totalPlanets // OK ✅

const moonsOfJupiter int32 = 95 // typed constant
// var z int64 = moonsOfJupiter // ❌ Error

// ❌ outside a function, every statement begins with a keyword
// (var, func, and so on) and so the := construct is not available
// flutter:=90

func main() {
	moreOnRunes()
	return
	// 1️⃣ VARIABLES
	println("Variables in GO Lang")
	v := true
	println("Bool variable : ", v)

	var i, j int = 0, 1
	i += 999
	println(i, j)

	fmt.Println("String and bool with default values : ", python, c)
	fmt.Println("int default values : ", java, kotlin)
	fmt.Println("Float Values: ", dart, js)

	fmt.Println(typescript)
	fmt.Println(totalPlanets)

	// 😎 Creating Variables using Short-Hand Notation :=
	println("Short Variables in GO Lang")

	{
		// Block Level Scope
		short := 78
		println("Short Notation: ", short)
	}

	x := 90.0988
	fmt.Println("Short Notation: ", x)
	/*
		The := operator can do one trick that you cannot do with `var`
		it allows you to assign values to existing variables, too.
		As long as there is one new variable on the lefthand side of the :=
		x := 10
		x, y := 30, "hello"
	*/
	myvar1, myvar2, myvar3 := 800, "Geeks", 47.56
	fmt.Println(myvar1, myvar2, myvar3)

	// 2️⃣ Data Types
	basicDataTypes()

	// calling a function from another package
	fmt.Println("Sum of two numbers : ", utils.Add(totalPlanets, kotlin))

	printAllSubStrings("Meow")

	differentWaysToReadInputFromUser()

	// Stringer interface
	runStringer()
}

// DATA TYPES
// Numbers, Strings, Boolean
/*
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64
*/

/*
uint8  : 0 to 255
uint16 : 0 to 65535
uint32 : 0 to 4294967295
uint64 : 0 to 18446744073709551615
int8   : -128 to 127
int16  : -32768 to 32767
int32  : -2147483648 to 2147483647
int64  : -9223372036854775808 to 9223372036854775807
*/

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	// Left Shift (a<<b) is equivalent to multiplying a with 2^b
)

func basicDataTypes() {
	// Non-ASCII literal. Go source is UTF-8.
	g := 'Σ' // rune type, an alias for int32, holds a unicode code point
	fmt.Printf("Type : %T and value : %v", g, g)
	moreOnRunes()

	var eightBitInt int8 = -128
	fmt.Println(eightBitInt)

	var oneByte byte = 255
	var Uint8Bits uint8 = 255
	Uint8Bits += 90 // (255 + 90) % 256 = 345 % 256 = 89
	fmt.Printf("oneByte Type : %T and value : %v\n", oneByte, oneByte)
	fmt.Printf("Uint8Bits Type : %T and value : %v\n", Uint8Bits, Uint8Bits)

	maxSignedInteger := 9223372036854775807
	// maxSignedInteger := 1<<63 - 1 // valid statement
	fmt.Println("Max Signed Integer : ", maxSignedInteger)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	n := int64(123)
	fmt.Println(strconv.FormatInt(n, 2)) // 1111011
}

func moreOnRunes() {
	var name = "नमस्ते"
	fmt.Println("len of name: ", len(name))
	/*
		-> len(name) returns the number of bytes in name
		-> each character may take more than a byte (UTF-8 Encoded)
		-> so use for-range loop to loop through a string instead of for..in
	*/

	for i, eachRune := range name {
		fmt.Printf("i : %v byte : %c and rune: %c\n", i, name[i], eachRune)
	}

	for i := 0; i < len(name); i++ {
		fmt.Printf("i : %v and char : %c\n", i, name[i])
	}

	// Another Example
	str := "A😊"
	fmt.Println("str len :", len(str)) // 5 bytes in UTF-8
	fmt.Println([]rune(str))           // [65 128522]
	fmt.Println(len([]rune(str)))      // 2 runes, each 4 bytes in memory
	/*
		"A" → 1 byte in UTF-8
		"😊" → 4 bytes in UTF-8
		But when converted to []rune, each rune takes 4 bytes, so []rune(s) uses 2 × 4 = 8 bytes in memory
	*/
}

func printAllSubStrings(name string) {
	fmt.Println("All Substrings of", name)
	len := len(name)
	for i := 0; i < len; i++ {
		for end := i; end < len; end++ {
			for loop := i; loop <= end; loop++ {
				fmt.Printf("%c", name[loop])
			}
			fmt.Println()
		}
	}
}
