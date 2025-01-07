package api

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WssResponseAbcTest struct {
}

func (srv *WebsocketServer) processGameWssMsg(conn *websocket.Conn, message []byte, mt int) error {
	var err error
	var response any
	defer func() {
		slog.Info(
			"processed websocket message",
			"messageType", mt,
			"received", string(message),
			"response", fmt.Sprintf("%v", response),
		)
	}()

	response = WssResponseAbcTest{}

	err = conn.WriteJSON(response)
	return err
}
