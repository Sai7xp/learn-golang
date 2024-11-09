package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getAllSymbols() []string {
	symbolsFreq := map[string]int{
		"🌻": 5,
		"👻": 10,
		"🍊": 15,
		"🎃": 20,
	}
	symbols := []string{}
	for k, v := range symbolsFreq {
		for v > 0 {
			symbols = append(symbols, k)
			v--
		}
	}
	return symbols
}

// performSpin returns a 2D array with 3 rows and each row contains 3 symbols
// basically a 3 x 3 array. if any of the rows contains same symbols then
// user balance will be multiplied according to the score map
// multiple row may contain same symbols
func performSpin(allSymbols []string) [][]string {
	rows, cols := 3, 3
	spinResult := [][]string{}
	// insert 3 rows
	for i := 0; i < rows; i++ {
		eachRow := []string{}
		selectedSymbolsIndex := map[int]bool{}
		for j := 0; j < cols; j++ {
			for {
				randIndex := getRandomNumber(0, len(allSymbols)-1)
				if !selectedSymbolsIndex[randIndex] {
					eachRow = append(eachRow, allSymbols[randIndex])
					selectedSymbolsIndex[randIndex] = true
					break
				}
			}
		}
		spinResult = append(spinResult, eachRow)
	}
	// now insert each symbol

	return spinResult
}

func main() {
	welcomeUser()

	// scores map
	scoreMultipliers := map[string]int{
		"🌻": 20,
		"👻": 10,
		"🍊": 5,
		"🎃": 2,
	}
	allSymbols := getAllSymbols()
	balance := uint(200)

	// main game logic
	for {
		bet := getBetFromUser(balance)
		if bet == 0 {
			fmt.Println("Good Bye.")
			break
		}
		balance -= bet
		spin := performSpin(allSymbols)
		printSpin(spin)
		wins := checkForWins(spin, scoreMultipliers)
		for i, multiplyBy := range wins {
			if multiplyBy > 0 {
				balance += (bet * uint(multiplyBy))
				fmt.Printf("Your bet multiplied by %d on line #%d.\n", multiplyBy, i+1)
			}
		}
	}

	fmt.Printf("You left with balance: %d.\n", balance)

}

func welcomeUser() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your name: ")
	sc.Scan()
	username := strings.TrimSpace(sc.Text())

	// get user's age
	var age int
	fmt.Print("Please input your age: ")
	_, err := fmt.Scan(&age)
	if err != nil || age < 18 {
		if err == nil {
			err = fmt.Errorf("sorry kiddo, you are not eligible to play the game. come back after %d years", 18-age)
		}
		log.Fatal("Error: ", err)
	}
	fmt.Printf("Hey %s, Welcome.\n", username)
}

func getBetFromUser(balance uint) uint {
	var bet uint
	for {
		if balance == 0 {
			return 0
		}
		fmt.Printf("Enter your bet, or 0 to quit the game (your balance = $%d): ", balance)
		fmt.Scan(&bet)
		if bet > balance {
			fmt.Println("Your bet can't be greater than the balance. Enter again")
		} else {
			break
		}
	}
	return bet
}
