package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/johannfh/go-utils/assert"
	"github.com/johannfh/informatik/simulator/game"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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

type BasicMessage struct {
	MsgKind string `json:"msgKind"`
}

func (srv *WebsocketServer) getWebsocketConn(w http.ResponseWriter, r *http.Request) {
	slog.Info("websocket connection request")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("failed to upgrade connection", "err", err)
		return
	}
	defer conn.Close()
	slog.Info("websocket successfully connected")

	alive := true
	stop := func() { alive = false }

	for alive {
		srv.handleConn(conn, stop)
	}
}
