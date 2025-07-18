package main

import (
	"fmt"
	"time"
)

// Demultiplex a stream of data into multiple concurrent streams
// https://www.karanpratapsingh.com/courses/go/advanced-concurrency-patterns

/*
Fan Out is a concurrency pattern where a single stream of data is distributed to multiple Goroutines
for processing them concurrently. This pattern is useful when you have large number of tasks that can
be processed independently. Goroutines are powerful and make your program run faster.
But with great power comes great... potential for an OOM (Out Of Memory) crash.

We create a fixed number of workers and then the workload will be shared among them. This approach is
more efficient than creating n goroutines for n jobs because:

 1. Better Resource Management
    Creating a goroutine for each job can lead to excessive resource usage(memory, CPU)

 2. Scalability
    We can adjust the number of workers based on the available resources (e.g., CPU cores)
*/
func FanOutImplementation() {
	fmt.Println("-------- Fan Out Concurrency Pattern Demonstration ----------")
	timeNow := time.Now()
	defer func() {
		fmt.Println("Finished in : ", time.Since(timeNow))
	}()

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	numsStream := createStream(nums) // original input stream

	output1 := demultiplexer("worker1", numsStream)
	output2 := demultiplexer("worker2", numsStream)

InfiniteLoop:
	// Infinite Loop - listen to workers, where each worker finishes it's job, set them to nil and when all the workers are nil, exit the loop
	for {
		if output1 == nil && output2 == nil {
			break InfiniteLoop
		}
		select {
		case _, ok := <-output1:
			if !ok {
				output1 = nil
			}

		case _, ok := <-output2:
			if !ok {
				output2 = nil
			}
		}
	}
}

func createStream(nums []int) <-chan int {
	stream := make(chan int) // create a stream and return it, then later below goroutine will add data into the stream

	go func() {
		for _, num := range nums {
			stream <- num
		}
		close(stream)
	}()

	return stream
}

func demultiplexer(workerName string, inputStream <-chan int) <-chan int {
	stopSignal := make(chan int)

	go func() {
		defer close(stopSignal)
		for data := range inputStream {
			time.Sleep(time.Second * 1)
			fmt.Println("Result from ", workerName, ": ", data*data)
		}
	}()

	return stopSignal
}
