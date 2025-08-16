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

func (g *Game) SwitchPlayer() {
	g.CurrentPlayerIndex = 1 - g.CurrentPlayerIndex
}
