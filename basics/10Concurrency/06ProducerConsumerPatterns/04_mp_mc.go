package main

import (
	"fmt"
	"sync"
)

func MultipleProducersMultipleConsumers() {
	fmt.Println("MP MC - Multiple Producers Multiple Consumers")
	ch := make(chan string)

	producersCount := 3
	consumersCount := 2

	var producersWg sync.WaitGroup
	var consumersWg sync.WaitGroup

	for i := 0; i < consumersCount; i++ {
		consumersWg.Add(1)
		go consumer4(i+1, ch, &consumersWg)
	}

	for i := 0; i < producersCount; i++ {
		producersWg.Add(1)
		go producer4(i+1, ch, &producersWg)
	}

	producersWg.Wait()
	close(ch)

	consumersWg.Wait()
}

func producer4(producerId int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("Message %d by Producer %d", i+1, producerId)
	}

}

func consumer4(consumerId int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// once the ch is closed and all messages is read from channel this for loop will exit and wg.Done() will be executed
	for eachMsg := range ch {
		// time.Sleep(time.Second)
		fmt.Printf("Consumed by %d: %s\n", consumerId, eachMsg)
	}
}
