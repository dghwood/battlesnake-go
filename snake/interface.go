package snake

import (
	models "github.com/dghwood/battlesnake-go/models"
)

type SnakeInterface interface {
	Start(models.GameState)
	Move(models.GameState) models.BattlesnakeMoveResponse
	End(models.GameState)
	Info() models.BattlesnakeInfoResponse
}
