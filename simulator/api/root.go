package api

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/johannfh/go-utils/assert"
	"github.com/johannfh/informatik/simulator/game"
)

type WebsocketServer struct {
	Context context.Context
	Game    *game.Game
}

func (srv WebsocketServer) CreateRouter() chi.Router {
	assert.NotNil(srv.Context, "missing Context in WebsocketServer")
	assert.NotNil(srv.Game, "missing *Game in WebsocketServer")

	r := chi.NewRouter()

	r.Get("/", srv.getWebsocketConn)

	return r
}
