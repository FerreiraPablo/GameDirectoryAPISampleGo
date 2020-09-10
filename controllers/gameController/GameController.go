/*
GameController

Controlador que se encarga de la definicion y comportamiento de los endpoints de la libreria de juegos.

*/

package gameController

import (
	"encoding/json"
	"gameDirectoryApi/helpers/responseHelper"
	"gameDirectoryApi/models/game"
	"gameDirectoryApi/services/gameService"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

// Estructura
type GameController struct {
	service gameService.GameService
}

// Constructor
func New(router *mux.Router, service gameService.GameService) GameController {
	gameController := GameController{
		service: service,
	}
	router.HandleFunc("/Game", gameController.Get).Methods(http.MethodGet)
	router.HandleFunc("/Game/{Id}", gameController.GetById).Methods(http.MethodGet)
	router.HandleFunc("/Game", gameController.Create).Methods(http.MethodPost)
	return gameController
}

// Metodos del controlador
func (controller *GameController) Get(writer http.ResponseWriter, request *http.Request) {
	responseHelper.New(writer).Ok(controller.service.All())
}

func (controller *GameController) GetById(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id, error := strconv.Atoi(parameters["Id"])
	if error != nil {
		responseHelper.New(writer).BadRequest()
		return
	}

	game := controller.service.GetById(id)
	if game.Id == 0 {
		responseHelper.New(writer).NotFound()
		return
	}

	responseHelper.New(writer).Ok(game)
}

func (controller *GameController) DeleteById(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id, error := strconv.Atoi(parameters["Id"])
	if error != nil {
		responseHelper.New(writer).BadRequest()
		return
	}

	game := controller.service.GetById(id)
	if game.Id == 0 {
		responseHelper.New(writer).NotFound()
		return
	}

	controller.service.DeleteById(id)
	responseHelper.New(writer).NoContent()
}

func (controller *GameController) Create(writer http.ResponseWriter, request *http.Request) {
	var game game.Game
	err := json.NewDecoder(request.Body).Decode(&game)
	if err != nil {
		responseHelper.New(writer).BadRequest()
		return
	}
	controller.service.Create(game)
	responseHelper.New(writer).Created(game)
}
