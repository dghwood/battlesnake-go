package snake

import (
	board "github.com/dghwood/battlesnake-go/board"
	models "github.com/dghwood/battlesnake-go/models"
)

type Snake struct {
	Version string
}

func (s Snake) Start(state models.GameState) {

}
func (s Snake) Move(state models.GameState) models.BattlesnakeMoveResponse {
	gameBoard := board.ParseState(state)
	you := state.You
	moves := gameBoard.AvaiableMoves(you.Head)

	// calc available moves

	return models.BattlesnakeMoveResponse{Move: board.PosToMove(moves[0], you.Head)}
}

func (s Snake) End(state models.GameState) {

}
func (s Snake) Info() models.BattlesnakeInfoResponse {
	return models.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "",        // TODO: Your Battlesnake username
		Color:      "#888888", // TODO: Choose color
		Head:       "default", // TODO: Choose head
		Tail:       "default", // TODO: Choose tail
	}
}
