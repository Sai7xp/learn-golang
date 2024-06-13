/*
* Created on 24 Feb 2024
* @author Sai Sumanth
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	// "strings"
	"time"
)

var workoutsList = []string{"Warm up 🏃🏻", "Bench press 💪", "Leg workout", "Hammer Curls"}

// /
type BodyBuilder struct {
	name string
	age  int
}

// Receiver function for BodyBuilder
func (man BodyBuilder) workout() {
	for _, v := range workoutsList {
		fmt.Printf("%v doing : %s\n", man.name, v)
		time.Sleep(1000 * time.Millisecond) // 1 sec of workout
	}
}

func (man BodyBuilder) smartWorkout(wg *sync.WaitGroup) {
	///
	defer wg.Done()
	for _, v := range workoutsList {
		fmt.Printf("%v doing : %s\n", man.name, v)
		time.Sleep(1000 * time.Millisecond) // 1 sec of workout
	}
}

var persons = []BodyBuilder{
	{name: "Sai", age: 22},
	{name: "Mohan", age: 24},
}

func main() {

	startingTime := time.Now()
	fmt.Println(" -------> GoRoutines in GO lang <------")
	fmt.Println("Two Persons Entered into GYM to WORKOUT")

	/// begin workout

	// 1️⃣ Method 1 of doing workout
	// persons[0].workout() // this person will take 4 secs to complete the workout
	// persons[1].workout() // takes 4 secs to complete the workout
	/// so 2 people will take 8 secs to complete the work (2nd person has to wait till p1 completes workout)

	/// 2️⃣ Method : Let's try to reduce the time by using two separate go routines #GoRoutines 😍
	go persons[0].workout()
	go persons[1].workout()
	/// surprisingly nothing will be printed to console and the program exits
	/// the main function does not wait for the Goroutines to finish their execution
	// before it exits. This means that the program may terminate before the
	// Goroutines have a chance to print anything to the console

	// 💡 Let's solve the above problem by forcefully stoping the main function from exiting for few secs
	time.Sleep(5000 * time.Millisecond)
	// 💡 or else we can remove go routine for 2nd person (then we can comment the above line)

	/// But still this is not the effective way to stop the main function to implement the go routines
	// let's make it more effective using sync.WaitGroup 🍒
	fmt.Println("This operation took: ", time.Since(startingTime))
	fmt.Println()

	// sync.WaitGroup
	goRoutinesAndWaitGroup()

	/// go routines and waitGroup example
	getStatusCodeOfWebsites()

	/// How we encounter race conditions while using go routines
	raceConditionInGo()

	/// how to solve the race conditions
	solveRaceCondUsingMutex()

}

/*
🔄  WaitGroup: it allows to block a specific code block to allow
a set of goroutines to complete execution.

It can be used to coordinate the execution of multiple coroutines
*/
func goRoutinesAndWaitGroup() {
	fmt.Println(" -------> sync.WaitGroup <------")
	startingTime := time.Now()

	var wg sync.WaitGroup

	go persons[0].smartWorkout(&wg)
	wg.Add(1) // add counter 1 to wait group (decrease this counter in go routine)
	go persons[1].smartWorkout(&wg)
	wg.Add(1)

	wg.Wait() // Wait blocks until the WaitGroup counter is zero
	fmt.Println("🥵 WORKOUT COMPLETED 🥵")
	fmt.Println("goRoutinesAndWaitGroup operation took: ", time.Since(startingTime))
	fmt.Println()
}

// /
// / Go Routines and WaitGroup in action
func getStatusCodeOfWebsites() {
	var getStatusWaitGroup sync.WaitGroup

	getWebsitesStatusStartTime := time.Now()

	websites := []string{
		"https://google.com",
		"https://github.com",
		"https://flutter.dev",
	}

	for _, eachUrl := range websites {
		// makeGetRequest(eachUrl) // this will trigger each url synchrounously one by one

		/// lets use go routines here to reduce the time
		go makeGetRequest(eachUrl, &getStatusWaitGroup)
		getStatusWaitGroup.Add(1)
	}
	getStatusWaitGroup.Wait()
	fmt.Println("getStatusCodeOfWebsites operation took: ", time.Since(getWebsitesStatusStartTime))

}

// make GET request using http package
func makeGetRequest(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	var successOrFailureEmoji string
	/// invoke url
	if res, err := http.Get(url); err != nil {
		/// 📝 Here res is of type *http.Response
		fmt.Println("Error while calling url : ", url)
		panic(err)
	} else {
		/// got response from url 🥳
		if res.StatusCode == http.StatusOK {
			successOrFailureEmoji = "✅"
			/// lets grab the body
			if dataBytes, err := io.ReadAll(res.Body); err == nil {
				fmt.Println(string(dataBytes[:250]))
				fmt.Println(res.ContentLength)
			}
		} else {
			successOrFailureEmoji = "😔"
		}
		defer res.Body.Close() // close the connection

		fmt.Printf("%s Status code for %s is %v\n", successOrFailureEmoji, url, res.StatusCode)
	}
}
