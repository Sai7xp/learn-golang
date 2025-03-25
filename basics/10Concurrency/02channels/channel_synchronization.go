package main

import (
	"fmt"
	"sync"
	"time"
)

func main2() {
	PrintEvenOddInSync()

	CreateDataStreamUsingChannels()
}

/*
Synchronize two go routines
*/
func PrintEvenOddInSync() {
	ch := make(chan bool) // Channel with buffer size 1
	var wg sync.WaitGroup

	wg.Add(1)
	go printEvenNumbers(ch, &wg)
	ch <- true // Initial signal to even numbers

	wg.Add(1)
	go printOddNumbers(ch, &wg)

	wg.Wait()
	close(ch)
	fmt.Println("Printing completed!")
}

func printEvenNumbers(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 10; i += 2 {
		<-ch // Wait for signal from odd numbers

		fmt.Println(i, "is even")
		time.Sleep(time.Millisecond * 500)

		// after printing 10 the last number of the series data will be sent to channel again here
		// and odd loop is already executed, so someone should listen to it. so add extra <-ch in printOdd
		ch <- true // Signal back to odd numbers
	}
}

func printOddNumbers(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i < 10; i += 2 {
		<-ch // Wait for signal from even numbers

		fmt.Println(i, "is odd")
		time.Sleep(time.Millisecond * 500)

		ch <- true // Signal back to even numbers
	}

	<-ch
	// listen to the last signal of even numbers fn, it sends signal after printing 10.
	// by the time odd numbers are already printed and above is executed
}

/*
 Creating Stream of data using Channels
*/

func CreateDataStreamUsingChannels() {
	nums := []int{2, 3, 4, 5, 6, 7, 8}
	dataStream := numbersAsStream(nums)

	output := processStream(dataStream)

	for d := range output {
		fmt.Println(d)
	}
}

// returns read only channel
func numbersAsStream(nums []int) <-chan int {
	stream := make(chan int)
	go func() {
		for _, x := range nums {
			stream <- x
		}
		close(stream)
	}()
	return stream
}

func processStream(stream <-chan int) (outputStream <-chan int) {
	out := make(chan int)
	go func() {
		defer close(out)
		for data := range stream {
			out <- data * data
			time.Sleep(time.Second * 1)
		}
	}()
	return out
}
