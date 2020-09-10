/*
GameService
Peque#a libreria para facilitar las respuestas a los request sobre un ResponseWrite. Similar a como ASP .NET Core lo hace.

Pablo Ferreira 20200910
*/

package gameService

import (
	"gameDirectoryApi/models/game"

	"gorm.io/gorm"
)

// Estructura
type GameService struct {
	database *gorm.DB
}

// Constructor
func New(database *gorm.DB) GameService {
	return GameService{
		database: database,
	}
}

// Metodos
func (gameService *GameService) Create(game game.Game) game.Game {
	gameService.database.Create(&game)
	gameService.database.Last(&game)
	return game
}

func (gameService *GameService) GetById(id int) game.Game {
	var game game.Game
	gameService.database.First(&game, id)
	return game
}

func (gameService *GameService) DeleteById(id int) {
	gameService.database.Delete(&game.Game{}, id)
}

func (gameService *GameService) All() []game.Game {
	var games []game.Game
	gameService.database.Find(&games)
	return games
}
