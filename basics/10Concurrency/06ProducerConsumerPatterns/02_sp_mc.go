package main

import (
	"fmt"
	"sync"
)

func SingleProducerMultipleConsumers() {
	fmt.Println("SP MC - Single Producer & Multiple Consumers")

	consumersCount := 3
	ch := make(chan string)

	go producer2(ch)

	var wg sync.WaitGroup

	// start multiple consumers
	for i := 0; i < consumersCount; i++ {
		wg.Add(1)
		go func() {
			consumer2(i+1, ch, &wg)
		}()
	}

	wg.Wait()
}

func producer2(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Message %d from Single Producer", i+1)
	}
	close(ch)
}

// spin up multiple consumers
func consumer2(id int, ch <-chan string, wg *sync.WaitGroup) {
	// listen to the messages produced by single producer
	for eachMsg := range ch {
		// time.Sleep(time.Second * 2)
		fmt.Printf("Consumed by %d. Message %s\n", id, eachMsg)
	}
	wg.Done()
}
