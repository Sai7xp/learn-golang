package main

import (
	"errors"
	"fmt"
)

type Board struct {
	Size int
	Grid [3][3]rune
}

func NewBoard() *Board {
	var grid [3][3]rune
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			grid[i][j] = ' '
		}
	}
	return &Board{
		Grid: grid,
		Size: 3,
	}
}

func (b Board) IsValidMove(row, col int) error {
	if row < 0 || row >= b.Size || col < 0 || col >= b.Size {
		return fmt.Errorf("row and col must be between 0 and %d. try again", b.Size-1)
	}

	if b.Grid[row][col] != ' ' {
		return errors.New("position already marked. try again")
	}

	return nil
}

func (b Board) Display() {
	fmt.Println()
	for i := 0; i < b.Size; i++ {
		for j := 0; j < b.Size; j++ {
			fmt.Printf(" %c ", b.Grid[i][j])
			if j < b.Size-1 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < b.Size-1 {
			fmt.Println("---+---+---")
		}
	}
	fmt.Println()
}
