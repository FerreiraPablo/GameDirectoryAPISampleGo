package gamesDatabase

import (
	"gameDirectoryApi/models/game"

	"gorm.io/gorm"
)

// Constructor
func New(dialect gorm.Dialector) *gorm.DB {
	database, error := gorm.Open(dialect, &gorm.Config{})
	if error != nil {
		return nil
	}

	database.AutoMigrate(&game.Game{})
	return database
}
