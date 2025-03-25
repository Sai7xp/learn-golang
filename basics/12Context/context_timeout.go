package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Another example for context.WithTimeout
func contextWithTimeout() {
	fmt.Println("contextWithTimeoutExample2 START")

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	/*
		// We can create a context with Deadline as well
		deadline := time.Parse(layout, "Feb 4, 2025 at 6:05pm (IST)")
		ctx, cancel := context.WithDeadline(context.Background(), deadline)
		defer cancel()
	*/
	userId := 76

	// calling third party API
	isUserActive, err := callThirdPartyAPI(ctx, userId)
	if err != nil {
		fmt.Println("Error while fetching ", userId, " active user status: ", err)

	}
	if isUserActive {
		fmt.Println(userId, "user is Active.")
	}
	fmt.Println("contextWithTimeoutExample2 EXIT")

}

func callThirdPartyAPI(ctx context.Context, userId int) (bool, error) {
	fmt.Println("Fetching active status for ", userId)
	select {
	case <-ctx.Done():
		return false, errors.New("context timeout")

	case <-time.After(time.Millisecond * 300):
		return true, nil
	}
}

// create a context with timeout and make list of api calls using go routines
func contextWithTimeoutExample2() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*3)
	defer cancel()

	urls := []string{
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/todos",
	}
	resChan := make(chan string)
	for _, v := range urls {
		go fetchAPI(ctx, v, resChan)
	}

	for range urls {
		fmt.Println(<-resChan)
	}
}

func fetchAPI(ctx context.Context, url string, result chan<- string) {
	// The http.NewRequestWithContext() function is used to create an HTTP request
	// with the provided context. If any of the API requests exceed the timeout duration,
	// the context's cancellation signal is propagated, canceling all other ongoing requests
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		result <- fmt.Sprintf("Error making request: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	result <- fmt.Sprintf("Response from %s: %d", url, resp.StatusCode)
}
