package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithCancel() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		fmt.Println("calling cancel fn() in 2 secs")
		time.Sleep(2 * time.Second)
		fmt.Println("context cancel fn called")
		cancelFunc()
	}()

	data, err := longRunning(ctx)
	if err != nil {
		fmt.Println("long running task exited with error", err)
		// return
	}

	fmt.Println("data from long running task: ", data)
	fmt.Scanln()
}

func longRunning(ctx context.Context) (int, error) {
	// Select blocks the code until one of its cases can run.
	// It chooses one at random if multiple are ready.
	select {
	case <-ctx.Done():
		fmt.Println("cancel fn called")
		return 0, ctx.Err()

	case <-time.After(6 * time.Second):
		fmt.Println("After 6 Seconds")
		return 10, nil
	}
}
