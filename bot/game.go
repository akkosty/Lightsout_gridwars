package bot

import "fmt"

// Game represents a Lights Out game state
type Game struct {
	ID      int
	Size    int
	Board   [][]bool
	Moves   int
	Player  string
	Started bool
}

// NewGame creates a new game with random board
func NewGame(size int) *Game {
	game := &Game{
		ID:     0,
		Size:   size,
		Board:  make([][]bool, size),
		Moves:  0,
		Player: "",
	}

	for i := 0; i < size; i++ {
		game.Board[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			game.Board[i][j] = false
		}
	}

	return game
}

// Toggle cell and its neighbors
func (g *Game) Toggle(row, col int) {
	if row >= 0 && row < g.Size && col >= 0 && col < g.Size {
		g.Board[row][col] = !g.Board[row][col]
	}

	if row > 0 {
		g.Board[row-1][col] = !g.Board[row-1][col]
	}
	if row < g.Size-1 {
		g.Board[row+1][col] = !g.Board[row+1][col]
	}
	if col > 0 {
		g.Board[row][col-1] = !g.Board[row][col-1]
	}
	if col < g.Size-1 {
		g.Board[row][col+1] = !g.Board[row][col+1]
	}
}

// IsSolved checks if all lights are off
func (g *Game) IsSolved() bool {
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.Board[i][j] {
				return false
			}
		}
	}
	return true
}

func (g *Game) String() string {
	result := fmt.Sprintf("Moves: %d\n", g.Moves)
	for i := 0; i < g.Size; i++ {
		for j := 0; j < g.Size; j++ {
			if g.Board[i][j] {
				result += "🔴 "
			} else {
				result += "🔵 "
			}
		}
		result += "\n"
	}
	return result
}
