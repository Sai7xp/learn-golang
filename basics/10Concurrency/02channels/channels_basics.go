/*
* Created on 13 March 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channels in Go Lang")

	ChannelsExample()

	SelectStatement()

	// we can read from the channel even after it is closed
	ch := make(chan int, 1) // buffered channel
	ch <- 1
	close(ch)
	data, ok := <-ch
	// this ok will return false where there's no data in channel and we still try to read it
	fmt.Println(data, " ", ok)

	BufferedChannels()

	UnBufferedChannels()

	// Channels Advanced
	PrintEvenOddInSync()

	// Generators : Create a stream of data using channels
	CreateDataStreamUsingChannels()
}

func BufferedChannels() {
	/// if we don't specify the buffer here as 1 - it'll throw fatal error: all goroutines are asleep - deadlock!
	dataChan := make(chan int, 1) // this is a buffered channel, we can listen to this channel in the same go routine
	dataChan <- 789               /// ADD Data to Channel
	fmt.Println(<-dataChan)       /// GET data from channel

	chanWithBuffer2 := make(chan string, 2)
	chanWithBuffer2 <- "Hello,"
	chanWithBuffer2 <- "World!"

	fmt.Println(<-chanWithBuffer2)
	fmt.Println(<-chanWithBuffer2)
	// fmt.Println(<-chanWithBuffer2) // this will cause an error : all goroutines are asleep - deadlock!
	// there are no goroutines running which can send the data to chanWithBuffer2 but it is still listening, so it's a deadlock
}

func UnBufferedChannels() {
	timeNow := time.Now()

	earnedMoney := make(chan float32) // Unbuffered channel, as soon as we add the data to chan, there should be someone who is listening to it

	go func() {
		for i := 0; i < 10; i++ {
			earnedMoney <- 34.9*0.2 + float32(i)
			time.Sleep(time.Second / 2)
		}
		// if we don't close the channel - we will get fatal error: all goroutines are asleep - deadlock!
		close(earnedMoney)
	}()

	for n := range earnedMoney {
		fmt.Println("New Earned Money: ", n)
	}
	/*
	 When will for range loops used on channel end ?
	 1. When the channel is closed using close()
	 2. All buffered values have been read after closing.
	*/

	fmt.Println("channelsWithGoRoutines finished in  ", time.Since(timeNow))
}

func ChannelsExample() {
	fmt.Println(1)

	goAheadChan := make(chan bool)
	go func() {
		time.Sleep(time.Second * 1)
		close(goAheadChan)
	}()
	<-goAheadChan // this will wait until there's a data in channel or channel closes

	fmt.Println(4)
}

/*
Select statement is used to wait on multiple channels simultaneously
Usecases - implement timeouts, cancelation signals

	select {
		case data := <-dataChan:
			fmt.Println("Got data:", data)
		case <-ctx.Done():
			fmt.Println("Context cancelled:", ctx.Err())
	}

	select {
		case res := <-someChannel:
			fmt.Println("Got result:", res)
		case <-time.After(2 * time.Second):
			fmt.Println("Timed out!")
	}
*/
func SelectStatement() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "channel 1 data"
	}()

	go func() {
		chan2 <- "channel 2 data"
	}()

	for i := 0; i < 2; i++ {
		// select statement will only wait until one of the case is executed,
		// so if we want to get the data from both the channels we have to use the above for loop
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
