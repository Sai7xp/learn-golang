package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Question struct {
	title  string
	answer string
}

func main() {
	fmt.Println("Gophercise - Quiz")

	// go run . -csv=problems.csv -tl=25
	questionFileName, timeLimit := readFlags()

	waitForKeyPress()

	csvData, err := readCsv(questionFileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	allQuestions := parseCsvLines(csvData)
	totalQuestionsLength := len(allQuestions)
	fmt.Println("Total records/question length :", totalQuestionsLength)

	/*
		🏁 🏁 🏁 Start Quiz 🏁 🏁 🏁
	*/
	fmt.Println("Quiz has Started and will end after", timeLimit, "seconds")
	var totalScore int
	scoreChan := make(chan int)
	quizDoneStatusChan := make(chan interface{})
	defer close(scoreChan)
	defer close(quizDoneStatusChan)

	done := false
	go startQuiz(allQuestions, scoreChan, quizDoneStatusChan, timeLimit)

	for !done {
		select {
		case score := <-scoreChan:
			totalScore += score
		case <-quizDoneStatusChan:
			done = true
		}
	}
	fmt.Printf("\nYou have answered %d out of %d questions correctly\n", totalScore, totalQuestionsLength)

}

// reads data from given csv file
func readCsv(filePath string) (questions [][]string, e error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error occurred while Opening the file: ", err)
		return nil, err

	}
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println("Error occurred while parsing the csv file ", err)
		return nil, err
	}
	if len(records) == 0 {
		return nil, fmt.Errorf("no questions found in the file")
	}
	return records, nil
}

// parse given 2d array csv data into list of [Question]s
func parseCsvLines(csvLines [][]string) []Question {
	questions := make([]Question, len(csvLines))
	for i, eachRecord := range csvLines {
		questions[i] = Question{title: eachRecord[0], answer: strings.TrimSpace(eachRecord[1])}
	}
	return questions
}

func startQuiz(records []Question, score chan int, quizDoneStatusChan chan interface{}, timeLimitInSeconds int) {
	scanner := bufio.NewScanner(os.Stdin)
	timer := time.NewTimer(time.Second * time.Duration(timeLimitInSeconds))

	for i, eachQuestion := range records {
		fmt.Printf("Problem %d: %s = ", i+1, eachQuestion.title)
		userResponseChan := make(chan string)
		defer close(userResponseChan)
		go func() {
			scanner.Scan()
			userResponse := strings.TrimSpace(scanner.Text())
			userResponseChan <- userResponse
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime is Up!!!")
			quizDoneStatusChan <- "Done"
			return
		case a := <-userResponseChan:
			if a == eachQuestion.answer {
				score <- 1
			}
		}
	}
	quizDoneStatusChan <- "Done"
}

func waitForKeyPress() {
	fmt.Print("Press enter/return to start the Quizz")
	var reader = bufio.NewReader(os.Stdin)
	reader.ReadRune()
}

// returns filename and time limit parsed from command line args
func readFlags() (string, int) {
	questionFileName := flag.String("csv", "problems.csv", "a csv file for questions and answers")
	timeLimit := flag.Int("tl", 30, "Specify the time limit in seconds")
	flag.Parse()
	return *questionFileName, *timeLimit
}
