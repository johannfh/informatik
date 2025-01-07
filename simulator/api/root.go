package api

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/johannfh/go-utils/assert"
	"github.com/johannfh/informatik/simulator/game"
)

type ApiServer struct {
	Context context.Context
	Game    *game.Game
}

func (srv ApiServer) CreateRouter(prefix string) chi.Router {
	assert.NotEmpty(prefix, "empty prefix")

	assert.NotNil(srv.Context, "missing Context in ApiServer")
	assert.NotNil(srv.Game, "missing *Game in ApiServer")

	r := chi.NewRouter()

	gr := GameController{
		Context: srv.Context,
		Game:    srv.Game,
	}.CreateRouter()

	r.Mount("/game", gr)
	r.Get("/hi", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi!"))
	})

	prefixedRouter := chi.NewRouter()
	prefixedRouter.Mount(prefix, r)
	return r
}
