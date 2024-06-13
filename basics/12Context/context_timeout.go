package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// Example usage of context with timeout
func contextWithTimeout() {
	fmt.Println("contextWithTimeout START")
	// Create a new context with a timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Start a goroutine that will run for more than 3 seconds
	go func() {
		time.Sleep(10 * time.Second) // set the time.sleep to 2 seconds and observe the output
		fmt.Println("Goroutine finished")
	}()

	// Wait for the goroutine to finish
	select {
	case <-ctx.Done(): // Done returns a channel
		fmt.Println("Goroutine timed out")
	}

	fmt.Println("contextWithTimeout EXIT")
}

// Another example for context.WithTimeout
func contextWithTimeoutExample2() {
	fmt.Println("contextWithTimeoutExample2 START")

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	userId := 76

	// calling third party API
	isUserActive, err := callThirdPartyAPI(ctx, userId)
	if err != nil {
		log.Fatal("Error while fetching ", userId, " active user status: ", err)

	}
	if isUserActive {
		fmt.Println(userId, "user is Active.")
	}
	fmt.Println("contextWithTimeoutExample2 EXIT")

}

func callThirdPartyAPI(ctx context.Context, userId int) (bool, error) {
	fmt.Println("Fetching active status for ", userId)

	time.Sleep(300 * time.Millisecond)
	if ctx.Err() == context.DeadlineExceeded {
		return false, errors.New("context timeout exceeded")
	}

	return true, nil
}
