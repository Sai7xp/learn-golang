package main

import (
	"fmt"
	"sync"
)

func SingleProducerSingleConsumer() {
	fmt.Println("SP SC - Single Producer & Single Consumer")
	pipe := make(chan string)

	/*
		We can implement this without using WaitGroup as well, first start a consumer1 goroutine
		and run the producer without any goroutine. Producer for loop will run, in each loop it adds a messages
		and paralley consumer go routine will consume it.

		✋ BUT BUT, the problem is the last message will be missed,
		producer will add last message to channel and then immediately func exits, before even consumer prints the msg to console
	*/

	go producer1(pipe)

	var wg sync.WaitGroup
	wg.Add(1) // once consumer completes consuming all the messages it should do wg.Done()
	go consumer1(pipe, &wg)

	wg.Wait()
}

func producer1(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Message %d from Single Producer", i+1)
	}
	close(ch)
}

func consumer1(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// below for loop will break when the ch is closed and all the data is read from it
	for eachMsg := range ch {
		// time.Sleep(time.Second)
		fmt.Println(eachMsg)
	}
}
