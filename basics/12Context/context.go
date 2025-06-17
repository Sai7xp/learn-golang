package main

import (
	"context"
	"fmt"
	"time"
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

	contextWithHTTP()

	contextWithCancel()

	contextWithTimeout()

	contextWithTimeoutExample2()

}

func contextWithValue(ctx context.Context, ch chan<- string) {
	type MyContextKey string
	ctxKey := MyContextKey("userId")
	ctx = context.WithValue(ctx, ctxKey, "76")

	go func(c context.Context) {
		userID := ctx.Value(ctxKey)
		time.Sleep(5 * time.Second)
		ch <- fmt.Sprintln("User ID:", userID)
	}(ctx)

}
