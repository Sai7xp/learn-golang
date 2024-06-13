/*
* Created on 24 Feb 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
	// "runtime"
	"sync"
)
// function as parameter
func filterSlice(slice []int, filterFn func(int) bool) []int {
	res := []int{}
	for _, v := range slice {
		if filterFn(v) {
			res = append(res, v)
		}
	}

	return res
}

func filterEvenNumbers() {
	fmt.Println("------> Anonymous Functions Example <--------")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	evenNumbers := filterSlice(nums, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println("Even Numbers are : ", evenNumbers)
	fmt.Println()
}

/*
below code produces a race condition because of the go routines usage
throws -> fatal error: concurrent map writes
because multiple go routines tries to write one map
*/
func raceConditionInGo() {
	/// anonymous functions example in go lang
	filterEvenNumbers()

	// let's see why we need sync.Mutex and how to solve race conditions using it
	fmt.Println("------> RACE CONDITION in GO <--------")

	wg := &sync.WaitGroup{}
	// created wait group as pointer so that we can pass it directly as refe to any other functions

	votingBox := make(map[string]int)

	wg.Add(3) // since we are using 3 go routines directly we can add 3 to Add
	go func(waitGroup *sync.WaitGroup) {
		defer wg.Done()
		/// this anonymous go routine will vote for x party
		/// 1000 people voted to x person and everytime votingBox is getting updated
		for i := 0; i < 500; i++ {
			votingBox["x"]++
		}
	}(wg)

	go func(waitGroup *sync.WaitGroup) {
		defer wg.Done()
		/// 1000 people voted to y person and everytime votingBox is getting updated
		for i := 0; i < 1000; i++ {
			votingBox["y"]++
		}

	}(wg)

	go func(waitGroup *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			votingBox["z"]++
		}
	}(wg)

	/// wait till all goroutines completes the execution
	wg.Wait()
	fmt.Printf("Vote Results are : %+v\n", votingBox)
	fmt.Println("raceConditionInGo - Multiple go routines completed their execution")
	fmt.Println()
}

// / 💡 lets solve the above Race Condition using Mutex
func solveRaceCondUsingMutex() {

	// noOfCpus := runtime.NumCPU()
	// fmt.Println("Number of CPUs ", noOfCpus)
	fmt.Println("------> MUTEX in GO <--------")

	// creating wait group as pointer so that we can pass it directly as ref to any other functions
	wg := &sync.WaitGroup{}
	var mx sync.Mutex

	votingBox := make(map[string]int)

	wg.Add(3) // since we are using 3 go routines directly we can add 3 to Add
	go func(waitGroup *sync.WaitGroup, mtx *sync.Mutex) {
		defer wg.Done()
		/// this anonymous go routine will vote for x party
		/// 1000 people voted to x person and everytime votingBox is getting updated
		for i := 0; i < 500; i++ {
			mtx.Lock()
			votingBox["x"]++
			mtx.Unlock()
		}
	}(wg, &mx)

	/// y votes
	go func(waitGroup *sync.WaitGroup, mtx *sync.Mutex) {
		defer wg.Done()
		/// 1000 people voted to y person and everytime votingBox is getting updated

		for i := 0; i < 1000; i++ {
			mtx.Lock()
			votingBox["y"]++
			mtx.Unlock()
		}

	}(wg, &mx)

	/// z party votes
	go func(waitGroup *sync.WaitGroup, mtx *sync.Mutex) {
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			mtx.Lock()
			votingBox["z"]++
			mtx.Unlock()

		}
	}(wg, &mx)

	/// wait till all goroutines completes the execution
	wg.Wait()
	fmt.Printf("Vote Results are : %+v\n", votingBox)
	fmt.Println("solveRaceCondUsingMutex - Multiple go routines completed their execution")
	fmt.Println()
}
