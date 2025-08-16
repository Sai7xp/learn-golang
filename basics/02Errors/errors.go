package main

import (
	"errors"
	"fmt"
)

type MyInt = int  // type alias
type MyIntPro int // type definition. completely new type

/*
Errors in Go -

	🌻 https://go.dev/blog/errors-are-values 🌻
*/
var ErrNotFound = errors.New("item not found")

func main() {
	// PanicRecover()

	/*
		Errors are just values in Go, represented using the built-in `error` interface
	*/
	err := errors.New("new error - age should be > 18")
	fmt.Println("error: ", err)

	/*
		Creating Custom Errors
	*/
	var myErr error
	myErr = MyAppError{Code: 101, Msg: "custom msg here"}
	fmt.Println("custom error: ", myErr)

	/*
		Error Wrapping
		Wrapping keeps the original error inside so you can inspect it later
	*/
	baseErr := doSomeOp()
	err = fmt.Errorf("error while calling doSomeOp %w", baseErr) // Wrapping error

	/*
		errors.Is() - Unwrapping errors
	*/
	if errors.Is(err, baseErr) {
		fmt.Println("yes err includes baseErr")
	}

	/*
		errors.As()
		As finds the first error in err's tree that matches target, and if one is found,
		sets target to that error value and returns true. Otherwise, it returns false.
	*/
	err = errors.Join(err, myErr)
	var checkMyError MyAppError
	if errors.As(err, &checkMyError) {
		fmt.Println("extracted the MyAppError from err tree - ", checkMyError)
	}
}

type MyAppError struct {
	Code int
	Msg  string
}

func (e MyAppError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Msg)
}

func doSomeOp() error {
	fmt.Println("doing some operation")
	return fmt.Errorf("failed to get record with pk - %s", "user123")
}
