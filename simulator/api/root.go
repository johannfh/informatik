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
	MsgTyp string `json:"messageType"`
}

// All messages from the server start with 'server.'
// All messages from the client start with 'client.'
const (
	// like 'a plan just died'
	MsgTypServerEvent = "server.event"
	// like '{"entities": 10, ...}'
	MsgTypServerStats = "server.stats"

	// restarts the game
	MsgTypClientRestartGame = "client.restartGame"
	// changes the water level in the game
	MsgTypClientUpdateWater = "client.updateWater"
)

// Contains every valid websocket message type
// together with their owner/user (client/server)
var MsgTypes = map[string]string{
	// server sent events
	MsgTypServerEvent: "server",
	MsgTypServerStats: "server",

	// client sent events
	MsgTypClientRestartGame: "client",
	MsgTypClientUpdateWater: "client",
}

func IsValidMsgTyp(msgTyp string) bool {
	for _, v := range MsgTypes {
		if v == msgTyp {
			return true
		}
	}
	return false
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

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("read websocket message error", "err", err)
			break
		}
		if mt == websocket.CloseMessage {
			slog.Info("websocket connection closed", "message", message)
			break
		}
		srv.processGameWssMsg(conn, message, mt)
	}

	return
}
