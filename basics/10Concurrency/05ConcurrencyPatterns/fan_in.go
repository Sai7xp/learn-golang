package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Imagine we have two generators and we want to consume the messages from both the channels.
*/
func FanInImplementation() {
	fmt.Println("FAN IN Pattern - Multiplexing")
	evenC := produceEvenNums()
	oddC := produceOddNums()

	// 👎 Naive Way of listening to both the streams, but odd chan will be blocked until data received from evenC
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-evenC)
	// 	fmt.Println(<-oddC)
	// }

	// 😌 Some workaround
	op1, op2 := true, true
InfiniteLoop:
	for {
		if !op1 && !op2 {
			break InfiniteLoop
		}
		select {
		case e, ok := <-evenC:
			if !ok {
				op1 = false
			} else {
				fmt.Println(e)
			}
		case o, ok := <-oddC:
			if !ok {
				op2 = false
			} else {
				fmt.Println(o)
			}
		}
	}

	fmt.Println(">>>>>>>>>>>>>>>><<<<<<<<<<<<<<<<<<")

	stream1 := produceEvenNums()
	stream2 := produceOddNums()
	stream3 := produceOddNums()

	/// 🌻 BUT, is there any better way of doing this ? YES, FAN IN Pattern
	// output := fanIn(stream1, stream2)
	output := fanInPro(stream1, stream2, stream3)
	for msg := range output {
		fmt.Println(msg)
	}

	fmt.Println("Bye 🙂")
}

// Multiplex channels into one channel, if you are not sure about channel count, create fanIn as variadic fn
func fanIn(c1 <-chan int, c2 <-chan int) <-chan int {
	multiplexedStream := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for data := range c1 {
			multiplexedStream <- data
		}
	}()

	go func() {
		defer wg.Done()
		for data := range c2 {
			multiplexedStream <- data
		}
	}()

	go func() {
		wg.Wait() // wait for both to complete and close the multiplexedStream
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func fanInPro(inputChannels ...<-chan int) <-chan int {
	multiplexedStream := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(inputChannels))

	for _, eachInput := range inputChannels {
		go func(c <-chan int) {
			defer wg.Done()
			for data := range c {
				multiplexedStream <- data
			}
		}(eachInput)
	}

	go func() {
		wg.Wait() // wait for both to complete and close the multiplexedStream
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func produceEvenNums() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for even := 2; even <= 10; even += 2 {
			c <- even
			time.Sleep(time.Second * 1)
		}
	}()
	return c
}

func produceOddNums() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for odd := 1; odd < 10; odd += 2 {
			c <- odd
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return c
}
