package board

import (
	models "github.com/dghwood/battlesnake-go/models"
)

type Square struct {
	HasFood         bool
	HasHazard       bool
	HasSnake        bool
	SnakeHeadLength int
}

type Board struct {
	Squares [][]Square
	Width   int
	Height  int
}

func (b *Board) GetSquare(pos models.Coord) (bool, Square) {
	if pos.X < 0 || pos.X >= b.Width || pos.Y < 0 || pos.Y >= b.Height {
		return true, Square{}
	}
	return false, b.Squares[pos.X][pos.Y]
}

func ParseState(state models.GameState) Board {
	height := state.Board.Height
	width := state.Board.Width

	// Initialize Board
	board := make([][]Square, height)
	for i := 0; i < height; i++ {
		board[i] = make([]Square, width)
	}

	// Add Food
	for _, coord := range state.Board.Food {
		board[coord.X][coord.Y].HasFood = true
	}

	// Add Hazard
	for _, coord := range state.Board.Hazards {
		board[coord.X][coord.Y].HasHazard = true
	}

	// Add Snake
	for _, snake := range state.Board.Snakes {
		for _, coord := range snake.Body {
			board[coord.X][coord.Y].HasSnake = true
		}
		// Add Snake Head & L
		board[snake.Head.X][snake.Head.Y].SnakeHeadLength = snake.Length
	}

	return Board{
		Width:   width,
		Height:  height,
		Squares: board,
	}
}

func (b *Board) AvaiableMoves(pos models.Coord) []models.Coord {
	moves := [4]models.Coord{
		{X: pos.X - 1, Y: pos.Y},
		{X: pos.X, Y: pos.Y - 1},
		{X: pos.X + 1, Y: pos.Y},
		{X: pos.X, Y: pos.Y + 1},
	}
	availableMoves := make([]models.Coord, 0)
	for _, coord := range moves {
		err, sq := b.GetSquare(coord)
		if !err && !sq.HasSnake {
			availableMoves = append(availableMoves, coord)
		}
	}
	return availableMoves
}

func PosToMove(move models.Coord, head models.Coord) string {
	if move.X > head.X {
		return "right"
	}
	if move.X < head.X {
		return "left"
	}
	if move.Y > head.Y {
		return "up"
	}
	if move.Y < head.Y {
		return "down"
	}
	return "?"
}
