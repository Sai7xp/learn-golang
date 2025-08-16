package main

type Game struct {
	Board              *Board
	Players            [2]*Player
	CurrentPlayerIndex int // 0 or 1
	TotalMoves         int
}

func NewGame() *Game {
	return &Game{
		Board: NewBoard(),
		Players: [2]*Player{
			NewPlayer("Tom", 'X'),
			NewPlayer("Cat", 'O'),
		},
		CurrentPlayerIndex: 0,
		TotalMoves:         0,
	}
}

// validates the move and switches to next player if the move is valid
func (g *Game) MakeMove(row, col int) error {
	if err := g.Board.IsValidMove(row, col); err != nil {
		return err
	}
	currentPlayer := g.Players[g.CurrentPlayerIndex]
	g.Board.Grid[row][col] = currentPlayer.Symbol
	g.TotalMoves++
	g.Board.Display()
	return nil
}

func (g *Game) IsDraw() bool {
	// for i := 0; i < g.Board.Size; i++ {
	// 	for j := 0; j < g.Board.Size; j++ {
	// 		if g.Board.Grid[i][j] == ' ' {
	// 			return false
	// 		}
	// 	}
	// }
	return g.TotalMoves == (g.Board.Size * g.Board.Size)
}

func (g *Game) CheckWin() bool {
	// check rows
	for row := 0; row < g.Board.Size; row++ {
		if g.Board.Grid[row][0] != ' ' && g.Board.Grid[row][0] == g.Board.Grid[row][1] && g.Board.Grid[row][1] == g.Board.Grid[row][2] {
			return true
		}
	}

	// check cols
	for col := 0; col < g.Board.Size; col++ {
		if g.Board.Grid[0][col] != ' ' && g.Board.Grid[0][col] == g.Board.Grid[1][col] && g.Board.Grid[1][col] == g.Board.Grid[2][col] {
			return true
		}
	}

	// Diagonals
	if g.Board.Grid[0][0] != ' ' && g.Board.Grid[0][0] == g.Board.Grid[1][1] && g.Board.Grid[1][1] == g.Board.Grid[2][2] {
		return true
	}

	if g.Board.Grid[0][2] != ' ' && g.Board.Grid[0][2] == g.Board.Grid[1][1] && g.Board.Grid[1][1] == g.Board.Grid[2][0] {
		return true
	}

	return false
}

// CheckWinPro checks whether the most recent move at (row, col)
// resulted in a winning condition for the current player
//
// Instead of scanning the entire board, this optimized method only validates:
//  1. The row where the move was placed.
//  2. The column where the move was placed.
//  3. The main diagonal (top-left → bottom-right), if the move lies on it.
//  4. The anti-diagonal (top-right → bottom-left), if the move lies on it.
//
// This reduces time complexity from O(N²) (full board scan) to O(N),
func (g *Game) CheckWinPro(row, col int) bool {
	curPlayerSymbol := g.Players[g.CurrentPlayerIndex].Symbol
	n := g.Board.Size
	// check the row
	rowWin := true
	for c := 0; c < n; c++ {
		if g.Board.Grid[row][c] != curPlayerSymbol {
			rowWin = false
			break
		}
	}
	if rowWin {
		return true
	}

	// check the col
	colWin := true
	for r := 0; r < n; r++ {
		if g.Board.Grid[r][col] != curPlayerSymbol {
			colWin = false
			break
		}
	}
	if colWin {
		return true
	}

	// check main diagonal (top left to bottom right)
	if row == col {
		diagWin := true
		for i := 0; i < n; i++ {
			if g.Board.Grid[i][i] != curPlayerSymbol {
				diagWin = false
				break
			}
		}
		if diagWin {
			return true
		}
	}

	// check anti diagonal, if moves lies in it
	if row+col == n-1 {
		antiDiagWin := true
		for i := 0; i < n; i++ {
			if g.Board.Grid[i][n-i-1] != curPlayerSymbol {
				antiDiagWin = false
				break
			}
		}
		if antiDiagWin {
			return true
		}
	}
	return false
}

func (g *Game) SwitchPlayer() {
	g.CurrentPlayerIndex = 1 - g.CurrentPlayerIndex
}
