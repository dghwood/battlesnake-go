package board

import (
	models "github.com/dghwood/battlesnake-go/models"
	"testing"
)

func getSampleBoard() Board {
	return ParseState(
		models.GameState{
			Board: models.Board{
				Height:  11,
				Width:   11,
				Food:    []models.Coord{{X: 0, Y: 0}, {X: 10, Y: 5}},
				Hazards: []models.Coord{{X: 5, Y: 5}, {X: 2, Y: 2}},
				Snakes: []models.Battlesnake{
					{
						Length: 3,
						Head:   models.Coord{X: 1, Y: 1},
						Body: []models.Coord{
							{X: 1, Y: 1},
							{X: 0, Y: 1},
							{X: 0, Y: 2},
						},
					},
				},
			},
		},
	)
}

func TestParseBoard(t *testing.T) {

	board := getSampleBoard()

	if !board.Squares[0][0].HasFood {
		t.Error("No food at 0,0")
	}
	if !board.Squares[10][5].HasFood {
		t.Error("No food at 10, 5")
	}
	if !board.Squares[1][1].HasSnake {
		t.Error("No Snake at 1,1")
	}
	if board.Squares[1][1].SnakeHeadLength != 3 {
		t.Error("No Snake Head Length 3 at 1,1")
	}

}

func TestAvailableMoves(t *testing.T) {
	board := getSampleBoard()
	moves := board.AvaiableMoves(models.Coord{X: 1, Y: 1})
	// {1 0} {2 1} {1 2}
	if len(moves) != 3 {
		t.Error("AvailableMoves returned <> 3 moves")
	}
}
