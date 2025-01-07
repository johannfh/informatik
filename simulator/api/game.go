package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/johannfh/go-utils/assert"
	"github.com/johannfh/informatik/simulator/game"
)

type GameController struct {
	Context context.Context
	Game    *game.Game
}

func (gc GameController) CreateRouter() chi.Router {
	assert.NotNil(gc.Context, "missing Context in GameController")
	assert.NotNil(gc.Game, "missing *Game in GameController")

	r := chi.NewRouter()

	r.Post("/water/change-by", Make(gc.handleAddWater))
	r.Get("/water/current", Make(gc.handleGetWater))

	return r
}

func (gc *GameController) handleAddWater(w http.ResponseWriter, r *http.Request) error {
	amountParam := r.URL.Query().Get("amount")

	if amountParam == "" {
		return InvalidRequestData(map[string]string{
			"amount": "missing parameter (float)",
		})
	}

	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		return InvalidRequestData(map[string]string{
			"amount": "invalid format (float)",
		})
	}

	oldAmount := gc.Game.GetWater()

	gc.Game.SetWater(gc.Game.GetWater() + amount)

	w.Write([]byte(fmt.Sprintf("Changed water by %f! Water is now at %f", gc.Game.GetWater()-oldAmount, gc.Game.GetWater())))
	return nil
}

type getWaterResponse struct {
	StatusCode int `json:"statusCode"`
	Data       int `json:"water"`
}

func (gc *GameController) handleGetWater(w http.ResponseWriter, r *http.Request) error {
	res := getWaterResponse{
		StatusCode: http.StatusOK,
		Data:       int(gc.Game.GetWater()),
	}
	return writeJSON(w, http.StatusOK, res)
}
