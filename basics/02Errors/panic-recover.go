package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
panic() is a built-in function. It means something terribly gone wrong unexpectedly. like runtime errors.
Program will stop it's execution when panic is called.

recover() : let's you recover from the panic situations, especially when we don't want our whole app to go down
due to run time errors.

NOTE: panic/recover is kind of similar to try/catch in other languages, but panic/recover is used only
for truly unrecoverable conditions like:
  - Nil pointer dereference
  - Any runtime issues (divide by zero, index out of bounds)

We should never use panic for normal error handling.
*/
func PanicRecover() {

	nums := [...]int{1, 2, 3}
	printArray(nums)
	// this print statement gets executed even though above printArray() caused a panic,
	// because we recovered from the panic *within* printArray itself.
	// otherwise if we put the recovery logic in this function, then the below print statement does not gets executed
	fmt.Println("Do some other Computation")

	// panic & recover real world example
	runServer()
}

func printArray(nums [3]int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()

	// panic: runtime error: index out of range [3] with length 3
	// i <= len(nums) array index out of bounds
	for i := 0; i <= len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Println("Printing of given is completed.")
}

/*
 One real world example for panic and recover.
 We don't want our entire server to crash just because of one error.
 Let's see how we can use recover to keep our server alive when something goes terribly wrong
*/

func runServer() {
	mux := http.ServeMux{}
	mux.HandleFunc("GET /data", withRecovery(handler))
	// http.ListenAndServe(":7878", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Simulating a panic
	panic("Something went wrong in the handler!")
}

func withRecovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next(w, r)
	}
}
