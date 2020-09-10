package main

import (
	"gameDirectoryApi/contexts/gamesDatabase"
	"gameDirectoryApi/controllers/gameController"
	"gameDirectoryApi/services/gameService"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
)

func main() {
	router := mux.NewRouter()
	database := gamesDatabase.New(sqlite.Open("./game.db"))
	api := router.PathPrefix("/api").Subrouter()

	gameService := gameService.New(database)
	gameController.New(api, gameService)

	log.Fatal(http.ListenAndServe(":89", router))
}
