package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func getRandomNumber(min int, max int) int {
	num := rand.IntN(max-min) + min
	return num
}

func printSpin(spin [][]string) {
	for _, row := range spin {
		fmt.Println(strings.Join(row, "|"))
	}
}

func checkForWins(spin [][]string, scoresMap map[string]int) []int {
	wins := []int{}
	for _, row := range spin {
		win := true
		firstSymbol := row[0]
		for _, eachSymbol := range row[1:] {
			if firstSymbol != eachSymbol {
				win = false
				break
			}
		}
		if win {
			wins = append(wins, scoresMap[firstSymbol])
		} else {
			wins = append(wins, 0)
		}
	}
	return wins
}
