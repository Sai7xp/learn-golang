package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	title  string
	answer string
}

func main() {
	fmt.Println("Gophercise - Quiz")

	// read arguments
	// go run . -csv=problems.csv
	questionFileName := flag.String("csv", "problems.csv", "a csv file for questions and answers")
	timeLimit := flag.Int("limit", 25, "Specify the time limit in seconds")
	fmt.Println("Quiz will end after", *timeLimit, "seconds")
	flag.Parse()

	csvData, err := readCsv(*questionFileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	allQuestions := parseCsvLines(csvData)
	totalQuestionsLength := len(allQuestions)
	fmt.Println("Total records/question length :", totalQuestionsLength)
	// fmt.Printf("First Record %v and its type %T\n", records[0], records[0])

	/*
		Start Quiz
	*/
	score := startQuiz(allQuestions)
	fmt.Printf("You have answered %d out of %d questions correctly\n", score, totalQuestionsLength)
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

func startQuiz(records []Question) (score int) {
	scanner := bufio.NewScanner(os.Stdin)

	for i, eachQuestion := range records {
		fmt.Printf("Problem %d: %s = ", i+1, eachQuestion.title)
		scanner.Scan()
		userResponse := strings.TrimSpace(scanner.Text())
		if userResponse == eachQuestion.answer {
			score++
		}
	}
	return score
}
