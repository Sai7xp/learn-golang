package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1️⃣.Sequential Approach to Calculate the Fibonacci
*/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	start := time.Now()
	fmt.Println(fib(46))
	fmt.Println(fib(46))
	fmt.Println("Sequential Approach - Time taken:", time.Since(start)) // takes >10 secs 🤯

	concurrentApproach()
}

/*
2️⃣. Concurrent approach
*/

func concurrentApproach() {
	// runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	start := time.Now()
	// will use max available cpu cores
	// Increment the WaitGroup counter for each goroutine
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println(fib(46))
	}()

	go func() {
		defer wg.Done()
		fmt.Println(fib(46))
	}()

	wg.Wait()

	fmt.Println("Concurrent Approach - Time Taken: ", time.Since(start)) // takes around 6 secs (because of two concurrent goroutines)

}
