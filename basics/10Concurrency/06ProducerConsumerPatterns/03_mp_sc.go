package main

import (
	"fmt"
	"sync"
)

func MultipleProducersSingleConsumer() {
	fmt.Println("MP SC - Multiple Producers Single Consumer")

	ch := make(chan string)
	producersCount := 3

	var consumerWg sync.WaitGroup
	consumerWg.Add(1)
	go consumer3(ch, &consumerWg)

	// start multiple producers
	var wg sync.WaitGroup
	for i := 0; i < producersCount; i++ {
		wg.Add(1)
		go producer3(i+1, ch, &wg)
	}

	wg.Wait()
	close(ch)

	consumerWg.Wait()
}

func producer3(producerId int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- fmt.Sprintf("Produced by %d. Message %d", producerId, i+1)
	}
}

func consumer3(ch <-chan string, consumerWg *sync.WaitGroup) {
	for eachMsg := range ch {
		// time.Sleep(time.Second * 1)
		fmt.Println(eachMsg)
	}
	consumerWg.Done()
}
