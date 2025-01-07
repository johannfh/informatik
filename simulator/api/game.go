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

	r.Post("/water/change-by", gc.handleAddWater)
	r.Get("/water/current", gc.handleGetWater)

	return r
}

func (gc *GameController) handleAddWater(w http.ResponseWriter, r *http.Request) {
	amountParam := r.URL.Query().Get("amount")

	if amountParam == "" {
		w.Write([]byte("Missing 'amount' parameter!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountParam, 64)
	if err != nil {
		w.Write([]byte("Invalid 'amount' parameter!"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	oldAmount := gc.Game.GetWater()

	gc.Game.SetWater(gc.Game.GetWater() + amount)

	w.Write([]byte(fmt.Sprintf("Changed water by %f! Water is now at %f", gc.Game.GetWater()-oldAmount, gc.Game.GetWater())))
	w.WriteHeader(http.StatusOK)
}
func (gc *GameController) handleGetWater(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Water currently at %f!", gc.Game.GetWater())))
	w.WriteHeader(http.StatusOK)
}
