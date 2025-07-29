package main

import (
	"fmt"
	"net/http"
	"strings"
)

/*
Type Alias tell the compiler : use 'Age' as just another name for 'int'
*/
type Age = int // Type Alias
type MyAge int // Type Definition

func PrintAge(x int, y Age) {
	fmt.Println("int age:", x)
	fmt.Println("Age Age:", y)
}

/*
Type Definition : Create a new type named 'Email' which behaves like a 'string', but is not the same.
*/
type Email string // Type Definition

func PrintEmail(e1 Email, e2 string) {
	fmt.Println("Email: ", e1)
	fmt.Println("normal string mail: ", e2)
}

// ✅ we can define receiver methods on our types (http.HandlerFunc is the best example for this)
func (e Email) isValid() bool {
	return strings.Contains(string(e), "@")
}

func RunTypeAliasAndTypeDef() {
	// Type Alias Usage
	var userAge Age = 9
	PrintAge(userAge, 4)
	PrintAge(4, userAge) // Age is just an alias for int, so we can pass anything

	// Type Definition usage
	var userEmail Email = "user@test.com" // or Email("user@test.com")
	PrintEmail(userEmail, "mailhere")
	// PrintEmail("mailhere", userEmail) // ❌ cannot use 'userEmail'  as string value, second param gives an error

	fmt.Println("is valid mail ? ", userEmail.isValid())

}

type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*1.8 + 32
}

/*
Best use case for type definitions - http.Handler and http.HandlerFunc
*/

// The MyHandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, MyHandlerFunc(f) is a
// [Handler] that calls f
//
// This mirrors how http.HandlerFunc works in the Go standard library
type MyHandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP allows MyHandlerFunc to satisfy the http.Handler interface
func (f MyHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

// GetUserController is a standard HTTP handler function
// It does not implement http.Handler directly, but can be adapted using MyHandlerFunc.
func GetUserController(w http.ResponseWriter, r *http.Request) {
	// your custom logic
}

func CustomListenAndServe(addr string, handler http.Handler) {}

func BestTypeDefExample() {
	// CustomListenAndServe accepts a handler which is basically an interface that has ServeHTTP(ResponseWriter, *Request) method
	// so if we want to call that method we need to create a struct and that struct should implement ServeHTTP method right ?
	// but is there any other way ? yes, type def

	// Convert our regular handler function to an http.Handler using MyHandlerFunc
	controller1 := MyHandlerFunc(GetUserController)
	CustomListenAndServe(":8999", controller1)

	// standard library provides http.HandlerFunc as a type adapter
	// a way to turn regular functions into something that satisfies the http.Handler interface

	_ = LoggerMiddleware(controller1) // middleware usage
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return MyHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Incoming request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

/*
Type Def another usecase
*/
type Dictionary map[string]string

func (d Dictionary) Search(word string) (meaning string) {
	if meaning, exists := d[word]; exists {
		return meaning
	}
	return "uh oh, word not found in the dictionary"
}

func TypeDefUsage() {
	dictionary := Dictionary{"volition": "power of using one's will."}
	dictionary.Search("volition")
}
