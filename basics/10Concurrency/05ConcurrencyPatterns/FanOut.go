package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
Use Fan Out pattern when a single producer want to distribute multiple jobs among n Goroutines
and process them concurrently. This pattern is useful when you have large number of tasks that can
be processed independently.

We create a fixed number of workers and then the workload will be shared among them.

This approach is more efficient than creating n goroutines for n jobs because:
1. Better Resource Management
  - Creating a goroutine for each job can lead to excessive resource usage(memory, CPU)

2. Scalability
  - we can adjust the number of workers based on the available resources (e.g., CPU cores)
*/

// jobsChan <-chan int means receive only channel, worker can't send data to jobs channel
func worker(id string, jobsChan <-chan int, results chan<- string) {
	for job := range jobsChan {
		fmt.Printf("worker %s received %d job\n", id, job)
		time.Sleep(time.Second * 2)
		results <- fmt.Sprintf("worker %s finished job %d", id, job)
	}
}
func FanOutExample() {
	jobsCount := 10
	workersCount := 3

	jobsChannel := make(chan int, jobsCount)
	results := make(chan string) // observer we have used unbuffered channel here

	// spin up the limited workers
	var wg sync.WaitGroup
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(workerId string) {
			defer wg.Done()
			worker(workerId, jobsChannel, results)
		}(strconv.Itoa(i + 1))
	}

	// distribute the jobs
	for job := 1; job <= jobsCount; job++ {
		jobsChannel <- job
	}
	close(jobsChannel)

	// This will cause a dead lock since results channel is a unbuffered channel and we are listening to results
	// at the end after all jobs are completed. So solve this problem use one more goroutine and put either
	// the below code in that or else put read listens login in goroutine
	// wg.Wait()
	// close(results)

	go func() {
		// waits till all the jobs are executed and then closes the results channel
		wg.Wait()
		close(results)
		fmt.Println("YAY! FINISHED EXECUTING ALL THE JOBS")
	}()

	for eachResult := range results {
		fmt.Println("Result : ", eachResult)
	}
}
