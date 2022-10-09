package snake

import (
	models "github.com/dghwood/battlesnake-go/models"
	"testing"
)

func getSampleState() models.GameState {
	return models.GameState{
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
	}
}

func TestMove(t *testing.T) {
	snake := Snake{}
	move := snake.Move(getSampleState()).Move
	if move != "up" && move != "down" && move != "left" && move != "right" {
		t.Error("Move is not up|down|left|right")
	}
}
