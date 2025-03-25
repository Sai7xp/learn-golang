package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func FanOutEaxmple2() {
	jobsCount := 10
	workersCount := 3

	jobsChannel := make(chan int, jobsCount) // jobs channel is buffered because even though all the workers are busy we should be able to add jobs to the queue
	results := make(chan string)             // observe we have used unbuffered channel here

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

	// close the buffered jobsChannel
	// observe we are closing this channel as soon as we added all the jobs to channel,
	// but still workers can read data from this channel even it is closed,
	// It just that we can't send any more data once the channel is closed
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
	/*
	 When will for range loops used on channel end ?
	 1. When the channel is closed using close()
	 2. All buffered values have been read after closing.
	*/
}

// jobsChan <-chan int means receive only channel, worker can't send data to jobs channel
func worker(id string, jobsChan <-chan int, results chan<- string) {
	for job := range jobsChan {
		fmt.Printf("worker %s received %d job\n", id, job)
		time.Sleep(time.Second * 2)
		results <- fmt.Sprintf("worker %s finished job %d", id, job)
	}
}
