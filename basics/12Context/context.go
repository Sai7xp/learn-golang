package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("Context Package in Go")

	/*
		Context With Value
	*/
	ctx := context.Background()
	contextWithValueChannel := make(chan string)
	contextWithValue(ctx, contextWithValueChannel)
	fmt.Println(<-contextWithValueChannel)

	contextWithCancel()

	contextWithTimeout()

	contextWithTimeoutExample2()

	contextWithHTTPServer() // 🌻
}

func contextWithValue(ctx context.Context, ch chan<- string) {
	type MyContextKey string
	ctxKey := MyContextKey("userId")
	ctx = context.WithValue(ctx, ctxKey, "76")

	go func(c context.Context) {
		userID := ctx.Value(ctxKey)
		ch <- fmt.Sprintln("User ID:", userID)
	}(ctx)

}
