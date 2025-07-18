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
		cancelFunc() // we can cal this cancel whenever we want
	}()

	msg, err := longRunningUploadTask(ctx)
	if err != nil {
		fmt.Println("❌ long running task exited with error: ", err)
		// return
	}
	fmt.Println("message from long running task: ", msg)
	time.Sleep(time.Second * 5)
}

func longRunningUploadTask(ctx context.Context) (string, error) {
	// Select blocks the code until one of its cases can run.
	// It chooses one at random if multiple are ready.

	for i := range 100 { // assume reading a file chunks by chunks and uploading
		select {
		case <-ctx.Done():
			fmt.Println("cancelling upload task....")
			return "", ctx.Err()
		default:
			time.Sleep(time.Millisecond * 100)
			fmt.Print(i, "..")
		}
	}
	return "Upload task completed", nil
}
