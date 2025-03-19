package main

import (
	"fmt"
	"sync"
)

func SingleProducerSingleConsumer() {
	fmt.Println("Single Producer & Single Consumer")
	pipe := make(chan string)

	/*
		We can implement this without using WaitGroup as well, first start a consumer1 goroutine
		and run the producer without any goroutine. Producer for loop will run, in each loop it adds a messages
		and paralley consumer go routine will consume it.

		✋ BUT BUT, the problem is the last message will be missed,
		producer will add last message to channel and then immediately func exits, before even consumer prints the msg to console
	*/
	var wg sync.WaitGroup

	wg.Add(1) // once producer completes producing all the message it should do wg.Done
	go producer1(pipe, &wg)

	go consumer1(pipe)

	wg.Wait()
	close(pipe)
}

func producer1(ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Message from Producer %d", i+1)
	}
	wg.Done()
}

func consumer1(ch <-chan string) {
	for eachMsg := range ch {
		fmt.Println(eachMsg)
	}
}
