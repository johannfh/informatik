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

const (
	MsgTypEvent = "event"
	MsgTypStats = "stats"
)

func (srv *ApiServer) getGameWebsocketConn(w http.ResponseWriter, r *http.Request) {
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

type WssResponseAbcTest struct {
}

func (srv *ApiServer) processGameWssMsg(conn *websocket.Conn, message []byte, mt int) error {
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
