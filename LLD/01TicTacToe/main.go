package main

import "fmt"

/*
Functional Requirements:
- 2 players can play the game, one after another
- 3 X 3 Board
- Row or Column or Diagonal with same symbol wins
- Check for win/draw
- Validate Input


Entities : Player, Board, Game
Player - name, symbol
Board  - size, matrix[size][size]
Game   - Board, currentPlayerIndex(0 or 1), []Player, totalMoves


Core functionalities or actions
- StartGame()
- makeMove(row,col), isValidMove()
- CheckWin(), checkDraw()

*/

func main() {

	game := NewGame()

	for {
		currentPlayer := game.Players[game.CurrentPlayerIndex]
		fmt.Println(currentPlayer.Name, "Turn", string(currentPlayer.Symbol))
		row := getInput("Enter Row (0-2): ")
		col := getInput("Enter Col (0-2): ")
		fmt.Println(row, " ", col)

		if err := game.MakeMove(row, col); err != nil {
			fmt.Println(err)
			continue
		}

		// check for win
		if game.CheckWin() {
			fmt.Println(currentPlayer.Name, "Won!!")
			break
		}

		// check for draw
		if game.IsDraw() {
			fmt.Println("DRAW!!!")
			break
		}
		game.SwitchPlayer()
	}
}

func getInput(prompt string) int {
	for {
		fmt.Print(prompt)
		var input int
		_, err := fmt.Scan(&input)
		if err == nil {
			return input
		}
		fmt.Println("Please enter valid input")
	}
}
